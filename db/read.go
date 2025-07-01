package db

import (
	"context"
	"dbaas/helpers"
	"fmt"
	"strings"
)

func Read(table string, condition map[string][]string, path string) (any, error) {
	// fmt.Println(strings.Split(table, "/"))
	conditions_list := strings.Split(path, "/")
	fmt.Println(condition)
	condition_query, err := helpers.Condition_extract(condition)
	if err != nil {
		return "", err
	}
	fmt.Println(conditions_list, condition)
	query := fmt.Sprintf("SELECT %s FROM %s %s", conditions_list[2], table, condition_query)
	fmt.Println(query)
	// err := DB.QueryRow(context.Background(), query).Scan(&data.userid, &data.name, &data.email)
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		return "", err
	}

	defer rows.Close()

	description := rows.FieldDescriptions()
	data := make([]any, len(description))
	dataptrs := make([]any, len(description))

	for i := range data {
		dataptrs[i] = &data[i]
	}

	var results []map[string]any
	for rows.Next() {
		if err := rows.Scan(dataptrs...); err != nil {
			return "", err
		}
		rowData := make(map[string]any)
		for i, desc := range description {
			rowData[string(desc.Name)] = data[i]
		}
		results = append(results, rowData)
	}
	if err := rows.Err(); err != nil {
		return "", err
	}

	// for _, row := range results {
	// 	fmt.Println(row)
	// }

	return results, nil
}
