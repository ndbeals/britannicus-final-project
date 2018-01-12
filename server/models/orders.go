package models

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/forms"
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
	Price              float64 `json:"item_price"`
	Note               string  `json:"note"`
}

//Order ...
type Order struct {
	ID         int         `json:"order_id"`
	Customer   Customer    `json:"customer"`
	OrderTime  int         `json:"order_time"`
	ItemList   *[]ItemList `json:"item_list"`
	TotalPrice float64     `json:"total_price"`
}

//OrderModel ...
type OrderModel struct{}

var (
	orderModel *OrderModel
)

//InitializeOrderModel ...
func InitializeOrderModel() *OrderModel {
	GetOrderModel()
	loadedOrders = make(map[int]Order)

	return orderModel
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

//GetOne ...
func (m OrderModel) GetOne(OrderID int) (order Order, err error) {
	rows, err := db.DB.Query("select order_id, customer_id, date_time from tblOrder WHERE tblOrder.order_id=$1", OrderID)
	if err != nil {
		return order, err
	}

	var cached bool

	for rows.Next() {
		var orderID, customerID, inventoryID, quantity, orderTime int
		err = rows.Scan(&orderID, &customerID, &orderTime)
		if err != nil {
			// panic(err)
		}

		customer, err := GetCustomerModel().GetOne(customerID)
		if err != nil {
			// panic(nil)
		}

		order, cached = getOrder(orderID)
		order.Customer = customer
		if cached == true {
			return order, err
		}
		order.OrderTime = orderTime

		rows, err := db.DB.Query("select jncOrderItems.inventory_id, jncOrderItems.quantity FROM jncOrderItems WHERE order_id=$1", orderID)
		if err != nil {
			// panic(err)
		}

		for rows.Next() {
			err = rows.Scan(&inventoryID, &quantity)
			if err != nil {
				// panic(err)
			}

			order.AddToInventory(inventoryID, quantity)
		}
	}

	loadedOrders[order.ID] = order

	return order, err
}

//GetAll ...
func (m OrderModel) GetList(Page int, Amount int) (orders []Order, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	rows, err := db.DB.Query("select tblOrder.order_id, tblOrder.customer_id, tblOrder.date_time from tblOrder ORDER BY tblOrder.order_id OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		return orders, err
	}

	var order Order
	var cached bool

	for rows.Next() {
		var orderID, customerID, inventoryID, quantity, orderTime int
		err = rows.Scan(&orderID, &customerID, &orderTime)
		if err != nil {
			// panic(err)
		}

		customer, err := GetCustomerModel().GetOne(customerID)
		if err != nil {
			// panic(nil)
		}

		order, cached = getOrder(orderID)
		order.Customer = customer
		if cached {
			orders = append(orders, order)
		} else {

			order.OrderTime = orderTime

			rows, err := db.DB.Query("select jncOrderItems.inventory_id, jncOrderItems.quantity FROM jncOrderItems WHERE order_id=$1", orderID)
			if err != nil {
				// panic(err)
			}

			for rows.Next() {
				err = rows.Scan(&inventoryID, &quantity)
				if err != nil {
					// panic(err)
				}

				order.AddToInventory(inventoryID, quantity)

				orders = append(orders, order)
			}
			loadedOrders[order.ID] = order
		}
	}

	return orders, err
}

//Create ...
func (m OrderModel) Create(form forms.CreateOrderForm) (order Order, err error) {
	// Vars to for inserting into the table
	var orderID int
	customerID := form.CustomerID
	dateTime := int(time.Now().Unix())

	row := db.DB.QueryRow("INSERT INTO public.tblOrder(customer_id, date_time) VALUES($1, $2) RETURNING order_id", customerID, dateTime)
	err = row.Scan(&orderID)

	if orderID > 0 && err == nil {
		for item, quantity := range form.ItemList {
			fmt.Println("ITEM QUANT", item, quantity)
			_, err := db.DB.Exec("INSERT INTO public.jncOrderItems(order_id, inventory_id,quantity) VALUES($1, $2, $3)", orderID, item, quantity)
			if err != nil {
				return order, err
			}
		}

		order, err = orderModel.GetOne(int(orderID))
		if err != nil {
			return Order{}, err
		}

		return order, nil
	}

	return Order{}, errors.New("Couldn't create Order record: " + err.Error())
}

func (o *Order) AddToInventory(inventoryID int, orderAmount int) error {
	inventory, err := GetInventoryModel().GetOne(inventoryID)
	if err != nil {
		return err
		// panic(err)
	}

	itemList := ItemList{inventory.ID, inventory.Product, inventory.InventoryCondition, inventory.ConditionString, orderAmount, inventory.Price, inventory.Note}

	o.TotalPrice = o.TotalPrice + (inventory.Price * float64(orderAmount))

	*o.ItemList = append(*o.ItemList, itemList)

	return nil
}

//Order Delete ...
func (this *Order) Delete() (bool, error) {
	_, err := db.DB.Query("DELETE FROM jncorderitems WHERE order_id=$1", this.ID)
	if err != nil {
		// // panic(err)
		return false, err
	}

	_, err = db.DB.Query("DELETE FROM tblorder WHERE order_id=$1", this.ID)

	fmt.Println("deleted order model")

	if err != nil {
		// // panic(err)
		return false, err
	}

	return true, err
}

// Update ...
func (this *Order) Update(newdata forms.UpdateOrderForm) (success bool, err error) {

	// stmt, err := db.DB.Prepare("update tblorder set order_condition=$2, amount=$3, price=$4, notes=$5 where order_id=$1")
	// if err != nil {
	// return false, err
	// }

	// _, err = stmt.Exec(this.ID, newdata.OrderCondition, newdata.Amount, newdata.Price, newdata.Note)

	// if err != nil {
	// return false, err
	// }
	// fmt.Println(newdata.ItemList)

	orderModel.GetOne(this.ID)

	return true, err
}
