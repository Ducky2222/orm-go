package duckorm

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func prepareDemo(db *sql.DB) {
	stmt, err := db.Prepare("SELECT NAME FROM User WHERE NAME = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = stmt.Close() }()
	rows, err := stmt.Query("AMBER")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = rows.Close() }()
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("We found %s\n", name)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func queryRowDemo(db *sql.DB) {
	var name string
	err := db.QueryRow("SELECT NAME FROM User WHERE NAME = ?", "ANDY").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("We found %s\n", name)
}

func main() {
	db, err := sql.Open("sqlite3", "aw.db")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = db.Close() }()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connection created successfully")
	}

	_, _ = db.Exec("DROP TABLE IF EXISTS User")
	_, err = db.Exec("CREATE TABLE User(NAME text NOT NULL);")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Table created successfully")
	}
	result, err := db.Exec("INSERT INTO User VALUES (?), (?), (?), (?)", "TOM", "ANDY", "BETTY", "AMBER")
	if err != nil {
		log.Fatal(err)
	} else {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}

	//QUERY
	rows, err := db.Query("SELECT NAME FROM User WHERE NAME = ? OR NAME = ?", "TOM", "BETTY")
	if err != nil {
		log.Fatal(err)
	} else {
		defer func() { _ = rows.Close() }()
		var name string
		for rows.Next() {
			err = rows.Scan(&name)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("We found %s\n", name)
			}
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}
	}

	//PREPARE
	prepareDemo(db)

	//QUERYROW
	queryRowDemo(db)


}
