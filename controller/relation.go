package controller

import (
	"douyin/global"
	"douyin/logic"
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
	//token := c.Query("token")
	p := new(models.ParamRelationAction)
	if err := c.ShouldBind(p); err != nil {
		global.Logger.Error("Error in getting relation action paras", zap.Error(err))
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "parameters invalid"})
		return
	} // 确认参数是否正确

	// 获取用户id
	uid := logic.GetUserByToken(p.Token).Id
	if err := logic.RelationActionLogic(uid, p.ToUserId, p.ActionType); err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "action failed"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0})

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
	followList, err := logic.GetFollowListLogic(p.Uid)
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
	followerList, err := logic.GetFollowerListLogic(p.Uid)
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
