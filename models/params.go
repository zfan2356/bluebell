package models

// 定义请求参数的结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登陆请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票请求参数
type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"` // 帖子ID
	Direction int8  `json:"direction" binding:"oneof=1 0 -1"`  // 赞成or反对(1 or -1)
}

// 获取query string参数, 因为不需要json传输, 而是query传输, 所以只需要打上form的tag

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	CommunityID int64  `form:"community_id"`
	Order       string `form:"order"`
}
