package forms

//UpdateProduct ...
type UpdateProductForm struct {
	// ProductID         int    `json:"product_id" `
	ID          int    `json:"product_id" form:"product_id" `
	ISBN        string `json:"product_isbn" form:"product_isbn" `
	ProductName string `json:"product_name" form:"product_name" `
	Author      string `json:"product_author" form:"product_author" `
	Genre       string `json:"product_genre" form:"product_genre" `
	Publisher   string `json:"product_publisher" form:"product_publisher" `
	// ProductType       int    `json:"-"`
	// ProductTypeString string `json:"product_type"`
	Description string `json:"product_description" form:"product_description" `
}

// //SignupForm ...
// type SignupForm struct {
// 	Name     string `form:"name" json:"name" binding:"required,max=100"`
// 	Email    string `form:"email" json:"email" binding:"required,email"`
// 	Password string `form:"password" json:"password" `
// }
