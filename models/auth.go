package models

//Auth 权限结构体
type Auth struct {
	ID       int    `grom:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

//CheckAuth 校验权限
func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{
		Username: username,
		Password: password,
	}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}
