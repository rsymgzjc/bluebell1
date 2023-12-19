package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"community_id":1,
		"title":"test",
		"content":"just a test"
	}`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	//方法1  判断响应内容是不是包含指定的字符串
	assert.Contains(t, w.Body.String(), "需要登录")
}
