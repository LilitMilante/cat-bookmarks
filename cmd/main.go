package main

import (
	"cat-bookmarks/api"
	"cat-bookmarks/bootstrap"
	"cat-bookmarks/domain/usecase"
	"log"
)

func main() {
	conf, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	bookmarkCase := usecase.NewBookmark()

	s := api.NewServer(conf.HttpPort, bookmarkCase)
	s.SetRoutes()
	log.Fatal(s.Start())
}
