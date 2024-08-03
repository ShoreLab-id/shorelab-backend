package dto

type UserCreateRequest struct {
	Name     string
	Email    string
	Password string
}

type AuthRequest struct {
	Email    string
	Password string
}
