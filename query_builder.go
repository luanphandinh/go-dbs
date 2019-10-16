package dbs

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	_SELECT   = "SELECT"
	_FROM     = "FROM"
	_WHERE    = "WHERE"
	_AND      = "AND"
	_OR       = "OR"
	_ORDER_BY = "ORDER BY"
)

// QueryBuilder create query builder
type QueryBuilder struct {
	schema  string
	pick    *clause
	from    *clause
	filters []*clause
	offset  int
	limit   int

	query string
	built bool
	errs  []error
}

// clause is a simple expression with args.
// @TODO: review performance of parsing each clause with args, could be faster if we merge all clauses then parse all args??
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
		prefix:     _SELECT,
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
		prefix:     _SELECT,
		expression: concatStrings(selections, ", "),
	}

	return builder
}

// From specify table that query will be executed on
// ex: 	new(QueryBuilder).Select(..).From("user")
// 		new(QueryBuilder).Select(..).From("user as u")
func (builder *QueryBuilder) From(expression string) *QueryBuilder {
	builder.from = &clause{
		prefix: _FROM,
		// @TODO: temporary hack for DB with schema like postgresql, mssql
		expression: _platform().getSchemaAccessName(builder.schema, expression),
	}

	return builder
}

// Where apply filter to query
// ex: builder.Where("name = '%s'", "Luan Phan")
func (builder *QueryBuilder) Where(expression string, args ...interface{}) *QueryBuilder {
	filter := &clause{
		prefix:     _WHERE,
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
		prefix:     _AND,
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
		prefix:     _OR,
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
		prefix:     _ORDER_BY,
		expression: concatStrings(expression, ", "),
		args:       nil,
		postfix:    "",
	}

	builder.filters = append(builder.filters, filter)
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
	declarations := make([]string, 0)
	declarations = append(declarations, builder.pick.build())
	if from := builder.from; from != nil {
		declarations = append(declarations, builder.from.build())
		for _, filter := range builder.filters {
			declarations = append(declarations, filter.build())
		}
	}

	declarations = append(declarations, _platform().getPagingDeclaration(builder.limit, builder.offset))

	return concatStrings(declarations, "\n")
}

func (clause *clause) build() string {
	expression := concatStrings([]string{clause.prefix, clause.expression, clause.postfix}, " ")

	if args := clause.args; len(args) > 0 {
		parsedArgs := make([]interface{}, 0)
		for _, arg := range args {
			parsedArgs = append(parsedArgs, clause.parseArg(arg))
		}

		return fmt.Sprintf(expression, parsedArgs...)
	}

	return expression
}

func (clause *clause) parseArg(arg interface{}) interface{} {
	rt := reflect.TypeOf(arg)
	switch rt.Kind() {
	case reflect.Slice:
		return getContentOutOfArraySyntax(arg)
	case reflect.Array:
		return getContentOutOfArraySyntax(arg)
	}

	return arg
}
