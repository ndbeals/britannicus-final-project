package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
type DB struct {
	*sql.DB
}

const (
	//DbUser ...
	DbUser = "brittanicus"
	//DbPassword ...
	DbPassword = "brittanicus"
	//DbName ...
	DbName = "brittanicus"
)

var (
	db *sql.DB
	// DBE *sql.DB
)

//Init ...
func Init() sql.DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DbUser, DbPassword, DbName)

	var err error
	// db, err = ConnectDB(dbinfo)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return *db
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
	return db
}
