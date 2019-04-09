package services

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	driverName     = "mssql"
	dataSourceName = "sqlserver://sa:sa@localhost:1433"
	dbName         = "Ugs"
	containerName  = "mssql17-test"
)

type Brand struct {
	Id     uuid.UUID `db:"Id"`
	Code   string    `db:"Code"`
	Status int       `db:"Status"`
}

func QueryBrands() {
	// this Pings the database trying to connect, panics on error
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv

	// Query the database, storing results in a []Person (wrapped in []interface{})
	brand := []Brand{}

	query := fmt.Sprintf(`
	USE [%s]
	SELECT Id,Code,Status FROM brands
	`, dbName)

	e := db.Select(&brand, query)
	if e != nil {
		log.Fatalln(err)
	}
	jason, _ := brand[0], brand[1]

	fmt.Printf("%#v", jason)
	fmt.Println(jason.Id.String())
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}
	defer db.Close()
}
