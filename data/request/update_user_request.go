package request

type UpdateUserRequest struct {
	Id       int    `validate:"required"`
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `validate:"required,min=1,max=200,email" json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
