package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fadlinux/sawitpro_assignment/service"

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

	// res, err := s.Service.CreateEstate(ctx.Request().Context(), payload)
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, err)
	// }

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": "",
	})
}

// This endpoint to get stats
func (s *Server) GetStats(ctx echo.Context) error {
	var request service.CreateTreeRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// res, err := s.Service.CreateEstate(ctx.Request().Context(), payload)
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, err)
	// }

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": "",
	})
}

// This endpoint to get drone plan
func (s *Server) GetDronePlan(ctx echo.Context) error {
	var request service.CreateTreeRequest

	if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	// res, err := s.Service.CreateEstate(ctx.Request().Context(), payload)
	// if err != nil {
	// 	return ctx.JSON(http.StatusInternalServerError, err)
	// }

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id": "",
	})
}
