package main

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestDumpToCreateTableSQLStatement(t *testing.T) {
	RegisterTestingT(t)

	Expect(dumpSQLStatement("create", "table_name")).To(Equal("CREATE TABLE IF NOT EXISTS `table_name` (\n);"))
}

func TestDumpToAlterTableSQLStatement(t *testing.T) {
	RegisterTestingT(t)

	Expect(dumpSQLStatement("alter", "table_name")).To(Equal("ALTER TABLE `table_name` ;"))
}

func TestDumpToDropTableSQLStatement(t *testing.T) {
	RegisterTestingT(t)

	Expect(dumpSQLStatement("drop", "table_name")).To(Equal("DROP TABLE IF EXISTS `table_name`;"))
}
