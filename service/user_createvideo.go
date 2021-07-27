package service

import (
	"go_web/connect"
	"go_web/model"
)

type CreateVideo struct {
	VideoName   string `json:"videoname" binding:"required,gt=3,lt=6"`
	Description string `json:"description" binding:"required,gt=3,lt=6"`
}

func NewVideo(newVideo *CreateVideo) {

	video := &model.Video{
		VideoName:   newVideo.VideoName,
		Description: newVideo.Description,
	}
	db := connect.InitDB()
	db.Create(&video)

}

func GetVideoList(name string) []model.VideoList {
	videolist := make([]model.VideoList,20,20)
	db := connect.InitDB()
	db.Raw("select videos.id,users.name,videos.video_name,videos.star,videos.likes,videos.retweets,videos.description,videos.created_at"+
		" from users join videos on users.id = videos.user_id where users.name = ? ;", name).Scan(&videolist)
	return videolist

}
func UpdateVideo() {

}
func DeleteVideo() {

}
