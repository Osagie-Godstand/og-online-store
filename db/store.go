package db

import (
	"context"

	"github.com/Osagie-Godstand/og-online-store/types"
)

type ProductStorer interface {
	Insert(context.Context, *types.Product) error
	GetByID(context.Context, string) (*types.Product, error)
	GetAll(context.Context) ([]*types.Product, error)
}
