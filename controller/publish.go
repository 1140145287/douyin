package controller

import (
	"douyin/logic"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VideoListResponse struct {
	Response
	VideoList []models.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	param := new(models.ParamPublishAction)
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CodeMap[CodeInvalidParam])
		return
	}
	//调用上传服务
	err = logic.DoPublish(param, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CodeMap[CodeInternalError])
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "视频上传成功"})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	param := new(models.ParamPublishList)
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CodeMap[CodeInvalidParam])
		return
	}
	//调用查询服务
	videos := logic.GetPublishList(param)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
