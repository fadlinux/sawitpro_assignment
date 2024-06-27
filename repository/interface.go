package repository

import (
	"context"

	"github.com/fadlinux/sawitpro_assignment/model"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	CreateEstate(ctx context.Context, request model.Estate) error
	CreateTree(ctx context.Context, request model.Tree) error
	CreateStats(ctx context.Context, request model.Stats) error
	ListStatsByEstateId(ctx context.Context, id uuid.UUID) ([]model.Stats, error)
	UpdateTree(ctx context.Context, request model.Tree) error
	FindTreeById(ctx context.Context, id uuid.UUID) (model.Tree, error)
	FindStatsByEstateId(ctx context.Context, id uuid.UUID) (FindStatsResponse, error)
	FindAllTreeByEstateId(ctx context.Context, estateId uuid.UUID) ([]model.Tree, error)
	FindEstateById(ctx context.Context, id uuid.UUID) (model.Estate, error)
}
