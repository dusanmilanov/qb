package qb

type Query struct {
	Alias
	From Targetable
	SelectList ExpressionList

	parentQuery *Query

	joinList []joinSpec
}

func (q *Query) Select(expressions ...Expression) {
	if q.parentQuery != nil {
		q.parentQuery.Select(expressions...)
	}
	q.SelectList = append(q.SelectList, expressions...)
}

func (q *Query) SelectAliased(alias string, expression Expression) {
	if q.parentQuery != nil {
		q.parentQuery.SelectAliased(alias, expression)
	}
	expression.SetAlias(alias)
	q.SelectList = append(q.SelectList, expression)
}

func (q *Query) joinTable(otherQuery *Query, thisKey Column, otherKey Column, options ...JoinOption) *Query{
	q.joinList = append(q.joinList, joinSpec{
		Alias:      Alias{},
		ThisKey:    thisKey,
		ThisQuery:  q,
		OtherKey:   otherKey,
		OtherQuery: otherQuery,
	})
	otherQuery.parentQuery = q
	for _, opt := range options {
		opt(otherQuery)
	}
	return otherQuery
}

func (q *Query) JoinQuery(otherQuery *Query, thisKey Column, otherKey Column, options ...JoinOption) *Query{
	q.joinList = append(q.joinList, joinSpec{
		Alias:      Alias{},
		ThisKey:    thisKey,
		ThisQuery:  q,
		OtherKey:   otherKey,
		OtherQuery: querySpec{otherQuery},
	})
	otherQuery.parentQuery = q
	for _, opt := range options {
		opt(otherQuery)
	}
	return otherQuery
}

type querySpec struct {
	Query *Query
}

func (q querySpec) BuildSql(params *buildParameters) string {
	return q.Query.BuildSelectSql(params)
}

func (q querySpec) FromSql(params *buildParameters) string {
	return parenthesis(q.Query.BuildSelectSql(params))
}

func (q querySpec) GetName() string {
	return "query"
}

func (q querySpec) Visit(visitor visitorFunc) {
	q.Query.visit(visitor)
}