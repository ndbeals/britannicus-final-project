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

var (
	transModel *TransactionModel
)

//GetTransactionModel
func GetTransactionModel() (model TransactionModel) {

	if transModel != nil {
		return *transModel
	}
	transModel = new(TransactionModel)
	model = *transModel

	return model
}

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
func (m TransactionModel) GetOne(TransactionID int) (transaction Transaction, err error) {
	// dbaa := db.Init()
	row := db.DB.QueryRow("SELECT transaction_id, customer_id, order_id, payment_method, amount_paid FROM tblTransaction WHERE transaction_id=$1", TransactionID)

	var transactionID, customerID, orderID, paymentMethod int
	var amountPaid float64

	err = row.Scan(&transactionID, &customerID, &orderID, &paymentMethod, &amountPaid)

	if err != nil {
		return Transaction{}, err
	}

	userID, err := GetCustomerModel().GetOne(int(customerID))

	transaction = Transaction{transactionID, userID, orderID, paymentMethod, amountPaid}

	return transaction, err
}

//GetAll ...
func (m TransactionModel) GetAllByCustomer(CustomerID int) (transactions []Transaction, err error) {

	rows, err := db.DB.Query("SELECT transaction_id, customer_id, order_id, payment_method, amount_paid FROM tblTransaction tblTransaction WHERE customer_id=$1", CustomerID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transactionID, customerID, orderID, paymentMethod int
		var amountPaid float64

		err = rows.Scan(&transactionID, &customerID, &orderID, &paymentMethod, &amountPaid)
		if err != nil {
			panic(err)
		}

		userID, err := GetCustomerModel().GetOne(int(customerID))
		if err != nil {
			panic(err)
		}

		transactions = append(transactions, Transaction{transactionID, userID, orderID, paymentMethod, amountPaid})

	}

	return transactions, err
}
