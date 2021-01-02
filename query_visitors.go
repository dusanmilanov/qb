package qb

type visitorFunc func(target Targetable)

func (q *Query) visit(visitor visitorFunc) {
	q.From.Visit(visitor)
	for _, join := range q.joinList {
		join.OtherQuery.visit(visitor)
	}
}

type joinVisitorFunc func (joinSpec)

func (q *Query) visitJoins(visitor joinVisitorFunc) {
	for _, join := range q.joinList {
		join.Visit(visitor)
		join.OtherQuery.visitJoins(visitor)
	}
}


