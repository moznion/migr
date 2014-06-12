package main

import "fmt"

func dumpSQLStatement(command, tableName string) string {
	switch command {
	case "create":
		return buildCreateSQL(tableName)
	case "alter":
		return buildAlterSQL(tableName)
	case "drop":
		return buildDropSQL(tableName)
	}

	return ""
}

func buildCreateSQL(tableName string) string {
	statement := ""
	statement += fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (\n);", tableName)
	return statement
}

func buildAlterSQL(tableName string) string {
	statement := ""
	statement += fmt.Sprintf("ALTER TABLE `%s`;", tableName)
	return statement
}

func buildDropSQL(tableName string) string {
	statement := ""
	statement += fmt.Sprintf("DROP TABLE IF EXISTS `%s`;", tableName)
	return statement
}
