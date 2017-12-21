package models

import (
	"database/sql"
	"math"

	"github.com/ndbeals/brittanicus-final-project/db"
)

//Order ...
type Order struct {
	ID          int         `db:"order_id, primarykey, autoincrement" json:"order_id"`
	CustomerID  int         `db:"order_name" json:"first_name"`
	Inventory   []Inventory `json:"last_name"`
	Email       string      `db:"order_email" json:"order_email"`
	PhoneNumber string      `db:"order_phone" json:"order_phone"`
	Address     string      `json:"order_address"`
	City        string      `json:"order_city"`
	State       string      `json:"order_state"`
	Country     string      `json:"order_country"`
	// UpdatedAt int  `db:"updated_at" json:"updated_at"`
	// CreatedAt int  `db:"created_at" json:"created_at"`
}

//OrderModel ...
type OrderModel struct{}

var (
	orderModel *OrderModel
)

//GetOrderModel ...
func GetOrderModel() (model OrderModel) {
	if orderModel != nil {
		return *orderModel
	}
	orderModel = new(OrderModel)
	model = *orderModel

	return model
}

//GetOne ...
func (m OrderModel) GetOne(OrderID int) (order Order, err error) {
	row := db.DB.QueryRow("SELECT order_id, first_name, last_name, email, phone_number FROM tblOrder WHERE Order_id=$1", OrderID)

	var orderID int
	var firstName, lastName, email, phoneNumber string
	var address, city, state, country sql.NullString

	err = row.Scan(&orderID, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

	if err != nil {
		panic(err)
	}

	// orders = append(orders, Order{orderID, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

	// if err != nil {
	// return Order{}, err
	// }

	// order = Order{orderID, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String}

	return order, err
}

//GetAll ...
func (m OrderModel) GetList(Page int, Amount int) (orders []Order, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	rows, err := db.DB.Query("SELECT order_id, first_name, last_name, email, phone_number, address, city, state, country FROM tblOrder OFFSET $1 LIMIT $2", Page, Amount)
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
