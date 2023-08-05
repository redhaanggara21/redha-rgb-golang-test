package request

type CreateReedemRequest struct {
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
	Amount int `json:"amount"`
}
