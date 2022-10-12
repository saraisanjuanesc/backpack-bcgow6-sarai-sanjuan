package file

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/saraisanjuanesc/backpack-bcgow6-sarai-sanjuan/Go_Bases/hackaton-go-bases-main/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() ([]service.Ticket, error) {
	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}
	var listTickets []service.Ticket
	datas := strings.Split(string(data), "\n")
	var ntick service.Ticket
	for _, line := range datas {
		if len(line) > 0 {
			row := strings.Split(line, ",")
			ntick.Id, _ = strconv.Atoi(row[0])
			ntick.Names = row[1]
			ntick.Email = row[2]
			ntick.Destination = row[3]
			ntick.Date = row[4]
			ntick.Price, _ = strconv.Atoi(row[0])
			listTickets = append(listTickets, ntick)
		}
	}
	return listTickets, nil
}

func (f *File) Write(tickets []service.Ticket) error {
	var text string
	for _, ticket := range tickets {
		text += fmt.Sprintf("%d,%s,%s,%s,%s,%d\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	}

	//wticket := fmt.Sprintf("\n%d,%s,%s,%s,%s,%d", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
	t := []byte(text)
	//opfile, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_APPEND, 0666)
	err := os.WriteFile(f.Path, t, 0644)
	if err != nil {
		return err
	}
	/*_, err = opfile.Write(t)
	if err != nil {
		return err
	}
	opfile.Close()*/
	return nil
}
