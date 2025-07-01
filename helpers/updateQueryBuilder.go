package helpers

import "fmt"

func UpdateQuery(table_name string, columns []string, placeholders []string) (string, error) {
	query := fmt.Sprintf("UPDATE %s SET ", table_name)
	if len(placeholders) != len(columns) {
		return "", fmt.Errorf("Bad Request")
	}
	for i := 0; i < len(placeholders); i++ {
		query += columns[i] + " = " + placeholders[i] + ","
	}
	query = query[:len(query)-1] + " "
	return query, nil
}
