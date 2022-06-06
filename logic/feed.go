package logic

import (
	"douyin/dao"
	"douyin/models"
)

// GetVideoList获取前30条视频,直接用于播放
func GetVideoList() ([]models.Video, error) {
	//获取播放视频列表
	videoList, err := dao.GetVideoList()
	if err != nil {
		return nil, err
	}
	//根据不同的视频获取不同作者信息
	list_len := len(videoList)
	for i := 0; i < list_len; i++ {
		authorInfo, err := dao.GetUserByID(videoList[i].UserId)
		//这里如果有一些视频没查到作者信息,直接continue了
		if err != nil {
			continue
		}
		videoList[i].Author = *authorInfo
	}
	return videoList, nil
}
