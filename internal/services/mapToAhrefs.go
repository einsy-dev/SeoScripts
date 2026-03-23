package services

import (
	"domains/internal/models"
	u "domains/internal/utils"
)

func MapToAhrefs(m map[string]any, target *models.Ahrefs) {
	target.DR = u.ToUint(m["DR"])
	target.Traffic = u.ToUint(m["Traffic"])
	target.Age = u.ToUint(m["Age"])
	target.Geo = u.ToString(m["Geo"])
	target.RefDomains = u.ToUint(m["RefDomains"])
	target.OutDomains = u.ToUint(m["OutDomains"])
}
