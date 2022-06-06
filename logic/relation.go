package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"go.uber.org/zap"
)

// GetFollowListLogic logic层获取关注用户列表
func GetFollowListLogic(uid int64) ([]models.User, error) {
	return dao.GetFollowList(uid)
}

// GetFollowerListLogic logic层获取粉丝列表
func GetFollowerListLogic(uid int64) ([]models.User, error) {
	return dao.GetFollowerList(uid)
}

func RelationActionLogic(uid int64, toUid int64, actionType int8) error {
	// 判断关注与被关注用户是否存在
	err1 := dao.CheckUserExistById(uid)
	if err1 != nil {
		global.Logger.Error("Error in setting relation action", zap.Error(err1))
		return err1
	}
	err2 := dao.CheckUserExistById(toUid)
	if err1 != nil {
		global.Logger.Error("Error in setting relation action", zap.Error(err2))
		return err2
	}

	// 执行关注/取关操作
	switch actionType {
	case 1:
		if err := dao.DoActionFollow(uid, toUid); err != nil {
			return err
		}
	case 2:
		if err := dao.DoActionUnfollow(uid, toUid); err != nil {
			return err
		}
	default:
		global.Logger.Error("Error: invalid action type")
	}
	return nil
}
