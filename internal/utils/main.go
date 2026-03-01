package utils

import (
	"encoding/json"
	"fmt"
)

func MapToStruct(m map[string]any, target any) (any, error) {
	temp, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(temp, target)
	fmt.Println(target)
	return target, nil
}
