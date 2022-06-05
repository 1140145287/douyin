package controller

import (
	"douyin/global"
	"douyin/logic"
	"douyin/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type CommentListResponse struct {
	Response
	CommentList []models.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment models.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	param := new(models.ParamCommentAction)
	user, _ := c.Get("auth")
	c.ShouldBindQuery(param)
	comment, err := logic.DoComment(param)
	if err != nil {
		global.Logger.Error("can't create comment", zap.Error(err))
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  CodeMap[CodeInternalError],
		})
		return
	}
	c.JSON(http.StatusOK, CommentActionResponse{
		Response: Response{StatusCode: 0},
		Comment: models.Comment{
			Id:         comment.Id,
			User:       user.(models.User),
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("01-03"),
		}},
	)
}

// CommentList Get all comments
func CommentList(c *gin.Context) {
	param := new(models.ParamCommentList)
	c.ShouldBindQuery(param)
	comments, err := logic.GetCommentList(param)
	if err != nil {
		global.Logger.Error("Can't get comments ", zap.Error(err))
		c.JSON(http.StatusOK, Response{
			StatusCode: 500,
			StatusMsg:  CodeInternalError.Msg(),
		})
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
