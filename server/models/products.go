package models

import (
	"database/sql"
	"math"

	"github.com/ndbeals/britannicus-final-project/db"
)

//Product ...
type Product struct {
	ID                int    `json:"product_id"`
	ISBN              string `json:"isbn"`
	ProductName       string `json:"product_name"`
	Author            string `json:"product_author"`
	Genre             string `json:"product_genre"`
	Publisher         string `json:"product_publisher"`
	ProductType       int    `json:"-"`
	ProductTypeString string `json:"product_type"`
	Description       string `json:"product_description"`
}

//ProductModel ...
type ProductModel struct{}

var (
	productModel   *ProductModel
	loadedProducts map[int]Product
	typeLookup     map[int]string
)

//InitializeProductModel ...
func InitializeProductModel() *ProductModel {
	GetProductModel()
	loadedProducts = make(map[int]Product)
	typeLookup = make(map[int]string)

	typeLookup[1] = "Soft Cover"
	typeLookup[2] = "Hard Cover"

	return productModel
}

//GetProductModel ...
func GetProductModel() (model ProductModel) {
	if productModel != nil {
		return *productModel
	}
	productModel = new(ProductModel)
	model = *productModel

	return model
}

//GetOne ...
func (m ProductModel) GetOne(ProductID int) (product Product, err error) {

	if (loadedProducts[ProductID] != Product{}) {
		return loadedProducts[ProductID], nil
	}

	row := db.DB.QueryRow("SELECT product_id, isbn, product_name, author, genre, publisher, product_type, description FROM tblProducts WHERE product_id=$1", ProductID)

	var productID, productType int
	var isbn, productName, author, genre, publisher, description sql.NullString

	err = row.Scan(&productID, &isbn, &productName, &author, &genre, &publisher, &productType, &description)
	if err != nil {
		return product, err
	}

	product = Product{productID, isbn.String, productName.String, author.String, genre.String, publisher.String, productType, typeLookup[productType], description.String}
	loadedProducts[productID] = product

	return product, err
}

//GetList ...
func (m ProductModel) GetList(Page int, Amount int) (products []Product, err error) {

	Page = int(math.Max(float64((Page-1)*Amount), 0))

	// dbaa := db.Init()
	rows, err := db.DB.Query("SELECT product_id, isbn, product_name, author, genre, publisher, product_type, description FROM tblProducts OFFSET $1 LIMIT $2", Page, Amount)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var productID, productType int
		// var price float64
		var isbn, productName, author, genre, publisher, description sql.NullString

		err = rows.Scan(&productID, &isbn, &productName, &author, &genre, &publisher, &productType, &description)
		// err = row.Scan(&uid, &ProductName)

		// fmt.Printf("WUT : %s \n", ProductFName)

		if err != nil {
			panic(err)
		}

		products = append(products, Product{productID, isbn.String, productName.String, author.String, genre.String, publisher.String, productType, typeLookup[productType], description.String})

		// products = append(products, Product{uid, firstName, lastName, email, phoneNumber, address.String, city.String, state.String, country.String})

		// fmt.Println("uid | username | department | created ")
		// fmt.Printf("%3v | %8v | %6v | %6v\n", uid, firstName, lastName, email)
	}

	return products, err
}
