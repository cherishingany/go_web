package model

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	ID          uint   `json:"id"`
	VideoName   string `json:"videoname" `
	Star        uint   `json:"star"`
	Likes       uint   `json:"likes"`
	Retweets    uint   `json:"retweets"`
	Description string `json:"description"`
	gorm.Model
}

type VideoList struct {
	ID          uint   `json:"id,omitempty" `
	VideoName   string `json:"videoname" `
	Star        uint   `json:"star"`
	Likes       uint   `json:"likes"`
	Retweets    uint   `json:"retweets"`
	Description string `json:"description"`
	CreatedAt   time.Time `json:"createdtime"`
}
