package qb

type Expression interface {
	Aliased
	GetText(params *[]interface{}) string
}

type ExpressionList []Expression

func (el ExpressionList) GetTexts(params *[]interface{}) []string {
	result := []string{}
	for _, e := range el {
		result = append(result, e.GetText(params))
	}
	return result
}
