package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

// getIDsFromKey 根据指定Key在redis中查询范围内的id
func getIDsFromKey(key string, page, size int64) ([]string, error) {
	st := (page - 1) * size
	ed := st + size - 1
	return rdb.ZRevRange(key, st, ed).Result()
}

// CreatePost 创建帖子
func CreatePost(pid, cid string) (err error) {
	pipeline := rdb.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(GetRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: pid,
	})
	// 帖子分数
	pipeline.ZAdd(GetRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  0,
		Member: pid,
	})
	// 添加到对应的社区
	pipeline.SAdd(GetRedisKey(KeyCommunitySetPrefix+cid), pid)
	_, err = pipeline.Exec()
	return
}

// GetPostIDsInOrder 获取排序的帖子的id
func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := GetRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = GetRedisKey(KeyPostScoreZSet)
	}
	return getIDsFromKey(key, p.Page, p.Size)
}

// GetPostVoteData 获取帖子投票的数据(多少赞成)
func GetPostVoteData(ids []string) ([]int64, error) {
	//for _, id := range ids {
	//    v1 := rdb.ZCount(GetRedisKey(KeyPostVotedZSetPrefix+id), "1", "1").Val()
	//    data = append(data, v1)
	//}
	// 这里可以使用pipline, 减少rtt
	pipline := rdb.Pipeline()
	for _, id := range ids {
		pipline.ZCount(GetRedisKey(KeyPostVotedZSetPrefix+id), "1", "1")
	}
	cmders, err := pipline.Exec()
	if err != nil {
		return nil, err
	}

	data := make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}

// GetCommunityPostIDsInOrder 获取排序的指定社区的帖子ID
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	// 使用zinterstore, 把分区帖子的set以及帖子分数的zset生成一个新的zset
	// 针对新的zset, 按照之前的逻辑取数据

	// 利用缓存key减少 zintersore执行的次数
	cid := strconv.Itoa(int(p.CommunityID))
	okey := GetRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		okey = GetRedisKey(KeyPostScoreZSet)
	}
	key := p.Order + cid
	if rdb.Exists(key).Val() < 1 {
		// 不存在缓存的key, 需要计算
		pipline := rdb.Pipeline()
		pipline.ZInterStore(key, redis.ZStore{
			Weights:   nil,
			Aggregate: "MAX",
		}, GetRedisKey(KeyCommunitySetPrefix+cid), okey)
		pipline.Expire(key, 60*time.Second) // 设置缓存超时时间
		_, err := pipline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFromKey(key, p.Page, p.Size)
}
