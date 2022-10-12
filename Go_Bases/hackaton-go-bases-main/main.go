package main

import (
	"fmt"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Bases/hackaton-go-bases-main/internal/file"
	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Bases/hackaton-go-bases-main/internal/service"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var tickets []service.Ticket
	fl := file.File{Path: "/Users/saraisanjuan/Github/backpack-bcgow6-sarai-sanjuan/Go_Bases/hackaton-go-bases-main/tickets.csv"}

	data, err := fl.Read()
	if err != nil {
		panic(err)
	}
	tickets = append(tickets, data...)

	// Funcion para obtener tickets del archivo csv
	booking := service.NewBookings(tickets)

	ticket1 := service.Ticket{
		Id:          1101,
		Names:       "Martin Tapia",
		Email:       "martin.tapia@gmail.com",
		Destination: "Marruecos",
		Date:        "10:15",
		Price:       972,
	}
	ticket, err := booking.Create(ticket1)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ticket creado en el archivo .csv \n", ticket)

	readticket, err := booking.Read(1101)
	if err != nil {
		panic(err)
	}
	fmt.Println(readticket)

	updatik := service.Ticket{
		Id:          1101,
		Names:       "Martin Tapia Herrera",
		Email:       "martin.tapia.h@gmail.com",
		Destination: "Rusia",
		Date:        "15:55",
		Price:       862,
	}
	updateticket, err := booking.Update(1101, updatik)
	fmt.Println("Se actualizó el Ticket con los siguientes valores: \n", updateticket)

	deleteticket, err := booking.Delete(1000)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se eliminó el tiket con el id: ", deleteticket)

	err = fl.Write(tickets)
	if err != nil {
		panic(err)
	}

}
