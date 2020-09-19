package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()
	createTb := `CREATE TABLE IF NOT EXISTS pairs (
      DEVICE_ID INTEGER NOT NULL,
       USER_ID INTEGER NOT NULL
   );`
	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("can't create table pairs", err)
	}
	fmt.Println("create table success.")
}
