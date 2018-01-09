package controllers

import (
	"strconv"

	"github.com/ndbeals/britannicus-final-project/forms"
	"github.com/ndbeals/britannicus-final-project/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CustomerController ...
type CustomerController struct{}

var customerModel = models.GetCustomerModel()

//getCustomerID ...
func getCustomerID(c *gin.Context) int {
	session := sessions.Default(c)
	CustomerID := session.Get("Customer_id")
	if CustomerID != nil {
		return models.ConvertToInt64(CustomerID)
	}
	return 0
}

// //getSessionCustomerInfo ...
// func getSessionCustomerInfo(c *gin.Context) (CustomerSessionInfo models.CustomerSessionInfo) {
// 	session := sessions.Default(c)
// 	CustomerID := session.Get("Customer_id")
// 	if CustomerID != nil {
// 		CustomerSessionInfo.ID = models.ConvertToInt64(CustomerID)
// 		CustomerSessionInfo.Name = session.Get("Customer_name").(string)
// 		CustomerSessionInfo.Email = session.Get("Customer_email").(string)
// 	}
// 	return CustomerSessionInfo
// }

//Signin ...
func (ctrl CustomerController) Signin(c *gin.Context) {
	var signinForm forms.SigninForm

	if c.BindJSON(&signinForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signinForm})
		c.Abort()
		return
	}

	Customer, err := customerModel.Signin(signinForm)
	if err == nil {
		session := sessions.Default(c)
		session.Set("Customer_id", Customer.ID)
		session.Set("Customer_email", Customer.Email)
		// session.Set("Customer_name", Customer.Name)
		session.Save()

		c.JSON(200, gin.H{"message": "Customer signed in", "Customer": Customer})
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}

}

//Signup ...
func (ctrl CustomerController) Signup(c *gin.Context) {
	var signupForm forms.SignupForm

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": signupForm})
		c.Abort()
		return
	}

	Customer, err := customerModel.Signup(signupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if Customer.ID > 0 {
		session := sessions.Default(c)
		session.Set("Customer_id", Customer.ID)
		session.Set("Customer_email", Customer.Email)
		// session.Set("Customer_name", Customer.Name)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "Customer": Customer})
	} else {
		c.JSON(406, gin.H{"message": "Could not signup this Customer", "error": err.Error()})
	}

}

//Signout ...
func (ctrl CustomerController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}

//GetTransactions
func (ctrl CustomerController) GetTransactions(c *gin.Context) {
	userid := c.Param("id")

	if userid, err := strconv.ParseInt(userid, 10, 32); err == nil {
		userid := int(userid)

		data, err := customerModel.GetTransactions(userid)
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

//GetOne ...
func (ctrl CustomerController) GetOne(c *gin.Context) {
	// CustomerID := getCustomerID(c)

	// if CustomerID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	userid := c.Param("id")

	if userid, err := strconv.ParseInt(userid, 10, 32); err == nil {
		userid := int(userid)

		data, err := customerModel.GetOne(userid)
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

//GetOne ...
func (ctrl CustomerController) GetList(c *gin.Context) {
	// CustomerID := getCustomerID(c)

	// if CustomerID == 0 {
	// 	c.JSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	page := c.Param("page")
	amount, err := strconv.ParseInt(c.Param("amount"), 10, 64)

	if err != nil {
		amount = 100
	}

	if page, err := strconv.ParseInt(page, 10, 32); err == nil {
		page, amount := int(page), int(amount)
		data, err := customerModel.GetList(page, amount)
		if err != nil {
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
