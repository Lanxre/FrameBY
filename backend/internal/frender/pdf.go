package frender

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jung-kurt/gofpdf"
	"github.com/lanxre/frameby/internal/models/db"
	"github.com/lanxre/frameby/internal/utils"
)

//go:embed fonts/LiberationSans-Regular.ttf
var fontRegular []byte

//go:embed fonts/LiberationSans-Bold.ttf
var fontBold []byte

//go:embed fonts/LiberationSans-Italic.ttf
var fontItalic []byte

//go:embed fonts/LiberationSans-BoldItalic.ttf
var fontBoldItalic []byte

const fontFamily = "LiberationSans"

type SquadPDF struct {
	Squad        *db.StudentSquadWithParticipants
	Participants []db.SquadParticipantInfo
}

func NewSquadPDF(squad *db.StudentSquadWithParticipants, participants []db.SquadParticipantInfo) *SquadPDF {
	return &SquadPDF{Squad: squad, Participants: participants}
}

func (r *SquadPDF) Name() string { return "pdf" }
func (r *SquadPDF) Ext() string  { return ".pdf" }

func (r *SquadPDF) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func newPDF() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8FontFromBytes(fontFamily, "", fontRegular)
	pdf.AddUTF8FontFromBytes(fontFamily, "B", fontBold)
	pdf.AddUTF8FontFromBytes(fontFamily, "I", fontItalic)
	pdf.AddUTF8FontFromBytes(fontFamily, "BI", fontBoldItalic)
	pdf.SetAutoPageBreak(true, 20)
	return pdf
}

