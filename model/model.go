package model

import (
	"time"

	"github.com/google/uuid"
)

type (
	//default response uuid id
	DefaultIdResponse struct {
		Id uuid.UUID `json:"id"`
	}

	//request payload create estate
	CreateEstateRequest struct {
		Id     uuid.UUID `json:"id" validate:"required"`
		Width  int       `json:"width" validate:"required,min=1,max=50000"`
		Length int       `json:"length" validate:"required,min=1,max=50000"`
	}

	//request payload create tree
	CreateTreeRequest struct {
		Id       uuid.UUID `json:"id" validate:"required"`
		EstateId uuid.UUID `json:"estate_id" validate:"required"`
		X        int       `json:"x" validate:"required,min=1,max=50000"`
		Y        int       `json:"y" validate:"required,min=1,max=50000"`
		Height   int       `json:"height" validate:"required,min=1,max=30"`
	}

	//response struct Tree stats
	TreeStatsByEstateIdResponse struct {
		Count  int `json:"count"`
		Max    int `json:"max"`
		Min    int `json:"min"`
		Median int `json:"median"`
	}

	//model struct estate repository
	Estate struct {
		Id         uuid.UUID `json:"id"`
		Width      int       `json:"width"`
		Length     int       `json:"length"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
	}

	//model struct tree repository
	Tree struct {
		Id         uuid.UUID `json:"id"`
		Width      int       `json:"width"`
		Length     int       `json:"length"`
		Height     int       `json:"height"`
		EstateId   uuid.UUID `json:"estate_id"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
	}

	Stats struct {
		TreeId     uuid.UUID `json:"tree_id"`
		Width      int       `json:"width"`
		Length     int       `json:"length"`
		Height     int       `json:"height"`
		EstateId   uuid.UUID `json:"estate_id"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
	}
)
