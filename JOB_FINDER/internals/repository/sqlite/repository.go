package sqlite

import "database/sql"

type VacancyRepo struct {
	db *sql.DB
}

func NewVacancyRepo(db *sql.DB) *VacancyRepo {
	return &VacancyRepo{db: db}
}