func (r *SquadPDF) Bytes() ([]byte, error) {
	s := r.Squad
	pdf := newPDF()
	pdf.AddPage()

	startY := pdf.GetY()
	pad := 3.0

	pdf.SetFont(fontFamily, "B", 18)
	pdf.SetTextColor(5, 150, 105)
	pdf.CellFormat(0, 12, s.Title, "", 1, "L", false, 0, "")

	pdf.SetFont(fontFamily, "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(0, 5, fmt.Sprintf("Создан: %s • %s", s.CreatedAt.Format("02.01.2006 15:04"), s.OrganizerName), "", 1, "L", false, 0, "")
	pdf.Ln(4)

	pdf.SetFont(fontFamily, "", 10)
	pdf.SetTextColor(80, 80, 80)

	if s.OrganizerPosition != nil && *s.OrganizerPosition != "" {
		writePDFInfoRow(pdf, "Должность", *s.OrganizerPosition)
	}
	if s.OrganizerPhone != nil && *s.OrganizerPhone != "" {
		writePDFInfoRow(pdf, "Телефон", *s.OrganizerPhone)
	}
	if s.OrganizerEnterpriseName != nil && *s.OrganizerEnterpriseName != "" {
		writePDFInfoRow(pdf, "Компания", *s.OrganizerEnterpriseName)
	}
	writePDFInfoRow(pdf, "Статус", utils.StatusLabel(s.StatusName))
	writePDFInfoRow(pdf, "Участники", fmt.Sprintf("%d / %d", s.CurrentCount, s.MaxParticipants))
	writePDFInfoRow(pdf, "Создан", s.CreatedAt.Format("02.01.2006 15:04"))
	if s.ApprovedByName != "" {
		writePDFInfoRow(pdf, "Одобрено", s.ApprovedByName)
	}

	pdf.Ln(4)

	if s.Description != nil && *s.Description != "" {
		pdf.SetFont(fontFamily, "B", 12)
		pdf.SetTextColor(4, 120, 87)
		pdf.CellFormat(0, 8, "Описание", "", 1, "L", false, 0, "")
		pdf.SetFont(fontFamily, "", 10)
		pdf.SetTextColor(50, 50, 50)
		pdf.MultiCell(0, 5, *s.Description, "", "L", false)
		pdf.Ln(3)
	}

	if s.Profile != nil && *s.Profile != "" {
		pdf.SetFont(fontFamily, "B", 12)
		pdf.SetTextColor(4, 120, 87)
		pdf.CellFormat(0, 8, "Требования к участникам", "", 1, "L", false, 0, "")
		pdf.SetFont(fontFamily, "", 10)
		pdf.SetTextColor(50, 50, 50)
		pdf.MultiCell(0, 5, *s.Profile, "", "L", false)
		pdf.Ln(3)
	}

	if len(r.Participants) > 0 {
		pdf.SetFont(fontFamily, "B", 12)
		pdf.SetTextColor(4, 120, 87)
		pdf.CellFormat(0, 8, fmt.Sprintf("Участники (%d)", len(r.Participants)), "", 1, "L", false, 0, "")
		pdf.Ln(2)
		colWidths := fullWidthCols(pdf, 8, []float64{30, 10, 25, 20, 15})
		headers := []string{"#", "Имя", "Телефон", "ВУЗ", "Специальность", "Сред. балл"}
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetFont(fontFamily, "B", 9)
		pdf.SetFillColor(240, 253, 244)
		pdf.SetTextColor(6, 95, 70)
		for i, h := range headers {
			pdf.CellFormat(colWidths[i], 7, h, "1", 0, "C", true, 0, "")
		}
		pdf.Ln(-1)

		pdf.SetFont(fontFamily, "", 9)
		pdf.SetTextColor(50, 50, 50)
		lineHt := 4.5
		for i, p := range r.Participants {
			if i%2 == 0 {
				pdf.SetFillColor(249, 250, 251)
			} else {
				pdf.SetFillColor(255, 255, 255)
			}
			phone := utils.DashIfNil(p.Phone)
			spec := utils.DashIfNil(p.Specialty)
			grade := ""
			if p.Grade != nil {
				grade = fmt.Sprintf("%.0f", *p.Grade)
			}
			if grade == "" {
				grade = "—"
			}
			cells := []string{fmt.Sprintf("%d", i+1), utils.SafeString(p.FullName), utils.SafeString(phone), utils.SafeString(p.University), utils.SafeString(spec), grade}
			wrappedTableRow(pdf, colWidths, cells, lineHt)
		}
	}

	endY := pdf.GetY()
	drawBox(pdf, startY-pad, endY+pad)

	pdf.SetFont(fontFamily, "I", 8)
	pdf.SetTextColor(50, 50, 50)
	pdf.Ln(5)
	pdf.CellFormat(0, 5, fmt.Sprintf("Отчет сгенерирован • %s", time.Now().Format("02.01.2006 15:04")), "", 1, "C", false, 0, "")

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("pdf output: %w", err)
	}
	return buf.Bytes(), nil
}

type SquadsPDF struct {
	Squads          []db.StudentSquadWithDetails
	ParticipantsMap map[uuid.UUID][]db.SquadParticipantInfo
	RequesterLabel  string
}

func NewSquadsPDF(squads []db.StudentSquadWithDetails, participantsMap map[uuid.UUID][]db.SquadParticipantInfo, requesterLabel string) *SquadsPDF {
	return &SquadsPDF{Squads: squads, ParticipantsMap: participantsMap, RequesterLabel: requesterLabel}
}

func (r *SquadsPDF) Name() string { return "pdf" }
func (r *SquadsPDF) Ext() string  { return ".pdf" }

func (r *SquadsPDF) Render(ctx context.Context, outputPath string) error {
	data, err := r.Bytes()
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, data, 0644)
}

