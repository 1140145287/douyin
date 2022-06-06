package dao

import (
	"douyin/global"
	"douyin/models"
	"gorm.io/gorm/clause"
)

// UpsertFavorite update the favorite record
// not exists : insert new record
// exists 	  : update the flag
func UpsertFavorite(favorite *models.Favorite) error {
	return global.MysqlEngine.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&favorite).Error
}

// DeleteFavorite delete the favorite record
func DeleteFavorite(favorite *models.Favorite) error {
	//soft delete, change the delete flag from 0 to 1
	err := global.MysqlEngine.
		Where("video_id = ? and user_id = ?", favorite.VideoId, favorite.UserId).
		Delete(&models.Favorite{}).Error
	return err
}

// GetFavoriteListByUserId function as it's name
func GetFavoriteListByUserId(userId int64) ([]models.Video, error) {
	var videos []models.Video = nil
	var err error = nil
	subQuery := global.MysqlEngine.Select("video_id").Where("user_id = ?", userId).Where("is_del", 0).Table("dy_favorite")
	if err = global.MysqlEngine.Where("video_id in (?)", subQuery).Find(&videos).Error; err != nil {
		return videos, err
	}
	for i := 0; i < len(videos); i++ {
		user, _ := GetUserByID(videos[i].UserId)
		videos[i].Author = *user
	}
	return videos, err
}
