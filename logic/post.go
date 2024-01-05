package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
	"strconv"
)

// CreatePost 创建帖子业务逻辑
func CreatePost(p *models.Post) (err error) {
	// 1. 生成ID
	p.ID = snowflake.GenID()
	// 2. 保存在数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return
	}
	return redis.CreatePost(strconv.Itoa(int(p.ID)), strconv.Itoa(int(p.CommunityID)))
}

// GetPostById 根据帖子id获取数据
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询Post信息之后, 组合成ApiPostDetail, 然后再返回
	var post *models.Post
	post, err = mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err),
		)
		return
	}
	// 根据作者ID查询作者信息
	user, err := mysql.GetUserByUserID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed",
			zap.Int64("author_id", post.AuthorID),
			zap.Error(err),
		)
		return
	}
	// 根据社区ID查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err),
		)
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user.UserName,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList1 返回帖子的列表, 无其他功能
func GetPostList1(page, size int64) (postdetails []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed",
			zap.Error(err),
		)
		return
	}
	postdetails = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByUserID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err),
			)
			continue
		}
		// 根据社区ID查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err),
			)
			continue
		}
		data := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			Post:            post,
			CommunityDetail: community,
		}
		postdetails = append(postdetails, data)
	}
	return
}

// GetPostListWithOrder 返回帖子的列表并按照分数或者时间排序
func GetPostListWithOrder(p *models.ParamPostList) (postdetails []*models.ApiPostDetail, err error) {
	pids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder(p) failed", zap.Error(err))
		return
	}

	if len(pids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) empty")
		return
	}
	posts, err := mysql.GetPostByIds(pids)
	if err != nil {
		return
	}
	voteData, err := redis.GetPostVoteData(pids)
	if err != nil {
		return
	}

	postdetails = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByUserID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err),
			)
			continue
		}
		// 根据社区ID查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err),
			)
			continue
		}
		data := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		postdetails = append(postdetails, data)
	}
	return
}

// GetCommunityPostList 获取指定社区的所有的帖子, 并按照分数或者时间排序
func GetCommunityPostList(p *models.ParamPostList) (postdetails []*models.ApiPostDetail, err error) {
	pids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder(p) failed", zap.Error(err))
		return
	}

	if len(pids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) empty")
		return
	}
	posts, err := mysql.GetPostByIds(pids)
	if err != nil {
		return
	}
	voteData, err := redis.GetPostVoteData(pids)
	if err != nil {
		return
	}

	postdetails = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		// 根据作者ID查询作者信息
		user, err := mysql.GetUserByUserID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByUserID(post.AuthorID) failed",
				zap.Int64("author_id", post.AuthorID),
				zap.Error(err),
			)
			continue
		}
		// 根据社区ID查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err),
			)
			continue
		}
		data := &models.ApiPostDetail{
			AuthorName:      user.UserName,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		postdetails = append(postdetails, data)
	}
	return
}

// GetPostList 获取指定社区的所有的帖子, 并按照分数或者时间排序, 没有指定就获取所有的社区
func GetPostList(p *models.ParamPostList) (postdetails []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		return GetPostListWithOrder(p)
	} else {
		return GetCommunityPostList(p)
	}
}
