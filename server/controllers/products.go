package controllers

import (
	"strconv"

	"github.com/ndbeals/brittanicus-final-project/models"

	"github.com/gin-gonic/gin"
)

//ProductController ...
type ProductController struct{}

var productModel = models.GetProductModel()

//GetOne ...
func (ctrl ProductController) GetOne(c *gin.Context) {
	productid := c.Param("id")

	if productid, err := strconv.ParseInt(productid, 10, 32); err == nil {
		productid := int(productid)

		data, err := productModel.GetOne(productid)
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
func (ctrl ProductController) GetList(c *gin.Context) {
	page := c.Param("page")
	amount, err := strconv.ParseInt(c.Param("amount"), 10, 32)

	if err != nil {
		amount = 100
	}

	if page, err := strconv.ParseInt(page, 10, 32); err == nil {
		page, amount := int(page), int(amount)
		data, err := productModel.GetList(page, amount)
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
