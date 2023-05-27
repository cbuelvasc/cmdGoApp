package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/cbuelvasc/cmdGoApp/datasources/mysql"
	"github.com/cbuelvasc/cmdGoApp/models"
)

func GetAllClients() {
	mysql.OpenConnection()

	sql := "SELECT id, name, email, phone FROM clients ORDER BY id DESC;"
	clients, err := mysql.ClientDB.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	defer mysql.CloseConnection()

	fmt.Println("\nGET ALL CLIENTS")
	fmt.Println("-----------------------------------------------------------")

	for clients.Next() {
		var client models.Client
		err := clients.Scan(&client.Id, &client.Name, &client.Email, &client.Phone)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %v | Name: %v | E-Mail %s | Phone %s \n", client.Id, client.Name, client.Email, client.Phone)
		fmt.Println("-----------------------------------------------------------")
	}
}

func GetClientById(id int) {
	mysql.OpenConnection()
	sql := "SELECT id, name, email, phone FROM clients WHERE id=?"

	clients, err := mysql.ClientDB.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	defer mysql.CloseConnection()

	fmt.Println("\nGET CLIENT BY ID")
	fmt.Println("-----------------------------------------------------------")

	for clients.Next() {

		var client models.Client
		err := clients.Scan(&client.Id, &client.Name, &client.Email, &client.Phone)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %v | Name: %v | E-Mail:%s | Phone: %s \n", client.Id, client.Name, client.Email, client.Phone)
		fmt.Println("-----------------------------------------------------------")
	}
}

func CreateClient(client models.Client) {
	mysql.OpenConnection()

	sql := "INSERT INTO clients VALUES(null, ?, ?, ?, ?);"

	result, err := mysql.ClientDB.Exec(sql, client.Name, client.Email, client.Phone, time.Now())
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	fmt.Println("Client created successfully")
}

func UpdateClient(client models.Client, id int) {
	mysql.OpenConnection()

	sql := "UPDATE clients SET name=?, email=?, phone=? WHERE id=?;"
	_, err := mysql.ClientDB.Exec(sql, client.Name, client.Email, client.Phone, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Client updated successfully")
}

func DeleteClient(id int) {
	mysql.OpenConnection()

	sql := "DELETE FROM clients WHERE id=?;"
	_, err := mysql.ClientDB.Exec(sql, id)

	if err != nil {
		panic(err)
	}

	fmt.Println("Client deleted successfully")
}
