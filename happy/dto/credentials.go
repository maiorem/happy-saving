package dto

type Credentials struct {
	Useremail string `form:"user_email"`
	Password  string `form:"password"`
}
