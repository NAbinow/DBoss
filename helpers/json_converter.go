package helpers

import (
	"encoding/json"
	"fmt"
)

func Tojson(val []map[string]any) any {
	result, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}
