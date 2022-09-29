package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./inventorylist.csv")
	if err != nil {
		fmt.Println("ERROR!\n Cannot read file ")
	} else {
		datas := strings.Split(string(data), "\n")
		for _, line := range datas {
			if len(line) > 0 {
				row := strings.Split(line, ";")
				fmt.Printf("%s\t%20s\t%20s\n", row[0], row[1], row[2])
			}
		}
	}
}
