package dbs

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	schema  string
	pick    *Clause
	from    *Clause
	filters []*Clause

	query string
	built bool
	errs  []error
}

// Clause is a simple expression with args.
type Clause struct {
	prefix     string
	expression string
	args       []interface{}
	postfix    string
}

func NewQueryBuilder() *QueryBuilder {
	builder := new(QueryBuilder)
	builder.pick = &Clause{
		prefix:     "SELECT",
		expression: "*",
	}

	return builder
}

// OnSchema specify schema that query will be executed on
// new(QueryBuilder).OnSchema("some_schema")
func (builder *QueryBuilder) OnSchema(schema string) *QueryBuilder {
	builder.schema = strings.Trim(schema, " ")

	return builder
}

// Select
// specify one or more columns to be query
// eg:
//  	new(QueryBuilder).
//   		Select("*", "something as something_else").
//  		Select("something as something_else")
// Apply only second Select() called
func (builder *QueryBuilder) Select(selections ...string) *QueryBuilder {
	builder.pick = &Clause{
		prefix:     "SELECT",
		expression: concatStrings(selections, ", "),
	}

	return builder
}

// From specify table that query will be executed on
// new(QueryBuilder).Select(..).From("user")
// new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.from = &Clause{
		prefix:     "FROM",
		expression: expression,
	}

	return builder
}

// Where
// eg:
// builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	clause := &Clause{
		prefix:     "WHERE",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, clause)
	return builder
}

// AndWhere
// builder.
// 	Where("name = '%s'", "Luan Phan").
//	AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(expression string, args ...interface{}) *QueryBuilder {
	clause := &Clause{
		prefix:     "AND",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, clause)
	return builder
}

// OrWhere
// Join where statement
func (builder *QueryBuilder) OrWhere(expression string, args ...interface{}) *QueryBuilder {
	clause := &Clause{
		prefix:     "OR",
		expression: expression,
		args:       args,
		postfix:    "",
	}

	builder.filters = append(builder.filters, clause)
	return builder
}

// BuildQuery
// use GetQuery to get SQL declaration
func (builder *QueryBuilder) BuildQuery() *QueryBuilder {
	builder.query = builder.buildQuery()
	builder.built = true
	return builder
}

// GetQuery returns a built query
func (builder *QueryBuilder) GetQuery() string {
	if ! builder.built {
		builder.BuildQuery()
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

func (clause *Clause) build() string {
	partials := make([]string, 0)
	partials = append(partials, clause.prefix, clause.expression, clause.postfix)
	expression := concatStrings(partials, " ")

	if len(clause.args) > 0 {
		return fmt.Sprintf(expression, clause.args...)
	}

	return expression
}
