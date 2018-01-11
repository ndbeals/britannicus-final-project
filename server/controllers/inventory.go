package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ndbeals/britannicus-final-project/forms"
	"github.com/ndbeals/britannicus-final-project/models"

	"github.com/gin-gonic/gin"
)

//InventoryController ...
type InventoryController struct{}

var inventoryModel = models.GetInventoryModel()

//GetOne ...
func (ctrl InventoryController) GetOne(c *gin.Context) {
	inventoryid := c.Param("id")

	if inventoryid, err := strconv.ParseInt(inventoryid, 10, 32); err == nil {
		inventoryid := int(inventoryid)

		data, err := inventoryModel.GetOne(inventoryid)
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

//GetList ...
func (ctrl InventoryController) GetList(c *gin.Context) {
	page := c.Param("page")
	amount, err := strconv.ParseInt(c.Param("amount"), 10, 32)

	if err != nil {
		amount = 100
	}

	if page, err := strconv.ParseInt(page, 10, 32); err == nil {
		page, amount := int(page), int(amount)
		data, err := inventoryModel.GetList(page, amount)
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

//Update ...
func (ctrl InventoryController) Update(c *gin.Context) {
	inventoryid := c.Param("id")
	fmt.Println(inventoryid)

	if inventoryid, err := strconv.ParseInt(inventoryid, 10, 32); err == nil {
		inventoryid := int(inventoryid)

		var updateForm forms.UpdateInventoryForm
		err := c.BindJSON(&updateForm)
		if err != nil {
			panic(err)
			c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
			c.Abort()
			return
		}

		inventory, err := inventoryModel.GetOne(inventoryid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Inventory not found", "error": err.Error()})
			c.Abort()
			return
		}

		inventory.Update(updateForm)

		fmt.Println("updated")

		c.IndentedJSON(200, gin.H{"data": inventory})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//Create ...
func (ctrl InventoryController) Create(c *gin.Context) {
	// inventoryid := c.Param("id")
	fmt.Println("create")

	// if inventoryid, err := strconv.ParseInt(inventoryid, 10, 32); err == nil {
	// 	inventoryid := int(inventoryid)

	var updateForm forms.UpdateInventoryForm
	err := c.BindJSON(&updateForm)
	if err != nil {
		panic(err)
		c.IndentedJSON(404, gin.H{"message": "Invalid form", "form": updateForm})
		c.Abort()
		return
	}
	// fmt.Println(updateForm.ISBN, updateForm.InventoryName, updateForm.Author, updateForm.Genre, updateForm.Publisher, 1, "Soft Cover", updateForm.Description)

	product, _ := productModel.GetOne(updateForm.ProductID)

	inventory := models.Inventory{-1, product, 4, "Good", updateForm.Amount, updateForm.Price, updateForm.Note}
	if err != nil {
		c.IndentedJSON(404, gin.H{"Message": "Inventory not found", "error": err.Error()})
		c.Abort()
		return
	}

	newid, _ := inventory.Create()

	fmt.Println("Createds", newid)

	c.IndentedJSON(200, gin.H{"data": inventory, "id": newid})
	// } else {
	// 	c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	// }
}

//Delete ...
func (ctrl InventoryController) Delete(c *gin.Context) {
	inventoryid := c.Param("id")
	fmt.Println("Delete")

	if inventoryid, err := strconv.ParseInt(inventoryid, 10, 32); err == nil {
		inventoryid := int(inventoryid)

		inventory, err := inventoryModel.GetOne(inventoryid)
		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Inventory not found", "error": err.Error()})
			c.Abort()
			return
		}

		_, err = inventory.Delete()

		fmt.Println("deleted inventory from api")

		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Failed to delete", "error": err.Error()})
			c.Abort()
			return
		}

		c.IndentedJSON(200, gin.H{"Message": "Inventory Sucessfully deleted"})
	}
}

//InventoryListingPage ...
func (ctrl InventoryController) InventoryListingPage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "inventorys.html", gin.H{
		"title": "Inventorys Page",
		"route": "/inventorys",
		"user":  user,
	})
}

//InventoryDetailPage ...
func (ctrl InventoryController) InventoryDetailPage(c *gin.Context) {
	inventoryid := c.Param("id")

	if inventoryid, err := strconv.ParseInt(inventoryid, 10, 32); err == nil {
		inventoryid := int(inventoryid)
		inventory, err := inventoryModel.GetOne(inventoryid)

		if err != nil {
			// c.String(200, "<body onload=\"history.back()\"></body>" )
			// c.IndentedJSON(404, gin.H{"Message": "Inventory not found", "error": err.Error()})
			// c.Abort()
			// return
		}

		user, _ := GetLoggedinUser(c)

		c.HTML(http.StatusOK, "item.html", gin.H{
			"title":       "Inventory Detail Page",
			"route":       "/inventory/get",
			"user":        user,
			"inventory":   inventory,
			"inventoryid": inventoryid,
		})
	} else {
		c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
	}
}

//InventoryCreatePage ...
func (ctrl InventoryController) InventoryCreatePage(c *gin.Context) {
	user, _ := GetLoggedinUser(c)

	c.HTML(http.StatusOK, "newitem.html", gin.H{
		"title": "New Item Page",
		"route": "/inventory/create",
		"user":  user,
		// "productid": 1,
		// "product":   product,
	})
}
