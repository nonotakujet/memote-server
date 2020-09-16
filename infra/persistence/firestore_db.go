package persistence

import (
	"cloud.google.com/go/firestore"
)

type DB struct {
	client *firestore.Client
}

func NewDB(client *firestore.Client) *DB {
	return &DB{
		client: client,
	}
}

func (db *DB) Close() error {
	return db.client.Close()
}
