package models

// Token - allows you to wrap the generated token
type Token struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	RolID  uint   `json:"role_id"`
	Token  string `json:"token"`
}