func (r *SquadsPDF) Bytes() ([]byte, error) {
	pdf := newPDF()
	pdf.AddPage()

	pdf.SetFont(fontFamily, "B", 18)
	pdf.SetTextColor(5, 150, 105)
	pdf.CellFormat(0, 12, fmt.Sprintf("Студенческие отряды • %s", time.Now().Format("02.01.2006 15:04")), "", 1, "L", false, 0, "")
	if r.RequesterLabel != "" {
		pdf.SetFont(fontFamily, "", 10)
		pdf.SetTextColor(100, 100, 100)
		pdf.CellFormat(0, 7, r.RequesterLabel, "", 1, "L", false, 0, "")
	}
	pdf.SetFont(fontFamily, "", 10)
	pdf.SetTextColor(80, 80, 80)
	pdf.CellFormat(0, 6, fmt.Sprintf("Всего: %d", len(r.Squads)), "", 1, "L", false, 0, "")
	pdf.Ln(4)

	for i, s := range r.Squads {
		parts := r.ParticipantsMap[s.ID]

		tableHt := 0.0
		if len(parts) > 0 {
			tableHt = 30.0 + float64(len(parts))*6.0
		}
		extra := 0.0
		if s.Description != nil && *s.Description != "" {
			extra += 15.0
		}
		if s.Profile != nil && *s.Profile != "" {
			extra += 15.0
		}
		needed := 50.0 + tableHt + extra

		if i > 0 && pdf.GetY()+needed > 267 {
			pdf.AddPage()
		}

		if i > 0 {
			left, _, right, _ := pdf.GetMargins()
			pageW, _ := pdf.GetPageSize()
			pdf.SetDrawColor(0, 0, 0)
			pdf.Line(left, pdf.GetY(), pageW-right, pdf.GetY())
			pdf.Ln(6)
		}

		pdf.SetFont(fontFamily, "B", 13)
		pdf.SetTextColor(4, 120, 87)
		pdf.CellFormat(0, 8, s.Title, "", 1, "L", false, 0, "")

		pdf.SetFont(fontFamily, "", 8)
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(0, 4, fmt.Sprintf("Создан: %s • %s", s.CreatedAt.Format("02.01.2006 15:04"), s.OrganizerName), "", 1, "L", false, 0, "")

		pdf.SetFont(fontFamily, "", 9)
		pdf.SetTextColor(60, 60, 60)
		if s.OrganizerPosition != nil && *s.OrganizerPosition != "" {
			writePDFInfoRow(pdf, "Должность", *s.OrganizerPosition)
		}
		if s.OrganizerPhone != nil && *s.OrganizerPhone != "" {
			writePDFInfoRow(pdf, "Телефон", *s.OrganizerPhone)
		}
		if s.OrganizerEnterpriseName != nil && *s.OrganizerEnterpriseName != "" {
			writePDFInfoRow(pdf, "Компания", *s.OrganizerEnterpriseName)
		}
		writePDFInfoRow(pdf, "Статус", utils.StatusLabel(s.StatusName))
		writePDFInfoRow(pdf, "Участники", fmt.Sprintf("%d / %d", s.CurrentCount, s.MaxParticipants))
		if s.ApprovedByName != "" {
			writePDFInfoRow(pdf, "Одобрено", s.ApprovedByName)
		}

		if s.Description != nil && *s.Description != "" {
			pdf.Ln(2)
			pdf.SetFont(fontFamily, "B", 10)
			pdf.SetTextColor(4, 120, 87)
			pdf.CellFormat(0, 6, "Описание", "", 1, "L", false, 0, "")
			pdf.SetFont(fontFamily, "", 9)
			pdf.SetTextColor(50, 50, 50)
			pdf.MultiCell(0, 4.5, *s.Description, "", "L", false)
			pdf.Ln(2)
		}

		if s.Profile != nil && *s.Profile != "" {
			pdf.Ln(1)
			pdf.SetFont(fontFamily, "B", 10)
			pdf.SetTextColor(4, 120, 87)
			pdf.CellFormat(0, 6, "Требования к участникам", "", 1, "L", false, 0, "")
			pdf.SetFont(fontFamily, "", 9)
			pdf.SetTextColor(50, 50, 50)
			pdf.MultiCell(0, 4.5, *s.Profile, "", "L", false)
			pdf.Ln(2)
		}

		if len(parts) > 0 {
			pdf.SetDrawColor(0, 0, 0)
			pdf.SetFont(fontFamily, "B", 9)
			pdf.SetFillColor(240, 253, 244)
			pdf.SetTextColor(6, 95, 70)
			colWidths := fullWidthCols(pdf, 8, []float64{30, 10, 25, 20, 15})
			pdf.CellFormat(colWidths[0], 6, "#", "1", 0, "C", true, 0, "")
			pdf.CellFormat(colWidths[1], 6, "Имя", "1", 0, "C", true, 0, "")
			pdf.CellFormat(colWidths[2], 6, "Телефон", "1", 0, "C", true, 0, "")
			pdf.CellFormat(colWidths[3], 6, "ВУЗ", "1", 0, "C", true, 0, "")
			pdf.CellFormat(colWidths[4], 6, "Специальность", "1", 0, "C", true, 0, "")
			pdf.CellFormat(colWidths[5], 6, "Сред. балл", "1", 0, "C", true, 0, "")
			pdf.Ln(-1)

			pdf.SetFont(fontFamily, "", 8)
			pdf.SetTextColor(50, 50, 50)
			lineHt := 4.0
			for i, p := range parts {
				if i%2 == 0 {
					pdf.SetFillColor(249, 250, 251)
				} else {
					pdf.SetFillColor(255, 255, 255)
				}
				phone := utils.DashIfNil(p.Phone)
				spec := utils.DashIfNil(p.Specialty)
				grade := ""
				if p.Grade != nil {
					grade = fmt.Sprintf("%.0f", *p.Grade)
				}
				if grade == "" {
					grade = "—"
				}
				cells := []string{fmt.Sprintf("%d", i+1), utils.SafeString(p.FullName), utils.SafeString(phone), utils.SafeString(p.University), utils.SafeString(spec), grade}
				wrappedTableRow(pdf, colWidths, cells, lineHt)
			}
		}
		pdf.Ln(4)
	}

	pdf.SetFont(fontFamily, "I", 8)
	pdf.SetTextColor(180, 180, 180)
	pdf.Ln(5)

	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("pdf output: %w", err)
	}
	return buf.Bytes(), nil
}

