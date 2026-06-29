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
