package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
// type DB struct {
// 	*sql.DB
// }

const (
	DB_USER     = "brittanicus"
	DB_PASSWORD = "brittanicus"
	DB_NAME     = "brittanicus"
)

var (
	DB *sql.DB
	// DBE *sql.DB
)

//Init ...
func Init() *sql.DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)

	var err error
	// db, err = ConnectDB(dbinfo)
	dbs, err := sql.Open("postgres", dbinfo)

	DB = dbs

	if err != nil {
		log.Fatal(err)
	}

	return dbs
}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	// DBE, err := sql.Open("postgres", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	// defer DBE.Close()

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *sql.DB {
	return DB
}
