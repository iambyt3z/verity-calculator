package api

// Solve Verity Outside Dissection Request Body
type SolveVerityOutsideDissectionRequestBody struct {
	InsideRoomLeftStatueSymbol   string
	InsideRoomMidStatueSymbol    string
	InsideRoomRightStatueSymbol  string
	OutsideRoomLeftStatueSymbol  string
	OutsideRoomMidStatueSymbol   string
	OutsideRoomRightStatueSymbol string
}

// Solve Verity Outside Dissection Response
type SolveVerityOutsideDissectionResponse struct {
	Code                   uint
	OutsideDissectionSteps []string
}

type Error struct {
	Code    uint
	Message string
}
