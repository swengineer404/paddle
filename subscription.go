package paddle

type GetSubscriptionResult struct {
	Success  bool `json:"success"`
	Response []struct {
		State string `json:"state"`
	} `json:"response"`
}

type GetSubscriptionParams struct {
	SubscriptionID string `json:"subscription_id"`
}

type CancelSubscriptionResult struct {
	Success bool `json:"success"`
}

type CancelSubscriptionParams struct {
	SubscriptionID string `json:"subscription_id"`
}

type SubscriptionService struct {
	client *Client
}

func NewSubscriptionService(client *Client) *SubscriptionService {
	return &SubscriptionService{
		client: client,
	}
}

func (s *SubscriptionService) Get(subID string) (*GetSubscriptionResult, error) {
	var result GetSubscriptionResult
	return &result, s.client.Do("POST", "/subscription/users", &GetSubscriptionParams{SubscriptionID: subID}, &result)
}

func (s *SubscriptionService) Cancel(subID string) (*CancelSubscriptionResult, error) {
	var result CancelSubscriptionResult
	return &result, s.client.Do("POST", "/subscription/users_cancel", &CancelSubscriptionParams{SubscriptionID: subID}, &result)
}

//func (s *SubscriptionService) Cancel(subID string) error {
//
//}
