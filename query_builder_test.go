package dbs

import (
	"regexp"
	"strings"
	"testing"
)

func removeSpaces(str string) string {
	space := regexp.MustCompile(`\s+`)
	return strings.TrimSpace(space.ReplaceAllString(str, " "))
}

func TestQueryBuilder_BuildQuery(t *testing.T) {
	SetPlatform(sqlite3, nil)
	query := NewQueryBuilder().
		Select("*, last_name as lname, fname").
		From("employee").
		GetQuery()

	assertStringEquals(t,
		"SELECT *, last_name as lname, fname FROM employee",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*, last_name as lname, fname").
		From("employee").
		Where("employee.id = ?").
		AndWhere("employee.name = ?").
		GetQuery()

	assertStringEquals(t,
		"SELECT *, last_name as lname, fname FROM employee WHERE employee.id = ? AND employee.name = ?",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("(id = ? AND name = ?)").
		OrWhere("department_id = ?").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE (id = ? AND name = ?) OR department_id = ?",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name = ?").
		OrderBy("id ASC, name").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name = ? ORDER BY id ASC, name",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("id IN (?) AND name = ?").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE id IN (?) AND name = ?",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (?)").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN (?)",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (?)").
		Offset("10").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN (?) OFFSET 10",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (?)").
		Limit("10").
		Offset("10").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN (?) LIMIT 10 OFFSET 10",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("room, COUNT(room) as c_room").
		From("storage").
		GroupBy("room").
		GetQuery()

	assertStringEquals(t,
		"SELECT room, COUNT(room) as c_room FROM storage GROUP BY room",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("room, COUNT(room) as c_room").
		From("storage").
		GroupBy("room").
		Having("COUNT(room) > ?", 1).
		GetQuery()

	assertStringEquals(t,
		"SELECT room, COUNT(room) as c_room FROM storage GROUP BY room HAVING COUNT(room) > ?",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee e").
		Join("department d").On("e.department_id = d.id").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee e JOIN department d ON e.department_id = d.id",
		removeSpaces(query),
	)
}
