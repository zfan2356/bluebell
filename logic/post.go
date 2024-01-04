package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

// CreatePost 创建帖子业务逻辑
func CreatePost(p *models.Post) (err error) {
	// 1. 生成ID
	p.ID = snowflake.GenID()
	// 2. 保存在数据库
	return mysql.CreatePost(p)
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

// GetPostList 返回帖子的列表
func GetPostList(page, size int64) (postdetails []*models.ApiPostDetail, err error) {
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
