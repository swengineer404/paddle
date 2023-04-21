package paddle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type BaseResult struct {
	Success bool `json:"success"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
}

type restClient struct {
	vendorID string
	apiKey   string
	baseURL  string
	c        *http.Client
}

func newRestClient(vendorID, apiKey string) *restClient {
	return &restClient{
		vendorID: vendorID,
		apiKey:   apiKey,
		baseURL:  "https://sandbox-vendors.paddle.com/api/2.0",
		c: &http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func (c *restClient) do(method, path string, dto, result any) error {
	var b bytes.Buffer
	if dto != nil {
		dtoBytes, err := json.Marshal(dto)
		if err != nil {
			return err
		}
		m := map[string]any{}
		if err := json.Unmarshal(dtoBytes, &m); err != nil {
			return err
		}
		m["vendor_id"] = c.vendorID
		m["vendor_auth_code"] = c.apiKey

		if err := json.NewEncoder(&b).Encode(m); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, &b)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "go-paddle-client")

	if dto != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.c.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	resb, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("status_code: %d | message: %s", res.StatusCode, resb)
	}

	var baseResult BaseResult
	if err := json.Unmarshal(resb, &baseResult); err != nil {
		return fmt.Errorf("failed to decode[%s]: %s", err, resb)
	}

	if !baseResult.Success {
		return errors.New(string(resb))
	}

	if err := json.Unmarshal(resb, result); err != nil {
		return fmt.Errorf("failed to decode[%s]: %s", err, resb)
	}

	return nil
}
