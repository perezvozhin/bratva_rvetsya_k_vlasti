package caller

import (
	"io/ioutil"
	"os"
	"os/exec"

	"go.uber.org/zap"
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

//TODO : fix cycle imports usr_service->caller
// func (c *Caller) WritetoFile(text string, name string) (usr_service.FileChoose, error) {
// 	var retVal usr_service.FileChoose
//
// 	file, err := os.OpenFile(c.path+name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//
// 	if err != nil {
// 		c.logger.Error(err)
// 		return retVal, err
// 	}
// 	defer file.Close()
//
// 	_, err = file.WriteString(text) // Запись текста в файл
// 	if err != nil {                 // Проверка, успешно ли прошла запись
// 		c.logger.Error(err)
// 		return retVal, err
// 	}
// 	c.logger.Info("Запись на сервере и диске прошла успешно")
//
// 	readFile, err := os.ReadFile(file.Name())
// 	if err != nil {
// 		c.logger.Error(err)
// 		return retVal, err
// 	}
// 	retVal.File = file
// 	retVal.Text = string(readFile)
// 	return retVal, nil
// }

func (c *Caller) OpenFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {

		c.logger.Error(err)
		return nil, err
	}
	c.logger.Info("file opened", zap.String("filename", path))
	return file, nil
}

func (c *Caller) SaveJSON(name string, v any) error {
	// json.MarshalIndent -> os.WriteFile(c.path+name)
	return nil
}
func (c *Caller) LoadJSON(name string, v any) error {
	return nil
} // os.ReadFile -> json.Unmarshal
