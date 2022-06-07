package logic

import (
	"douyin/dao"
	"douyin/global"
	"douyin/models"
)

// GetVideoList获取前30条视频,直接用于播放
func GetVideoList(token string) ([]models.Video, error) {
	//获取播放视频列表
	videoList, err := dao.GetVideoList()
	if err != nil {
		return nil, err
	}
	//存储用户喜欢的视频列表
	mp := map[int64]bool{}
	if len(token) != 0 {
		if ExistsKey(global.TokenPrefix + token) {
			user_id := GetUserByToken(token).Id
			favoriteList, err := dao.GetFavoriteListByUserId(user_id)
			if err != nil {
				return nil, err
			}
			for i := 0; i < len(favoriteList); i++ {
				mp[favoriteList[i].Id] = true
			}
		}
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
		if mp[videoList[i].Id]{
			videoList[i].IsFavorite = true
		}
	}
	//如果用户已经登录，获取用户对应的视频信息
	return videoList, nil
}
