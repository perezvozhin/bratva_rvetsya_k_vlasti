package usr_service

import (
	"JOB_FINDER/internals/FS_config"
	"time"

	"go.uber.org/zap"
)

/*
Функция проверяет установку даты в джсоне
Если дата пуста (вперые открывается приложение)
-устанавливает текующую дату в формате
2006-01-02 15:04:05

От установки текущей даты должно пройти 7 дней, затем:
рассылаем резюме в 10 корп -> снова вызываем SetTimer
*/
func (s *Usr_service) SetTimer(cfg *FS_config.Config, log *zap.SugaredLogger) {
	// if configTime == "" {
	// 	configTime = time.Now().Format("2006-01-02 15:04:05")
	// }

	loc, err := time.LoadLocation(cfg.TimeZone)
	if err != nil {
		log.Errorw("failed to load timezone", zap.Error(err))
		loc = time.UTC
	}

	now := time.Now().In(loc)
	log.Infof("Таймер запущен. Текущее время: %s", now.Format(time.RFC3339))
}

// Добавил что просил и немного отрефакторил код, твой закомментил,
// если че не так, то сделаешь как надо

// func (s *Usr_service) UpdateTimer(timeGot string) *FS_config.Config {
// 	//ВЫЗВАТЬ В АПИ-ОБРАЩЕНИИ К МИРОВОМУ ВРЕМЕНИ
// 	//логику отделить

// 	return &FS_config.Config{
// 		TimeStamp: timeGot,
// 	}
// }
