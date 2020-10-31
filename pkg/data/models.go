package data

type ResponseBasic struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type Users struct {
	Id        string `form:"id" json:"id"`
	Username string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
}