package dbs

import (
	"fmt"
	"reflect"
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
	args []interface{}

	// sql for select query
	sql []byte
}

// Using byte to concat string is faster
func (builder *QueryBuilder) appendClause(clause Clause, expression string) {
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

// Select specify one or more columns to be query
// eg: new(QueryBuilder).Select("*, something as something_else").
func (builder *QueryBuilder) Select(selections string) *QueryBuilder {
	builder.appendClause(SELECT, selections)

	return builder
}

// From specify table that query will be executed on
// ex: 	new(QueryBuilder).Select(..).From("user")
// 		new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.appendClause(FROM, expression)

	return builder
}

// Where apply filter to query
// ex: builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(WHERE, expression)
	builder.args = append(builder.args, args...)

	return builder
}

// AndWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(AND, expression)
	builder.args = append(builder.args, args...)

	return builder
}

// OrWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").OrWhere("age > %d", 10)
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(OR, expression)
	builder.args = append(builder.args, args...)

	return builder
}

// GroupBy apply group by in query
// ex: builder.GroupBy("name")
func (builder *QueryBuilder) GroupBy(expression string) *QueryBuilder {
	builder.appendClause(GROUP_BY, expression)

	return builder
}

// Having apply having clause in query
// ex: builder.Having("age > 20")
func (builder *QueryBuilder) Having(expression string, args ...interface{}) *QueryBuilder {
	builder.appendClause(HAVING, expression)
	builder.args = append(builder.args, args...)

	return builder
}

// OrderBy apply order in query
// ex: builder.OrderBy("id ASC, name")
func (builder *QueryBuilder) OrderBy(expression string) *QueryBuilder {
	builder.appendClause(ORDER_BY, expression)

	return builder
}

// Offset apply offset in query
// ex: builder.Offset(10)
func (builder *QueryBuilder) Offset(offset string) *QueryBuilder {
	builder.appendClause(OFFSET, offset)

	return builder
}

// Limit apply limit in query
// ex: builder.Limit(10)
func (builder *QueryBuilder) Limit(limit string) *QueryBuilder {
	builder.appendClause(LIMIT, limit)

	return builder
}

// GetQuery returns a built query
func (builder *QueryBuilder) GetQuery() string {
	return builder.buildSql()
}

// This func should be call at the very end of building process
// This converts a slice of builder.sql bytes to a string without incurring overhead
func (builder *QueryBuilder) sqlByteToString() string {
	return bytesToString(builder.sql)
}

func (builder *QueryBuilder) buildSql() string {
	// Using this cause a really bad performance
	// TODO: Need a faster solution
	if args := builder.args; len(args) > 0 {
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
