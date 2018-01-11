package models

import (
	// "github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/db"
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

	rows, err := db.DB.Query("SELECT transaction_id, customer_id, order_id, payment_method, amount_paid FROM tblTransaction tblTransaction WHERE customer_id=$1 ORDER BY transaction_id", CustomerID)
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
