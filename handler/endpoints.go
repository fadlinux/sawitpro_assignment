package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fadlinux/sawitpro_assignment/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// This endpoint to create estate
func (s *Server) CreateEstate(ctx echo.Context) error {
	var request service.CreateEstateRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	res, err := s.Service.CreateEstate(ctx.Request().Context(), request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": res,
	})
}

// This endpoint to create tree
func (s *Server) CreateTree(ctx echo.Context) error {
	var request service.CreateTreeRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	id := ctx.Param("id")
	request.EstateId = uuid.MustParse(id)

	res, err := s.Service.CreateTree(ctx.Request().Context(), request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": res,
	})
}

func (s *Server) UpdateTree(ctx echo.Context) error {
	var request service.PayloadUpdateTree

	if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	id := ctx.Param("id")
	request.Id = uuid.MustParse(id)

	res, err := s.Service.UpdateTree(ctx.Request().Context(), request)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": res,
	})
}

func (s *Server) TreeStatsByEstateId(ctx echo.Context) error {
	id := ctx.Param("id")
	uuidId := uuid.MustParse(id)

	res, err := s.Service.TreeStatsByEstateId(ctx.Request().Context(), uuidId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (s Server) Distance(ctx echo.Context) error {
	estateId := ctx.Param("id")
	uuidEstateId := uuid.MustParse(estateId)

	res, err := s.Service.DroneDistance(ctx.Request().Context(), uuidEstateId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"distance": res,
	})
}
