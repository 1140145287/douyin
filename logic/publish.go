package logic

import (
	"bytes"
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func DoPublish(param *models.ParamPublishAction, c *gin.Context) error {
	//获取用户和视频文件
	userId := GetUserByToken(param.Token).Id
	data, err := c.FormFile("data")
	if err != nil {
		return err
	}
	//上传视频到阿里云，封面截取
	filename := filepath.Base(data.Filename)
	videoName := fmt.Sprintf("%d_%s", userId, filename) //video命名
	//videoPath := filepath.Join("./public/", videoName)  //video的public路径

	fileHandle, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "打开文件错误",
		})
		return err
	}
	defer fileHandle.Close()
	fileByte, _ := ioutil.ReadAll(fileHandle)

	//if err := c.SaveUploadedFile(data, videoPath); err != nil { //先上传到public文件夹
	//	return err
	//}
	videoUrl, coverUrl, err := SaveVideoToOSS(videoName, fileByte) //再上传到OSS
	if err != nil {
		return err
	}
	//video及cover路径存入视频列表
	video := models.Video{
		UserId:   userId,
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl,
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

func SaveVideoToOSS(videoName string, data []byte) (videoUrl string, coverUrl string, err error) {
	//上传 参数1为上传地址 参数2为本地文件地址
	//err = global.OSSEngine.PutObjectFromFile(global.OSSetting.TargetPath+videoName, "public/"+videoName)
	err = global.OSSEngine.PutObject(global.OSSetting.TargetPath+videoName, bytes.NewReader(data))
	if err != nil {
		return "", "", err
	}
	//回传视频和封面地址
	videoUrl = fmt.Sprintf("%s/%s%s", global.OSSetting.TargetURL, global.OSSetting.TargetPath, videoName)
	coverUrl = fmt.Sprintf("https://kauizhaotan.oss-cn-shanghai.aliyuncs.com/%s%s"+
		"?x-oss-process=video/snapshot,t_5000,m_fast", global.OSSetting.TargetPath, videoName)
	return videoUrl, coverUrl, nil
}
