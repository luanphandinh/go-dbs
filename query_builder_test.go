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

	query = NewQueryBuilder().
		OnSchema("company").
		From("employee").
		Where("(id = %d AND name = '%s')", 1, "Luan Phan").
		OrWhere("department_id = %d", 1)

	assertStringEquals(t,
		`SELECT *
FROM employee
WHERE (id = 1 AND name = 'Luan Phan')
OR department_id = 1`,
		query.buildQuery(),
	)
}
