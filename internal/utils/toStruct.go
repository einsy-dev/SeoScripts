package utils

import (
	"encoding/json"
	"fmt"
)

func ToTarget(m map[string]any, target *any) {
	temp, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(temp, target)
}
