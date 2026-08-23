package caller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

func (c *Caller) GetWithNoOpts(url string) (string, error) {
	/*
		Отправляет GET запрос на выбранный url
		без дополнительных параметров
	*/
	resp, err := c.client.Get(url)
	if err != nil {
		c.logger.Warn(err)
		return "", err
	}
	defer resp.Body.Close()
	defer c.client.CloseIdleConnections()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Warn(err)
		return "", err

	}
	return string(body), nil
}
func (c *Caller) GetWithResponse(url string) (map[string]interface{}, error) {
	/*
		Отправляет GET запрос на выбранный url
		возвращает ответ в формате map[string]interface{}
	*/
	var respMap map[string]interface{}
	resp, err := c.client.Get(url)
	if err != nil {
		c.logger.Warn(err)
		return respMap, err
	}
	defer resp.Body.Close()
	defer c.client.CloseIdleConnections()

	body, err := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &respMap)
	if err != nil {
		c.logger.Warn(err)
		return respMap, err

	}
	return respMap, nil
}
func (c *Caller) PostWithNoOpts(url string, body string) (string, error) {
	/*
		Отправляет POST запрос на выбранный url
		без дополнительных параметров
		content-type  по дефолту укзан, как application/json
	*/

	BodyToIoreader := strings.NewReader(body)
	response, err := c.client.Post(url, "application/json", BodyToIoreader)
	if err != nil {
		c.logger.Warn(err)
		return "", err
	}
	defer response.Body.Close()
	defer c.client.CloseIdleConnections()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.logger.Warn(err)
		return "", err
	}
	return string(responseBody), nil
}
