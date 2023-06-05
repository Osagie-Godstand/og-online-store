package fixtures

import (
	"context"
	"fmt"
	"log"

	"github.com/Osagie-Godstand/og-online-store/db"
	"github.com/Osagie-Godstand/og-online-store/types"
)

func AddUser(store *db.Store, fn, ln string, admin bool) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     fmt.Sprintf("%s@%s.com", fn, ln),
		FirstName: fn,
		LastName:  ln,
		Password:  fmt.Sprintf("%s_%s", fn, ln),
	})
	if err != nil {
		log.Fatal(err)
	}
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddProduct(store *db.Store, sku string, name string, price float64) *types.Product {
	product := types.Product{
		SKU:   sku,
		Name:  name,
		Price: price,
	}

	insertedProduct, err := store.Product.InsertProduct(context.TODO(), &product)
	if err != nil {
		log.Fatal(err)
	}

	return insertedProduct

}
