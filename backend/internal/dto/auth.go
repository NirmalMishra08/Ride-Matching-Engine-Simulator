package dto

type AuthRequest struct {
    Provider string `json:"provider"`
    Email    string `json:"email"`
    Phone    string `json:"phone"`
    FullName string `json:"full_name"`
    Password string `json:"password"`
}