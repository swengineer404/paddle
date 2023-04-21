package paddle

type GeneratePayLinkParams struct {
	ProductID          int      `json:"product_id,omitempty"`
	Title              string   `json:"title,omitempty"`
	WebhookURL         string   `json:"webhook_url,omitempty"`
	Prices             []string `json:"prices,omitempty"`
	RecurringPrices    []string `json:"recurring_prices,omitempty"`
	TrialDays          int      `json:"trial_days"`
	CustomMessage      string   `json:"custom_message,omitempty"`
	CouponCode         string   `json:"coupon_code,omitempty"`
	Discountable       int      `json:"discountable"`
	ImageURL           string   `json:"image_url,omitempty"`
	ReturnURL          string   `json:"return_url,omitempty"`
	QuantityVariable   int      `json:"quantity_variable"`
	Quantity           int      `json:"quantity"`
	Expires            string   `json:"expires,omitempty"`
	MarketingConsent   int      `json:"marketing_consent"`
	CustomerEmail      string   `json:"customer_email,omitempty"`
	CustomerCountry    string   `json:"customer_country,omitempty"`
	CustomerPostalCode string   `json:"customer_postcode,omitempty"`
	IsRecoverable      int      `json:"is_recoverable"`
	Passthrough        string   `json:"passthrough,omitempty"`
	VatNumber          string   `json:"vat_number,omitempty"`
	VatCompanyName     string   `json:"vat_company_name,omitempty"`
	VatCountry         string   `json:"vat_country,omitempty"`
	VatState           string   `json:"vat_state,omitempty"`
	VatCity            string   `json:"vat_city,omitempty"`
	VatStreet          string   `json:"vat_street,omitempty"`
	VatPostalCode      string   `json:"vat_postcode,omitempty"`
}

type GeneratePayLinkResult struct {
	Success  bool `json:"success"`
	Response struct {
		URL string `json:"url"`
	} `json:"response"`
}

type ProductService struct {
	client *Client
}

func NewProductService(client *Client) *ProductService {
	return &ProductService{
		client: client,
	}
}

func (c *ProductService) GeneratePayLink(params *GeneratePayLinkParams) (*GeneratePayLinkResult, error) {
	var result GeneratePayLinkResult
	return &result, c.client.Do("POST", "/product/generate_pay_link", params, &result)
}
