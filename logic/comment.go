package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"gorm.io/gorm"
)

/**
* 评论接口Service层
 */

// DoComment 用户评论
func DoComment(param *models.ParamCommentAction) (models.Comment, error) {
	var err error
	comment := models.Comment{
		Id:      param.CommentId,
		UserId:  GetUserByToken(param.Token).Id,
		VideoId: param.VideoId,
		Content: param.CommentText,
	}
	//ActionType == 1 : add new comment
	//ActionType == 2 : delete the old comment
	err = global.MysqlEngine.Transaction(func(tx *gorm.DB) error {
		if param.ActionType == 1 {
			comment.IsDel = 0
			err = dao.CreateComment(&comment)
		} else {
			err = dao.DeleteComment(&comment)
		}
		err = dao.UpdateVideoComment(&comment)
		return err
	})
	return comment, err
}

// GetCommentList 获取视频评论列表
func GetCommentList(param *models.ParamCommentList) ([]models.Comment, error) {
	comments, err := dao.GetCommentsByVideoId(param.VideoId)
	for i := 0; i < len(comments); i++ {
		user, _ := dao.GetUserByID(comments[i].UserId)
		comments[i].User = *user
	}
	return comments, err
}
