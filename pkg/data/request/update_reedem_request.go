package request

type UpdateReedemRequest struct {
	Id     int `gorm:"type:int;primary_key"`
	UserID int `json:"user_id"`
	GiftID int `json:"gift_id"`
	Amount int `json:"amount"`
}
