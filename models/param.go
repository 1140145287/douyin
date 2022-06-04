package models

type ParamRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

//登录参数
type ParamLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

//获取个人信息参数
type ParamInfo struct {
	Uid   int64  `form:"user_id" json:"user_id" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
}
