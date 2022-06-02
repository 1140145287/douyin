package controller

import (
	"douyin/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	if user, exist := usersLoginInfo[token]; exist {
		comment := Comment{
			Id:      commentId,
			UserId:  user.Id,
			VideoId: videoId,
		}
		if actionType == "1" {
			//add the comment
			commentText := c.Query("comment_text")
			comment.Content = commentText
			global.DBEngine.Create(&comment)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         comment.Id,
					User:       user,
					Content:    commentText,
					CreateDate: comment.CreatedAt.Format("01-03"),
				}})
		} else {
			//delete the comment
			global.DBEngine.Delete(&comment)
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	var comments []Comment
	if _, exist := usersLoginInfo[token]; exist {
		global.DBEngine.Where("video_id = ?", videoId).Find(&comments)
		for i := 0; i < len(comments); i++ {
			comments[i].CreateDate = comments[i].CreatedAt.Format("01-02")
		}
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0},
			CommentList: comments,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
