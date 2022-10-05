package main

import (
	"fmt"
	"math/rand"
	"os"
)

const path = "./customers.txt"

type Customer struct {
	Legajo        int
	Name_Lastname string
	DNI           string
	Phone         string
	Address       string
}

func generateID() int {
	id := rand.Intn(200)
	return id
}

func newCustomer(id int, name string, dni string, tel string, addrs string) (cust Customer) {
	cust = Customer{
		Legajo:        id,
		Name_Lastname: name,
		DNI:           dni,
		Phone:         tel,
		Address:       addrs,
	}
	return
}

func readFile() (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		panic("Error: the indicated file was not found or is corrupt")
	}
	return
}

func validate_Existence(cust *Customer) (bool, error) {
	_, err := readFile()
	if err != nil {
		return false, err
	}
	return true, nil

}

func validate_dataCustomer(cust *Customer) (err error) {
	if cust.Name_Lastname == "" {
		err = fmt.Errorf("The Name cannot be empty")
	} else if cust.DNI == "" {
		err = fmt.Errorf("The DNI cannot be empty")
	} else if cust.Phone == "" {
		err = fmt.Errorf("The Phone cannot be empty")
	} else if cust.Address == "" {
		err = fmt.Errorf("The Address cannot be empty")
	}
	return
}
func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Fin de la ejecucion")
			fmt.Println("Se detectaron varios errores en tiempo de ejecucion")
			fmt.Println("No han quedado archivos abiertos")
		}
	}()

	id := generateID()
	if id <= 0 {
		panic("Cannot generate Legajo")
	}
	customer1 := newCustomer(id, "Martin Rojas", "", "2721975643", "Av. Morelos no.31 El aguila Nogales Ver.")
	status, errorvalidate := validate_Existence(&customer1)
	if errorvalidate != nil {
		fmt.Println(errorvalidate)
	}
	if status {
		fmt.Println("Customer already exists")
	}

	errordata := validate_dataCustomer(&customer1)
	if errordata != nil {
		panic(errordata)
	}
}
