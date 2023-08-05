package response

type ReedemResponse struct {
	Id     int `json:"id"`
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
	Amount int `json:"amount"`
}
