package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/alyoshka/memory-db/db"
)

func main() {
	database := db.NewDatabase()
	consolereader := bufio.NewReader(os.Stdin)
	for {
		input, err := consolereader.ReadString('\n')
		if err != nil {
			fmt.Printf("Failed to read string, error: %s", err)
			return
		}
		input = strings.Trim(input, "\n")
		args := strings.Split(input, " ")
		switch args[0] {
		case "BEGIN":
			database.Begin()
		case "COMMIT":
			err := database.Commit()
			if err != nil {
				fmt.Println(err)
			}
		case "GET":
			result, err := database.Get(args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(result)
			}
		case "NUMEQUALTO":
			if len(args) == 2 {
				value, err := strconv.ParseInt(args[1], 10, 64)
				if err != nil {
					fmt.Println("Failed to parse argument: ", err)
				}
				result := database.NumEqualTo(value)
				fmt.Println(result)
			} else {
				fmt.Println("Not enough arguments")
			}
		case "ROLLBACK":
			err := database.Rollback()
			if err != nil {
				fmt.Println(err)
			}
		case "SET":
			if len(args) == 3 {
				value, err := strconv.ParseInt(args[2], 10, 64)
				if err != nil {
					fmt.Println("Failed to parse argument: ", err)
				}
				database.Set(args[1], value)
			} else {
				fmt.Println("Not enough arguments")
			}
		case "UNSET":
			if len(args) == 2 {
				database.Unset(args[1])
			} else {
				fmt.Println("Not enough arguments")
			}
		case "END":
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
