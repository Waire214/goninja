package main

import (
	"fmt"

	"github.com/Waire214/goninja/models"
	"github.com/Waire214/goninja/router"
)

type Customers struct {
	id    int
	name  string
	email string
}

var customer Customers
var customers []Customers

func LogData() {
	rows, err := models.Db.Query("SELECT * FROM customersinfo")
	if err != nil {
		fmt.Println("error loading table")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&customer.id, &customer.name, &customer.email)
		if err != nil {
			fmt.Println("error loading data")
		}
		customers = append(customers, customer)
	}
	fmt.Println(customers)
}

func main() {
	models.Database()
	LogData()
	router.ProductEndPoint()
}
