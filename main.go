package main

import (
	"github.com/gallifreyCar/gallifrey-wechat/database"
	"github.com/gallifreyCar/gallifrey-wechat/router"
)

func main() {
	database.InitDB()
	r := router.Router()
	r.Run()
}
