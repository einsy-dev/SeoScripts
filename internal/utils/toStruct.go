package utils

import (
	"encoding/json"
	"log"
)

func ToTarget(m map[string]any, target *any) {
	temp, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(temp, target)
}
