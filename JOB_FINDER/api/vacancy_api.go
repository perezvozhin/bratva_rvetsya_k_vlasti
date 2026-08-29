package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"JOB_FINDER/internals/domain"
	"JOB_FINDER/usr_service/parser"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type VacancyHandler struct {
	parser *parser.ParserService
	logger *zap.SugaredLogger
}

func NewVacancyHandler(p *parser.ParserService, l *zap.SugaredLogger) *VacancyHandler {
	return &VacancyHandler{parser: p, logger: l}
}

func (h *VacancyHandler) RegisterRoutes(r chi.Router) {
	r.Get("/api/vacancies", h.list)
	r.Post("/api/vacancies/search", h.search)
}

func (h *VacancyHandler) list(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))
	text := r.URL.Query().Get("q")

	vacancies, err := h.parser.ListByText(text, limit, offset)
	if err != nil {
		h.logger.Errorw("list vacancies failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	total, err := h.parser.CountByText(text)
	if err != nil {
		h.logger.Errorw("count vacancies failed", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, h.logger, struct {
		Vacancies []vacancyDTO `json:"vacancies"`
		Total     int          `json:"total"`
	}{
		Vacancies: toVacancyDTOs(vacancies),
		Total:     total,
	})
}

func (h *VacancyHandler) search(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Text       string `json:"text"`
		Area       string `json:"area"`
		Experience string `json:"experience"`
		OnlyRemote bool   `json:"onlyRemote"`
		Page       int    `json:"page"`
		PerPage    int    `json:"perPage"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Errorw("decode search request failed", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.parser.Search(r.Context(), domain.SearchQuery{
		Text:       req.Text,
		Area:       req.Area,
		Experience: req.Experience,
		OnlyRemote: req.OnlyRemote,
		Page:       req.Page,
		PerPage:    req.PerPage,
	})

	writeJSON(w, h.logger, struct {
		Vacancies []vacancyDTO          `json:"vacancies"`
		Sources   []parser.SourceResult `json:"sources"`
	}{
		Vacancies: toVacancyDTOs(result.Vacancies),
		Sources:   result.Sources,
	})
}

type vacancyDTO struct {
	ID             string `json:"id"`
	Source         string `json:"source"`
	Title          string `json:"title"`
	Company        string `json:"company"`
	URL            string `json:"url"`
	Experience     string `json:"experience,omitempty"`
	Salary         string `json:"salary"`
	SalaryFrom     int    `json:"salaryFrom,omitempty"`
	SalaryTo       int    `json:"salaryTo,omitempty"`
	Currency       string `json:"currency,omitempty"`
	Requirement    string `json:"requirement,omitempty"`
	Responsibility string `json:"responsibility,omitempty"`
	Remote         bool   `json:"remote"`
}

func toVacancyDTOs(list []domain.Vacancy) []vacancyDTO {
	out := make([]vacancyDTO, 0, len(list))

	for _, v := range list {
		dto := vacancyDTO{
			ID:             v.ID,
			Source:         v.Source,
			Title:          v.Name,
			Company:        v.EmployerName,
			URL:            v.URL,
			Experience:     v.Experience,
			Salary:         formatSalary(v.Salary),
			Requirement:    v.Requirement,
			Responsibility: v.Responsibility,
			Remote:         v.Remote,
		}

		if v.Salary != nil {
			dto.SalaryFrom = v.Salary.From
			dto.SalaryTo = v.Salary.To
			dto.Currency = v.Salary.Currency
		}

		out = append(out, dto)
	}

	return out
}

func writeJSON(w http.ResponseWriter, logger *zap.SugaredLogger, payload any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Errorw("encode response failed", "error", err)
	}
}
