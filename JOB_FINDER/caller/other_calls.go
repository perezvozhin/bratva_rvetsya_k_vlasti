package caller

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"

	"go.uber.org/zap"
)

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) ReadFromFileName(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		c.logger.Error(err)
		return "", fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		c.logger.Error(err)
		return "", fmt.Errorf("failed to read the file: %w", err)
	}

	return string(content), nil
}

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) CallWinApplication(name string) {
	cmd := exec.Command(name)
	if err := cmd.Run(); err != nil {
		c.logger.Error(err)
		return
	}

	if err := cmd.Wait(); err != nil {
		c.logger.Error(err)
		return
	}

	c.logger.Info(name + " открыт")
}

// TODO (DEPRECATED): Метод временно не используется
func (c *Caller) WriteToFile(filepath string, text string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	if _, err = file.WriteString(text); err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to write string to file: %w", err)
	}
	c.logger.Info("Text succesfully written to file", zap.String("filepath", filepath))
	return nil
}

func (c *Caller) OpenFile(filepath string) (*os.File, error) {
	file, err := os.Open(filepath)
	if err != nil {
		c.logger.Error(err)
		return nil, fmt.Errorf("failed to open the file: %w", err)
	}

	c.logger.Info("file opened", zap.String("filename", filepath))
	return file, nil
}

func (c *Caller) SaveJSON(filepath string, data interface{}) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to encode data: %w", err)
	}

	c.logger.Info("JSON succesfully loaded", zap.String("filepath", filepath))
	return nil
}

func (c *Caller) LoadJSON(filepath string, v interface{}) error {
	file, err := c.OpenFile(filepath)
	if err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to open the file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(v); err != nil {
		c.logger.Error(err)
		return fmt.Errorf("failed to decode data: %w", err)
	}

	c.logger.Info("JSON succesfully loaded", zap.String("filepath", filepath))
	return nil
}
