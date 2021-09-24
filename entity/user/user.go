package user

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name"`
	SocialID   string    `json:"social_id"` // CPF
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	ValidEmail bool      `json:"valid_email"` // true, valid email
	Address    string    `json:"address,omitempty"`
	Age        int       `json:"age,omitempty"`
	Sex        string    `json:"sex,omitempty"`
	Sign       string    `json:"sign,omitempty"` // Example: Aquarius
}
