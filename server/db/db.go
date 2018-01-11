package db

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
)

//DB ...
// type DB struct {
// 	*sql.DB
// }

const (
	DB_HOST     = "192.168.0.13"
	DB_USER     = "britannicus"
	DB_PASSWORD = "britannicus"
	DB_NAME     = "britannicus"
)

var (
	DB *sql.DB
	// DBE *sql.DB
)

//Init ...
func Init() *sql.DB {

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)

	fmt.Println(dbinfo)

	var err error
	// db, err = ConnectDB(dbinfo)
	dbs, err := sql.Open("postgres", "postgres://britannicus:britannicus@192.168.0.13/britannicus?sslmode=disable")

	DB = dbs

	if err != nil {
		// panic(err)
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
