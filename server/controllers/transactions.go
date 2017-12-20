package controllers

import (
	"strconv"

	"github.com/ndbeals/brittanicus-final-project/forms"
	"github.com/ndbeals/brittanicus-final-project/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//TransactionController ...
type TransactionController struct{}

var transactionModel = new(models.TransactionModel)

//getTransactionID ...
func getTransactionID(c *gin.Context) int64 {
	session := sessions.Default(c)
	transactionID := session.Get("transaction_id")
	if transactionID != nil {
		return models.ConvertToInt64(transactionID)
	}
	return 0
}

//getSessionTransactionInfo ...
// func getSessionTransactionInfo(c *gin.Context) (transactionSessionInfo models.TransactionSessionInfo) {
// 	session := sessions.Default(c)
// 	transactionID := session.Get("transaction_id")
// 	if transactionID != nil {
// 		transactionSessionInfo.ID = models.ConvertToInt64(transactionID)
// 		transactionSessionInfo.Name = session.Get("transaction_name").(string)
// 		transactionSessionInfo.Email = session.Get("transaction_email").(string)
// 	}
// 	return transactionSessionInfo
// }

//Signin ...
func (ctrl TransactionController) Signin(c *gin.Context) {
	var signinForm forms.SigninForm

	if c.BindJSON(&signinForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signinForm})
		c.Abort()
		return
	}

	transaction, err := transactionModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("transaction_id", transaction.ID)
		// session.Set("transaction_email", transaction.Email)
		// session.Set("transaction_name", transaction.Name)
		session.Save()

		c.JSON(200, gin.H{"message": "Transaction signed in", "transaction": transaction})
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}

}

//Signup ...
func (ctrl TransactionController) Signup(c *gin.Context) {
	var signupForm forms.SignupForm

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}

	transaction, err := transactionModel.Signup(signupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if transaction.ID > 0 {
		session := sessions.Default(c)
		session.Set("transaction_id", transaction.ID)
		// session.Set("transaction_email", transaction.Email)
		// session.Set("transaction_name", transaction.Name)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "transaction": transaction})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this transaction", "error": err.Error()})
	}

}

//Signout ...
func (ctrl TransactionController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}

//GetOne ...
func (ctrl TransactionController) GetOne(c *gin.Context) {
	// transactionID := getTransactionID(c)

	// if transactionID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	uid := c.Param("id")

	if uid, err := strconv.ParseInt(uid, 10, 64); err == nil {

		data, err := transactionModel.GetOne(uid)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
