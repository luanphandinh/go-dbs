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
		Where("employee.id = %d", 10).
		AndWhere("employee.name = '%s'", "Luan").
		GetQuery()

	assertStringEquals(t,
		"SELECT *, last_name as lname, fname FROM employee WHERE employee.id = 10 AND employee.name = 'Luan'",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("(id = %d AND name = '%s')", 1, "Luan").
		OrWhere("department_id = %d", 1).
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE (id = 1 AND name = 'Luan') OR department_id = 1",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name = '%s'", "Luan").
		OrderBy("id ASC, name").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name = 'Luan' ORDER BY id ASC, name",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("id IN (%v) AND name = '%s'", []int{1, 2}, "Luan").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE id IN (1, 2) AND name = 'Luan'",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (%v)", []string{"Luan", "Phan"}).
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN ('Luan', 'Phan')",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (%v)", []string{"Luan", "Phan"}).
		Offset("10").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN ('Luan', 'Phan') OFFSET 10",
		removeSpaces(query),
	)

	query = NewQueryBuilder().
		Select("*").
		From("employee").
		Where("name IN (%v)", []string{"Luan", "Phan"}).
		Limit("10").
		Offset("10").
		GetQuery()

	assertStringEquals(t,
		"SELECT * FROM employee WHERE name IN ('Luan', 'Phan') LIMIT 10 OFFSET 10",
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
		Having("COUNT(room) > %d", 1).
		GetQuery()

	assertStringEquals(t,
		"SELECT room, COUNT(room) as c_room FROM storage GROUP BY room HAVING COUNT(room) > 1",
		removeSpaces(query),
	)
}
