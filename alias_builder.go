package qb

import (
	"strconv"
	"strings"
)

type aliasToNextIndex map[string]int

func buildAliases(q *Query) {
	aliasMap := make(aliasToNextIndex)
	collectExistingAliases(aliasMap, q)
	createAliases(aliasMap, q)
}

func collectExistingAliases(aliasMap map[string]int, q *Query) {
	visitor := func(target Targetable) {
		if target.GetAlias() != "" {
			aliasMap[target.GetAlias()] = 1
		}
	}
	q.visit(visitor)
}

// Creates an alias for a given table name (with underscores)
func aliasForTableName(tableName string) string {
	parts := TableNamingStrategy.Split(tableName)
	result := ""
	for _, part := range parts {
		result = result + strings.ToLower(part[0:1])
	}
	return result
}

func createAliases(aliasMap map[string]int, q *Query) {
	visitor := func(target Targetable) {
		alias := target.GetAlias()
		if alias == "" {
			alias = aliasForTableName(target.GetName())
			index, ok := aliasMap[alias]
			if !ok {
				index = 1
			}
			aliasMap[alias] = index + 1

			target.SetAlias(alias + strconv.Itoa(index))
		}
	}
	q.visit(visitor)
}

