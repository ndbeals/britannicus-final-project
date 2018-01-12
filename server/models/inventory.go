package models

import (
	"database/sql"
	"math"

	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/forms"
)

//Inventory ...
type Inventory struct {
	ID                 int     `json:"inventory_id"`
	Product            Product `json:"product"`
	InventoryCondition int     `json:"-"`
	ConditionString    string  `json:"inventory_condition"`
	Amount             int     `json:"amount"`
	Price              float64 `json:"item_price"`
	Note               string  `json:"note"`
}

//InventoryModel ...
type InventoryModel struct{}

var (
	inventoryModel  *InventoryModel
	conditionLookup map[int]string
)

//InitializeInventoryModel ...
func InitializeInventoryModel() *InventoryModel {
	GetInventoryModel()
	conditionLookup = make(map[int]string)

	rows, err := db.DB.Query("SELECT condition_id, condition FROM tblInventoryConditions")
	if err != nil {
		// panic(err)
	}

	for rows.Next() {
		var conditionID int
		var condition sql.NullString

		err = rows.Scan(&conditionID, &condition)
		if err != nil {
			// panic(err)
		}
		conditionLookup[conditionID] = condition.String
	}

	return inventoryModel
}

//GetInventoryModel ...
func GetInventoryModel() (model InventoryModel) {
	if inventoryModel != nil {
		return *inventoryModel
	}

	inventoryModel = new(InventoryModel)
	model = *inventoryModel

	return model
}

//GetOne ...
func (m InventoryModel) GetOne(InventoryID int) (inventory Inventory, err error) {
	row := db.DB.QueryRow("SELECT inventory_id, product_id, inventory_condition, amount, price, notes FROM tblInventory WHERE inventory_id=$1", InventoryID)

	var inventoryID, productID, inventoryCondition, amount int
	var price float64
	var notes sql.NullString

	err = row.Scan(&inventoryID, &productID, &inventoryCondition, &amount, &price, &notes)
	if err != nil {
		// panic(err)
	}

	product, err := GetProductModel().GetOne(productID)
	if err != nil {
		// panic(err)
	}

	inventory = Inventory{inventoryID, product, inventoryCondition, conditionLookup[inventoryCondition], amount, price, notes.String}

	return inventory, err
}

//GetList ...
func (m InventoryModel) GetList(Page int, Amount int) (inventoryList []Inventory, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	rows, err := db.DB.Query("SELECT inventory_id, product_id, inventory_condition, amount, price, notes FROM tblInventory ORDER BY  inventory_id OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		return inventoryList, err
	}

	for rows.Next() {
		var inventoryID, productID, inventoryCondition, amount int
		var price float64
		var notes sql.NullString

		err = rows.Scan(&inventoryID, &productID, &inventoryCondition, &amount, &price, &notes)
		if err != nil {
			// // panic(err)
		}

		product, err := GetProductModel().GetOne(productID)
		if err != nil {
			// // panic(err)
		}

		inventoryList = append(inventoryList, Inventory{inventoryID, product, inventoryCondition, conditionLookup[inventoryCondition], amount, price, notes.String})
	}

	return inventoryList, err
}

//Inventory Delete ...
func (this *Inventory) Delete() (bool, error) {
	_, err := db.DB.Query("DELETE FROM tblInventory WHERE inventory_id=$1", this.ID)

	if err != nil {
		return false, err
	}

	return true, err
}

// Update ...
func (this *Inventory) Update(newdata forms.UpdateInventoryForm) (bool, error) {
	stmt, err := db.DB.Prepare("update tblinventory set inventory_condition=$2, amount=$3, price=$4, notes=$5 where inventory_id=$1")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(this.ID, newdata.InventoryCondition, newdata.Amount, newdata.Price, newdata.Note)

	if err != nil {
		return false, err
	}

	inventoryModel.GetOne(this.ID)

	return true, err
}

// Create ...
func (this *Inventory) Create() (int, error) {
	stmt, err := db.DB.Prepare("insert into tblinventory(product_id, inventory_condition, amount, price, notes) values( $1, $2, $3, $4, $5 ) RETURNING inventory_id")

	if err != nil {
		return 0, err
	}

	results := stmt.QueryRow(this.Product.ID, this.InventoryCondition, this.Amount, this.Price, this.Note)

	var newid int
	err = results.Scan(&newid)

	if err != nil {
		return 0, err
	}

	inventoryModel.GetOne(int(newid))

	return int(newid), err
}
