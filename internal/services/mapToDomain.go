package services

import (
	"domains/internal/models"
	u "domains/internal/utils"
)

func MapToDomain(m map[string]any, target *models.Domain) {
	target.Comment = u.ToString(m["Comment"])
	target.Type = u.ToString(m["Type"])
	target.Rel = u.ToString(m["Rel"])

	if a, ok := m["Ahrefs"].(map[string]any); ok {
		MapToAhrefs(a, &target.Ahrefs)
	}
	if s, ok := m["Semrush"].(map[string]any); ok {
		MapToSemrush(s, &target.Semrush)
	}
	if maj, ok := m["Majestic"].(map[string]any); ok {
		MapToMajestic(maj, &target.Majestic)
	}
}
