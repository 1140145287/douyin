package models

type Favorite struct {
	Favorite_id int64 `json:"favorite_id,omitempty"`
	User_id     int64 `json:"user_id"`
	Video_id    int64 `json:"video_id"`
	Isdeleted   int64 `json:"isdeleted"`
}
