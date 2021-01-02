package qb

import (
	"fmt"
	"strings"
)

type buildParameters struct {
	Parameters []interface{}
}

type Buildable interface {
	FromSql(params *buildParameters) string
	BuildSql(params *buildParameters) string
}

type Named interface {
	GetName() string
}

type Targetable interface {
	Named
	Aliased
	Buildable
	Visit(visitor visitorFunc)
}

func (q *Query) BuildSelectSql(params *buildParameters) string {
	buildAliases(q)

	selectList := q.buildSelectExpression()
	targetList := q.buildJoinList()

	const sqlSelect = "SELECT %s FROM %s"
	sql := fmt.Sprintf(sqlSelect, selectList, targetList)
	return sql
}

func (q *Query) buildSelectExpression() string {
	selectList := q.SelectList.GetTexts(nil)
	return strings.Join(selectList, ", ")
}

func (q *Query) buildJoinList() string {
	result := q.From.FromSql(nil)

	joinVisitor := func (join joinSpec) {
		result = result + " " + join.BuildSelectSql(nil)
	}

	q.visitJoins(joinVisitor)
	return result
}
