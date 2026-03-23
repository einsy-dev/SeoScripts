package services

import (
	"domains/internal/models"
	u "domains/internal/utils"
)

func MapToMajestic(m map[string]any, target *models.Majestic) {
	target.TF = u.ToUint(m["TF"])
	target.CF = u.ToUint(m["CF"])
	target.Topic = u.ToString(m["Topic"])
}
