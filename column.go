package qb

type ColumnType int

const (
	ctInt ColumnType = iota
	ctText
)

type ColumnDef struct {
	Name string
	Type ColumnType
}

type ColumnDefMap map[string]ColumnDef

type Column struct {
	ColumnDef
	Aliased
	Target Targetable
}

func (c Column) GetText(params *[]interface{}) string {
	return nameWithParentAlias(c.Name, c.Target.GetAlias())
}

func (c Column) GetSelectText() string {
	return nameWithParentAlias(c.Name, c.Target.GetAlias())
}