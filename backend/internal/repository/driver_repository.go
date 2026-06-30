package repository

import (
	"backend/db"
	"context"
)

type DriverRepository struct {
	q *db.Queries
}

func NewDriverRepository(q *db.Queries) *DriverRepository {
	return &DriverRepository{
		q: q,
	}
}

func (q *DriverRepository) CreateUser(ctx context.Context, params db.FindOrCreateParams) (db.FindOrCreateRow, error) {
	return q.q.FindOrCreate(ctx, params)
}
