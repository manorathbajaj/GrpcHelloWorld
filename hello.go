//Package main not a grpc service, just a a test func to test and try database connectivity and stuff
package main

import (
	"database/sql"
	"fmt"

	// import needed for postgrtes
	_ "github.com/lib/pq"
)

func main() {
	fmt.Printf("hello, World \n")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	//fmt.Printf(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Printf("in error")
		panic(err)
	}
	defer db.Close()

	test, err := db.Query("select * from COMPANY")
	if err != nil {
		fmt.Printf("/n Error fetching rown /n")
		panic(err)
	}
	for test.Next() {
		err := test.Scan(&id, &name, &age, &address, &salary)
		if err != nil {
			fmt.Printf("/n Error mapping rows /n")
			panic(err)
		}
		fmt.Printf(name + "\n")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
}

const (
	host     = "localhost"
	port     = 5432
	user     = "dbuser"
	password = "password"
	dbname   = "test"
)

var (
	id      int
	name    string
	age     int
	address string
	salary  int
)
