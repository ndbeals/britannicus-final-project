package controllers

import (
	"fmt"
	"net/http"
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

//GetOne ...
func (ctrl CustomerController) GetOne(c *gin.Context) {
	// CustomerID := getCustomerID(c)

	// if CustomerID == 0 {
	// 	c.IndentedJSON(403, gin.H{"message": "Please login first"})
	// 	c.Abort()
	// 	return
	// }

	userid := c.Param("id")

	if userid, err := strconv.ParseInt(userid, 10, 32); err == nil {
		userid := int(userid)

		data, err := customerModel.GetOne(userid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, data)
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//GetOne ...
func (ctrl CustomerController) GetList(c *gin.Context) {
	// CustomerID := getCustomerID(c)

	// if CustomerID == 0 {
	// 	c.IndentedJSON(403, gin.H{"message": "Please login first"})
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
			c.IndentedJSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.IndentedJSON(200, data)
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Create ...
func (ctrl CustomerController) Create(c *gin.Context) {
	// customerid := c.Param("id")
	// fmt.Println(customerid)

	// if customerid, err := strconv.ParseInt(customerid, 10, 32); err == nil {
	// customerid := int(customerid)

	var updateForm forms.UpdateCustomerForm
	err := c.BindJSON(&updateForm)
	if err != nil {
		// // panic(err)
		c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
		c.Abort()
		return
	}

	customer := models.Customer{-1, updateForm.FirstName, updateForm.LastName, updateForm.Email, updateForm.PhoneNumber, updateForm.Address, updateForm.City, updateForm.State, updateForm.Country}

	newid, err := customer.Create()

	fmt.Println("creapdated")

	// c.Redirect(302, "/customer/get/"+string(newid))

	c.IndentedJSON(200, gin.H{"data": customer, "id": newid})
	// } else {
	// 	c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	// }

	fmt.Println("Createds")

	// c.IndentedJSON(200, gin.H{"data": newid})
	// } else {
	// 	c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	// }
}

//Update ...
func (ctrl CustomerController) Update(c *gin.Context) {
	customerid := c.Param("id")
	fmt.Println(customerid)

	if customerid, err := strconv.ParseInt(customerid, 10, 32); err == nil {
		customerid := int(customerid)

		var updateForm forms.UpdateCustomerForm
		err := c.BindJSON(&updateForm)
		if err != nil {
			// panic(err)
			c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
			c.Abort()
			return
		}

		customer, err := customerModel.GetOne(customerid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Customer not found", "error": err.Error()})
			c.Abort()
			return
		}

		_, err = customer.Update(updateForm)

		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Customer not Updated", "error": err.Error()})
			c.Abort()
			return
		}

		fmt.Println("updated")

		c.IndentedJSON(200, gin.H{"data": customer})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Delete ...
func (ctrl CustomerController) Delete(c *gin.Context) {
	customerid := c.Param("id")
	fmt.Println("Delete")

	if customerid, err := strconv.ParseInt(customerid, 10, 32); err == nil {
		customerid := int(customerid)

		customer, err := customerModel.GetOne(customerid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Customer not found", "error": err.Error()})
			c.Abort()
			return
		}

		_, err = customer.Delete()

		fmt.Println("deleted customer from api")

		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Failed to delete", "error": err.Error()})
			c.Abort()
			return
		}

		c.IndentedJSON(200, gin.H{"Message": "Customer Sucessfully deleted"})
	}
}

//CustomersListingPage ...
func (ctrl CustomerController) CustomersListingPage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "customers.html", gin.H{
		"title": "Customers Page",
		"route": "/customers",
		"user":  user,
	})
}

//CustomerDetailPage ...
func (ctrl CustomerController) CustomerDetailPage(c *gin.Context) {
	customerid := c.Param("id")

	if customerid, err := strconv.ParseInt(customerid, 10, 32); err == nil {
		customerid := int(customerid)
		customer, _ := customerModel.GetOne(customerid)

		// if err != nil {
		// 	c.IndentedJSON(404, gin.H{"Message": "customer not found", "error": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		user, _ := GetLoggedinUser(c)

		c.HTML(http.StatusOK, "customer.html", gin.H{
			"title":    "Customer Detail Page",
			"route":    "/customer",
			"user":     user,
			"customer": customer,
			"custid":   customerid,
		})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//CustomerCreatePage ...
func (ctrl CustomerController) CustomerCreatePage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "newcustomer.html", gin.H{
		"title": "New Customer Page",
		"route": "/customer/create",
		"user":  user,
		// "productid": 1,
		// "product":   product,
	})
}
