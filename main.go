package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/nonotakujet/memote-server/handler"
	"github.com/nonotakujet/memote-server/registry"
)

func main() {
	// registryの生成
	repo := registry.NewRepository()

	// handlers.
	positionHandler := handler.NewPositionHandler(repo)

	//ルーティングの設定
	router := httprouter.New()
	router.POST("/api/positions", positionHandler.Post)

	//サーバー起動
	port := ":8080" //"3000"だとエラーになる
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, router))
}
