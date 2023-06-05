package db

import (
	"context"
	"fmt"

	"github.com/Osagie-Godstand/og-online-store/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductStorer interface {
	InsertProduct(context.Context, *types.Product) (*types.Product, error)
	GetByID(context.Context, string) (*types.Product, error)
	GetAll(context.Context) ([]*types.Product, error)
}

type MongoProductStore struct {
	db   *mongo.Database
	coll string
}

func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:   db,
		coll: ("products"),
	}
}

func (s *MongoProductStore) InsertProduct(ctx context.Context, p *types.Product) (*types.Product, error) {
	res, err := s.db.Collection(s.coll).InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	InsertedIDString := p.ID.Hex()
	fmt.Println(InsertedIDString)

	return p, err
}

func (s *MongoProductStore) GetAll(ctx context.Context) ([]*types.Product, error) {
	cursor, err := s.db.Collection(s.coll).Find(ctx, map[string]any{})
	if err != nil {
		return nil, err
	}

	products := []*types.Product{}
	err = cursor.All(ctx, &products)
	return products, err
}

func (s *MongoProductStore) GetByID(ctx context.Context, id string) (*types.Product, error) {
	var (
		objID, _ = primitive.ObjectIDFromHex(id)
		res      = s.db.Collection(s.coll).FindOne(ctx, bson.M{"_id": objID})
		p        = &types.Product{}
		err      = res.Decode(p)
	)
	return p, err
}
