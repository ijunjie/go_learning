package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	flag.Parse()
	argCount := flag.NArg()
	if argCount != 1 {
		fmt.Println("Usage: flag <md5>\nExample: flag c4ca4238a0b923820dcc509a6f75849b")
		return
	}
	md5 := flag.Arg(0)

	db, err1 := sql.Open("sqlite3", "./data.db")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT * from t WHERE m LIKE '%s'", md5+"%")
	rows, err2 := db.Query(query)
	if err2 != nil {
		log.Fatal(err2)
	}

	for rows.Next() {
		var v string
		var m string
		var t string
		err := rows.Scan(&v, &m, &t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("md5: %-32s value: %-10s type: %-1s", m, v, t))
	}
}
