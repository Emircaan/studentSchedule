package model

type LoginRequest struct {
	Eposta string `json:"email" form:"email" validate:"required,email"`
	Sifre  string `json:"password" form:"password" validate:"required"`
}
