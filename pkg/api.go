package auth

type email string

type User struct {
	ID          int    `json:"id"`
	Email       email  `json:"email"`
	Password    string `json:"password"`
	Birthday    string `json:"birthday"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Token       string `json:"session_token"`
	IsAdmin     bool   `json:"is_admin"`
}

type Film struct {
	Name           string `json:"name"`
	Category       string `json:"category"`
	Type           string `json:"type"`
	AgeRestriction string `json:"age_restriction"`
	Year           int    `json:"year"`
	Lenght         int    `json:"lenght"`
	KeyWords       string `json:"key_words"`
	Description    string `json:"description"`
	Director       string `json:"director"`
}

type LoginData struct {
	Email    email  `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	User    User   `json:"user"`
}
