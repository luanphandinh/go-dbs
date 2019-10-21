package dbs

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Clause int

const (
	// This is used as a shortcut for "sqlClauses" declaration
	// Make it faster to query and concat sql string
	// Indicates the index of specific clause in sqlClauses array
	EMPTY Clause = iota - 1
	SELECT
	FROM
	WHERE
	AND
	OR
	GROUP_BY
	HAVING
	ORDER_BY
	LIMIT
	OFFSET
)

// All possible sqlClauses that are supported in this packages
// Access through constants defined above
var sqlClauses = [10][]byte{
	[]byte(" SELECT "),
	[]byte(" FROM "),
	[]byte(" WHERE "),
	[]byte(" AND "),
	[]byte(" OR "),
	[]byte(" GROUP BY "),
	[]byte(" HAVING "),
	[]byte(" ORDER BY "),
	[]byte(" LIMIT "),
	[]byte(" OFFSET "),
}

// QueryBuilder create query builder
type QueryBuilder struct {
	schema     string
	offset     string
	limit      string
	filterArgs []interface{}
	havingArgs []interface{}

	// sql for select query
	sql []byte
}

// Using byte to concat string is faster
func (builder *QueryBuilder) appendClause(clause Clause, expression []byte) {
	if clause > EMPTY {
		builder.sql = append(builder.sql[:], sqlClauses[clause][:]...)
	}

	builder.sql = append(builder.sql[:], expression[:]...)
}

// NewQueryBuilder make new(QueryBuilder) along with some default config
func NewQueryBuilder() *QueryBuilder {
	builder := new(QueryBuilder)

	// Since all sqlClauses have len of 64
	// It better that we initialize length for sql as 64 * 2 = 128
	builder.sql = make([]byte, 0, 128)

	return builder
}

// OnSchema specify schema that query will be executed on
// ex: new(QueryBuilder).OnSchema("some_schema")
// postgresql and mssql required schema
// This function is used particularly for query that involve schema access
// See more on From()
func (builder *QueryBuilder) OnSchema(schema string) *QueryBuilder {
	builder.schema = schema

	return builder
}

// Select specify one or more columns to be query
// eg:
//  	new(QueryBuilder).
//   		Select("*, something as something_else").
// Apply only second Select() called
func (builder *QueryBuilder) Select(selections string) *QueryBuilder {
	builder.appendClause(SELECT, []byte(selections))

	return builder
}

// From specify table that query will be executed on
// ex: 	new(QueryBuilder).Select(..).From("user")
// 		new(QueryBuilder).Select(..).From("user as u")
// Some platforms like postgresql and mssql is using schema access name
// Using _platform().getSchemaAccessName(builder.schema, expression)
// is a temporary hack for accessing correct resource
// but you can use new(QueryBuilder).Select(..).From("<schema_name>.<table_name> as <alias>") as a replace
// and don't call OnSchema().
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.appendClause(FROM, []byte(_platform().getSchemaAccessName(builder.schema, expression)))

	return builder
}

// Where apply filter to query
// ex: builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(WHERE, []byte(expression))
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// AndWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(AND, []byte(expression))
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// OrWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").OrWhere("age > %d", 10)
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(OR, []byte(expression))
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// GroupBy apply group by in query
// ex: builder.GroupBy("name")
func (builder *QueryBuilder) GroupBy(expression string) *QueryBuilder {
	builder.appendClause(GROUP_BY, []byte(expression))

	return builder
}

// Having apply having clause in query
// ex: builder.Having("age > 20")
func (builder *QueryBuilder) Having(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(HAVING, []byte(expression))
	builder.havingArgs = args

	return builder
}

// OrderBy apply order in query
// ex: builder.OrderBy("id ASC, name")
func (builder *QueryBuilder) OrderBy(expression string) *QueryBuilder {
	builder.appendClause(ORDER_BY, []byte(expression))

	return builder
}

// Offset apply offset in query
// ex: builder.Offset(10)
func (builder *QueryBuilder) Offset(offset string) *QueryBuilder {
	builder.appendClause(OFFSET, []byte(offset))

	return builder
}

// Limit apply limit in query
// ex: builder.Limit(10)
func (builder *QueryBuilder) Limit(limit string) *QueryBuilder {
	builder.appendClause(LIMIT, []byte(limit))

	return builder
}

// GetQuery returns a built query
func (builder *QueryBuilder) GetQuery() string {
	return builder.buildSql()
}

// This func should be call at the very end of building process
// This converts a slice of builder.sql bytes to a string without incurring overhead
func (builder *QueryBuilder) sqlByteToString() string {
	return *(*string)(unsafe.Pointer(&builder.sql))
}

func (builder *QueryBuilder) buildSql() string {
	// Using this cause a really bad performance
	// TODO: Need a faster solution
	if args := append(builder.filterArgs, builder.havingArgs...); len(args) > 0 {
		return fmt.Sprintf(builder.sqlByteToString(), parseArgs(args[0:])...)
	}

	return builder.sqlByteToString()
}

func parseArgs(args []interface{}) []interface{} {
	for i := 0; i < len(args); i++ {
		args[i] = parseArg(args[i])
	}

	return args
}

func parseArg(arg interface{}) interface{} {
	rt := reflect.TypeOf(arg)
	switch rt.Kind() {
	case reflect.Slice:
		return getContentOutOfArraySyntax(arg)
	case reflect.Array:
		return getContentOutOfArraySyntax(arg)
	}

	return arg
}
