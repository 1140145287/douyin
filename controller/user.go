package controller

import (
	"douyin/dao"
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"douyin/pkg/jwt"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User *models.User `json:"user"`
}

func Register(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")

	// token := username + password

	// if _, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
	// 	})
	// } else {
	// 	atomic.AddInt64(&userIdSequence, 1)
	// 	newUser := User{
	// 		Id:   userIdSequence,
	// 		Name: username,
	// 	}
	// 	usersLoginInfo[token] = newUser
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   userIdSequence,
	// 		Token:    username + password,
	// 	})
	// }

	//1、获取参数以及参数校验
	p := new(models.ParamRegister)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.Logger.Error("Register with client problems", zap.Error((err)))
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeServerBusy]},
			})
		}
		global.Logger.Error("Register with invalid param", zap.Error((err)))
		if errors.Is(err, dao.ErrorUserExit) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeInvalidParam]},
			})
		}
		return
	}
	//手动对请求参数进行相信业务校验
	//2、业务处理
	user, err := logic.Register(p)
	if err != nil {
		global.Logger.Error("创建用户失败!", zap.Error(err))
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeUserExit]},
		})
		return
	}
	//3、返回响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    user.Token,
	})
}

func Login(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")

	// token := username + password

	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 0},
	// 		UserId:   user.Id,
	// 		Token:    token,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserLoginResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// 	})
	// }

	//1、获取请求参数以及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.Logger.Error("Register with client problems", zap.Error((err)))
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeServerBusy]},
			})
		}
		global.Logger.Error("Register with invalid param", zap.Error((err)))
		if errors.Is(err, dao.ErrorUserExit) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeInvalidParam]},
			})
		}
		return
	}
	//2、业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		global.Logger.Error("logic.Login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, dao.ErrorInvalidPassword) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeInvalidPassword]},
			})
		} else if errors.Is(err, dao.ErrorUserNotExit) {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeUserNotExit]},
			})
		}
		return
	}
	//3、返回响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.Id,
		Token:    user.Token,
	})
}

func UserInfo(c *gin.Context) {
	// if user, exist := usersLoginInfo[token]; exist {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 0},
	// 		User:     user,
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, UserResponse{
	// 		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	// 	})
	// }
	//1、参数校验
	token := c.Query("token")
	mc, err := jwt.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeInvalidAuth]},
		})
	}
	//2、业务逻辑处理
	user, err := logic.UserInfo(mc.UserID)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeUserNotExit]},
		})
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

}
