package controller

import (
	"douyin/logic"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction favorite operation
func FavoriteAction(c *gin.Context) {
	param := new(models.ParamFavoriteAction)
	c.ShouldBindQuery(param)
	if err := logic.DoFavorite(param); err != nil {
		c.JSON(http.StatusInternalServerError, CodeMap[CodeInternalError])
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	param := new(models.ParamFavoriteList)
	c.ShouldBindQuery(param)
	videos := logic.GetFavoriteList(param)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}
