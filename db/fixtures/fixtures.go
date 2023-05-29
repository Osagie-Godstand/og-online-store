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
	//user.IsAdmin = admin
	insertedUser, err := store.User.InsertUser(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return insertedUser
}

func AddProduct(store *db.Store, SKU string, Name string, Slug string) *types.Product {
	product, err := types.NewProductFromRequest(&types.CreateProductRequest{
		SKU:  fmt.Sprintf("%s", SKU),
		Name: Name,
		Slug: Slug,
	})
	if err != nil {
		log.Fatal(err)
	}

	return product

}
