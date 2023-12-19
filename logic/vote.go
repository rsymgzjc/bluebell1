package logic

import (
	"bluebell1/dao/redis"
	"bluebell1/models"
	"strconv"

	"go.uber.org/zap"
)

//投票功能
//1.用户投票的数据
//

// 使用简化版投票分数
//投一票加432分，86400/200 ->200张赞成票可以给你的帖子续一天

/*
	投票的几种情况

direction=1时，有两种情况：

	1.之前没有投过票，现在投赞成票
	2.之前投反对票，现在投赞成票

direction=0时，有两种情况：

	1.之前投过赞成票，现在要取消投票
	2.之前投过反对票，现在取消投票

direction=-1时，有两种情况：

	1.之前没有投过票，现在投反对票
	2.之前投赞成票，现在投反对票

投票的限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了

	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除那个 KeyPostVotedZSetPF
*/
func VoteForPost(userID int64, p *models.ParamVoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("userID", userID), zap.Int64("postID", p.PostID), zap.Int8("direction", p.Direction))
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
