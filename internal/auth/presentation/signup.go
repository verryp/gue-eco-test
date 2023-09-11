package presentation

type (
	SignUpRequest struct {
		Name     string `json:"name" valid:"required"`
		Email    string `json:"email" valid:"required,email"`
		Password string `json:"password" valid:"required"`
	}
)
