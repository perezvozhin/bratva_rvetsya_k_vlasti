package caller

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func (c *Caller) ReadFromFileName(name string) (string, error) {
	/*
		чтение из указаного файла
		возвращает содержимое в формате строки
	*/

	file, err := os.Open(c.path + name)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		c.logger.Error(err)
		return "", err
	}
	return string(content), nil
}

func (c *Caller) CallWinApplication(name string) {
	/*
		Функция для запуска Windows
		приложения из командной строки

	*/

	cmd := exec.Command(name)
	if err := cmd.Run(); err != nil {
		c.logger.Error(err)
		return
	}
	err := cmd.Wait()
	if err != nil {
		c.logger.Error(err)
		return
	}
	c.logger.Info(name + " закрыт")
}

func (c *Caller) WritetoFile(text string, name string) error {
	file, err := os.OpenFile(c.path+name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		c.logger.Error(err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(text) // Запись текста в файл
	if err != nil {                 // Проверка, успешно ли прошла запись
		c.logger.Error(err)
		return err
	}
	c.logger.Info("Запись на сервере и диске прошла успешно")
	return nil
}
