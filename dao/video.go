package dao

import (
	"douyin/global"
	"douyin/models"
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
