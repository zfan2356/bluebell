package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"strconv"
)

// VoteForPost 投票相关业务处理
func VoteForPost(uid int64, p *models.ParamVoteData) (err error) {
	// 投票限制, 每个帖子发布一星期之内可以允许投票, 之后就不允许投票
	// 到期之后将redis中保存的赞成票数以及反对票数存储到mysql的表中
	// 到期之后删除 keyKeyPostVotedZSetPrefix
	return redis.VoteForPost(strconv.Itoa(int(uid)), strconv.Itoa(int(p.PostID)), float64(p.Direction))
}
