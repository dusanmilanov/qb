package qb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SelectColumnsOnTable(t *testing.T) {
	productQuery := Product.Query()
	productQuery.Select(productQuery.Id, productQuery.Name)

	require.EqualValues(t, "SELECT Id, Name FROM Product", productQuery.BuildSelectSql(nil))

	productQuery.From.SetAlias("p")
	require.EqualValues(t, "SELECT p.Id, p.Name FROM Product p", productQuery.BuildSelectSql(nil))
}

func Test_JoinTable(t *testing.T) {
	productQuery := Product.Query()
	productQuery.Select(productQuery.Id, productQuery.Name)
	orderItemQuery := productQuery.Join_OrderItem()
	orderItemQuery.Select(orderItemQuery.Id)

	require.EqualValues(t,
		"SELECT p1.Id, p1.Name, oi1.Id FROM Product p1 INNER JOIN OrderItem oi1 ON oi1.ProductId = p1.Id",
		productQuery.BuildSelectSql(nil))

	orderQuery := orderItemQuery.Join_Order()
	orderQuery.Select(orderQuery.Total)
	require.EqualValues(t,
		"SELECT p1.Id, p1.Name, oi1.Id, o1.Total FROM Product p1 INNER JOIN OrderItem oi1 ON oi1.ProductId = p1.Id INNER JOIN Order o1 ON o1.Id = oi1.OrderId",
		productQuery.BuildSelectSql(nil))
}

func Test_JoinQuery(t *testing.T) {
	productQuery := Product.Query()
	productQuery.Select(productQuery.Id, productQuery.Name)

	orderItemQuery := OrderItem.Query()
	orderItemQuery.Select(orderItemQuery.Id)

	productQuery.JoinQuery(&orderItemQuery.Query, productQuery.Id, orderItemQuery.ProductId)
	require.EqualValues(t,
		"SELECT p1.Id, p1.Name, oi1.Id FROM Product p1 INNER JOIN (SELECT oi1.Id FROM OrderItem oi1) q1 ON q1.ProductId = p1.Id",
		productQuery.BuildSelectSql(nil))
}