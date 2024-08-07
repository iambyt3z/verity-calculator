package verity_calculator

import "github.com/iambyt3z/verity-calculator/api"

func GetTargetStatues(
	insideRoomLeftStatueSymbol api.Shape2d,
	insideRoomMidStatueSymbol api.Shape2d,
	insideRoomRightStatueSymbol api.Shape2d,
	isChallengePhaseTwo bool,
) [3]Statue {
	var targetStatueMap map[api.Shape2d]Statue
	var result [3]Statue

	if isChallengePhaseTwo {
		targetStatueMap = map[api.Shape2d]Statue{
			api.Circle:   {2, 0, 0},
			api.Square:   {0, 2, 0},
			api.Triangle: {0, 0, 2},
		}

		result = [3]Statue{
			targetStatueMap[insideRoomMidStatueSymbol],
			targetStatueMap[insideRoomRightStatueSymbol],
			targetStatueMap[insideRoomLeftStatueSymbol],
		}
	} else {
		targetStatueMap = map[api.Shape2d]Statue{
			api.Circle:   {0, 1, 1},
			api.Square:   {1, 0, 1},
			api.Triangle: {1, 1, 0},
		}

		result = [3]Statue{
			targetStatueMap[insideRoomLeftStatueSymbol],
			targetStatueMap[insideRoomMidStatueSymbol],
			targetStatueMap[insideRoomRightStatueSymbol],
		}
	}

	return result
}
