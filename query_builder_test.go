package dbs

import "testing"

func TestQueryBuilder_BuildQuery(t *testing.T) {
	query := new(QueryBuilder).
		OnSchema("company").
		Select("*", "last_name as lname", "fname").
		From("employee")

	SetPlatform(mssql, nil)
	assertStringEquals(t, "SELECT *, last_name as lname, fname \nFROM company.employee", query.buildQuery())

	SetPlatform(postgres, nil)
	assertStringEquals(t, "SELECT *, last_name as lname, fname \nFROM company.employee", query.buildQuery())

	SetPlatform(sqlite3, nil)
	assertStringEquals(t, "SELECT *, last_name as lname, fname \nFROM employee", query.buildQuery())

	SetPlatform(mysql, nil)
	assertStringEquals(t, "SELECT *, last_name as lname, fname \nFROM employee", query.buildQuery())
	assertStringEquals(t, "SELECT * \nFROM employee", query.Select("*").buildQuery())
}
