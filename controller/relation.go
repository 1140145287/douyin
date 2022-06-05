package controller

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []models.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	// 获取request参数
	p := new(models.ParamInfo)
	if err := c.ShouldBind(p); err != nil {
		global.Logger.Error("Error in getting follow list", zap.Error(err))
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
			UserList: []models.User{},
		})
		return
	}

	// 业务处理
	followList, err := dao.GetFollowList(p.Uid)
	if err != nil {
		global.Logger.Error("获取关注用户失败", zap.Error(err))
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
			UserList: []models.User{},
		})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followList,
	})
}

// FollowerList 获取关注者列表
func FollowerList(c *gin.Context) {
	// 获取request参数
	p := new(models.ParamInfo)
	if err := c.ShouldBind(p); err != nil {
		global.Logger.Error("Error in getting follower list", zap.Error(err))
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
			UserList: []models.User{},
		})
		return
	}

	// 业务处理
	followerList, err := dao.GetFollowerList(p.Uid)
	if err != nil {
		global.Logger.Error("获取粉丝用户失败", zap.Error(err))
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
			},
			UserList: []models.User{},
		})
		return
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: followerList,
	})
}
