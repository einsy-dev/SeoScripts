package utils

import (
	"domains/internal/models"
	"encoding/json"
	"fmt"
)

func MapToTarget(m map[string]any, target *models.Domain) {
	temp, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(temp, target)
}

func StructToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	return result, err
}
