package dto

type (
	RegisterRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Username        string `json:"username" validate:"required,min=3,max=32"`
		Password        string `json:"password" validate:"required,min=8,max=32"`
		PasswordConfirm string `json:"password_confirm" validate:"required,min=8,max=32,eqfield=Password"`
	}

	RegisterResponse struct {
		UserId string `json:"user_id"`
	}
)

type (
	LoginRequest struct {
		Username string `json:"username" validate:"required,min=3,max=32"`
		Password string `json:"password" validate:"required,min=8,max=32"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
