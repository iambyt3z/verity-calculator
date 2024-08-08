package verity_calculator

import "github.com/iambyt3z/verity-calculator/api"

type Statue struct {
	circles   uint
	squares   uint
	triangles uint
}

type KeyType [3]uint // {circles, squares, triangles}

var shapeNameMap2d = map[KeyType]api.Shape2d{
	{1, 0, 0}: api.Circle,
	{0, 1, 0}: api.Square,
	{0, 0, 1}: api.Triangle,
}

var shapeNameMap3d = map[KeyType]api.Shape3d{
	{2, 0, 0}: api.Sphere,
	{0, 2, 0}: api.Cube,
	{0, 0, 2}: api.Pyramid,
	{1, 1, 0}: api.Cylinder,
	{0, 1, 1}: api.Prism,
	{1, 0, 1}: api.Cone,
}

var statueFromShapeMap2d = map[api.Shape2d]Statue{
	api.Circle:   {1, 0, 0},
	api.Square:   {0, 1, 0},
	api.Triangle: {0, 0, 1},
}

var statueFromShapeMap3d = map[api.Shape3d]Statue{
	api.Sphere:   {2, 0, 0},
	api.Cube:     {0, 2, 0},
	api.Pyramid:  {0, 0, 2},
	api.Cylinder: {1, 1, 0},
	api.Prism:    {0, 1, 1},
	api.Cone:     {1, 0, 1},
}

func (s *Statue) GetShapeName() interface{} {
	key := KeyType{s.circles, s.squares, s.triangles}

	if s.circles+s.squares+s.triangles == 1 {
		if shape, found := shapeNameMap2d[key]; found {
			return shape
		}
	} else if s.circles+s.squares+s.triangles == 2 {
		if shape, found := shapeNameMap3d[key]; found {
			return shape
		}
	}

	return nil
}

func GetStatueFromShapeName2d(shapeName api.Shape2d) Statue {
	return statueFromShapeMap2d[shapeName]
}

func GetStatueFromShapeName3d(shapeName api.Shape3d) Statue {
	return statueFromShapeMap3d[shapeName]
}

func SubtractStatues(s1 *Statue, s2 *Statue) ShapesExcess {
	return ShapesExcess{
		int(s1.circles) - int(s2.circles),
		int(s1.squares) - int(s2.squares),
		int(s1.triangles) - int(s2.triangles),
	}
}
