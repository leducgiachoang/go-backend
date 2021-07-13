package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go-backend/model"
	"log"
)

type SQL struct {
	Db *sqlx.DB
}

func (s *SQL) Connect(cfg model.Config) {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.DbHost,
		cfg.Database.DbPort,
		cfg.Database.DbUserName,
		cfg.Database.DbPassword,
		cfg.Database.DbName)
	s.Db = sqlx.MustConnect("postgres", dataSource)
	if err := s.Db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connect database successfully")
}

func (s *SQL) Close() {

}
