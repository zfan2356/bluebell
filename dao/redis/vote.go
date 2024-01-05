package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票的分数
)

var (
	ErrorVoteTimeExpire = errors.New("超出投票时间限制")
	ErrorVoteRepeated   = errors.New("不允许重复投票")
)

// VoteForPost 投票
func VoteForPost(userID, postID string, value float64) (err error) {
	// 1. 判断投票限制
	// 先获取帖子的发布时间
	postTime := rdb.ZScore(GetRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrorVoteTimeExpire
	}
	// 2. 更新帖子的分数
	// 先查当前用户给当前帖子投票的记录
	o := rdb.ZScore(GetRedisKey(KeyPostVotedZSetPrefix+postID), userID).Val()
	diff := math.Abs(o - value)

	pipline := rdb.TxPipeline()
	if value > o {
		pipline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), diff*scorePerVote, postID)
	} else if value < o {
		pipline.ZIncrBy(GetRedisKey(KeyPostScoreZSet), -diff*scorePerVote, postID)
	} else {
		// 和之前数据一致, 返回错误
		return ErrorVoteRepeated
	}
	// 3. 记录用户为该帖子投票的数据
	if value == 0 {
		pipline.ZRem(GetRedisKey(KeyPostVotedZSetPrefix+postID), userID)
	} else {
		pipline.ZAdd(GetRedisKey(KeyPostVotedZSetPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipline.Exec()
	return
}
