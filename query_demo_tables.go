package qb

func buildColumnDefMap(columnDefList ...ColumnDef) ColumnDefMap {
	result := make(ColumnDefMap)
	for _, cd := range columnDefList {
		result[cd.Name] = cd
	}
	return result
}

type ProductTableDef TableDef
var Product = buildProductTableDef()

func (t ProductTableDef) Query() *ProductQuery {
	result := &ProductQuery{
		Query: Query{
			From: &tableSpec{Table: TableDef(Product)},
		},
		Id: Column{ColumnDef: t.Columns["Id"]},
		Name: Column{ColumnDef:t.Columns["Name"]},
	}
	result.Id.Target = result.From
	result.Name.Target = result.From
	return result
}

type ProductQuery struct {
	Query
	Id Column
	Name Column
}

func (q* ProductQuery) Join_OrderItem(joinOptions ...JoinOption) *OrderItemQuery {
	result := OrderItem.Query()
	q.joinTable(&result.Query, q.Id, result.ProductId, joinOptions...)
	return result
}

func buildProductTableDef() ProductTableDef {
	return ProductTableDef{
		Name: "Product",
		Columns: buildColumnDefMap (
			ColumnDef{Name:"Id", Type: ctInt},
			ColumnDef{Name: "Name", Type: ctText},
		),
	}
}

type OrderItemTableDef TableDef
var OrderItem = buildOrderItemTableDef()

func (t OrderItemTableDef) Query() *OrderItemQuery {
	result := &OrderItemQuery{
		Query: Query{
			From: &tableSpec{Table: TableDef(OrderItem)},
		},
		Id: Column{ColumnDef: t.Columns["Id"]},
		ProductId: Column{ColumnDef:t.Columns["ProductId"]},
		OrderId: Column{ColumnDef:t.Columns["OrderId"]},
	}
	result.Id.Target = result.From
	result.ProductId.Target = result.From
	result.OrderId.Target = result.From
	return result
}

type OrderItemQuery struct {
	Query
	Id Column
	OrderId Column
	ProductId Column
}

func buildOrderItemTableDef() OrderItemTableDef {
	return OrderItemTableDef{
		Name: "OrderItem",
		Columns: buildColumnDefMap (
			ColumnDef{Name:"Id", Type: ctInt},
			ColumnDef{Name: "ProductId", Type: ctInt},
			ColumnDef{Name: "OrderId", Type: ctInt},
		),
	}
}

func (q* OrderItemQuery) Join_Product(joinOptions ...JoinOption) *ProductQuery {
	result := Product.Query()
	q.joinTable(&result.Query, q.ProductId, result.Id, joinOptions...)
	return result
}

func (q* OrderItemQuery) Join_Order(joinOptions ...JoinOption) *OrderQuery {
	result := Order.Query()
	q.joinTable(&result.Query, q.OrderId, result.Id, joinOptions...)
	return result
}

type OrderTableDef TableDef
var Order = buildOrderTableDef()

func (t OrderTableDef) Query() *OrderQuery {
	result := &OrderQuery{
		Query: Query{
			From: &tableSpec{Table: TableDef(Order)},
		},
		Id: Column{ColumnDef: t.Columns["Id"]},
		Total: Column{ColumnDef:t.Columns["Total"]},
	}
	result.Id.Target = result.From
	result.Total.Target = result.From
	return result
}

type OrderQuery struct {
	Query
	Id Column
	Total Column
}

func buildOrderTableDef() OrderTableDef {
	return OrderTableDef{
		Name: "Order",
		Columns: buildColumnDefMap (
			ColumnDef{Name:"Id", Type: ctInt},
			ColumnDef{Name: "Total", Type: ctInt},
		),
	}
}

func (q* OrderQuery) Join_OrderItem(joinOptions ...JoinOption) *OrderItemQuery {
	result := OrderItem.Query()
	q.joinTable(&result.Query, q.Id, result.OrderId, joinOptions...)
	return result
}

