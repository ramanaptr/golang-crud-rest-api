package entities

type User struct {
	ID       int16    `json:"id"`
	FullName string   `json:"fullname"`
	Email    string   `json:"email"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Time     BaseTime `json:"time"`
}
