package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	videoName := fmt.Sprintf("%d_%s", userId, filename)         //video命名
	videoPath := filepath.Join("./public/", videoName)          //video的public路径
	if err := c.SaveUploadedFile(data, videoPath); err != nil { //先上传到public文件夹
		return err
	}
	videoUrl, coverUrl, err := SaveVideoToOSS(videoName) //再上传到OSS
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

func SaveVideoToOSS(videoName string) (videoUrl string, coverUrl string, err error) {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endpoint := "https://oss-accelerate.aliyuncs.com"
	accessKeyId := "LTAI4FysLakF4dQbPPJakWia"
	accessKeySecret := "TbSZ2mcpDLYvDocbd5s949MThLcUYX"
	bucketName := "kauizhaotan"
	//上传路径
	targetPath := "douyin/video/"
	// 创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return "", "", err
	}
	// 使用 特定的 bucket
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", "", err
	}
	//上传 参数1为上传地址 参数2为本地文件地址
	err = bucket.PutObjectFromFile(targetPath+videoName, "public/"+videoName)
	if err != nil {
		return "", "", err
	}
	//回传视频和封面地址
	videoUrl = fmt.Sprintf("https://kauizhaotan.oss-cn-shanghai.aliyuncs.com/%s%s", targetPath, videoName)
	coverUrl = fmt.Sprintf("https://kauizhaotan.oss-cn-shanghai.aliyuncs.com/%s%s"+
		"?x-oss-process=video/snapshot,t_5000,m_fast", targetPath, videoName)
	return videoUrl, coverUrl, nil
}
