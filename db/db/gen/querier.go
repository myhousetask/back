// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"context"
)

type Querier interface {
	CreateFamily(ctx context.Context, arg CreateFamilyParams) error
	GetLastFiveFamilies(ctx context.Context) ([]Family, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

var _ Querier = (*Queries)(nil)
