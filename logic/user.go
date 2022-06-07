package logic

import (
	"context"
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"douyin/pkg/jwt"
	"encoding/json"
	"time"
)

// Register 存放注册逻辑的代码
func Register(p *models.ParamRegister) (*models.User, error) {
	// 1、判断用户是否存在
	if err := dao.CheckUserExist(p.Username); err != nil {
		return nil, err
	}
	// 2、密码加密
	password := global.Md5(p.Password)
	// 3、构造一个用户实例
	user := &models.User{
		Name:     p.Username,
		PassWord: password,
	}
	// 4、保存进数据库
	if err := dao.InsertUser(user); err != nil {
		return nil, err
	}
	// 5、有要求说返回token,所以我们要返回token
	token, err := jwt.GetToken(user.Id, user.Name)
	user.Token = token
	// 6、将token加入redis中，过期时间是24小时, 键是token, 值是用户对象
	userJson, _ := json.Marshal(*user)
	if err = global.RedisEngine.Set(context.Background(), global.TokenPrefix+token, userJson, 24*time.Hour).Err(); err != nil {
		global.Logger.Error("缓存用户时出错出错！")
	}
	return user, err
}

// Login 存放登录逻辑的代码
func Login(p *models.ParamLogin) (*models.User, error) {
	user := &models.User{
		Name:     p.Username,
		PassWord: p.Password,
	}
	//传递的是指针，通过在login里面操作，因此能够拿到user,UserID
	if err := dao.Login(user); err != nil {
		return nil, err
	}
	//生成JWT
	// fmt.Println(jwt.GetToken(user.UserID, user.UserName))
	token, err := jwt.GetToken(user.Id, user.Name)
	user.Token = token

	//将token加入redis中，过期时间是24小时, 键是token, 值是用户对象
	userJson, _ := json.Marshal(*user)
	if err = global.RedisEngine.Set(context.Background(), global.TokenPrefix+token, userJson, 24*time.Hour).Err(); err != nil {
		global.Logger.Error("缓存用户时出错出错！")
	}
	return user, err
}

//UserInfo 存放用户信息的代码
func UserInfo(param *models.ParamInfo) (*models.User, error) {
	//1、判断用户信息，并返回用户的基本信息
	return dao.GetUserByID(param.Uid)
}

//func GetUserByUserId(userId int64) (*models.User, error) {
//	var user *models.User
//	var err error
//	// 从缓存查找
//	userJson, _ := global.RedisEngine.Get(global.Ctx, global.UserPrefix+string(userId)).Result()
//	if len(userJson) == 0 {
//		//不存在
//		user, err = dao.GetUserByID(userId)
//		userJson, _ := json.Marshal(*user)
//		global.RedisEngine.Set(global.Ctx, global.UserPrefix+string(userId), userJson, 24*time.Hour)
//	} else {
//		//存在
//		user = &models.User{}
//		json.Unmarshal([]byte(userJson), user)
//	}
//	return user, err
//}

// GetUserByToken 通过token获取对象
// token : 用户token值
func GetUserByToken(token string) *models.User {
	var user = models.User{}
	userJson, _ := global.RedisEngine.Get(global.Ctx, global.TokenPrefix+token).Result()
	err := json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		global.Logger.Error("错误")
	}
	return &user
}

func ExistsKey(token string) bool {
	exists, _ := global.RedisEngine.Exists(global.Ctx, token).Result()
	if exists > 0 {
		return true
	}
	return false
}
