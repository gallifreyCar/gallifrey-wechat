package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromID    uint   `json:"from_id"`
	ToID      uint   `json:"to_id"`
	Content   string `json:"content"`
	ChatType  int    `json:"chat_type"`  // 0: 单聊 1: 群聊 2: 广播
	MediaType int    `json:"media_type"` // 0: 文字 1: 图片 2: 语音 3: 视频 4: 文件
	Picture   string `json:"picture"`
	Audio     string `json:"audio"`
	Url       string `json:"url"`
}
