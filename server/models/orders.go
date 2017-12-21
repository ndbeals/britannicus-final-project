package models

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/ndbeals/brittanicus-final-project/db"
	"github.com/ndbeals/brittanicus-final-project/forms"
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
	// fmt.Printf("\n\nOrder: %+v \n\n", o)
	// fmt.Printf("\n\nOrder ITEMLIST: %+v \n\n", o.ItemList)

	itemList := ItemList{inventory.ID, inventory.Product, inventory.InventoryCondition, inventory.ConditionString, orderAmount, inventory.Note}

	// fmt.Println("ADDED?")
	// fmt.Printf("GOT ITEMLIST: %+v \n", itemList)

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

//Create ...
func (m OrderModel) Create(form forms.CreateOrderForm) (order Order, err error) {
	// // turn item list into actual inventory list
	// for _,items := range form.ItemList {

	// }

	// Vars to for inserting into the table
	customerID := form.CustomerID
	dateTime := int(time.Now().Unix())

	res, err := db.DB.Exec("INSERT INTO public.tblOrder(customer_id, date_time) VALUES($1, $2 )", customerID, dateTime)

	log.Println(err)

	if res != nil && err == nil {
		orderID, err := res.LastInsertId()

		fmt.Printf("order id? %d \n\n", orderID)

		for item, quantity := range form.ItemList {
			db.DB.Exec("INSERT INTO public.jncOrderItems(order_id, inventory_id,quantity) VALUES($1, $2 )", orderID, item, quantity)

		}

		order, err = orderModel.GetOne(int(orderID))
		fmt.Printf("\n\nOrder saf DATA: %+v \n\n", order)
		if err != nil {
			return Order{}, err
		}

		return order, nil
	}

	return Order{}, errors.New("Couldn't create Order record for some reason " + err.Error())
}

//GetOne ...
func (m OrderModel) GetOne(OrderID int) (order Order, err error) {
	rows, err := db.DB.Query("select order_id, customer_id, date_time from tblOrder WHERE tblOrder.order_id=$1", OrderID)
	if err != nil {
		log.Fatal(err)
	}

	var cached bool

	for rows.Next() {
		var orderID, customerID, inventoryID, quantity, orderTime int
		err = rows.Scan(&orderID, &customerID, &orderTime)
		if err != nil {
			panic(err)
		}

		customer, err := GetCustomerModel().GetOne(customerID)
		if err != nil {
			panic(nil)
		}

		order, cached = getOrder(orderID)
		if cached == true {
			return order, err
		}
		order.Customer = customer
		order.OrderTime = orderTime

		rows, err := db.DB.Query("select jncOrderItems.inventory_id, jncOrderItems.quantity FROM jncOrderItems WHERE order_id=$1", orderID)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			err = rows.Scan(&inventoryID, &quantity)
			if err != nil {
				log.Fatal(err)
			}

			order.AddToInventory(inventoryID, quantity)
		}
	}

	return order, err
}

//GetAll ...
func (m OrderModel) GetList(Page int, Amount int) (orders []Order, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	rows, err := db.DB.Query("select tblOrder.order_id, tblOrder.customer_id, tblOrder.date_time from tblOrder OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		log.Fatal(err)
	}

	var order Order
	var cached bool

	for rows.Next() {
		var orderID, customerID, inventoryID, quantity, orderTime int
		err = rows.Scan(&orderID, &customerID, &orderTime)
		if err != nil {
			panic(err)
		}

		customer, err := GetCustomerModel().GetOne(customerID)
		if err != nil {
			panic(nil)
		}

		order, cached = getOrder(orderID)
		if cached {
			orders = append(orders, order)
		} else {

			order.Customer = customer
			order.OrderTime = orderTime

			rows, err := db.DB.Query("select jncOrderItems.inventory_id, jncOrderItems.quantity FROM jncOrderItems WHERE order_id=$1", orderID)
			if err != nil {
				log.Fatal(err)
			}

			for rows.Next() {
				err = rows.Scan(&inventoryID, &quantity)
				if err != nil {
					log.Fatal(err)
				}

				order.AddToInventory(inventoryID, quantity)
				// }

				// if len(*order.ItemList) < 2 {
				orders = append(orders, order)
				// fmt.Println("appended")
				// }
			}

		}
	}

	return orders, err
}

func getOrder(OrderID int) (order Order, cached bool) {
	if loadedOrders[OrderID].ID > 0 {
		return loadedOrders[OrderID], true
	}
	order = Order{}
	order.ID = OrderID
	order.ItemList = &[]ItemList{}
	loadedOrders[OrderID] = order
	return order, false
}
