package SQLFUNC

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type users struct { // Модель данных
	id int
	name string
	age int
	role string
}

func Create(db *sql.DB, name string, age int) {
	res, err := db.Exec("insert into users (name, age) values ($1, $2)", name, age)
	if err != nil {
    	panic(err)
	}
    fmt.Print("\n\tResult: ")
    fmt.Print(res.LastInsertId())
}

func Read(db *sql.DB) { // Стоит пределать
	rows, err := db.Query("select * from users")
	if err != nil {
        panic(err)
	}
    var data []users = []users{}
    for rows.Next() {
        usrs := users{}
        err := rows.Scan(&usrs.id, &usrs.name, &usrs.age, &usrs.role)
        if err != nil{
            fmt.Println(err)
            continue
        }
        data = append(data, usrs)
    }
    rows.Close()
	for _, d := range data {
		fmt.Printf("\n\tid: %d \t name: %s \t age: %d \t role: %s", d.id, d.name, d.age, d.role)
	}
}

func Update(db *sql.DB, name string, age int, id int) {
	res, err := db.Exec("update users set name = $1, age = $2 where id = $3", name, age, id)
	if err != nil {
		panic(err)
	}
	count, error := res.RowsAffected()
	if error != nil {
		panic(error)
	}
	if count == 0 {
		fmt.Print("\n\tError: Invalid ID")
	} else {
		fmt.Print("\n\tResult: ")
		fmt.Print(count)
	}
}

func Delete(db *sql.DB, id int) {
	res, err := db.Exec("delete from users where id = $1", id)
	if err != nil {
		panic(err)
	}
	fmt.Print("\n\tResult: ")
	fmt.Print(res.RowsAffected())
}