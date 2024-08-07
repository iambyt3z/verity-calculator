package verity_calculator

import "github.com/iambyt3z/verity-calculator/api"

type ShapesExcess struct {
	circles   int
	squares   int
	triangles int
}

func (se *ShapesExcess) HasShapesExcess() (api.Shape2d, bool) {
	switch {
	case se.circles > 0:
		return api.Circle, true

	case se.squares > 0:
		return api.Square, true

	case se.triangles > 0:
		return api.Triangle, true
	}

	return "", false
}

func (se *ShapesExcess) HasShapesDeficiency(shapeName api.Shape2d) bool {
	switch {
	case shapeName == api.Circle && se.circles < 0:
		return true

	case shapeName == api.Square && se.squares < 0:
		return true

	case shapeName == api.Triangle && se.triangles < 0:
		return true
	}

	return false
}

func (se *ShapesExcess) IsBalanced() bool {
	return (se.circles == 0 &&
		se.squares == 0 &&
		se.triangles == 0)
}

func (se *ShapesExcess) IncreamentShapeExcess(shapeName api.Shape2d) {
	switch shapeName {
	case api.Circle:
		se.circles += 1

	case api.Square:
		se.squares += 1

	case api.Triangle:
		se.triangles += 1
	}
}

func (se *ShapesExcess) DecreamentShapeExcess(shapeName api.Shape2d) {
	switch shapeName {
	case api.Circle:
		se.circles -= 1

	case api.Square:
		se.squares -= 1

	case api.Triangle:
		se.triangles -= 1
	}
}
