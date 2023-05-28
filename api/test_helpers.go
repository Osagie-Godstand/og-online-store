package api

import (
	"context"
	"os"
	"testing"

	"github.com/Osagie-Godstand/og-online-store/db"
	"go.mongodb.org/mongo-driver/mongo"
)

type testdb struct {
	client *mongo.Client
	*db.Store
}

func (tdb *testdb) teardown(t *testing.T) {
	dbname := os.Getenv(db.MongoDBNameEnvName)
	if err := tdb.client.Database(dbname).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}
