package userRepo

type UserData struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type User struct {
	UserData
	Email string `json:"email"`
}

type UserCreate struct {
	Name  string `json:"name" db:"name"`
	Email string `json:"email"`
}
