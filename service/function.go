package service

import (
	"context"

	"github.com/fadlinux/sawitpro_assignment/helpers"
	"github.com/fadlinux/sawitpro_assignment/model"
	"github.com/fadlinux/sawitpro_assignment/repository"

	"github.com/google/uuid"
)

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) ServiceInterface {
	return Service{
		repo: repo,
	}
}

func (s Service) CreateEstate(ctx context.Context, request CreateEstateRequest) (uuid.UUID, error) {
	id := uuid.New()
	request.Id = id

	if err := helpers.Validator(request); err != nil {
		return uuid.Nil, err
	}

	if err := s.repo.CreateEstate(ctx, model.Estate{
		Id:     request.Id,
		Width:  request.Width,
		Length: request.Length,
	}); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (s Service) CreateTree(ctx context.Context, payload CreateTreeRequest) (uuid.UUID, error) {
	id := uuid.New()
	payload.Id = id

	// //, err := s.repo.FindEstateById(ctx, payload.EstateId)
	// if err != nil {
	// 	return uuid.Nil, err
	// }

	return id, nil
}
