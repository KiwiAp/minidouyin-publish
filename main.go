package main

import (
	"log"
	publish "publish/kitex_gen/publish/publishservice"
	"publish/model/db"
)

func main() {
	db.Init()
	svr := publish.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
