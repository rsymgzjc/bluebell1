package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

/*
	投票的几种情况

direction=1时，有两种情况：

	1.之前没有投过票，现在投赞成票         差值的绝对值 1 +432
	2.之前投反对票，现在投赞成票           差值的绝对值 2  +432*2

direction=0时，有两种情况：

	1.之前投过赞成票，现在要取消投票       差值的绝对值 1 -432
	2.之前投过反对票，现在取消投票         差值的绝对值 1 +432

direction=-1时，有两种情况：

	1.之前没有投过票，现在投反对票         差值的绝对值 1  -432
	2.之前投赞成票，现在投反对票           差值的绝对值 2  -432*2

投票的限制：
每个帖子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了

	1.到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
	2.到期之后删除那个 KeyPostVotedZSetPF
*/

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(PostID int64) (err error) {
	pipeline := rdb.TxPipeline()
	//帖子时间
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: PostID,
	})
	//帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: PostID,
	})
	_, err = pipeline.Exec()
	return
}
func VoteForPost(userID string, postID string, value float64) error {
	//判断投票限制
	//取贴子发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	//更新帖子分数
	//先查之前的投票纪录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedZsetPrefix+postID), userID).Val()
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	//放到事务中
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), diff*op*scorePerVote, postID)
	//记录用户为该帖子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZsetPrefix+postID), userID).Result()
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZsetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
