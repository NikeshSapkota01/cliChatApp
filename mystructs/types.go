package mystructs

type UserInfo struct {
	Username string
	Email    string
	Password string
}

type User struct {
	Id       string
	Username string
	Password string
}

type Chat struct {
	User    string `json:"username"`
	Message string `json:"message"`
}
