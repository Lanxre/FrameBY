package utils

import "strings"

func StatusLabel(status string) string {
	switch status {
	case "recruitment_open":
		return "Идёт набор"
	case "pending":
		return "Ожидает подтверждения"
	case "rejected":
		return "Отклонена"
	case "closed":
		return "Закрыта"
	default:
		return status
	}
}

func SanitizeFilename(name string) string {
	r := strings.NewReplacer("/", "_", "\\", "_", ":", "_", "*", "_", "?", "_", "\"", "_", "<", "_", ">", "_", "|", "_", " ", "_")
	return r.Replace(name)
}

func SafeString(s interface{}) string {
	if s == nil {
		return "—"
	}
	switch v := s.(type) {
	case string:
		if v == "" {
			return "—"
		}
		return v
	case *string:
		if v == nil || *v == "" {
			return "—"
		}
		return *v
	}
	return "—"
}

func DashIfNil(s *string) string {
	if s == nil || *s == "" {
		return "—"
	}
	return *s
}

func Truncate(s string, maxLen int) string {
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	return string(runes[:maxLen-1]) + "…"
}
