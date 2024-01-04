package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList  查找到所有的community, 并且返回
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 查找到对应id的社区
func GetCommunityDetail(communityID int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(communityID)
}
