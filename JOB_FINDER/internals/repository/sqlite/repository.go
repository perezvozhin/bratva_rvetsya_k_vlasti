package sqlite

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type VacancyRepo struct {
	db *sql.DB
}

func NewVacancyRepo(dsn string) (*VacancyRepo, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping sqlite: %w", err)
	}

	return &VacancyRepo{db: db}, nil
}

func (r *VacancyRepo) Close() error {
	return r.db.Close()
}
