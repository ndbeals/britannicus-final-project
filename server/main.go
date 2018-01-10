package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ndbeals/britannicus-final-project/controllers"
	"github.com/ndbeals/britannicus-final-project/db"
	"github.com/ndbeals/britannicus-final-project/models"
)

var wannatest = "cross"

var (
	userModel           *models.UserModel
	userController      *controllers.UserController
	productModel        *models.ProductModel
	productController   *controllers.ProductController
	orderModel          *models.OrderModel
	orderController     *controllers.OrderController
	inventoryModel      *models.InventoryModel
	inventoryController *controllers.InventoryController
	customerController  *controllers.CustomerController
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())

	store := sessions.NewCookieStore([]byte("secret"))
	//sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("britannicus-session", store))

	dbs := db.Init()
	defer dbs.Close()

	inventoryModel = models.InitializeInventoryModel()
	productModel = models.InitializeProductModel()
	orderModel = models.InitializeOrderModel()

	userController = new(controllers.UserController)
	customerController = new(controllers.CustomerController)

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

	// fmt.Println("main delme", userController.Delme)

	r.LoadHTMLGlob("./public/html/templates/*")
	// r.LoadHTMLGlob("./public/html/templates/*/*")

	// r.Static("/public", "./public")
	r.Static("/js", "./public/js")
	r.Static("/css", "./public/css")

	initializeAPIRoutes(r)

	initializeBasicRoutes(r)

	r.Use(AuthenticationMiddleware())

	r.GET("/products", AuthenticationMiddleware(), func(c *gin.Context) {
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
