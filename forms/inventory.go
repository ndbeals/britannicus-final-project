package forms

//UpdateInventory ...
type UpdateInventoryForm struct {
	// InventoryID         int    `json:"inventory_id" `
	ID                 int     `json:"inventory_id" form:"inventory_id" `
	ProductID          int     `json:"product_id" form:"product_id"`
	InventoryCondition int     `json:"inventory_condition" form:"inventory_condition" `
	Amount             int     `json:"inventory_amount" form:"inventory_amount" `
	Price              float64 `json:"inventory_price" form:"inventory_price" `
	Note               string  `json:"inventory_note" form:"inventory_note" `
}
