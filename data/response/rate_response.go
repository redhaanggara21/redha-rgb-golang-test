package response

type RateResponse struct {
	Id     int `json:"id"`
	Rated  int `json:"rated"`
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
}
