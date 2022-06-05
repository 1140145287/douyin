package dao

import (
	"douyin/global"
	"douyin/models"
)

// CreateComment create new comment
func CreateComment(comment *models.Comment) error {
	if err := global.MysqlEngine.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

//DeleteComment delete the comment
//Note : Soft Delete that this operation only change the delete flag
func DeleteComment(comment *models.Comment) error {
	if err := global.MysqlEngine.Delete(comment).Error; err != nil {
		return err
	}
	return nil
}

// GetCommentsByVideoId the function details is as it's name :)
func GetCommentsByVideoId(commentId int64) ([]models.Comment, error) {
	var comments []models.Comment
	global.MysqlEngine.Where("video_id = ?", commentId).Order("created_at DESC").Find(&comments)
	for i := 0; i < len(comments); i++ {
		comments[i].CreateDate = comments[i].CreatedAt.Format("01-02")
	}
	return comments, nil
}
