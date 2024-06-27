package service

import (
	"context"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	CreateEstate(ctx context.Context, payload CreateEstateRequest) (uuid.UUID, error)
	CreateTree(ctx context.Context, payload CreateTreeRequest) (uuid.UUID, error)
	UpdateTree(ctx context.Context, payload PayloadUpdateTree) (uuid.UUID, error)
	TreeStatsByEstateId(ctx context.Context, estateId uuid.UUID) (TreeStatsByEstateIdResponse, error)
	DroneDistance(ctx context.Context, estateId uuid.UUID) (int, error)
}
