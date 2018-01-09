package controllers

import (
	"strconv"

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
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, gin.H{"data": data})
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
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
			c.JSON(404, gin.H{"Message": "Article not found", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(200, data)
	} else {
		c.JSON(404, gin.H{"Message": "Invalid parameter"})
	}
}
