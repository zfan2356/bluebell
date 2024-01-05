package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// 测试时选择自己重新创建一个路由, 然后进行测试
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	// 首先向这个路由发送POST请求, 这个时候测试的是创建帖子, 所以要写一段创建帖子的json数据作为body
	body := `{
        "title": "test",
        "content": "just a test",
        "community_id": "1"
    }`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 判断状态码
	assert.Equal(t, 200, w.Code)
	// 还要判断响应的内容是不是按照预期, 返回的是需要登陆这个错误
	assert.Contains(t, w.Body.String(), "需要登陆")
	// 或者将body反序列化到ResponseData中, 然后判断是否与预期一致
}
