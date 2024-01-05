package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() {
	dbCfg := settings.MySQLConfig{
		Host:         "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "123456",
		Dbname:       "web",
		MaxOpenConns: 10,
		MaxIdleConns: 10,
	}
	err := InitMySQL(&dbCfg)
	if err != nil {
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		ID:          114514,
		AuthorID:    2356,
		CommunityID: 1,
		Title:       "test",
		Content:     "just a test",
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost insert into mysql failed, err: %#v\n", err)
	}
	t.Logf("CreatePost insert into mysql success\n")
}
