package session

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	testDB, err := sql.Open("sqlite3", "../aw.db")
	if err != nil {
		log.Fatal("In TestMain, database open failed ")
		return
	}
	code := m.Run()
	testDB.Close()
	os.Exit(code)
}

func NewSession() (s *Session) {
	return &Session{db: testDB}
}

func TestSession_Exec(t *testing.T) {
	s := NewSession()

	_, _ = s.db.Exec("DROP TABLE IF EXISTS User")
	_, err := s.db.Exec("CREATE TABLE User(NAME text NOT NULL);")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("In TestSession_Exec, Table created successfully")
	}
}