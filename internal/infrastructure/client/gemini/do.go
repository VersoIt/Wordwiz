package gemini

import (
	"bytes"
	"context"
	"encoding/json"
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
		return "", err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", model.ErrAPIFetch
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
