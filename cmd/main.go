package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"vendors/pkg/SQLFUNC"
)

var command string
var isExit bool = true

type users struct { // Модель данных
	id int
	name string
	age int
	role string
}

func main() {
	fmt.Printf("\n\tConnection with Data Base...\n")
	db, err := sql.Open("sqlite3", "../db/users.db")
	if err != nil {
        panic(err)
    }
    defer db.Close()
    fmt.Printf("\tData Base Connected!\n\n")
    
	for isExit {
		inputCommand()
		switch command {
			case "exit":
				exit()
				fmt.Println("")
			case "help":
				help()
				fmt.Println("")
			case "create":
				var newUser users
				fmt.Print("\n\tName: ")
				fmt.Scan(&newUser.name)
				fmt.Print("\n\tAge: ")
				fmt.Scan(&newUser.age)
				SQLFUNC.Create(db, newUser.name, newUser.age)
                fmt.Print("\n\n")
			case "read":
				SQLFUNC.Read(db)
				fmt.Print("\n\n")
			case "update":
				var updateUser users
				fmt.Print("\n\tId: ")
				fmt.Scan(&updateUser.id)
				fmt.Print("\n\tName: ")
				fmt.Scan(&updateUser.name)
				fmt.Print("\n\tAge: ")
				fmt.Scan(&updateUser.age)
				SQLFUNC.Update(db, updateUser.name, updateUser.age, updateUser.id)
				fmt.Print("\n\n")
			case "delete":
				var deleteId int
				fmt.Print("\n\tId: ")
				fmt.Scan(&deleteId)
				SQLFUNC.Delete(db, deleteId)
				fmt.Print("\n\n")
			default:
				fmt.Printf("\n\tUnknown command\n\n")
		}
	}
}

func inputCommand() {
	fmt.Print("enter: ")
	fmt.Scan(&command)
}

func exit() {
	fmt.Printf("\n\tExit from programm...\n")
	isExit = false
}

func help() {
	fmt.Printf("\n\tNot create =)\n")
}