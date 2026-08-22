package sqlite

import (
	"JOB_FINDER/internals/domain"
	"fmt"

	"go.uber.org/zap"
)

func (r *VacancyRepo) SaveVacancy(v domain.Vacancy, log *zap.SugaredLogger) error {
	var salaryFrom, salaryTo, currency any

	if v.Salary != nil {
		salaryFrom = v.Salary.From
		salaryTo = v.Salary.To
		currency = v.Salary.Currency
	}

	query := `
		INSERT INTO vacancies (id, title, company, url, experience, salary_from, salary_to, currency)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO NOTHING;
	`
	if _, err := r.db.Exec(query,
		v.ID,
		v.Name,
		v.EmployerName,
		v.URL,
		v.Requirement,
		salaryFrom,
		salaryTo,
		currency,
	); err != nil {
		return fmt.Errorf("failed to save a vacancy: %w", err)
	}

	return nil
}
