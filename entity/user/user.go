package user

type User struct {
	FistName   string `json:"fist_name"`
	LastName   string `json:"last_name"`
	ID         string `json:"id"`
	SocialID   string `json:"social_id"` // CPF
	Age        int    `json:"age"`
	Address    string `json:"address"`
	Sex        string `json:"sex"`
	Sign       string `json:"sign"` // Example: Aquarius
	Email      string `json:"email"`
	ValidEmail bool   `json:"valid_email"` // true, valid email
	Phone      string `json:"phone"`
}
