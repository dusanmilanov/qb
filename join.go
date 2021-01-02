package qb

import "fmt"

type JoinType string

const (
	jtLeftOuter JoinType = "LEFT OUTER"
	jtInner JoinType = "INNER"
)

type JoinOption func(j *joinSpec)

func LeftOuter(j *joinSpec) {
	j.joinType = jtLeftOuter
}

/*func As(alias string) func(q *Query) {
	return func(q *Query) {
		q.SetAlias(alias)
	}
}*/

type joinSpec struct {
	Alias
	ThisKey    Column
	ThisQuery  *Query
	OtherKey   Column
	OtherQuery *Query
	joinType JoinType
}

func (j joinSpec) Visit(visitor joinVisitorFunc) {
	visitor(j)
}

func (j *joinSpec) BuildSelectSql(params *buildParameters) string {
	joinType := jtInner
	if j.joinType != "" {
		joinType = j.joinType
	}

	var target string
	target = j.OtherQuery.From.FromSql(params)
	join := fmt.Sprintf("%s JOIN %s ON %s.%s = %s.%s",
		joinType,
		target,
		j.OtherQuery.GetAlias(), j.OtherKey.Name,
		j.ThisQuery.From.GetAlias(), j.ThisKey.Name)
	return join
}

