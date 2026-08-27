package usr_service

import (
	"JOB_FINDER/internals/FS_config"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
)

type TimeResponse struct {
	Datetime string `json:"datetime"`
}

func Synctime(
	logs *zap.SugaredLogger) {

	resp, err := http.Get("http://worldtimeapi.org") // Отправляем HTTP GET-запрос к внешнему API мирового времени

	if err != nil {
		return
	}
	defer resp.Body.Close() // Закрытие потока данных, во избежание переполнения памяти

	body, err := io.ReadAll(resp.Body) // Читаем сырое бинарное тело ответа и переводим его в байтовый срез (текст)

	var timeData TimeResponse
	err = json.Unmarshal(body, &timeData) // Распаковываем сырой JSON-текст в структуру-трафарет TimeResponse

	if err != nil {
		return
	}
	fmt.Println("Время от API:", timeData.Datetime)

	apiTime, err = time.Parse(time.RFC3339, timeData.Datetime) // Превращаем текстовую дату от API в системный тип времени для сравнения
	if err != nil {
		return
	}

	var cfgTime time.Time

	cfg := FS_config.Init(logs)
	cfgTime, err = time.Parse(time.RFC3339, cfg.ConfigTime) // Парсим текстовую дату из конфига в системный тип времени

	if err != nil {
		return
	}

	// Проверяем, прошло ли 7 дней с момента последнего обновления конфигурации
	if apiTime.After(cfgTime.AddDate(0, 0, 7)) {
		logs.Info("Прошло 7 дней. Пора обновлять время конфигурации!")

		// Обновляем дату в структуре конфигурации в памяти на самую свежую от API
		cfg.ConfigTime = timeData.Datetime

		// Кодируем обновлённую структуру обратно в JSON-текст с красивыми отступами
		updateJSON, err := json.MarshalIndent(cfg, "", " ")

		// Перезаписываем файл config.json на диске новыми данными
		_ = os.WriteFile("../config/config.json", updateJSON, 0644)
	} else {
		logs.Info("7 дней еще не прошло, обновление не требуется")
	}
}
