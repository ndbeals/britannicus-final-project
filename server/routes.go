package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ndbeals/britannicus-final-project/controllers"
)

//
func initializeAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1", AuthenticationMiddleware())
	{
		/*** START USER ***/

		v1.GET("/user/:id", userController.GetOne)

		// v1.POST("/user/signup", user.Signup)
		// v1.GET("/user/signout", user.Signout)

		/*** START CUSTOMER ***/
		v1.GET("/customer/:id", customerController.GetOne)
		v1.PATCH("/customer/:id", customerController.Update)
		v1.DELETE("/customer/:id", customerController.Delete)
		v1.POST("/customer", customerController.Create)
		v1.GET("/customers/:page/:amount", customerController.GetList)
		v1.GET("/customer/:id/transactions", customerController.GetTransactions)

		/*** START INVENTORY ***/
		inventory := new(controllers.InventoryController)

		v1.GET("/inventory/:id", inventory.GetOne)
		v1.GET("/inventories/:page/:amount", inventory.GetList)

		/*** START PRODUCTS ***/
		v1.GET("/product/:id", productController.GetOne)
		v1.PATCH("/product/:id", productController.Update)
		v1.DELETE("/product/:id", productController.Delete)
		v1.POST("/product", productController.Create)
		v1.GET("/products/:page/:amount", productController.GetList)

		/*** START ORDERS ***/
		order := new(controllers.OrderController)

		v1.GET("/order/:id", order.GetOne)
		v1.GET("/orders/:page/:amount", order.GetList)

		v1.POST("order/new", order.CreateOrder)

		/*** START TRANSACTIONS ***/
		transaction := new(controllers.TransactionController)

		v1.GET("/transaction/:id", transaction.GetOne)
		// v1.GET("/customers/:page/:amount", order.GetList)

		/*** START Article ***/
		// article := new(controllers.ArticleController)

		// v1.POST("/article", article.Create)
		// v1.GET("/articles", article.All)
		// v1.GET("/article/:id", article.One)
		// v1.PUT("/article/:id", article.Update)
		// v1.DELETE("/article/:id", article.Delete)
	}
}

func initializeBasicRoutes(r *gin.Engine) {
	r.POST("/user/signin", userController.Signin)
	r.GET("/user/logout", userController.Signout)
	r.GET("/", func(c *gin.Context) {
		userID := controllers.GetUserID(c)
		fmt.Println(c.Get("user"))
		fmt.Println(sessions.Default(c).Get("user"))

		fmt.Println("ASF", userID, c.Request.URL)

		if userID == 0 {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Login Page",
			})
		} else {
			user, _ := controllers.GetLoggedinUser(c)

			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Home Page",
				"route": "/",
				"user":  user,
			})
		}
	})
}

func initializeControlRoutes(r *gin.Engine) {
	r.Use(AuthenticationMiddleware())

	r.GET("/products", productController.ProductListingPage)

	r.GET("/product/get", func(c *gin.Context) {
		product, err := productModel.GetOne(1)

		if err != nil {
			c.IndentedJSON(404, gin.H{"Message": "Product not found", "error": err.Error()})
			c.Abort()
			return
		}

		user, _ := controllers.GetLoggedinUser(c)

		c.Redirect(http.StatusTemporaryRedirect, "/product/get/1")

		c.HTML(http.StatusOK, "product.html", gin.H{
			"title":   "Product Detail Page",
			"route":   "/product",
			"user":    user,
			"product": product,
		})
	})

	r.GET("/product/get/:id", func(c *gin.Context) {
		productid := c.Param("id")

		if productid, err := strconv.ParseInt(productid, 10, 32); err == nil {
			productid := int(productid)
			product, err := productModel.GetOne(productid)

			if err != nil {
				// c.String(200, "<body onload=\"history.back()\"></body>" )
				// c.IndentedJSON(404, gin.H{"Message": "Product not found", "error": err.Error()})
				// c.Abort()
				// return
			}

			user, _ := controllers.GetLoggedinUser(c)

			c.HTML(http.StatusOK, "product.html", gin.H{
				"title":   "Product Detail Page",
				"route":   "/product",
				"user":    user,
				"product": product,
				"prodid":  productid,
			})
		} else {
			c.IndentedJSON(404, gin.H{"Message": "Invalid parameter"})
		}
	})

	r.GET("/product/delete/:id", func(c *gin.Context) {
		productid := c.Param("id")

		fmt.Println("delete,", productid)

		if productid, err := strconv.ParseInt(productid, 10, 32); err == nil {
			productid := int(productid)
			product, err := productModel.GetOne(productid)

			if err != nil {
				c.IndentedJSON(404, gin.H{"Message": "Product not found", "error": err.Error()})
				c.Abort()
				return
			}

			deleted, err := product.Delete()
			fmt.Println(deleted)
			fmt.Println(err)

			if deleted {
				user, _ := controllers.GetLoggedinUser(c)

				c.HTML(http.StatusOK, "products.html", gin.H{
					"title": "Products Page",
					"route": "/products",
					"user":  user,
				})
			} else {
				c.JSON(404, gin.H{"Message": "Could not delete product, as other data relies on this Record."})
			}

		} else {
			c.IndentedJSON(406, gin.H{"Message": "Invalid parameter"})
		}
	})

	r.GET("/product/create", func(c *gin.Context) {
		// product, err := productModel.GetOne(1)

		// if err != nil {
		// 	c.IndentedJSON(404, gin.H{"Message": "Product not found", "error": err.Error()})
		// 	c.Abort()
		// 	return
		// }

		user, _ := controllers.GetLoggedinUser(c)

		c.HTML(http.StatusOK, "newproduct.html", gin.H{
			"title": "Product Create Page",
			"route": "/product/create",
			"user":  user,
			// "productid": 1,
			// "product":   product,
		})
	})

	// Inventory routes
	r.GET("/inventory", func(c *gin.Context) {
		user, _ := controllers.GetLoggedinUser(c)

		c.HTML(http.StatusOK, "inventory.html", gin.H{
			"title": "Inventory Page",
			"route": "/inventory",
			"user":  user,
		})
	})

	// Customer routes
	r.GET("/customers", customerController.CustomersListingPage)

	r.GET("/customer/get", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/customer/get/1")
	})

	r.GET("/customer/get/:id", customerController.CustomerDetailPage)

	r.GET("/customer/create", func(c *gin.Context) {
		user, _ := controllers.GetLoggedinUser(c)

		c.HTML(http.StatusOK, "newcustomer.html", gin.H{
			"title": "New Customer Page",
			"route": "/customer/create",
			"user":  user,
			// "productid": 1,
			// "product":   product,
		})
	})
}
