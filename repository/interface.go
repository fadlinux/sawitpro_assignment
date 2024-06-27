package repository

import (
	"context"

	"github.com/fadlinux/sawitpro_assignment/model"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	CreateEstate(ctx context.Context, payload model.Estate) error
	CreateTree(ctx context.Context, payload model.Tree) error
	CreateStats(ctx context.Context, payload model.Stats) error
	UpdateTree(ctx context.Context, payload model.Tree) error
	FindTreeById(ctx context.Context, id uuid.UUID) (model.Tree, error)
	FindStatsByEstateId(ctx context.Context, id uuid.UUID) (FindStatsResponse, error)
	ListStatsByEstateId(ctx context.Context, id uuid.UUID) ([]model.Stats, error)
	FindAllTreeByEstateId(ctx context.Context, estateId uuid.UUID) ([]model.Tree, error)
}
