package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
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

// GetPostByIds 根据给定的一些帖子ID查询对应的所有帖子
func GetPostByIds(pids []string) (postList []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time
    from post
    where post_id in (?)
    order by FIND_IN_SET(post_id, ?)`

	query, args, err := sqlx.In(sqlStr, pids, strings.Join(pids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}

// GetPostList 返回所有的帖子
func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
    post_id, title, content, author_id, community_id, create_time
    from post
	ORDER BY create_time
	DESC 
	limit ?, ?`

	postList = make([]*models.Post, 0, 2)
	err = db.Select(&postList, sqlStr, (page-1)*size, size)
	return
}
