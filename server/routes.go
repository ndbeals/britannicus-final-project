package main

import (
	"net/http"

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
		v1.GET("/inventory/:id", inventoryController.GetOne)
		v1.PATCH("/inventory/:id", inventoryController.Update)
		v1.DELETE("/inventory/:id", inventoryController.Delete)
		v1.POST("/inventory", inventoryController.Create)
		v1.GET("/inventories/:page/:amount", inventoryController.GetList)

		/*** START PRODUCTS ***/
		v1.GET("/product/:id", productController.GetOne)
		v1.PATCH("/product/:id", productController.Update)
		v1.DELETE("/product/:id", productController.Delete)
		v1.POST("/product", productController.Create)
		v1.GET("/products/:page/:amount", productController.GetList)

		/*** START ORDERS ***/
		v1.GET("/order/:id", orderController.GetOne)
		// v1.PATCH("/order/:id", orderController.Update)
		// v1.DELETE("/order/:id", orderController.Delete)
		v1.POST("/order", orderController.CreateOrder)

		v1.GET("/orders/:page/:amount", orderController.GetList)

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
	r.GET("/", func(c *gin.Context) {
		userID := controllers.GetUserID(c)
		// fmt.Println(c.Get("user"))
		// fmt.Println(sessions.Default(c).Get("user"))

		// fmt.Println("ASF", userID, c.Request.URL)

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

	r.POST("/user/signin", userController.Signin)
	r.GET("/user/logout", userController.Signout)
}

func initializeControlRoutes(r *gin.Engine) {
	r.Use(AuthenticationMiddleware())

	// Product routes
	r.GET("/products", productController.ProductListingPage)
	r.GET("/product/get", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/product/get/1")
	})
	r.GET("/product/get/:id", productController.ProductDetailPage)
	r.GET("/product/create", productController.ProductCreatePage)

	// Customer routes
	r.GET("/customers", customerController.CustomersListingPage)
	r.GET("/customer/get", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/customer/get/1")
	})
	r.GET("/customer/get/:id", customerController.CustomerDetailPage)
	r.GET("/customer/create", customerController.CustomerCreatePage)

	// Inventory routes
	r.GET("/inventory", func(c *gin.Context) {
		user, _ := controllers.GetLoggedinUser(c)

		c.HTML(http.StatusOK, "inventory.html", gin.H{
			"title": "Inventory Page",
			"route": "/inventory",
			"user":  user,
		})
	})
	r.GET("/inventory/get", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/inventory/get/1")
	})
	r.GET("/inventory/get/:id", inventoryController.InventoryDetailPage)
	r.GET("/inventory/create", inventoryController.InventoryCreatePage)

	// Order Routes
	r.GET("/orders", orderController.OrderListingPage)
	r.GET("/order/get", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/order/get/1")
	})
	r.GET("/order/get/:id", orderController.OrderDetailPage)
	r.GET("/order/create", orderController.OrderCreatePage)
}
