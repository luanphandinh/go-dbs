package dbs

import "testing"

func TestQueryBuilder_BuildQuery(t *testing.T) {
	SetPlatform(sqlite3, nil)
	query := NewQueryBuilder().
		OnSchema("company").
		Select("*", "last_name as lname", "fname").
		From("employee").
		Where("employee.id = %d", 10).
		AndWhere("employee.name = '%s'", "Luan Phan")

	assertStringEquals(t,
		`SELECT *, last_name as lname, fname 
FROM employee 
WHERE employee.id = 10
AND employee.name = 'Luan Phan'`,
		query.buildQuery(),
	)
}
