package dbs

import (
	"fmt"
	"reflect"
	"strings"
)

// QueryBuilder create query builder
type QueryBuilder struct {
	schema     string
	selections string
	from       string
	filters    string
	groupBy    string
	order      string
	offset     int
	limit      int
	args       []interface{}

	query string
	built bool
	errs  []error
}

// NewQueryBuilder make new(QueryBuilder) along with some default config
func NewQueryBuilder() *QueryBuilder {
	builder := new(QueryBuilder)
	builder.selections = "SELECT *"

	return builder
}

// OnSchema specify schema that query will be executed on
// ex: new(QueryBuilder).OnSchema("some_schema")
func (builder *QueryBuilder) OnSchema(schema string) *QueryBuilder {
	builder.schema = strings.Trim(schema, " ")

	return builder
}

// Select specify one or more columns to be query
// eg:
//  	new(QueryBuilder).
//   		Select("*, something as something_else").
// Apply only second Select() called
func (builder *QueryBuilder) Select(selections string) *QueryBuilder {
	builder.selections = "SELECT " + selections

	return builder
}

// From specify table that query will be executed on
// ex: 	new(QueryBuilder).Select(..).From("user")
// 		new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.from = "FROM " + _platform().getSchemaAccessName(builder.schema, expression)

	return builder
}

// Where apply filter to query
// ex: builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	builder.filters += "WHERE " + expression
	builder.args = append(builder.args, args...)

	return builder
}

// AndWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.filters += " AND " + expression
	builder.args = append(builder.args, args...)

	return builder
}

// OrWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").OrWhere("age > %d", 10)
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.filters += " OR " + expression
	builder.args = append(builder.args, args...)

	return builder
}

// GroupBy apply group by in query
// ex: builder.GroupBy("name")
func (builder *QueryBuilder) GroupBy(expression string) *QueryBuilder {
	builder.groupBy = "GROUP BY " + expression

	return builder
}

// OrderBy apply order in query
// ex: builder.OrderBy("id ASC, name")
func (builder *QueryBuilder) OrderBy(expression string) *QueryBuilder {
	builder.order = "ORDER BY " + expression

	return builder
}

// Offset apply offset in query
// ex: builder.Offset(10)
func (builder *QueryBuilder) Offset(offset int) *QueryBuilder {
	builder.offset = offset

	return builder
}

// Limit apply limit in query
// ex: builder.Limit(10)
func (builder *QueryBuilder) Limit(limit int) *QueryBuilder {
	builder.limit = limit

	return builder
}

// GetQuery returns a built query
func (builder *QueryBuilder) GetQuery() string {
	if ! builder.built {
		builder.query = builder.buildQuery()
		builder.built = true
	}

	return builder.query
}

func (builder *QueryBuilder) logError(err error) *QueryBuilder {
	if err != nil {
		builder.errs = append(builder.errs, err)
	}

	return builder
}

// build SQL Query declaration
func (builder *QueryBuilder) buildQuery() string {
	declaration := builder.selections + " " +
		builder.from + " " +
		builder.filters + " " +
		builder.groupBy + " " +
		builder.order + " " +
		_platform().getPagingDeclaration(builder.limit, builder.offset)

	if args := builder.args; len(args) > 0 {
		return fmt.Sprintf(declaration, parseArgs(args[0:])...)
	}

	return declaration
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
