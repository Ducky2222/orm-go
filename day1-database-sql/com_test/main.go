package main

import (
	"duckorm"
	"fmt"
)

func main()  {
	engine, err := duckorm.NewEngine("sqlite3", "aw.db")
	if err != nil {
		return
	}
	defer engine.Close()
	session := engine.NewSession()
	_, _ = session.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, err = session.Raw("CREATE TABLE User(NAME text NOT NULL);").Exec()
	if err != nil {
		return
	}
	_, _ = session.Raw("CREATE TABLE User(NAME text NOT NULL);").Exec()
	result, err := session.Raw("INSERT INTO User VALUES (?), (?), (?), (?)", "TOM", "ANDY",
		"BETTY", "AMBER").Exec()
	if err != nil {
		return
	}
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}