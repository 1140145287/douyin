package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"

	"douyin/pkg/jwt"
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
	return user, err
}

//UserInfo 存放用户信息的代码
func UserInfo(id int64) (*models.User, error) {
	return dao.GetUserByID(id)
}
