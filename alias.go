package qb

type Aliased interface {
	GetAlias() string
	SetAlias(string)
}

type Alias struct {
	alias string
}

func (a Alias) GetAlias() string {
	return a.alias
}

func (a *Alias) SetAlias(alias string) {
	a.alias = alias
}
