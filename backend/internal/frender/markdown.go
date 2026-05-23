package frender

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"text/template"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/utils"
)

//go:embed templates/squad.md.tmpl
var squadMDTemplate string

//go:embed templates/all_squads.md.tmpl
var allSquadsMDTemplate string

type SquadMarkdown struct {
	Squad        *db.StudentSquadWithParticipants
	Participants []db.SquadParticipantInfo
}

func NewSquadMarkdown(squad *db.StudentSquadWithParticipants, participants []db.SquadParticipantInfo) *SquadMarkdown {
	return &SquadMarkdown{Squad: squad, Participants: participants}
}

func (r *SquadMarkdown) Name() string { return "markdown" }
func (r *SquadMarkdown) Ext() string  { return ".md" }

func (r *SquadMarkdown) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func (r *SquadMarkdown) Bytes() ([]byte, error) {
	s := r.Squad

	orgName := s.OrganizerName
	if orgName == "" {
		orgName = "—"
	}

	parts := make([]mdParticipant, len(r.Participants))
	for i, p := range r.Participants {
		parts[i] = mdParticipant{
			FullName: p.FullName,
			PhoneStr: utils.SafeString(utils.DashIfNil(p.Phone)),
			UnivStr:  utils.SafeString(p.University),
			SpecStr:  utils.SafeString(utils.DashIfNil(p.Specialty)),
			GradeStr: gradeStr(p.Grade),
		}
	}

	data := squadMDData{
		Title:           s.Title,
		OrgName:         orgName,
		OrgPosition:     ptrStr(s.OrganizerPosition),
		OrgPhone:        ptrStr(s.OrganizerPhone),
		OrgEnterprise:   ptrStr(s.OrganizerEnterpriseName),
		StatusLabel:     utils.StatusLabel(s.StatusName),
		CurrentCount:    s.CurrentCount,
		MaxParticipants: s.MaxParticipants,
		CreatedAt:       s.CreatedAt.Format("02.01.2006 15:04"),
		ApprovedByName:  s.ApprovedByName,
		Description:     ptrStr(s.Description),
		Profile:         ptrStr(s.Profile),
		Participants:    parts,
	}

	tmpl, err := template.New("squad").Funcs(mdFuncMap()).Parse(squadMDTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type mdParticipant struct {
	FullName   string
	PhoneStr   string
	UnivStr    string
	SpecStr    string
	GradeStr   string
}

type squadMDData struct {
	Title           string
	OrgName         string
	OrgPosition     string
	OrgPhone        string
	OrgEnterprise   string
	StatusLabel     string
	CurrentCount    int
	MaxParticipants int
	CreatedAt       string
	ApprovedByName  string
	Description     string
	Profile         string
	Participants    []mdParticipant
}

type SquadsMarkdown struct {
	Squads          []db.StudentSquadWithDetails
	ParticipantsMap map[uuid.UUID][]db.SquadParticipantInfo
	RequesterLabel  string
}

func NewSquadsMarkdown(squads []db.StudentSquadWithDetails, participantsMap map[uuid.UUID][]db.SquadParticipantInfo, requesterLabel string) *SquadsMarkdown {
	return &SquadsMarkdown{Squads: squads, ParticipantsMap: participantsMap, RequesterLabel: requesterLabel}
}

func (r *SquadsMarkdown) Name() string { return "markdown" }
func (r *SquadsMarkdown) Ext() string  { return ".md" }

func (r *SquadsMarkdown) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func (r *SquadsMarkdown) Bytes() ([]byte, error) {
	squads := make([]allSquadEntry, len(r.Squads))
	for i, s := range r.Squads {
		orgName := s.OrganizerName
		if orgName == "" {
			orgName = "—"
		}
		parts := r.ParticipantsMap[s.ID]
		mdParts := make([]mdParticipant, len(parts))
		for j, p := range parts {
			mdParts[j] = mdParticipant{
				FullName: p.FullName,
				PhoneStr: utils.SafeString(utils.DashIfNil(p.Phone)),
				UnivStr:  utils.SafeString(p.University),
				SpecStr:  utils.SafeString(utils.DashIfNil(p.Specialty)),
				GradeStr: gradeStr(p.Grade),
			}
		}

		squads[i] = allSquadEntry{
			Title:           s.Title,
			OrgName:         orgName,
			OrgPosition:     ptrStr(s.OrganizerPosition),
			OrgPhone:        ptrStr(s.OrganizerPhone),
			OrgEnterprise:   ptrStr(s.OrganizerEnterpriseName),
			StatusLabel:     utils.StatusLabel(s.StatusName),
			CurrentCount:    s.CurrentCount,
			MaxParticipants: s.MaxParticipants,
			CreatedAt:       s.CreatedAt.Format("02.01.2006 15:04"),
			ApprovedByName:  s.ApprovedByName,
			Description:     ptrStr(s.Description),
			Profile:         ptrStr(s.Profile),
			Participants:    mdParts,
		}
	}

	data := allSquadsMDData{
		Total:          len(r.Squads),
		Squads:         squads,
		RequesterLabel: r.RequesterLabel,
	}
	tmpl, err := template.New("allSquads").Funcs(mdFuncMap()).Parse(allSquadsMDTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type allSquadEntry struct {
	Title           string
	OrgName         string
	OrgPosition     string
	OrgPhone        string
	OrgEnterprise   string
	StatusLabel     string
	CurrentCount    int
	MaxParticipants int
	CreatedAt       string
	ApprovedByName  string
	Description     string
	Profile         string
	Participants    []mdParticipant
}

type allSquadsMDData struct {
	Total          int
	Squads         []allSquadEntry
	RequesterLabel string
}

func mdFuncMap() template.FuncMap {
	return template.FuncMap{
		"add1": func(i int) int { return i + 1 },
	}
}

func ptrStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func gradeStr(g *float64) string {
	if g == nil {
		return "—"
	}
	return fmt.Sprintf("%.0f", *g)
}

var _ Render = (*SquadMarkdown)(nil)
var _ Render = (*SquadsMarkdown)(nil)
