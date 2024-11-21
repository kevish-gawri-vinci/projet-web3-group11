package request

type UserRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserRoleRequest struct {
	IsAdmin  bool   `json:"isadmin"`
	Username string `json:"username"`
}
