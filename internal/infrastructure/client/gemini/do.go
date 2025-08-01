package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wordwiz/internal/domain/model"
	"wordwiz/internal/infrastructure/client/gemini/dto"
)

func (c *Client) Do(ctx context.Context, request string) (string, error) {
	req := dto.Request{Contents: []dto.RequestContent{
		{
			Parts: []dto.RequestPart{
				{
					Text: request,
				},
			},
		},
	}}

	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	body := bytes.NewReader(bodyBytes)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cfg.AI.Host, body)
	if err != nil {
		return "", fmt.Errorf("%w: %w", model.ErrAPIFetch, err)
	}

	httpReq.Header.Add("content-type", "application/json")
	httpReq.Header.Add("X-goog-api-key", c.cfg.AI.APIKey)

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("%w: %w", model.ErrAPIFetch, err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%w: status code %d is not 200", model.ErrAPIFetch, resp.StatusCode)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response dto.Response
	if err := json.Unmarshal(respBytes, &response); err != nil {
		return "", nil
	}

	if len(response.Candidates) == 0 {
		return "", model.ErrAPIFetch
	}

	parts := response.Candidates[0].Content.Parts

	if len(parts) == 0 {
		return "", model.ErrAPIFetch
	}

	return parts[0].Text, nil
}
