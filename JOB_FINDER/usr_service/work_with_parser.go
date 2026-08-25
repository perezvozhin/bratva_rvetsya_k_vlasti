package usr_service

// функция инициализации парсера для примера
func (s *Usr_service) WorkWithParser() {
	//вызываем метод запуска парсера в методе сервиса
	//пользователя
	s.ParserService.RunParser()
}
