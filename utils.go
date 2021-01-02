package qb

import "fmt"

func nameWithParentAlias(name string, parentAlias string) string {
	if parentAlias != "" {
		return parentAlias + "." + name
	}
	return name
}

func nameWithAlias(name string, alias string) string {
	if alias != "" {
		return fmt.Sprintf("%s %s", name, alias)
	}
	return name
}

func parenthesis(value string) string {
	return fmt.Sprintf("(%s)", value)
}
