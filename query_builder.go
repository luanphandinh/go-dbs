package dbs

import (
	"errors"
	"fmt"
	"strings"
)

type QueryBuilder struct {
	schema     string
	from       string
	selections []string
	// This is all where clauses are put.
	filter     *Filter

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
// Select, Update condition
func (builder *QueryBuilder) Where(clause string, args ...interface{}) *QueryBuilder {
	builder.filter.add(&Clause{clause, args})

	return builder
}

// AndWhere
// Join where statement
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
	declarations = append(declarations, "SELECT")
	if len(builder.selections) == 0 {
		declarations = append(declarations, "*")
	} else {
		declarations = append(declarations, concatStrings(builder.selections, ", "))
	}

	if builder.from == "" {
		builder.logError(errors.New("no table provided, please use From()"))
	} else {
		declarations = append(declarations, "\nFROM")
		declarations = append(declarations, _platform().getSchemaAccessName(builder.schema, builder.from))
	}

	if clauses := builder.filter.clauses; len(clauses) > 0 {
		declarations = append(declarations, "\nWHERE")
		whereClauses := make([]string, 0)

		for _, clause := range clauses {
			whereClauses = append(whereClauses, fmt.Sprintf(clause.expression, clause.args...))
		}

		declarations = append(declarations, concatStrings(whereClauses, "\n" + builder.filter.action + " "))
	}

	return concatStrings(declarations, " ")
}
