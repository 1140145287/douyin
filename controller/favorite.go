package controller

import (
	"douyin/global"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

// FavoriteAction favorite operation
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		var favoriteCountChange int8
		videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 8)
		if actionType == 1 {
			favoriteCountChange = 1
		} else {
			favoriteCountChange = -1
		}
		favorite := Favorite{
			UserId:  user.Id,
			VideoId: videoId,
			IsDel:   0,
		}
		global.DBEngine.Transaction(func(tx *gorm.DB) error {

			//insert or update or remove favorite record
			if actionType == 2 {
				//soft delete, change the delete flag from 0 to 1
				global.DBEngine.
					Where("video_id = ? and user_id = ?", favorite.VideoId, favorite.UserId).
					Delete(&Favorite{})
			} else {
				//insert new record or update the record
				global.DBEngine.Clauses(clause.OnConflict{
					UpdateAll: true,
				}).Create(&favorite)
			}

			//update video favorite_count plus 1
			if err := global.DBEngine.Model(&Video{}).Where("video_id", videoId).
				Update("favorite_count", gorm.Expr("favorite_count + ?", favoriteCountChange)).Error; err != nil {
				return err
			}
			return nil
		})
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if user, exist := usersLoginInfo[token]; exist {
		var videos []Video
		subQuery := global.DBEngine.Select("video_id").Where("user_id = ?", user.Id).Where("is_del", 0).Table("dy_favorite")
		if err := global.DBEngine.Where("video_id in (?)", subQuery).Find(&videos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "Error"})
		} else {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 0,
				},
				VideoList: videos,
			})
		}

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
