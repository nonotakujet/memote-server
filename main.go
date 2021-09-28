package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/DeNA/aelog"
	"github.com/DeNA/aelog/middleware"
	"github.com/go-chi/chi"
	"github.com/nonotakujet/memote-server/domain/model"
	"github.com/nonotakujet/memote-server/handler"
	"github.com/nonotakujet/memote-server/registry"
)

func main() {
	// Get a Firestore client.
	ctx := context.Background()
	client := createFirestoreClient(ctx)
	if client == nil {
		os.Exit(1)
	}
	defer client.Close()

	// registryの生成
	repo := registry.NewRepository(ctx, client)

	// handlers.
	//	positionHandler := handler.NewPositionHandler(repo)
	recordHandler := handler.NewRecordHandler(repo)
	recommendedRecordsHandler := handler.NewRecommendedRecordsHandler(repo)
	fixedRecordsHandler := handler.NewFixedRecordsHandler(repo)

	// routing.
	r := chi.NewRouter()

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(setupLogger)
		r.Use(verifyFirebaseToken)
		r.Post("/records", recordHandler.Post)
		r.Get("/recommended_records", recommendedRecordsHandler.Get)
		r.Get("/fixed_records/{recordId}", fixedRecordsHandler.Get)
		r.Put("/fixed_records/{recordId}", fixedRecordsHandler.Update)
		r.Get("/fixed_records", fixedRecordsHandler.GetAll)
	})

	// サーバー起動
	port := ":8080"
	aelog.Infof(ctx, `Server Start >> http://localhost:%d`, port)
	http.ListenAndServe(port, r)
}

func createFirestoreClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		aelog.Criticalf(ctx, "GOOGLE_CLOUD_PROJECT must be set")
		return nil
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		aelog.Criticalf(ctx, "Failed to create client: %v", err)
		return nil
	}

	// Verify that we can communicate and authenticate with the Firestore
	err = client.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		return nil
	})
	if err != nil {
		aelog.Criticalf(ctx, "firestoredb: could not connect: %v", err)
		return nil
	}

	return client
}

func verifyFirebaseToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Initialize default app
		ctx := context.Background()

		app, err := firebase.NewApp(ctx, nil)
		if err != nil {
			aelog.Errorf(ctx, "error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		auth, err := app.Auth(ctx)
		if err != nil {
			aelog.Errorf(ctx, "error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(ctx, idToken)
		if err != nil {
			aelog.Errorf(ctx, "error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		newContext := model.ContextWithUID(r.Context(), model.NewUID(token.UID))
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}

func setupLogger(next http.Handler) http.Handler {
	// https://pkg.go.dev/github.com/DeNA/aelog#section-readme
	return middleware.AELogger("ServeHTTP")(next)
}
