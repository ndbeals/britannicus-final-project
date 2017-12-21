package models

import (
	"database/sql"
	"math"

	"github.com/ndbeals/brittanicus-final-project/db"
)

//InventoryCondition ...
type InventoryCondition struct {
	ID                 int    `json:"inventory_id"`
	ProductID          int    `json:"product_id"`
	InventoryCondition int    `json:"inventory_condition"`
	Amount             int    `json:"amount"`
	Note               string `json:"note"`
}

//Inventory ...
type Inventory struct {
	ID                 int    `json:"inventory_id"`
	ProductID          int    `json:"product_id"`
	InventoryCondition int    `json:"-"`
	ConditionString    string `json:"inventory_condition"`
	Amount             int    `json:"amount"`
	Note               string `json:"note"`
}

//InventoryModel ...
type InventoryModel struct{}

var (
	inventoryModel  *InventoryModel
	conditionLookup map[int]string
)

//InitializeInventoryModel ...
func InitializeInventoryModel() {
	GetInventoryModel()
	conditionLookup = make(map[int]string)

	rows, err := db.DB.Query("SELECT condition_id, condition FROM tblInventoryConditions")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var conditionID int
		var condition sql.NullString

		err = rows.Scan(&conditionID, &condition)
		if err != nil {
			panic(err)
		}

		conditionLookup[conditionID] = condition.String
	}

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
	row := db.DB.QueryRow("SELECT inventory_id, product_id, inventory_condition, amount, notes FROM tblInventory WHERE inventory_id=$1", InventoryID)

	var inventoryID, productID, inventoryCondition, amount int
	var notes sql.NullString

	err = row.Scan(&inventoryID, &productID, &inventoryCondition, &amount, &notes)
	if err != nil {
		panic(err)
	}

	inventory = Inventory{inventoryID, productID, inventoryCondition, conditionLookup[inventoryCondition], amount, notes.String}

	return inventory, err
}

//GetList ...
func (m InventoryModel) GetList(Page int, Amount int) (inventoryList []Inventory, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	rows, err := db.DB.Query("SELECT inventory_id, product_id, inventory_condition, amount, notes FROM tblInventory OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var inventoryID, productID, inventoryCondition, amount int
		var notes sql.NullString

		err = rows.Scan(&inventoryID, &productID, &inventoryCondition, &amount, &notes)
		if err != nil {
			panic(err)
		}

		inventoryList = append(inventoryList, Inventory{inventoryID, productID, inventoryCondition, conditionLookup[inventoryCondition], amount, notes.String})
	}

	return inventoryList, err
}
