package api

import (
	"errors"
)

type Shape2d string
type Shape3d string

const (
	Circle   Shape2d = "Circle"
	Square   Shape2d = "Square"
	Triangle Shape2d = "Triangle"

	Cone     Shape3d = "Cone"
	Cylinder Shape3d = "Cylinder"
	Cube     Shape3d = "Cube"
	Prism    Shape3d = "Prism"
	Pyramid  Shape3d = "Pyramid"
	Sphere   Shape3d = "Sphere"
)

func (s2d Shape2d) Validate() error {
	switch s2d {
	case Circle, Square, Triangle:
		return nil
	}

	return errors.New("can have only Circle, Square, or Triangle as it's value")
}

func (s3d Shape3d) Validate() error {
	switch s3d {
	case Cone, Cylinder, Cube, Prism, Pyramid, Sphere:
		return nil
	}

	return errors.New("can have only Cone, Cylinder, Cube, Prism, Pyramid, or Sphere as it's value")
}

type SolveVerityRequestBody struct {
	InsideRoomLeftStatueSymbol   Shape2d
	InsideRoomMidStatueSymbol    Shape2d
	InsideRoomRightStatueSymbol  Shape2d
	OutsideRoomLeftStatueSymbol  Shape3d
	OutsideRoomMidStatueSymbol   Shape3d
	OutsideRoomRightStatueSymbol Shape3d
	IsChallengePhaseTwo          bool
}

func (req SolveVerityRequestBody) Validate() error {
	// Empty Value Validations
	if req.InsideRoomLeftStatueSymbol == "" {
		return errors.New("InsideRoomLeftStatueSymbol field is not present")
	}

	if req.InsideRoomMidStatueSymbol == "" {
		return errors.New("InsideRoomMidStatueSymbol field is not present")
	}

	if req.InsideRoomRightStatueSymbol == "" {
		return errors.New("InsideRoomRightStatueSymbol field is not present")
	}

	if req.OutsideRoomLeftStatueSymbol == "" {
		return errors.New("OutsideRoomLeftStatueSymbol field is not present")
	}

	if req.OutsideRoomMidStatueSymbol == "" {
		return errors.New("OutsideRoomMidStatueSymbol field is not present")
	}

	if req.OutsideRoomRightStatueSymbol == "" {
		return errors.New("OutsideRoomRightStatueSymbol field is not present")
	}

	// Wrong value validations
	if err := req.InsideRoomLeftStatueSymbol.Validate(); err != nil {
		return errors.New("InsideRoomLeftStatueSymbol field " + err.Error())
	}

	if err := req.InsideRoomMidStatueSymbol.Validate(); err != nil {
		return errors.New("InsideRoomMidStatueSymbol field " + err.Error())
	}

	if err := req.InsideRoomRightStatueSymbol.Validate(); err != nil {
		return errors.New("InsideRoomRightStatueSymbol field " + err.Error())
	}

	if err := req.OutsideRoomLeftStatueSymbol.Validate(); err != nil {
		return errors.New("OutsideRoomLeftStatueSymbol field " + err.Error())
	}

	if err := req.OutsideRoomMidStatueSymbol.Validate(); err != nil {
		return errors.New("OutsideRoomMidStatueSymbol field " + err.Error())
	}

	if err := req.OutsideRoomRightStatueSymbol.Validate(); err != nil {
		return errors.New("OutsideRoomRightStatueSymbol field " + err.Error())
	}

	return nil
}

// Solve Verity Outside Dissection Response
type SolveVerityResponse struct {
	OutsideDissectionSteps        []string
	OutsideTargetStatueShapeNames [3]Shape3d
	InsideDissectionSteps         [][]string
}

type Error struct {
	Code    uint
	Message string
}
