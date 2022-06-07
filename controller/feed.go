package controller

import (
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FeedResponse struct {
	Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//处理业务逻辑
	videoList, err := logic.GetVideoList()
	if err != nil {
		global.Logger.Error("videoList not exit", zap.Error(err))
		c.JSON(http.StatusBadRequest, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "获取视频列表失败"},
			NextTime: time.Now().Unix(),
		})
	}
	//TODO:如果用户存在Token，那么需要设置是否喜欢状态。也就是需要设置 is_favorite字段
	//token := c.Query("token")
	//if len(token) != 0 {
	//	//存在token且未过期，用户处在登陆状态
	//	if logic.ExistsKey(global.TokenPrefix + token){
	//
	//	}
	//
	//}
	//返回响应
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
