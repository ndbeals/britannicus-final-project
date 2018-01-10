package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ndbeals/britannicus-final-project/controllers"
	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/models"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// authentication middleware
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := controllers.GetUserID(c)

		if userID == 0 {

			c.AbortWithStatus(404)
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	store := sessions.NewCookieStore([]byte("secret"))
	//sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("britannicus-session", store))

	dbs := db.Init()
	defer dbs.Close()

	models.InitializeInventoryModel()
	models.InitializeProductModel()
	models.InitializeOrderModel()

	// db.DB = dbs

	// dba := db.Init()

	// fmt.Println("# Inserting values")

	// var lastInsertId int
	// err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
	// checkErr(err)
	// fmt.Println("last inserted id =", lastInsertId)

	// fmt.Println("# Updating")
	// stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)

	// res, err := stmt.Exec("astaxieupdate", lastInsertId)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect, "rows changed")

	// fmt.Println("# Deleting")
	// stmt, err = db.Prepare("delete from userinfo where uid=$1")
	// checkErr(err)

	// res, err = stmt.Exec(lastInsertId)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect, "rows changed")

	user := new(controllers.UserController)
	v1 := r.Group("/v1", AuthenticationMiddleware())
	{
		/*** START USER ***/

		v1.GET("/user/:id", user.GetOne)

		// v1.POST("/user/signup", user.Signup)
		// v1.GET("/user/signout", user.Signout)

		/*** START CUSTOMER ***/
		customer := new(controllers.CustomerController)

		v1.GET("/customer/:id", customer.GetOne)
		v1.GET("/customers/:page/:amount", customer.GetList)
		v1.GET("/customer/:id/transactions", customer.GetTransactions)

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

	r.LoadHTMLGlob("./public/html/templates/*")
	// r.LoadHTMLGlob("./public/html/templates/*/*")

	// r.Static("/public", "./public")
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")

	r.POST("/user/signin", user.Signin)
	r.GET("/user/signout", user.Signout)
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

	r.GET("/products", func(c *gin.Context) {
		c.HTML(http.StatusOK, "products.html", gin.H{
			"title": "Products Page",
			"route": "/products",
		})
	})

	r.GET("/inventory", func(c *gin.Context) {
		c.HTML(http.StatusOK, "inventory.html", gin.H{
			"title": "Inventory Page",
			"route": "/inventory",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
