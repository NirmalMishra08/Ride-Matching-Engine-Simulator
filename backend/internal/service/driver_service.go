package service

import (
	"backend/db"
	"backend/internal/dto"
	"backend/internal/firebase"
	"backend/internal/repository"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type DriverService struct {
	DriverRepository *repository.DriverRepository
}

func NewDriverService(DriverRepository *repository.DriverRepository) *DriverService {
	return &DriverService{
		DriverRepository: DriverRepository,
	}
}

func (s *DriverService) CreateNewUser(ctx context.Context, payload firebase.FirebasePayload, req dto.AuthRequest) (db.FindOrCreateRow, error) {
	params := db.FindOrCreateParams{
		Email:        req.Email,
		Provider:     db.NullProvider{Provider: db.Provider(req.Provider), Valid: true},
		Phone:        pgtype.Text{String: req.Phone},
		Name:         req.FullName,
		PasswordHash: pgtype.Text{String: req.Password},
	}

	return s.DriverRepository.CreateUser(ctx, params)
}
