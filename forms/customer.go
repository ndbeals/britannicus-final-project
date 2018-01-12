package forms

//Updatecustomer ...
type UpdateCustomerForm struct {
	// customerID         int    `json:"customer_id" `
	ID          int    `json:"customer_id" form:"customer_id" `
	FirstName   string `json:"customer_firstname" form:"customer_firstname" `
	LastName    string `json:"customer_lastname" form:"customer_lastname" `
	Email       string `json:"customer_email" form:"customer_email" `
	PhoneNumber string `json:"customer_phonenumber" form:"customer_phonenumber" `
	Address     string `json:"customer_address" form:"customer_address" `
	City        string `json:"customer_city" form:"customer_city"`
	State       string `json:"customer_state" form:"customer_state"`
	Country     string `json:"customer_country" form:"customer_country"`
}
