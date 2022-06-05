package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"gorm.io/gorm"
)

// DoFavorite 点赞service操作
func DoFavorite(param *models.ParamFavoriteAction) error {
	favorite := &models.Favorite{
		VideoId: param.VideoId,
		UserId:  GetUserByToken(param.Token).Id,
	}
	err := global.MysqlEngine.Transaction(func(tx *gorm.DB) error {
		var err error
		if param.ActionType == 1 {
			favorite.IsDel = 0
			err = dao.UpsertFavorite(favorite)
		} else {
			favorite.IsDel = 1
			err = dao.DeleteFavorite(favorite)
		}
		err = dao.UpdateVideoFavorite(favorite)
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

// GetFavoriteList 获取感兴趣列表
func GetFavoriteList(param *models.ParamFavoriteList) []models.Video {
	return dao.GetFavoriteListByUserId(param.UserId)
}
