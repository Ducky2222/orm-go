package duckorm

import (
	"database/sql"
	"duckorm/log"
	"duckorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver string, dataSource string) (e *Engine, err error) {
	db, err := sql.Open(driver, dataSource)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Database connected successfully!")
	return &Engine{db: db}, err
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error(err)
		return
	}
	log.Info("Database closed successfully!")
}
func (e *Engine) NewSession() (s *session.Session) {
	return session.New(e.db)
}

