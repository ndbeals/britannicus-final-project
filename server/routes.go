package main

import (
	"fmt"
	"net/http"

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
		v1.GET("/customers/:page/:amount", customerController.GetList)
		v1.GET("/customer/:id/transactions", customerController.GetTransactions)

		/*** START INVENTORY ***/
		inventory := new(controllers.InventoryController)

		v1.GET("/inventory/:id", inventory.GetOne)
		v1.GET("/inventories/:page/:amount", inventory.GetList)

		/*** START PRODUCTS ***/
		product := new(controllers.ProductController)

		v1.GET("/product/:id", product.GetOne)
		v1.GET("/products/:page/:amount", product.GetList)

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
