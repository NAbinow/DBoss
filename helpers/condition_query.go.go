package helpers

import (
	"fmt"
	"strings"
)

var operatorMap = map[string]string{
	"eq":   "=",
	"ne":   "<>",
	"neq":  "<>",
	"gt":   ">",
	"lt":   "<",
	"gte":  ">=",
	"lte":  "<=",
	"like": "LIKE",
	"in":   "IN",
	"nin":  "NOT IN",
}

func Condition_extract(query map[string][]string) (string, error) {
	var conditions []string

	for key, values := range query {
		parts := strings.SplitN(key, "_", 2)
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid filter key: %s", key)
		}

		column := parts[0]
		opKey := parts[1]
		op, ok := operatorMap[opKey]
		if !ok {
			return "", fmt.Errorf("unsupported operator: %s", opKey)
		}

		raw := values[0]

		if op == "IN" || op == "NOT IN" {
			items := strings.Split(raw, ",")
			var quotedItems []string
			for _, item := range items {
				quotedItems = append(quotedItems, fmt.Sprintf("'%s'", strings.TrimSpace(item)))
			}
			conditions = append(conditions, fmt.Sprintf("%s %s (%s)", column, op, strings.Join(quotedItems, ", ")))
		} else {
			conditions = append(conditions, fmt.Sprintf("%s %s '%s'", column, op, raw))
		}
	}

	if len(conditions) == 0 {
		return "", nil
	}

	return "WHERE " + strings.Join(conditions, " AND "), nil
}
