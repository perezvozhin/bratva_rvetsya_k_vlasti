package caller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) GetWithNoOpts(rawURL string) (string, error) {
	resp, err := c.client.Get(rawURL)
	if err != nil {
		c.logger.Warn(err)
		return "", fmt.Errorf("failed to send GET-request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Warn(err)
		return "", fmt.Errorf("failed to read body: %w", err)
	}
	return string(body), nil
}

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) GetWithResponse(rawURL string, target interface{}) error {
	resp, err := c.client.Get(rawURL)
	if err != nil {
		c.logger.Warn(err)
		return fmt.Errorf("failed to send GET-request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logger.Warn(err)
		return fmt.Errorf("failed to read body: %w", err)
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		c.logger.Warn(err)
		return err
	}
	return nil
}

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) PostWithNoOpts(rawURL string, body string) (string, error) {
	postBody := strings.NewReader(body)
	response, err := c.client.Post(rawURL, "application/json", postBody)
	if err != nil {
		c.logger.Warn(err)
		return "", fmt.Errorf("failed to send POST-request: %w", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		c.logger.Warn(err)
		return "", fmt.Errorf("failed to read body: %w", err)
	}
	return string(responseBody), nil
}

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) URLParseNoOpts(rawURL string) (url.URL, error) {
	urlm, err := url.Parse(rawURL)
	if err != nil {
		c.logger.Warn(err)
		return url.URL{}, fmt.Errorf("failed to parse URL: %w", err)
	}
	return *urlm, nil
}
