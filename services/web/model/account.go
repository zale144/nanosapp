package model

type Account struct {
	Username      string  `json:"username" form:"username" query:"username"`
	Password       string  `json:"password" form:"password" query:"password"`
	ConfirmPassword string `json:"confirmPassword" form:"confirmPassword" query:"confirmPassword"`
}
