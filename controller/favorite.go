package controller

import (
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// FavoriteAction favorite operation
func FavoriteAction(c *gin.Context) {
	param := new(models.ParamFavoriteAction)
	if err := c.ShouldBindQuery(param); err != nil {
		global.Logger.Error("绑定错误", zap.Error(err))
		c.JSON(http.StatusOK, CodeMap[CodeInvalidParam])
		return
	}
	if err := logic.DoFavorite(param); err != nil {
		global.Logger.Error("Error in favorite operation", zap.Error(err))
		c.JSON(http.StatusOK, CodeMap[CodeInternalError])
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "点赞成功!",
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	param := new(models.ParamFavoriteList)
	if err := c.ShouldBindQuery(param); err != nil {
		global.Logger.Error("错误", zap.Error(err))
		c.JSON(http.StatusOK, CodeMap[CodeInvalidParam])
		return
	}
	videos, err := logic.GetFavoriteList(param)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "获取感兴趣列表失败！",
		})
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "获取感兴趣列表成功！",
		},
		VideoList: videos,
	})
}
