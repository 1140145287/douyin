package controller

import (
	"douyin/dao"
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User *models.User `json:"user"`
}

// Register register a new user
func Register(c *gin.Context) {
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

// Login
func Login(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBind(p); err != nil {
		//请求参数有误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.Logger.Error("Login with client problems", zap.Error((err)))
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeServerBusy]},
			})
		}
		global.Logger.Error("Login with invalid param", zap.Error((err)))
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

// UserInfo get user info
func UserInfo(c *gin.Context) {
	//1、参数校验
	p := new(models.ParamInfo)
	if err := c.ShouldBindQuery(p); err != nil {
		//请求参数有误
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			global.Logger.Error("UserInfo with client problems", zap.Error((err)))
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeServerBusy]},
			})
		}
		global.Logger.Error("UserInfo with invalid param", zap.Error((err)))
		return
	}
	//2、业务逻辑处理
	user, err := logic.UserInfo(p)
	if errors.Is(err, dao.ErrorInvalidID) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: CodeMap[CodeUserNotExit]},
		})
	}
	//3、返回响应
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})

}
