package services

import (
	"domains/internal/models"
	u "domains/internal/utils"
)

func MapToSemrush(m map[string]any, target *models.Semrush) {
	target.AS = u.ToUint(m["AS"])
	target.Traffic = u.ToUint(m["Traffic"])
	target.RefDomains = u.ToUint(m["RefDomains"])
	target.OutDomains = u.ToUint(m["OutDomains"])
	target.LinkFarm = u.ToString(m["LinkFarm"])
}
