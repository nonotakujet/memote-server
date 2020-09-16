package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/julienschmidt/httprouter"

	"github.com/nonotakujet/memote-server/handler"
	"github.com/nonotakujet/memote-server/registry"
)

func main() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	// registryの生成
	repo := registry.NewRepository(ctx, client)

	// handlers.
	positionHandler := handler.NewPositionHandler(repo)

	//ルーティングの設定
	router := httprouter.New()
	router.GET("/api/positions", positionHandler.Post)
	router.POST("/api/positions", positionHandler.Post)

	//サーバー起動
	port := ":8080" //"3000"だとエラーになる
	fmt.Println(`Server Start >> http:// localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, router))
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal("GOOGLE_CLOUD_PROJECT must be set")
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Verify that we can communicate and authenticate with the Firestore
	err = client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		return nil
	})
	if err != nil {
		log.Fatalf("firestoredb: could not connect: %v", err)
	}

	// Close client when done with
	// defer client.Close()
	return client
}
