package dbs

import (
	"fmt"
	"strings"
)

// QueryBuilder create query builder
type QueryBuilder struct {
	schema  string
	pick    *clause
	from    *clause
	filters []*clause

	query string
	built bool
	errs  []error
}

// clause is a simple expression with args.
type clause struct {
	prefix     string
	expression string
	args       []interface{}
	postfix    string
}

// NewQueryBuilder make new(QueryBuilder) along with some default config
func NewQueryBuilder() *QueryBuilder {
	builder := new(QueryBuilder)
	builder.pick = &clause{
		prefix:     "SELECT",
		expression: "*",
	}

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
//   		Select("*", "something as something_else").
//  		Select("something as something_else")
// Apply only second Select() called
func (builder *QueryBuilder) Select(selections ...string) *QueryBuilder {
	builder.pick = &clause{
		prefix:     "SELECT",
		expression: concatStrings(selections, ", "),
	}

	return builder
}

// From specify table that query will be executed on
// ex: 	new(QueryBuilder).Select(..).From("user")
// 		new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.from = &clause{
		prefix: "FROM",
		// @TODO: temporary hack for DB with schema like postgresql, mssql
		expression: _platform().getSchemaAccessName(builder.schema, expression),
	}

	return builder
}

// Where apply filter to query
// ex: builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	filter := &clause{
		prefix:     "WHERE",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, filter)
	return builder
}

// AndWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	filter := &clause{
		prefix:     "AND",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, filter)
	return builder
}

// OrWhere chaining filter on query
// ex: builder.Where("name = '%s'", "Luan Phan").OrWhere("age > %d", 10)
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	filter := &clause{
		prefix:     "OR",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, filter)
	return builder
}

// OrderBy apply order in query
// ex: builder.OrderBy("id ASC", "name")
func (builder *QueryBuilder) OrderBy(expression ...string) *QueryBuilder {
	filter := &clause{
		prefix:     "ORDER BY",
		expression: concatStrings(expression, ", "),
		args:       nil,
		postfix:    "",
	}

	builder.filters = append(builder.filters, filter)
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
	declarations := make([]string, 0)
	declarations = append(declarations, builder.pick.build())
	if from := builder.from; from != nil {
		declarations = append(declarations, builder.from.build())
		for _, filter := range builder.filters {
			declarations = append(declarations, filter.build())
		}
	}

	return concatStrings(declarations, "\n")
}

func (clause *clause) build() string {
	partials := make([]string, 0)
	partials = append(partials, clause.prefix, clause.expression, clause.postfix)
	expression := concatStrings(partials, " ")

	if len(clause.args) > 0 {
		return fmt.Sprintf(expression, clause.args...)
	}

	return expression
}
