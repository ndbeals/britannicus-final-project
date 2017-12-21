package models

import (
	"database/sql"
	"fmt"
	"log"
	"math"

	"github.com/ndbeals/brittanicus-final-project/db"
)

var (
	loadedOrders map[int]Order
)

type ItemList struct {
	ID                 int     `json:"inventory_id"`
	Product            Product `json:"product"`
	InventoryCondition int     `json:"-"`
	ConditionString    string  `json:"inventory_condition"`
	OrderAmount        int     `json:"order_amount"`
	Note               string  `json:"note"`
}

//Order ...
type Order struct {
	ID        int         `json:"order_id"`
	Customer  Customer    `json:"customer"`
	OrderTime int         `json:"order_time"`
	ItemList  *[]ItemList `json:"item_list"`
	// UpdatedAt int  `db:"updated_at" json:"updated_at"`
	// CreatedAt int  `db:"created_at" json:"created_at"`
}

func (o *Order) AddToInventory(inventoryID int, orderAmount int) {
	inventory, err := GetInventoryModel().GetOne(inventoryID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n\nOrder: %+v \n\n", o)
	fmt.Printf("\n\nOrder ITEMLIST: %+v \n\n", o.ItemList)

	itemList := ItemList{inventory.ID, inventory.Product, inventory.InventoryCondition, inventory.ConditionString, orderAmount, inventory.Note}

	fmt.Println("ADDED?")
	fmt.Printf("GOT ITEMLIST: %+v \n", itemList)

	*o.ItemList = append(*o.ItemList, itemList)
}

//OrderModel ...
type OrderModel struct{}

var (
	orderModel *OrderModel
)

//InitializeOrderModel ...
func InitializeOrderModel() {
	GetOrderModel()
	loadedOrders = make(map[int]Order)
}

//GetOrderModel ...
func GetOrderModel() (model OrderModel) {
	if orderModel != nil {
		return *orderModel
	}
	orderModel = new(OrderModel)
	model = *orderModel

	return model
}

func getOrder(OrderID int) (order Order) {
	if loadedOrders[OrderID].ID > 0 {
		fmt.Println("cached order")
		fmt.Printf("%+v\n\n", loadedOrders[OrderID].ItemList)
		return loadedOrders[OrderID]
	}
	order = Order{}
	order.ID = OrderID
	order.ItemList = &[]ItemList{}
	loadedOrders[OrderID] = order
	return order
}

//GetOne ...
func (m OrderModel) GetOne(OrderID int) (order Order, err error) {
	rows, err := db.DB.Query("select tblOrder.order_id, tblOrder.customer_id, jncOrderItems.inventory_id, jncOrderItems.quantity, tblOrder.date_time from tblOrder LEFT OUTER JOIN jncOrderItems ON tblOrder.order_id = jncOrderItems.order_id WHERE tblOrder.order_id=$1", OrderID)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var orderID, customerID, inventoryID, quantity, orderTime int
		// var firstName, lastName, email, phoneNumber string
		// var address, city, state, country sql.NullString

		err = rows.Scan(&orderID, &customerID, &inventoryID, &quantity, &orderTime)
		if err != nil {
			panic(err)
		}

		customer, err := GetCustomerModel().GetOne(customerID)
		if err != nil {
			panic(nil)
		}

		// if err != nil {
		// return Order{}, err
		// }

		order = getOrder(orderID)
		order.Customer = customer
		order.OrderTime = orderTime

		order.AddToInventory(inventoryID, quantity)
	}

	return order, err
}

//GetAll ...
func (m OrderModel) GetList(Page int, Amount int) (orders []Order, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	rows, err := db.DB.Query("select tblOrder.order_id, tblOrder.customer_id, jncOrderItems.inventory_id, jncOrderItems.quantity from tblOrder INNER JOIN jncOrderItems ON tblOrder.order_id = jncOrderItems.order_id OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var orderID int
		var firstName, lastName, email, phoneNumber string
		var address, city, state, country sql.NullString

		err = rows.Scan(&orderID, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

		if err != nil {
			panic(err)
		}

		// orders = append(orders, Order{orderID, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

		// fmt.Println("orderID | username | department | created ")
		// fmt.Printf("%3v | %8v | %6v | %6v\n", orderID, firstName, lastName, email)
	}

	return orders, err
}
