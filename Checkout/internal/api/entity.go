package api

type ApplyRewardRequest struct {
	OrderID string `json:"order_id,omitempty"`
	Rewards int    `json:"rewards,omitempty"`
}

type PayRequest struct {
	OrderID string `json:"order_id,omitempty"`
}
