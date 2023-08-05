package request

type CreateRateRequest struct {
	Rated  int `validate:"required,min=1,max=25" json:"rated"`
	UserID int `validate:"required,min=1,max=200" json:"user_id"`
	GiftID int `json:"gift_id"`
}
