package models

type ParamRegister struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

// ParamLogin 登录参数
type ParamLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=12"`
}

// ParamInfo 获取个人信息参数
type ParamInfo struct {
	Uid   int64  `form:"user_id" json:"user_id" binding:"required"`
	Token string `form:"token" json:"token" binding:"required"`
}

// ParamFavoriteAction 用户点赞请求
type ParamFavoriteAction struct {
	Token      string `form:"token" json:"token" binding:"required"`
	VideoId    int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType int8   `form:"action_type" json:"action_type" binding:"required"`
}

// ParamFavoriteList  用户获取点赞请求列表
type ParamFavoriteList struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserId int64  `form:"user_id" json:"user_id" binding:"required"`
}

// ParamCommentAction 用户获取评论请求列表
type ParamCommentAction struct {
	Token       string `form:"token" json:"token" binding:"required"`
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int8   `form:"action_type" json:"action_type" binding:"required"`
	CommentText string `form:"comment_text" json:"comment_text"`
	CommentId   int64  `form:"comment_id" json:"comment_id" json:"comment_id"`
}

// ParamCommentList  用户获取评论请求列表
type ParamCommentList struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoId int64  `form:"video_id" json:"video_id" binding:"required"`
}
