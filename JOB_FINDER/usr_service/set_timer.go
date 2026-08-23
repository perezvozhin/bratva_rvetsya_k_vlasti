package usr_service

import (
	"JOB_FINDER/internals/FS_config"
	"time"
)

/*
Функция проверяет установку даты в джсоне
Если дата пуста (вперые открывается приложение)
-устанавливает текующую дату в формате
2006-01-02 15:04:05

От установки текущей даты должно пройти 7 дней, затем:
рассылаем резюме в 10 корп -> снова вызываем SetTimer
*/
func (s *Usr_service) SetTimer(configTime string) {
	if configTime == "" {
		configTime = time.Now().Format("2006-01-02 15:04:05")
	}
}

func (s *Usr_service) UpdateTimer(timeGot string) *FS_config.Config {
	//ВЫЗВАТЬ В АПИ-ОБРАЩЕНИИ К МИРОВОМУ ВРЕМЕНИ
	//логику отделить

	return &FS_config.Config{
		TimeStamp: timeGot,
	}
}
