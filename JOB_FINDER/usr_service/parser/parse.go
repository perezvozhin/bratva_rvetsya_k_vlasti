package parser

import (
	"go.uber.org/zap"
)

func (s *ParserService) RunParserProcess(log *zap.SugaredLogger) {
	vacancies, err := s.client.ParseHH(log)
	if err != nil {
		log.Errorw("failed to get data from HH", zap.Error(err))
		return
	}

	for _, v := range vacancies {
		err = s.repo.SaveVacancy(v, log)
		if err != nil {
			log.Errorw("failed to save vacancy", zap.Error(err), "vacancy_id", v.ID)
			continue
		}
	}
}
