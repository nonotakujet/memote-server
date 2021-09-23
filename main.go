package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/go-chi/chi"

	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/handler"
	"github.com/nonotakujet/memote-server/registry"
)

func main() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	defer client.Close()

	// registryの生成
	repo := registry.NewRepository(ctx, client)

	// handlers.
	//	positionHandler := handler.NewPositionHandler(repo)
	recordHandler := handler.NewRecordHandler(repo)

	// routing.
	r := chi.NewRouter()

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(verifyFirebaseToken)
		r.Post("/records", recordHandler.Post)
	})

	//サーバー起動
	port := ":8080" //"3000"だとエラーになる
	fmt.Println(`Server Start >> http://localhost:%d`, port)
	log.Fatal(http.ListenAndServe(port, r))
}

func createFirestoreClient(ctx context.Context) *firestore.Client {
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

func verifyFirebaseToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Initialize default app
		ctx := context.Background()

		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			log.Printf("error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		auth, err := app.Auth(ctx)
		if err != nil {
			log.Printf("error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			log.Printf("error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		newContext := model.ContextWithUID(r.Context(), model.NewUID(token.UID))
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}
