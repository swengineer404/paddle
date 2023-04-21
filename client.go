package paddle

import "fmt"

type Client struct {
	rest    *restClient
	Product *ProductService
}

func New(vendorID, apiKey string) *Client {
	rc := newRestClient(vendorID, apiKey)

	c := &Client{
		rest: rc,
	}

	c.Product = NewProductService(c)

	return c
}

func (c *Client) Do(method, path string, dto, result any) error {
	if err := c.rest.do(method, path, dto, result); err != nil {
		return fmt.Errorf("paddle_api_error: %w", err)
	}

	return nil
}
