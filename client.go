package paddle

import "fmt"

type Client struct {
	rest         *restClient
	Product      *ProductService
	Subscription *SubscriptionService
}

func New(vendorID, apiKey string, sandbox bool) *Client {
	rc := newRestClient(vendorID, apiKey, sandbox)

	c := &Client{
		rest: rc,
	}

	c.Product = NewProductService(c)
	c.Subscription = NewSubscriptionService(c)

	return c
}

func (c *Client) Do(method, path string, dto, result any) error {
	if err := c.rest.do(method, path, dto, result); err != nil {
		return fmt.Errorf("paddle_api_error: %w", err)
	}

	return nil
}
