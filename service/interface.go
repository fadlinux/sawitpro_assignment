package service

import (
	"context"

	"github.com/google/uuid"
)

type ServiceInterface interface {
	CreateEstate(ctx context.Context, payload CreateEstateRequest) (uuid.UUID, error)
	CreateTree(ctx context.Context, payload CreateTreeRequest) (uuid.UUID, error)
}
