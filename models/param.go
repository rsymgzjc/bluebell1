package models

// 定义请求参数的结构体
const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	UserName       string `json:"username" binding:"required"`
	PassWord       string `json:"password" binding:"required"`
	RepeatPassWord string `json:"repeatpassword" binding:"required,eqfield=PassWord"`
	Email          string `json:"email" binding:"required,email"`
}
type User struct {
	UserID   int64  `db:"user_id"`
	Password string `db:"password"`
	UserName string `db:"username"`
	Email    string `db:"email"`
}

type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	PostID    int64 `json:"post_id,string" binding:"required"`       //帖子id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票还是反对票
}

// 获取帖子列表query string 参数
type ParamPostList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
