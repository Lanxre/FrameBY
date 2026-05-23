package frender

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/utils"
)

//go:embed templates/squad.html.tmpl
var squadHTMLTemplate string

//go:embed templates/all_squads.html.tmpl
var allSquadsHTMLTemplate string

type SquadHTML struct {
	Squad        *db.StudentSquadWithParticipants
	Participants []db.SquadParticipantInfo
}

func NewSquadHTML(squad *db.StudentSquadWithParticipants, participants []db.SquadParticipantInfo) *SquadHTML {
	return &SquadHTML{Squad: squad, Participants: participants}
}

func (r *SquadHTML) Name() string { return "html" }
func (r *SquadHTML) Ext() string  { return ".html" }

func (r *SquadHTML) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func (r *SquadHTML) Bytes() ([]byte, error) {
	data := squadHTMLData{
		Squad:        r.Squad,
		Participants: r.Participants,
		StatusLabel:  utils.StatusLabel(r.Squad.StatusName),
		CreatedAt:    r.Squad.CreatedAt.Format(time.RFC822),
	}
	if r.Squad.OrganizerPosition != nil {
		data.OrgPosition = *r.Squad.OrganizerPosition
	}
	if r.Squad.OrganizerPhone != nil {
		data.OrgPhone = *r.Squad.OrganizerPhone
	}
	if r.Squad.OrganizerEnterpriseName != nil {
		data.OrgEnterprise = *r.Squad.OrganizerEnterpriseName
	}
	if r.Squad.Description != nil {
		data.Description = *r.Squad.Description
	}
	if r.Squad.Profile != nil {
		data.Profile = *r.Squad.Profile
	}

	tmpl, err := template.New("squad").Funcs(funcMap()).Parse(squadHTMLTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type squadHTMLData struct {
	Squad         *db.StudentSquadWithParticipants
	Participants  []db.SquadParticipantInfo
	StatusLabel   string
	CreatedAt     string
	OrgPosition   string
	OrgPhone      string
	OrgEnterprise string
	Description   string
	Profile       string
}

type SquadsHTML struct {
	Squads          []db.StudentSquadWithDetails
	ParticipantsMap map[uuid.UUID][]db.SquadParticipantInfo
	RequesterLabel  string
}

func NewSquadsHTML(squads []db.StudentSquadWithDetails, participantsMap map[uuid.UUID][]db.SquadParticipantInfo, requesterLabel string) *SquadsHTML {
	return &SquadsHTML{Squads: squads, ParticipantsMap: participantsMap, RequesterLabel: requesterLabel}
}

func (r *SquadsHTML) Name() string { return "html" }
func (r *SquadsHTML) Ext() string  { return ".html" }

func (r *SquadsHTML) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func (r *SquadsHTML) Bytes() ([]byte, error) {
	data := allSquadsHTMLData{
		Squads:          r.Squads,
		ParticipantsMap: r.ParticipantsMap,
		Total:           len(r.Squads),
		RequesterLabel:  r.RequesterLabel,
	}
	tmpl, err := template.New("allSquads").Funcs(funcMap()).Parse(allSquadsHTMLTemplate)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type allSquadsHTMLData struct {
	Squads          []db.StudentSquadWithDetails
	ParticipantsMap map[uuid.UUID][]db.SquadParticipantInfo
	Total           int
	RequesterLabel  string
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"add1":      func(i int) int { return i + 1 },
		"orDash":    func(s *string) string { return utils.DashIfNil(s) },
		"gradeStr": func(g *float64) string {
			if g == nil {
				return "—"
			}
			return fmt.Sprintf("%.0f", *g)
		},
		"formatTime": func(t time.Time) string {
			return t.Format(time.RFC822)
		},
		"now": func() string {
			return time.Now().Format(time.RFC822)
		},
		"statusLabel": utils.StatusLabel,
	}
}

var _ Render = (*SquadHTML)(nil)
var _ Render = (*SquadsHTML)(nil)
