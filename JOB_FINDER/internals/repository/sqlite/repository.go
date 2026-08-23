package sqlite

import (
	"database/sql"
	"log"
)

type VacancyRepo struct {
	db *sql.DB
}

func NewVacancyRepo(dsn string) *VacancyRepo {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Fatal(err)

	}
	return &VacancyRepo{db: db}
}
