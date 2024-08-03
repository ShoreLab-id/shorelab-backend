package dto

type AuthResponse struct {
	Token string `json:"token"`
	Name  string `json:"name"`
}
