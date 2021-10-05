package user

import "github.com/google/uuid"

// User contains all information about client
// This information can be used to sign in
type User struct {
	UserID     uuid.UUID `json:"user_id"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name"`
	SocialID   string    `json:"social_id"` // CPF
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	ValidEmail bool      `json:"valid_email"` // true, valid email
	City       string    `json:"city"`
	Country    string    `json:"country"`
	Address    string    `json:"address,omitempty"`
	Zipcode    string    `json:"zipcode"`
	Age        int       `json:"age,omitempty"`
	Sex        string    `json:"sex,omitempty"`
	Sign       string    `json:"sign,omitempty"` // Example: Aquarius
}

func NewUser() *User {
	user := User{
		UserID: uuid.New(),
	}
	return &user
}
