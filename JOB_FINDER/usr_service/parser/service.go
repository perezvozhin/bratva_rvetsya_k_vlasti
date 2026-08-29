package parser

import (
	"context"
	"sort"
	"sync"

	"JOB_FINDER/internals/domain"

	"go.uber.org/zap"
)

type ParserService struct {
	providers []domain.Provider
	repo      domain.VacancyRepository
	logger    *zap.SugaredLogger
}

func NewParserService(repo domain.VacancyRepository, logger *zap.SugaredLogger, providers ...domain.Provider) *ParserService {
	return &ParserService{
		providers: providers,
		repo:      repo,
		logger:    logger,
	}
}

type SourceResult struct {
	Source string `json:"source"`
	Found  int    `json:"found"`
	Error  string `json:"error,omitempty"`
}

type SearchResult struct {
	Vacancies []domain.Vacancy `json:"vacancies"`
	Sources   []SourceResult   `json:"sources"`
}

func (s *ParserService) Search(ctx context.Context, q domain.SearchQuery) SearchResult {
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		all     []domain.Vacancy
		sources []SourceResult
	)

	for _, p := range s.providers {
		wg.Add(1)

		go func(p domain.Provider) {
			defer wg.Done()

			found, err := p.Search(ctx, q)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				s.logger.Warnw("provider failed", "source", p.Name(), "error", err)
				sources = append(sources, SourceResult{Source: p.Name(), Error: err.Error()})
				return
			}

			all = append(all, found...)
			sources = append(sources, SourceResult{Source: p.Name(), Found: len(found)})
		}(p)
	}

	wg.Wait()

	if err := s.repo.SaveVacancies(all); err != nil {
		s.logger.Errorw("failed to save vacancies", "count", len(all), "error", err)
	}

	sort.Slice(sources, func(i, j int) bool { return sources[i].Source < sources[j].Source })

	return SearchResult{Vacancies: all, Sources: sources}
}

func (s *ParserService) List(limit, offset int) ([]domain.Vacancy, error) {
	return s.repo.ListVacancies(limit, offset)
}

func (s *ParserService) ListByText(text string, limit, offset int) ([]domain.Vacancy, error) {
	if text == "" {
		return s.repo.ListVacancies(limit, offset)
	}
	return s.repo.SearchVacancies(text, limit, offset)
}

func (s *ParserService) CountByText(text string) (int, error) {
	if text == "" {
		return s.repo.CountVacancies()
	}
	return s.repo.CountSearch(text)
}

func (s *ParserService) Count() (int, error) {
	return s.repo.CountVacancies()
}

func (s *ParserService) Get(id string) (domain.Vacancy, error) {
	return s.repo.GetVacancy(id)
}
