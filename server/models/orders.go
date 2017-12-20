package models

import (
	"database/sql"
	"fmt"
	"math"

	"github.com/ndbeals/brittanicus-final-project/db"
)

//Order ...
type Order struct {
	ID          int    `db:"order_id, primarykey, autoincrement" json:"order_id"`
	FirstName   string `db:"order_name" json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `db:"order_email" json:"order_email"`
	PhoneNumber string `db:"order_phone" json:"order_phone"`
	Address     string `json:"order_address"`
	City        string `json:"order_city"`
	State       string `json:"order_state"`
	Country     string `json:"order_country"`
	// UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	// CreatedAt int64  `db:"created_at" json:"created_at"`
}

//OrderModel ...
type OrderModel struct{}

//GetOne ...
func (m OrderModel) GetOne(OrderID int64) (order Order, err error) {

	fmt.Printf("GET ONE: %d \n", OrderID)

	// dbaa := db.Init()
	row := db.DB.QueryRow("SELECT order_id, first_name, last_name, email, phone_number FROM tblOrder WHERE Order_id=$1", OrderID)

	var uid int
	var OrderFName, OrderLName string
	var OrderEmail string
	var OrderPassword string

	err = row.Scan(&uid, &OrderFName, &OrderLName, &OrderEmail, &OrderPassword)
	// err = row.Scan(&uid, &OrderName)

	fmt.Printf("WUT : %s \n", OrderFName)

	if err != nil {
		return Order{}, err
	}

	order = Order{uid, OrderFName, OrderLName, OrderEmail, OrderPassword, "", "", "", ""}

	fmt.Printf("GOT Order: %+v \n", order)

	return order, err
}

//GetAll ...
func (m OrderModel) GetSet(Page int64, Amount int64) (orders []Order, err error) {

	Page = int64(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	rows, err := db.DB.Query("SELECT order_id, first_name, last_name, email, phone_number, address, city, state, country FROM tblOrder OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var uid int
		var firstName, lastName, email, phoneNumber string
		var address, city, state, country sql.NullString

		err = rows.Scan(&uid, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

		if err != nil {
			panic(err)
		}

		orders = append(orders, Order{uid, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

		// fmt.Println("uid | username | department | created ")
		// fmt.Printf("%3v | %8v | %6v | %6v\n", uid, firstName, lastName, email)
	}

	return orders, err
}
