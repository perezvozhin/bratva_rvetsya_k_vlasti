package parser

import (
	"net/url"

	"go.uber.org/zap"
)

func (s *ParserService) RunParserProcess(log *zap.SugaredLogger) {
	//rework
	vacancies, err := s.client.UrlParseNoOpts("url")
	if err != nil {
		log.Errorw("failed to get data from HH", zap.Error(err))
		return
	}
	var vacanciespars url.URL
	vacanciespars = vacancies
	for _, v := range vacanciespars {
		err = s.repo.SaveVacancy(v, log)
		if err != nil {
			log.Errorw("failed to save vacancy", zap.Error(err), "vacancy_id", v.ID)
			continue
		}
	}
}
