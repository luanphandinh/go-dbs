package dbs

import "testing"

// It's suck
// go test -cpu 1 -run none -bench . -benchtime 3s
// goos: darwin
// goarch: amd64
// pkg: github.com/luanphandinh/go-dbs
// BenchmarkQueryBuilder           20000000               293 ns/op
// BenchmarkQueryBuilderComplex     1000000              5278 ns/op
// BenchmarkRawQuery               5000000000               0.29 ns/op
// PASS
// ok      github.com/luanphandinh/go-dbs  13.599s

func doQueryBuilder() string {
	return NewQueryBuilder().
		OnSchema("company").
		Select("*, last_name as lname, fname").
		From("employee").
		GetQuery()
}

func doQueryBuilderComplex() string {
	return NewQueryBuilder().
		OnSchema("company").
		Select("*, last_name as lname, fname").
		From("employee").
		Where("id > %d", 1).
		AndWhere("name IN (%v)", []string{"Luan", "Phan"}).
		GetQuery()
}

func doRawQuery() string {
	return "SELECT *, last_name as lname, fname FROM employee"
}

// Quite good, but need to improve a lots more
func BenchmarkQueryBuilder(b *testing.B) {
	SetPlatform(sqlite3, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doQueryBuilder()
	}
}

// This benchmark is bad
// Running through the whole convert params is exhausted
func BenchmarkQueryBuilderComplex(b *testing.B) {
	SetPlatform(sqlite3, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doQueryBuilderComplex()
	}
}

// Essentially when you input your query directly
// It's way more faster
func BenchmarkRawQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doRawQuery()
	}
}
