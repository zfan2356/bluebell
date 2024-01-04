package mysql

import (
	"bluebell/models"
)

// CreatePost 创建帖子并保存在mysql之中
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
    post_id, title, content, author_id, community_id) 
    values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据帖子ID查询数据
func GetPostById(pid int64) (data *models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time
    from post
    where post_id = ?`

	data = new(models.Post)
	err = db.Get(data, sqlStr, pid)
	return
}

// GetPostList 返回所有的帖子
func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time
    from post limit ?, ?`

	postList = make([]*models.Post, 0, 2)
	err = db.Select(&postList, sqlStr, (page-1)*size, size)
	return
}
