package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os/exec"
	"path/filepath"
	"strings"
)

func DoPublish(param *models.ParamPublishAction, c *gin.Context) error {
	//获取用户和视频文件
	userId := GetUserByToken(param.Token).Id
	data, err := c.FormFile("data")
	if err != nil {
		return err
	}

	//保存视频至public文件夹(这块不用了)
	filename := filepath.Base(data.Filename)                  //a.mp4
	videoName := fmt.Sprintf("video_%d_%s", userId, filename) //video_user_a.mp4
	coverName := fmt.Sprintf("cover_%d_%s", userId, filename) //cover_user_a.mp4
	coverName = strings.Split(coverName, ".")[0] + ".jpg"
	videoUrl := filepath.Join("./public/", videoName) //video路径
	coverUrl := filepath.Join("./public/", coverName) //cover路径
	if err := c.SaveUploadedFile(data, videoUrl); err != nil {
		return err
	}
	GetCoverFromVideo(videoUrl, coverUrl, 1) //获取视频封面

	//上传视频到阿里云，封面截取(待更新)

	//video及cover路径存入视频列表
	video := models.Video{
		UserId:   userId,
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl, //封面地址待确认
		Title:    param.Title,
	}
	err = global.MysqlEngine.Transaction(func(tx *gorm.DB) error {
		return dao.PublishVideo(video)
	})
	if err != nil {
		return err
	}
	return nil
}

func GetPublishList(param *models.ParamPublishList) []models.Video {
	return dao.GetPublishListByUserId(param.UserId)
}

//GetCoverFromVideo 获取视频帧
func GetCoverFromVideo(videoUrl string, coverUrl string, frameNum int) {
	dirAbs, _ := filepath.Abs("./")
	videoAbsUrl := dirAbs + "\\" + videoUrl
	coverAbsUrl := dirAbs + "\\" + coverUrl
	cmdArgs := []string{"-i", videoAbsUrl, "-y", "-f", "image2", "-ss", fmt.Sprint(frameNum), coverAbsUrl}
	cmd := exec.Command(dirAbs+"\\pkg\\ffmpeg", cmdArgs...)
	err := cmd.Run()
	if err != nil {
		return
	}
}
