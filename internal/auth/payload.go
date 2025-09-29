package auth

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
	Position string `json:"position"`
	Filials  string `json:"filials"`
	Brand    string `json:"brand" validate:"required"`
}

type RegisterResponse struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	Token       string `json:"token,omitempty"`
	MainCabinet string `json:"main_cabinet"`
	UserID      uint   `json:"user_id"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Status      string `json:"status"`
	Token       string `json:"token"`
	MainCabinet string `json:"main_cabinet"`
	Role        string `json:"role"`
	UserID      uint   `json:"user_id"`
	Email       string `json:"email"`
}
