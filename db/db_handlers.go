package db

import (
	"context"
	"dbaas/helpers"
	"dbaas/model"
	"fmt"
	"strings"
)

func extraxt_value(data map[string]any) ([]string, []string, []any) {
	var values []any
	var columns []string
	var placeholders []string
	for column, column_value := range data {
		columns = append(columns, column)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)+1))
		values = append(values, column_value)
	}
	return columns, placeholders, values
}

func Insert(table string, data map[string]any) error {
	// Query example:
	// 	(INSERT INTO TABLE (userid,name,age) VALUES ($1 $2 $3)),values
	columns, placeholders, values := extraxt_value(data)
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, " ,"), strings.Join(placeholders, " ,"))
	fmt.Println(query)
	_, err := DB.Exec(context.Background(), query, values...)
	return err

}

func Create_Table(table_name string, table_details map[string]string) error {
	fmt.Println(table_name)
	query := "CREATE TABLE " + table_name + "("

	for column_name, data_type := range table_details {
		sql_data_type, exists := model.SimpleNameToSQL[data_type]
		if exists == false {
			return fmt.Errorf("DataTypes not valid")
		}
		query += fmt.Sprintf("%s %s,", column_name, sql_data_type)
		fmt.Println(column_name, sql_data_type, exists)

	}
	query = query[:len(query)-1]
	query += ");"
	fmt.Println(query)
	_, err := DB.Exec(context.Background(), query)
	return err

}

func Delete_table(table_name string) error {
	query := "DROP TABLE gopgx_schema." + table_name
	_, err := DB.Exec(context.Background(), query)
	return err
}

func DeleteRow(table_name string, condition map[string][]string) error {
	condition_query, err := helpers.Condition_extract(condition)
	if err != nil {
		return err
	}
	fmt.Println(condition_query)
	query := "DELETE FROM " + table_name + " " + condition_query
	fmt.Println(query)
	commandTag, err := DB.Exec(context.Background(), query)
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("No Rows Affected")
	}
	if err != nil {
		return err
	}
	return nil
}

func UpdateRow(table_name string, condition map[string][]string, changes map[string]any) error {
	column, placeholder, values := extraxt_value(changes)
	query, err := helpers.UpdateQuery(table_name, column, placeholder)
	if err != nil {
		return err
	}
	cndn_query, err := helpers.Condition_extract(condition)
	if err != nil {
		return err
	}
	query += cndn_query
	fmt.Println(query, values)
	_, err = DB.Exec(context.Background(), query, values...)
	if err != nil {
		return err
	}
	return nil
}
