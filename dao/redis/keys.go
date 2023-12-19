package redis

const (
	KeyPrefix              = "bluebell1:"
	KeyPostTimeZSet        = "post:time"   //zset 帖子发帖时间为分数
	KeyPostScoreZSet       = "post:score"  //zset 帖子及投票分数
	KeyPostVotedZsetPrefix = "post:voted:" //zset 记录用户及投票类型 参数是post id
	KeyCommunitySetPF      = "community:"  //set 保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
