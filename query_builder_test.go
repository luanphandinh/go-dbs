package dbs

import "testing"

func TestQueryBuilder_BuildQuery(t *testing.T) {
	SetPlatform(sqlite3, nil)
	query := NewQueryBuilder().
		OnSchema("company").
		Select("*", "last_name as lname", "fname").
		From("employee").
		Where("employee.id = %d", 10).
		AndWhere("employee.name = '%s'", "Luan").
		GetQuery()

	assertStringEquals(t,
		`SELECT *, last_name as lname, fname
FROM employee
WHERE employee.id = 10
AND employee.name = 'Luan'`,
		query,
	)

	query = NewQueryBuilder().
		OnSchema("company").
		From("employee").
		Where("(id = %d AND name = '%s')", 1, "Luan").
		OrWhere("department_id = %d", 1).
		GetQuery()

	assertStringEquals(t,
		`SELECT *
FROM employee
WHERE (id = 1 AND name = 'Luan')
OR department_id = 1`,
		query,
	)

	query = NewQueryBuilder().
		OnSchema("company").
		From("employee").
		Where("name = '%s'", "Luan").
		OrderBy("id ASC", "name").
		GetQuery()

	assertStringEquals(t,
		`SELECT *
FROM employee
WHERE name = 'Luan'
ORDER BY id ASC, name`,
		query,
	)
}
