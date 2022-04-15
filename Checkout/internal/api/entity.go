package api

type ApplyRewardRequest struct {
	OrderID int `json:"order_id,omitempty"`
	Rewards int `json:"rewards,omitempty"`
}

type PayRequest struct {
	OrderID int `json:"order_id,omitempty"`
}
