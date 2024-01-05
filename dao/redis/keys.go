package redis

// 存储redis中的key

// redis的key尽量使用命名空间的方式, 方便查询和拆分
// 如果当前常量是一个前缀, 也就是需要后续指定参数, 一般后面会带一个Prefix
const (
	KeyPrefix              = "bluebell:"
	KeyPostTimeZSet        = "post:time"      // 帖子及发帖时间
	KeyPostScoreZSet       = "post:score"     // 帖子及投票的分数
	KeyPostVotedZSetPrefix = "post:voted:"    // 用户及投票类型
	KeyCommunitySetPrefix  = "community:post" // 帖子id集合
)

func GetRedisKey(key string) string {
	return KeyPrefix + key
}
