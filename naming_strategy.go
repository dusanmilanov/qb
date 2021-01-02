package qb

import (
	"strings"
	"unicode"
)

type NamingStrategy interface {
	Split(string)[]string
}

type CamelCaseNamingStrategy struct {}

func (cc CamelCaseNamingStrategy) Split(name string) []string {
	parts := []string{}
	if name =="" {
		return parts
	}
	currentPart := name[0:1]
	for _, c := range name[1:] {
		if unicode.IsUpper(c) {
			parts = append(parts, currentPart)
			currentPart = string(c)
		} else {
			currentPart = currentPart + string(c)
		}
	}
	parts = append(parts, currentPart)
	return parts
}

type SnakeCaseNamingStrategy struct {}

func (sc SnakeCaseNamingStrategy) Split(name string) []string {
	return strings.Split(name, "_")
}

var TableNamingStrategy = CamelCaseNamingStrategy{}