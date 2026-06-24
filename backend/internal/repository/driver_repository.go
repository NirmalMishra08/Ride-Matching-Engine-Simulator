package repository

import (
	"backend/db"
	"context"

	"github.com/gofrs/uuid"
)

type DriverRepository struct {
	q *db.Queries
}

func NewDriverRepository(q *db.Queries) *DriverRepository {
	return &DriverRepository{
		q: q,
	}
}

func (r *DriverRepository) GetAllUser(ctx context.Context, id uuid.UUID) ([]db.User, error) {
	return r.q.GetAllUser(ctx)
}