func drawBox(pdf *gofpdf.Fpdf, top, bottom float64) {
	left, _, right, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	x := left
	w := pageW - left - right
	pdf.SetDrawColor(200, 200, 200)
	pdf.Rect(x, top, w, bottom-top, "D")
}

func wrappedTableRow(pdf *gofpdf.Fpdf, colWidths []float64, texts []string, lineHt float64) {
	maxLines := 1
	for i, txt := range texts {
		lines := pdf.SplitLines([]byte(txt), colWidths[i])
		if len(lines) > maxLines {
			maxLines = len(lines)
		}
	}
	rowHt := float64(maxLines) * lineHt

	startX := pdf.GetX()
	startY := pdf.GetY()
	pdf.SetDrawColor(0, 0, 0)

	for i, txt := range texts {
		x := startX
		for j := 0; j < i; j++ {
			x += colWidths[j]
		}

		pdf.Rect(x, startY, colWidths[i], rowHt, "DF")
		pdf.SetXY(x+0.5, startY+0.5)
		pdf.MultiCell(colWidths[i]-1, lineHt, txt, "", "C", false)
	}

	pdf.SetXY(startX, startY+rowHt)
}

func tableWidth(pdf *gofpdf.Fpdf) float64 {
	left, _, right, _ := pdf.GetMargins()
	pageW, _ := pdf.GetPageSize()
	return pageW - left - right
}

func fullWidthCols(pdf *gofpdf.Fpdf, fixed float64, ratios []float64) []float64 {
	avail := tableWidth(pdf) - fixed
	totalRatio := 0.0
	for _, r := range ratios {
		totalRatio += r
	}
	cols := make([]float64, 1+len(ratios))
	cols[0] = fixed
	for i, r := range ratios {
		cols[i+1] = avail * r / totalRatio
	}
	return cols
}

func writePDFInfoRow(pdf *gofpdf.Fpdf, label, value string) {
	pdf.SetFont(fontFamily, "B", 9)
	pdf.SetTextColor(80, 80, 80)
	pdf.CellFormat(35, 5, label+":", "", 0, "L", false, 0, "")
	pdf.SetFont(fontFamily, "", 9)
	pdf.SetTextColor(50, 50, 50)
	pdf.CellFormat(0, 5, value, "", 1, "L", false, 0, "")
}

var _ Render = (*SquadPDF)(nil)
var _ Render = (*SquadsPDF)(nil)
