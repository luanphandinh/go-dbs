package dbs

import "testing"

// It's suck
// go test -cpu 1 -run none -bench . -benchtime 3s
// goos: darwin
// goarch: amd64
// pkg: github.com/luanphandinh/go-dbs
// BenchmarkQueryBuilder           50000000                99.8 ns/op           128 B/op          1 allocs/op
// BenchmarkQueryBuilderComplex    10000000               646 ns/op             576 B/op          5 allocs/op
// BenchmarkRawQuery               5000000000               0.29 ns/op            0 B/op          0 allocs/op
// PASS
// ok      github.com/luanphandinh/go-dbs  14.068s

func doQueryBuilder() string {
	return NewQueryBuilder().
		Select("*, last_name as lname, fname").
		From("employee").
		GroupBy("last_name").
		OrderBy("id DESC").
		Offset("10").
		Limit("10").
		buildSql()
}

func doQueryBuilderComplex() string {
	return NewQueryBuilder().
		Select("*, last_name as lname, fname").
		From("employee").
		GroupBy("last_name").
		OrderBy("id DESC").
		Where("id > %d", 1).
		AndWhere("name = '%s'", "Luan").
		Offset("10").
		Limit("10").
		buildSql()
}

func doRawQuery() string {
	return "SELECT *, last_name as lname, fname FROM employee"
}

// Quite good, but need to improve a lots more
func BenchmarkQueryBuilder(b *testing.B) {
	SetPlatform(sqlite3, nil)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		doQueryBuilder()
	}
}

// This benchmark is bad
// Running through the whole convert params is exhausted
// The more custom param, the more resource consuming
func BenchmarkQueryBuilderComplex(b *testing.B) {
	SetPlatform(sqlite3, nil)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		doQueryBuilderComplex()
	}
}

// Essentially when you input your query directly
// It's way more faster
func BenchmarkRawQuery(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		doRawQuery()
	}
}
