package qb

type TableDef struct {
	Name string
	Columns ColumnDefMap
}

type tableSpec struct {
	Alias
	Table TableDef
}

func (t tableSpec) BuildSql(params *buildParameters) string {
	return nameWithAlias(t.Table.Name, t.GetAlias())
}

func (t tableSpec) FromSql(params *buildParameters) string {
	return nameWithAlias(t.Table.Name, t.GetAlias())
}

func (t *tableSpec) Visit(visitor visitorFunc) {
	visitor(t)
}

func (t tableSpec) GetName() string {
	return t.Table.Name
}