package dbs

import "strings"

type QueryBuilder struct {
	picks  []string
	schema string
	from   string
}

// OnSchema specify schema that query will be executed on
// new(QueryBuilder).OnSchema("some_schema")
func (builder *QueryBuilder) OnSchema(schema string) *QueryBuilder {
	builder.schema = strings.Trim(schema, " ")

	return builder
}

// Select specify one or more columns to be query
// new(QueryBuilder).Select("*", "something as something_else")
func (builder *QueryBuilder) Select(context ...string) *QueryBuilder {
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
