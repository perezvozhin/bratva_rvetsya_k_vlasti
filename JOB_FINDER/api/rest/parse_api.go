package rest

import (
	"JOB_FINDER/usr_service"
	"net/http"
)

func ParseApi(userservice *usr_service.Usr_service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//при инициализации этого api (нажали кнопку на фронтенде)
		//запускается процесс парсера
		userservice.WorkWithParser()
		userservice.WorkWithParserPW()
	}
}
