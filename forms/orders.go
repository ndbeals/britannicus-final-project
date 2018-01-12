package forms

//CreateOrderForm ...
type CreateOrderForm struct {
	CustomerID int         `form:"customer_id" json:"customer_id" binding:"required"`
	ItemList   map[int]int `form:"item_list" json:"item_list" binding:"required"`
}

//UpdateOrderForm ...
type UpdateOrderForm struct {
	OrderID    int         `form:"order_id" json:"order_id" binding:"required"`
	CustomerID int         `form:"customer_id" json:"customer_id" binding:"required"`
	ItemList   map[int]int `form:"item_list" json:"item_list" binding:"required"`
}
