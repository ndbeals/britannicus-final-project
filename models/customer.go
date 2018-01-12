package models

import (
	"database/sql"
	"math"

	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/forms"
)

//Customer ...
type Customer struct {
	ID          int    `db:"customer_id, primarykey, autoincrement" json:"customer_id"`
	FirstName   string `db:"customer_name" json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `db:"customer_email" json:"customer_email"`
	PhoneNumber string `db:"customer_phone" json:"customer_phone"`
	Address     string `json:"customer_address"`
	City        string `json:"customer_city"`
	State       string `json:"customer_state"`
	Country     string `json:"customer_country"`
}

//CustomerModel ...``
type CustomerModel struct {
}

var (
	loadedCustomers map[int]Customer
	customerModel   *CustomerModel
)

//InitializeOrderModel ...
func InitializeCustomerModel() *CustomerModel {
	GetCustomerModel()
	loadedCustomers = make(map[int]Customer)

	return customerModel
}

//CustomerModel ...
func GetCustomerModel() (model CustomerModel) {

	if customerModel != nil {
		return *customerModel
	}
	customerModel = new(CustomerModel)
	model = *customerModel

	return model
}

//GetOne ...
func (m CustomerModel) GetOne(CustomerID int) (customer Customer, err error) {
	if len(loadedCustomers) > 0 {
		if (loadedCustomers[CustomerID] != Customer{}) {
			return loadedCustomers[CustomerID], nil
		}
	} else {
		loadedCustomers = make(map[int]Customer)
	}

	// dbaa := db.Init()
	row := db.DB.QueryRow("SELECT customer_id, first_name, last_name, email, phone_number, customer_address, customer_city, customer_state, customer_country FROM tblCustomer WHERE Customer_id=$1", CustomerID)

	var customerID int
	var firstName, lastName, email, phoneNumber string
	var address, city, state, country sql.NullString
	err = row.Scan(&customerID, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

	if err != nil {
		return customer, err
	}

	customer = Customer{customerID, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String}
	loadedCustomers[customerID] = customer

	return customer, err
}

//GetList ...
func (m CustomerModel) GetList(Page int, Amount int) (customers []Customer, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	rows, err := db.DB.Query("SELECT customer_id, first_name, last_name, email, phone_number, customer_address, customer_city, customer_state, customer_country FROM tblCustomer ORDER BY customer_id OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		return customers, err
	}

	defer rows.Close()
	for rows.Next() {
		var uid int
		var firstName, lastName, email, phoneNumber string
		var address, city, state, country sql.NullString

		err = rows.Scan(&uid, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

		if err != nil {
			// panic(err)
		}

		customers = append(customers, Customer{uid, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})
	}

	return customers, err
}

// Create ...
func (this *Customer) Create() (int, error) {
	stmt, err := db.DB.Prepare("insert into tblcustomer(first_name, last_name, email, phone_number, customer_address, customer_city, customer_state, customer_country) values( $1, $2, $3, $4, $5, $6, $7, $8 ) RETURNING customer_id")

	if err != nil {

		return 0, err
	}

	results := stmt.QueryRow(this.FirstName, this.LastName, this.Email, this.PhoneNumber, this.Address, this.City, this.State, this.Country)

	var newid int
	err = results.Scan(&newid)

	if err != nil {
		return 0, err
	}

	delete(loadedCustomers, this.ID)
	customerModel.GetOne(int(newid))

	return int(newid), err
}

// Update ...
func (this *Customer) Update(newdata forms.UpdateCustomerForm) (bool, error) {

	stmt, err := db.DB.Prepare("update tblcustomer set first_name=$2, last_name=$3, email=$4, phone_number=$5, customer_address=$6, customer_city=$7, customer_state=$8, customer_country=$9 where customer_id=$1")
	if err != nil {
		return false, err
	}

	_, err = stmt.Exec(this.ID, newdata.FirstName, newdata.LastName, newdata.Email, newdata.PhoneNumber, newdata.Address, newdata.City, newdata.State, newdata.Country)

	if err != nil {
		return false, err
	}

	delete(loadedCustomers, this.ID)
	customerModel.GetOne(this.ID)

	return true, err
}

//Customer Delete ...
func (this *Customer) Delete() (bool, error) {
	_, err := db.DB.Query("DELETE FROM tblCustomer WHERE customer_id=$1", this.ID)

	if err != nil {
		return false, err
	}

	delete(loadedCustomers, this.ID)

	return true, err
}
