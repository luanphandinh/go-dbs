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
	having     string
	order      string
	offset     string
	limit      string
	filterArgs []interface{}
	havingArgs []interface{}

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
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// AndWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.filters += " AND " + expression
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// OrWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").OrWhere("age > %d", 10)
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	builder.filters += " OR " + expression
	builder.filterArgs = append(builder.filterArgs, args...)

	return builder
}

// GroupBy apply group by in query
// ex: builder.GroupBy("name")
func (builder *QueryBuilder) GroupBy(expression string) *QueryBuilder {
	builder.groupBy = "GROUP BY " + expression

	return builder
}

// Having apply having clause in query
// ex: builder.Having("age > 20")
func (builder *QueryBuilder) Having(expression string, args ...interface{}) *QueryBuilder {
	builder.having = "HAVING " + expression
	builder.havingArgs = args

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
func (builder *QueryBuilder) Offset(offset string) *QueryBuilder {
	builder.offset = offset

	return builder
}

// Limit apply limit in query
// ex: builder.Limit(10)
func (builder *QueryBuilder) Limit(limit string) *QueryBuilder {
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
		builder.having + " " +
		builder.order + " " +
		_platform().getPagingDeclaration(builder.limit, builder.offset)

	// Using this cause a really bad performance
	if args := append(builder.filterArgs, builder.havingArgs...); len(args) > 0 {
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
