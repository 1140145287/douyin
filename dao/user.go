package dao

import (
	"douyin/global"
	"douyin/models"

	"go.uber.org/zap"
)

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) error {
	user := []models.User{}
	global.MysqlEngine.Where("name=?", username).Find(&user)
	if len(user) > 0 {
		return ErrorUserExit
	}
	return nil
}

// InsertUser 插入新用户
func InsertUser(u *models.User) error {
	//执行sql语句入库
	u.Salt = global.Secrete
	if err := global.MysqlEngine.Create(u).Error; err != nil {
		return err
	}
	// if err := global.MysqlEngine.Where("name", u.Name).Find(&u).Error; err != nil {
	// 	return err
	// }
	return nil
}

// Login 登录
func Login(p *models.User) error {
	user := []models.User{}
	global.MysqlEngine.Where("name=?", p.Name).Find(&user)
	if len(user) <= 0 {
		return ErrorUserNotExit
	}
	if global.Md5(p.PassWord) != user[0].PassWord {
		return ErrorInvalidPassword
	}
	*p = user[0]
	return nil
}

// GetUserByID 根据作者的ID号得到作者信息
func GetUserByID(uid int64) (*models.User, error) {
	user := new(models.User)
	err := global.MysqlEngine.Where("user_id = ?", uid).Find(user).Error
	if err != nil {
		global.Logger.Warn("GetUserByID Failed!", zap.Int64("uid:", uid), zap.Error(err))
		err = ErrorInvalidID
	}
	return user, err
}

// CheckUserExistById 根据id判断用户是否存在
func CheckUserExistById(uid int64) error {
	var user []models.User
	global.MysqlEngine.Where("user_id=?", uid).Find(&user)
	if len(user) == 0 {
		return ErrorUserNotExit
	}
	return nil
}
