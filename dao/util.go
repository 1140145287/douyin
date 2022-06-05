package dao

import (
	"douyin/global"
	"douyin/models"
	"go.uber.org/zap"
)

// GetFollowList 获取关注用户列表
func GetFollowList(uid int64) ([]models.User, error) {
	var followList []models.User
	sqlStr := "select dy_user.user_id, dy_user.name, dy_user.follower_count, dy_user.is_follow, null, null, dy_user.follow_count FROM dy_user LEFT JOIN dy_relation ON dy_user.user_id=dy_relation.follower_id WHERE dy_relation.following_id=? AND dy_relation.is_del=0"
	err := global.MysqlEngine.Raw(sqlStr, uid).Scan(&followList).Error

	if err != nil {
		global.Logger.Warn("GetFollowList Failed!", zap.Int64("uid:", uid), zap.Error(err))
		return nil, err
	}

	return followList, err
}

// GetFollowerList 获取关注者用户列表
func GetFollowerList(uid int64) ([]models.User, error) {
	var followerList []models.User
	sqlStr := "select dy_user.user_id, dy_user.name, dy_user.follower_count, dy_user.is_follow, null, null, dy_user.follow_count FROM dy_user LEFT JOIN dy_relation ON dy_user.user_id=dy_relation.following_id WHERE dy_relation.follower_id=? AND dy_relation.is_del=0"
	err := global.MysqlEngine.Raw(sqlStr, uid).Scan(&followerList).Error

	if err != nil {
		global.Logger.Warn("GetFollowList Failed!", zap.Int64("uid:", uid), zap.Error(err))
		return nil, err
	}

	return followerList, err
}
