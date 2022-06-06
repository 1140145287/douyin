package dao

import (
	"douyin/global"
	"douyin/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
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

// DoActionFollow 添加关注关系
func DoActionFollow(uid int64, toUid int64) error {
	// 如果已经关注，则不进行任何操作
	var relation []models.Relation
	global.MysqlEngine.Where("follower_id=? and following_id=? and is_del=?", toUid, uid, 0).Find(&relation)
	if len(relation) > 0 {
		global.Logger.Warn("already followed")
		return nil
	}

	// 未关注，则插入数据，并更新follow_count和follower_count
	newRelation := models.Relation{FollowerId: toUid, FollowingId: uid, IsDel: 0}
	if err := global.MysqlEngine.Create(&newRelation).Error; err != nil {
		global.Logger.Error("Insert New Relation Failed!", zap.Error(err))
		return err
	}

	// 博主粉丝数量+1
	if err := global.MysqlEngine.Model(&models.User{}).Where("user_id=?", toUid).Update("follower_count", gorm.Expr("follower_count+ ?", 1)).Error; err != nil {
		global.Logger.Error("update follower_count Failed!", zap.Error(err))
		return err
	}

	// 粉丝关注数量+1
	if err := global.MysqlEngine.Model(&models.User{}).Where("user_id=?", uid).Update("follow_count", gorm.Expr("follow_count+ ?", 1)).Error; err != nil {
		global.Logger.Error("update follow_count Failed!", zap.Error(err))
		return err
	}

	return nil
}

// DoActionUnfollow 删除关注关系
func DoActionUnfollow(uid int64, toUid int64) error {
	// 如果已经取关，则不再进行任何操作
	var relation []models.Relation
	global.MysqlEngine.Where("follower_id=? and following_id=? and is_del=?", toUid, uid, 0).Find(&relation)
	if len(relation) == 0 {
		global.Logger.Warn("already unfollow")
		return nil
	}

	// 仍在关注，则更新条目
	if err := global.MysqlEngine.Model(&models.Relation{}).Where("follower_id=? and following_id=? and is_del=?", toUid, uid, 0).Update("is_del", 1).Error; err != nil {
		global.Logger.Error("update unfollow info Failed!", zap.Error(err))
		return err
	}

	// 博主粉丝数量-1
	if err := global.MysqlEngine.Model(&models.User{}).Where("user_id=?", toUid).Update("follower_count", gorm.Expr("follower_count- ?", 1)).Error; err != nil {
		global.Logger.Error("update follower_count Failed!", zap.Error(err))
		return err
	}

	// 粉丝关注数量-1
	if err := global.MysqlEngine.Model(&models.User{}).Where("user_id=?", uid).Update("follow_count", gorm.Expr("follow_count- ?", 1)).Error; err != nil {
		global.Logger.Error("update follow_count Failed!", zap.Error(err))
		return err
	}
	return nil
}
