package request

type CreateGiftRequest struct {
	Name        string `validate:"required,min=1,max=25" json:"name" binding:"required"`
	Description string `validate:"required,min=1,max=200" json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Image       string `json:"image"`
}
