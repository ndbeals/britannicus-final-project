package controllers

import (
	"strconv"

	"github.com/ndbeals/britannicus-final-project/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//TransactionController ...
type TransactionController struct{}

var transactionModel = new(models.TransactionModel)

//getTransactionID ...
func getTransactionID(c *gin.Context) int {
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

//Signout ...
func (ctrl TransactionController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.IndentedJSON(200, gin.H{"message": "Signed out..."})
}

//GetOne ...
func (ctrl TransactionController) GetOne(c *gin.Context) {
	transactionid := c.Param("id")

	if transactionid, err := strconv.ParseInt(transactionid, 10, 32); err == nil {
		transactionid := int(transactionid)

		data, err := transactionModel.GetOne(transactionid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, gin.H{"data": data})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
