package dao

import (
	"douyin/global"
	"douyin/models"
	"gorm.io/gorm/clause"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

/**
  视频接口 业务层
*/

// UpdateVideoFavorite 更新用户点赞量
func UpdateVideoFavorite(favorite *models.Favorite) error {
	var favoriteCountChange int
	if favorite.IsDel == 1 {
		favoriteCountChange = -1
	} else {
		favoriteCountChange = 1
	}
	if err := global.MysqlEngine.Model(&models.Video{}).Where("video_id = ?", favorite.VideoId).
		Update("favorite_count", gorm.Expr("favorite_count + ?", favoriteCountChange)).Error; err != nil {
		return err
	}
	return nil
}

// UpdateVideoComment 更新用户评论量
func UpdateVideoComment(comment *models.Comment) error {
	var commentCountChange int
	if comment.IsDel == 1 {
		commentCountChange = -1
	} else {
		commentCountChange = 1
	}
	if err := global.MysqlEngine.Model(&models.Video{}).Where("video_id = ?", comment.VideoId).
		Update("comment_count", gorm.Expr("comment_count + ?", commentCountChange)).Error; err != nil {
		return err
	}
	return nil
}

// GetVideoList 获取视频列表, 用于feed接口
func GetVideoList() (videoList []models.Video, err error) {
	err = global.MysqlEngine.Order("create_date desc").Limit(30).Find(&videoList).Error
	if err != nil {
		global.Logger.Warn("GetVideoList Failed!", zap.Error(err))
	}
	return
}

// PublishVideo 将视频和封面链接保存于数据库
func PublishVideo(video models.Video) error {
	return global.MysqlEngine.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&video).Error
}

// GetPublishListByUserId 查询用户已发布视频
func GetPublishListByUserId(userId int64) []models.Video {
	var videos []models.Video = nil
	//查询用户已发布视频
	err := global.MysqlEngine.Where("user_id = ?", userId).Find(&videos).Error
	if err != nil {
		return nil
	}
	//查询视频作者信息
	for i := 0; i < len(videos); i++ {
		user, _ := GetUserByID(videos[i].UserId)
		videos[i].Author = *user
	}
	return videos
}
