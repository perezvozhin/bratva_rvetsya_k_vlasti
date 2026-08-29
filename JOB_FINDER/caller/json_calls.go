package caller

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Caller) GetJSON(ctx context.Context, rawURL string, headers map[string]string, target interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return fmt.Errorf("failed to build request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
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

	if resp.StatusCode >= http.StatusBadRequest {
		c.logger.Warnw("bad status from source", "url", rawURL, "status", resp.StatusCode)
		return fmt.Errorf("bad status %d: %s", resp.StatusCode, string(body))
	}

	if err := json.Unmarshal(body, target); err != nil {
		c.logger.Warn(err)
		return fmt.Errorf("failed to decode json: %w", err)
	}
	return nil
}
