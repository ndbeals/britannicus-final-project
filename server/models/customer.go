package models

import (
	"database/sql"
	"errors"
	"math"

	// "github.com/ndbeals/britannicus-final-project/db"
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
	// UpdatedAt int  `db:"updated_at" json:"updated_at"`
	// CreatedAt int  `db:"created_at" json:"created_at"`
}

//CustomerModel ...``
type CustomerModel struct {
}

var (
	loadedCustomers map[int]Customer
	custModel       *CustomerModel
)

//CustomerModel ...
func GetCustomerModel() (model CustomerModel) {

	if custModel != nil {
		return *custModel
	}
	custModel = new(CustomerModel)
	model = *custModel

	return model
}

//Signin ...
func (m CustomerModel) Signin(form forms.SigninForm) (customer Customer, err error) {

	// err = db.GetDB().SelectOne(&Customer, "SELECT id, email, password, name, updated_at, created_at FROM public.Customer WHERE email=LOWER($1) LIMIT 1", form.Email)

	// if err != nil {
	// 	return Customer, err
	// }

	// bytePassword := []byte(form.Password)
	// byteHashedPassword := []byte(Customer.Password)

	// err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	// if err != nil {
	// 	return Customer, errors.New("Invalid password")
	// }

	return customer, nil
}

//Signup ...
func (m CustomerModel) Signup(form forms.SignupForm) (customer Customer, err error) {
	// getDb := db.GetDB()

	// checkCustomer, err := getDb.SelectInt("SELECT count(id) FROM public.Customer WHERE email=LOWER($1) LIMIT 1", form.Email)

	// if err != nil {
	// 	return Customer, err
	// }

	// if checkCustomer > 0 {
	// 	return Customer, errors.New("Customer exists")
	// }

	// bytePassword := []byte(form.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := getDb.Exec("INSERT INTO public.Customer(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())

	// if res != nil && err == nil {
	// 	err = getDb.SelectOne(&Customer, "SELECT id, email, name, updated_at, created_at FROM public.Customer WHERE email=LOWER($1) LIMIT 1", form.Email)
	// 	if err == nil {
	// 		return Customer, nil
	// 	}
	// }

	return customer, errors.New("Not registered")
}

//GetTransactions ...
func (m CustomerModel) GetTransactions(CustomerID int) (transactions []Transaction, err error) {

	return GetTransactionModel().GetAllByCustomer(CustomerID)
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
		panic(err)
	}

	customer = Customer{customerID, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String}
	loadedCustomers[customerID] = customer

	return customer, err
}

//GetList ...
func (m CustomerModel) GetList(Page int, Amount int) (customers []Customer, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	rows, err := db.DB.Query("SELECT customer_id, first_name, last_name, email, phone_number, customer_address, customer_city, customer_state, customer_country FROM tblCustomer OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var uid int
		var firstName, lastName, email, phoneNumber string
		var address, city, state, country sql.NullString

		err = rows.Scan(&uid, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

		if err != nil {
			panic(err)
		}

		customers = append(customers, Customer{uid, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

		// fmt.Println("uid | username | department | created ")
		// fmt.Printf("%3v | %8v | %6v | %6v\n", uid, firstName, lastName, email)
	}

	return customers, err
}
