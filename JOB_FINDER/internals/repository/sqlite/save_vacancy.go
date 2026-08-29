package sqlite

import (
	"database/sql"
	"fmt"

	"JOB_FINDER/internals/domain"
)

const insertQuery = `
	INSERT INTO vacancies (
		id, source, title, company, url, experience,
		salary_from, salary_to, currency, responsibility, remote
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO NOTHING;
`

func salaryArgs(s *domain.Salary) (any, any, any) {
	if s == nil {
		return nil, nil, nil
	}
	return s.From, s.To, s.Currency
}

func (r *VacancyRepo) SaveVacancy(v domain.Vacancy) error {
	salaryFrom, salaryTo, currency := salaryArgs(v.Salary)

	if _, err := r.db.Exec(insertQuery,
		v.ID,
		v.Source,
		v.Name,
		v.EmployerName,
		v.URL,
		v.Experience,
		salaryFrom,
		salaryTo,
		currency,
		v.Responsibility,
		v.Remote,
	); err != nil {
		return fmt.Errorf("failed to save a vacancy: %w", err)
	}

	return nil
}

func (r *VacancyRepo) SaveVacancies(list []domain.Vacancy) error {
	if len(list) == 0 {
		return nil
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(insertQuery)
	if err != nil {
		return fmt.Errorf("prepare insert: %w", err)
	}
	defer stmt.Close()

	for _, v := range list {
		salaryFrom, salaryTo, currency := salaryArgs(v.Salary)

		if _, err := stmt.Exec(
			v.ID, v.Source, v.Name, v.EmployerName, v.URL, v.Experience,
			salaryFrom, salaryTo, currency, v.Responsibility, v.Remote,
		); err != nil {
			return fmt.Errorf("failed to save a vacancy: %w", err)
		}
	}

	return tx.Commit()
}

const selectColumns = `
	id, source, title, company, url, experience,
	salary_from, salary_to, currency, responsibility, remote
`

func (r *VacancyRepo) ListVacancies(limit, offset int) ([]domain.Vacancy, error) {
	if limit <= 0 {
		limit = 50
	}

	rows, err := r.db.Query(
		`SELECT `+selectColumns+` FROM vacancies ORDER BY created_at DESC, rowid DESC LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list vacancies: %w", err)
	}
	defer rows.Close()

	var out []domain.Vacancy
	for rows.Next() {
		v, err := scanVacancy(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}

	return out, rows.Err()
}

func (r *VacancyRepo) SearchVacancies(text string, limit, offset int) ([]domain.Vacancy, error) {
	if limit <= 0 {
		limit = 50
	}

	like := "%" + text + "%"

	rows, err := r.db.Query(
		`SELECT `+selectColumns+` FROM vacancies
		 WHERE title LIKE ? OR company LIKE ? OR responsibility LIKE ?
		 ORDER BY created_at DESC, rowid DESC LIMIT ? OFFSET ?`,
		like, like, like, limit, offset,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to search vacancies: %w", err)
	}
	defer rows.Close()

	var out []domain.Vacancy
	for rows.Next() {
		v, err := scanVacancy(rows)
		if err != nil {
			return nil, err
		}
		out = append(out, v)
	}

	return out, rows.Err()
}

func (r *VacancyRepo) CountSearch(text string) (int, error) {
	like := "%" + text + "%"

	var n int
	err := r.db.QueryRow(
		`SELECT COUNT(*) FROM vacancies
		 WHERE title LIKE ? OR company LIKE ? OR responsibility LIKE ?`,
		like, like, like,
	).Scan(&n)
	if err != nil {
		return 0, fmt.Errorf("failed to count search: %w", err)
	}
	return n, nil
}

func (r *VacancyRepo) CountVacancies() (int, error) {
	var n int
	if err := r.db.QueryRow(`SELECT COUNT(*) FROM vacancies`).Scan(&n); err != nil {
		return 0, fmt.Errorf("failed to count vacancies: %w", err)
	}
	return n, nil
}

func (r *VacancyRepo) GetVacancy(id string) (domain.Vacancy, error) {
	row := r.db.QueryRow(`SELECT `+selectColumns+` FROM vacancies WHERE id = ?`, id)

	v, err := scanVacancy(row)
	if err == sql.ErrNoRows {
		return domain.Vacancy{}, fmt.Errorf("vacancy %s not found", id)
	}
	return v, err
}

type scanner interface {
	Scan(dest ...any) error
}

func scanVacancy(s scanner) (domain.Vacancy, error) {
	var (
		v              domain.Vacancy
		experience     sql.NullString
		responsibility sql.NullString
		salaryFrom     sql.NullInt64
		salaryTo       sql.NullInt64
		currency       sql.NullString
	)

	err := s.Scan(
		&v.ID, &v.Source, &v.Name, &v.EmployerName, &v.URL, &experience,
		&salaryFrom, &salaryTo, &currency, &responsibility, &v.Remote,
	)
	if err != nil {
		return domain.Vacancy{}, err
	}

	v.Experience = experience.String
	v.Responsibility = responsibility.String

	if salaryFrom.Valid || salaryTo.Valid {
		v.Salary = &domain.Salary{
			From:     int(salaryFrom.Int64),
			To:       int(salaryTo.Int64),
			Currency: currency.String,
		}
	}

	return v, nil
}
