package dto

type Credentials struct {
	Username  string `form:"username"`
	Useremail string `form:"useremail"`
	Password  string `form:"password"`
}
