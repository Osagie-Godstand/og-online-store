package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Osagie-Godstand/online-shop-apiv1/api"
	"github.com/Osagie-Godstand/online-shop-apiv1/db"
	"github.com/Osagie-Godstand/online-shop-apiv1/db/fixtures"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	var (
		ctx           = context.Background()
		mongoEndpoint = os.Getenv("MONGO_DB_URL")
		mongoDBName   = os.Getenv("MONGO_DB_NAME")
	)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(mongoDBName).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Product: *db.NewMongoProductStore(client.Database("og-online-store")),
	}

	user := fixtures.AddUser(store, "osagie", "desmond", false)
	fmt.Println("osagie ->", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", api.CreateTokenFromUser(admin))
	product := fixtures.AddProduct(store, "22446688", "ball", 11.1)
	fmt.Println("product ->", product.ID)

}
