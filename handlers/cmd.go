package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/cbuelvasc/cmdGoApp/models"
)

var (
	id    int
	name  string
	email string
	phone string
)

func ExecuteCMD() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Choise an option: ")
	fmt.Println("1- Get all clients")
	fmt.Println("2- Get client by Id")
	fmt.Println("3- Create a client")
	fmt.Println("4- Edit a client")
	fmt.Println("5- Delete a client")

	if scanner.Scan() {
		for {

			if scanner.Text() == "1" {
				GetAllClients()
				return
			}

			if scanner.Text() == "2" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					id, _ = strconv.Atoi(scanner.Text())
				}
				GetClientById(id)
				return
			}

			if scanner.Text() == "3" {

				fmt.Println("Ingrese nombre : ")
				if scanner.Scan() {
					name = scanner.Text()
				}
				fmt.Println("Ingrese E-Mail : ")
				if scanner.Scan() {
					email = scanner.Text()
				}
				fmt.Println("Ingrese Teléfono : ")
				if scanner.Scan() {
					phone = scanner.Text()
				}
				client := models.Client{
					Name:  name,
					Email: email,
					Phone: phone,
				}
				CreateClient(client)
				return
			}

			if scanner.Text() == "4" {

				fmt.Println("Ingrese ID cliente : ")
				if scanner.Scan() {
					id, _ = strconv.Atoi(scanner.Text())
				}
				fmt.Println("Ingrese nombre : ")
				if scanner.Scan() {
					name = scanner.Text()
				}
				fmt.Println("Ingrese E-Mail : ")
				if scanner.Scan() {
					email = scanner.Text()
				}
				fmt.Println("Ingrese Teléfono : ")
				if scanner.Scan() {
					phone = scanner.Text()
				}
				client := models.Client{
					Name:  name,
					Email: email,
					Phone: phone,
				}
				UpdateClient(client, id)
				return

			}

			if scanner.Text() == "5" {
				fmt.Println("Ingrese el ID del cliente : ")
				if scanner.Scan() {
					id, _ = strconv.Atoi(scanner.Text())
				}
				DeleteClient(id)
				return
			}
		}
	}
}
