package dbs

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	schema     string
	from       string
	selections []string
	// This is all where clauses are put.
	filter *Filter

	query string
	built bool
	errs  []error
}

// Filter is used as a placeholder for all where clauses and their conditions.
type Filter struct {
	action  string // AND | OR
	clauses []*Clause
}

// Clause is a simple expression with args.
type Clause struct {
	expression string
	args       []interface{}
}

func NewQueryBuilder() *QueryBuilder {
	builder := new(QueryBuilder)
	builder.filter = new(Filter)

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
func (builder *QueryBuilder) Select(context ...string) *QueryBuilder {
	builder.selections = make([]string, 0)
	builder.selections = append(builder.selections, context...)

	return builder
}

// From specify table that query will be executed on
// new(QueryBuilder).Select(..).From("user")
// new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(context string) *QueryBuilder {
	builder.from = strings.Trim(context, " ")

	return builder
}

func (filter *Filter) add(clauses ...*Clause) {
	filter.clauses = append(filter.clauses, clauses...)
}

func (filter *Filter) setAction(action string) error {
	filter.action = action

	return nil
}

// Where
// eg:
// builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(clause string, args ...interface{}) *QueryBuilder {
	builder.filter.add(&Clause{clause, args})

	return builder
}

// AndWhere
// builder.
// 	Where("name = '%s'", "Luan Phan").
//	AndWhere("age > %d", 10)
func (builder *QueryBuilder) AndWhere(clause string, args ...interface{}) *QueryBuilder {
	if err := builder.filter.setAction("AND"); err != nil {
		builder.logError(err)
	}

	return builder.Where(clause, args...)
}

// OrWhere
// Join where statement
func (builder *QueryBuilder) OrWhere(clause string, args ...interface{}) *QueryBuilder {
	if err := builder.filter.setAction("OR"); err != nil {
		builder.logError(err)
	}

	return builder.Where(clause, args...)
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
	declarations = append(declarations, builder.selectClause())
	declarations = append(declarations, builder.fromTableClause())
	declarations = append(declarations, builder.whereClauses())

	return concatStrings(declarations, "\n")
}

func (builder *QueryBuilder) selectClause() string {
	if len(builder.selections) == 0 {
		return "SELECT *"
	}

	return "SELECT " + concatStrings(builder.selections, ", ")
}

func (builder *QueryBuilder) fromTableClause() string {
	if builder.from != "" {
		return "FROM " + _platform().getSchemaAccessName(builder.schema, builder.from)
	}

	return ""
}

// all where clause
func (builder *QueryBuilder) whereClauses() string {
	if clauses := builder.filter.clauses; len(clauses) > 0 && builder.from != "" {
		where := make([]string, 0)

		for _, clause := range clauses {
			where = append(where, fmt.Sprintf(clause.expression, clause.args...))
		}

		return "WHERE " + concatStrings(where, "\n"+builder.filter.action+" ")
	}

	return ""
}
