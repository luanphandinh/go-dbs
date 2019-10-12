package dbs

import (
	"strings"
)

type QueryBuilder struct {
	picks  []string
	schema string
	from   string
	query  string
	built  bool
	errs   []string
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
	builder.picks = make([]string, 0)
	builder.picks = append(builder.picks, context...)

	return builder
}

// From specify table that query will be executed on
// new(QueryBuilder).Select(..).From("user")
// new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(context string) *QueryBuilder {
	builder.from = strings.Trim(context, " ")

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

func (builder *QueryBuilder) logError(err string) *QueryBuilder {
	builder.errs = append(builder.errs, err)

	return builder
}

// build SQL Query declaration
func (builder *QueryBuilder) buildQuery() string {
	declarations := make([]string, 0)
	declarations = append(declarations, "SELECT")
	if len(builder.picks) == 0 {
		declarations = append(declarations, "*")
	} else {
		declarations = append(declarations, concatStrings(builder.picks, ", "))
	}

	if builder.from == "" {
		builder.logError("no table provided, please use From()")
	}
	declarations = append(declarations, "\nFROM")
	declarations = append(declarations, _platform().getSchemaAccessName(builder.schema, builder.from))

	return concatStrings(declarations, " ")
}
