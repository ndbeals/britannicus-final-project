package models

import (
	"errors"

	// "github.com/ndbeals/brittanicus-final-project/db"
	"github.com/ndbeals/brittanicus-final-project/db"
	"github.com/ndbeals/brittanicus-final-project/forms"
)

//Transaction ...
type Transaction struct {
	ID            int      `json:"transaction_id"`
	Customer      Customer `json:"customer"`
	OrderID       int      `json:"order_id"`
	PaymentMethod int      `json:"payment_method"`
	AmountPaid    float64  `json:"amount_paid"`
}

//TransactionModel ...
type TransactionModel struct{}

//Signin ...
func (m TransactionModel) Signin(form forms.SigninForm) (transaction Transaction, err error) {

	// err = db.GetDB().SelectOne(&Transaction, "SELECT id, email, password, name, updated_at, created_at FROM public.Transaction WHERE email=LOWER($1) LIMIT 1", form.Email)

	// if err != nil {
	// 	return Transaction, err
	// }

	// bytePassword := []byte(form.Password)
	// byteHashedPassword := []byte(Transaction.Password)

	// err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)

	// if err != nil {
	// 	return Transaction, errors.New("Invalid password")
	// }

	return transaction, nil
}

//Signup ...
func (m TransactionModel) Signup(form forms.SignupForm) (transaction Transaction, err error) {
	// getDb := db.GetDB()

	// checkTransaction, err := getDb.SelectInt("SELECT count(id) FROM public.Transaction WHERE email=LOWER($1) LIMIT 1", form.Email)

	// if err != nil {
	// 	return Transaction, err
	// }

	// if checkTransaction > 0 {
	// 	return Transaction, errors.New("Transaction exists")
	// }

	// bytePassword := []byte(form.Password)
	// hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }

	// res, err := getDb.Exec("INSERT INTO public.Transaction(email, password, name, updated_at, created_at) VALUES($1, $2, $3, $4, $5) RETURNING id", form.Email, string(hashedPassword), form.Name, time.Now().Unix(), time.Now().Unix())

	// if res != nil && err == nil {
	// 	err = getDb.SelectOne(&Transaction, "SELECT id, email, name, updated_at, created_at FROM public.Transaction WHERE email=LOWER($1) LIMIT 1", form.Email)
	// 	if err == nil {
	// 		return Transaction, nil
	// 	}
	// }

	return transaction, errors.New("Not registered")
}

//GetOne ...
func (m TransactionModel) GetOne(TransactionID int64) (transaction Transaction, err error) {
	// dbaa := db.Init()
	row := db.DB.QueryRow("SELECT transaction_id, customer_id, order_id, payment_method, amount_paid FROM tblTransaction WHERE transaction_id=$1", TransactionID)

	var transactionID, customerID, orderID, paymentMethod int
	var amountPaid float64

	err = row.Scan(&transactionID, &customerID, &orderID, &paymentMethod, &amountPaid)

	if err != nil {
		return Transaction{}, err
	}

	userID, err := NewCustomerModel().GetOne(int64(customerID))

	transaction = Transaction{transactionID, userID, orderID, paymentMethod, amountPaid}

	return transaction, err
}

//GetAll ...
func (m TransactionModel) GetAllByUser(UserID int64) (transactions []Transaction, err error) {

	// Page = int64(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	// rows, err := dbaa.Query("SELECT transaction_id, first_name, last_name, email, phone_number, address, city, state, country FROM tblTransaction OFFSET $1 LIMIT $2", Page, Amount)
	// if err != nil {
	// 	panic(err)
	// }

	// for rows.Next() {
	// 	var uid int
	// 	var firstName, lastName, email, phoneNumber string
	// 	var address, city, state, country sql.NullString

	// 	err = rows.Scan(&uid, &firstName, &lastName, &email, &phoneNumber, &address, &city, &state, &country)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	transactions = append(transactions, Transaction{uid, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

	// 	// fmt.Println("uid | username | department | created ")
	// 	// fmt.Printf("%3v | %8v | %6v | %6v\n", uid, firstName, lastName, email)
	// }

	return transactions, err
}
