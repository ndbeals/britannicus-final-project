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

// //SignupForm ...
// type SignupForm struct {
// 	Name     string `form:"name" json:"name" binding:"required,max=100"`
// 	Email    string `form:"email" json:"email" binding:"required,email"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }
