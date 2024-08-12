package verity_calculator

import (
	"fmt"
	"log"

	"github.com/iambyt3z/verity-calculator/api"
)

func SimulateOutsideDissection(
	initialOutsideRoomStatues *[3]Statue,
	targetOutsideRoomStatues *[3]Statue,
) []string {
	var statuePosition [3]string = [3]string{"Left", "Middle", "Right"}
	var simulation []string = make([]string, 0, 16)

	var shapesExcesses [3]ShapesExcess = [3]ShapesExcess{
		SubtractStatues(&(*initialOutsideRoomStatues)[0], &(*targetOutsideRoomStatues)[0]),
		SubtractStatues(&(*initialOutsideRoomStatues)[1], &(*targetOutsideRoomStatues)[1]),
		SubtractStatues(&(*initialOutsideRoomStatues)[2], &(*targetOutsideRoomStatues)[2]),
	}

	for !shapesExcesses[0].IsBalanced() ||
		!shapesExcesses[1].IsBalanced() ||
		!shapesExcesses[2].IsBalanced() {

		var dissectionStep string

		var dissectionStatueOneIndex = 0
		var dunkOneShape api.Shape2d = ""

		var dissectionStatueTwoIndex = 0
		var dunkTwoShape api.Shape2d = ""

		for i := 0; i < 3; i++ {
			excessShapeName, hasExcess := shapesExcesses[i].HasShapesExcess()

			if hasExcess {
				dissectionStatueOneIndex = i
				dunkOneShape = excessShapeName
				break
			}
		}

		for i := 0; i < 3; i++ {
			hasDeficiency := shapesExcesses[i].HasShapesDeficiency(dunkOneShape)

			if i != dissectionStatueOneIndex && hasDeficiency {
				dissectionStatueTwoIndex = i
				break
			}
		}

		dunkTwoShape, _ = shapesExcesses[dissectionStatueTwoIndex].HasShapesExcess()
		dissectionStep = fmt.Sprintf("Dunk %s on the %v Statue", dunkOneShape, statuePosition[dissectionStatueOneIndex])
		simulation = append(simulation, dissectionStep)

		shapesExcesses[dissectionStatueOneIndex].DecreamentShapeExcess(dunkOneShape)
		shapesExcesses[dissectionStatueTwoIndex].IncreamentShapeExcess(dunkOneShape)

		dissectionStep = fmt.Sprintf("Dunk %s on the %v Statue", dunkTwoShape, statuePosition[dissectionStatueTwoIndex])
		simulation = append(simulation, dissectionStep)

		shapesExcesses[dissectionStatueTwoIndex].DecreamentShapeExcess(dunkTwoShape)
		shapesExcesses[dissectionStatueOneIndex].IncreamentShapeExcess(dunkTwoShape)
	}

	return simulation
}

func SolveOutsideDissection(
	insideRoomLeftStatueSymbol api.Shape2d,
	insideRoomMidStatueSymbol api.Shape2d,
	insideRoomRightStatueSymbol api.Shape2d,
	outsideRoomLeftStatueSymbol api.Shape3d,
	outsideRoomMidStatueSymbol api.Shape3d,
	outsideRoomRightStatueSymbol api.Shape3d,
	isChallengePhaseTwo bool,
) ([]string, [3]api.Shape3d) {
	var initialStatues [3]Statue = [3]Statue{
		GetStatueFromShapeName3d(outsideRoomLeftStatueSymbol),
		GetStatueFromShapeName3d(outsideRoomMidStatueSymbol),
		GetStatueFromShapeName3d(outsideRoomRightStatueSymbol),
	}

	targetStatues := GetTargetStatues(
		insideRoomLeftStatueSymbol,
		insideRoomMidStatueSymbol,
		insideRoomRightStatueSymbol,
		isChallengePhaseTwo,
	)

	var targetShapeNames [3]api.Shape3d

	for i := 0; i < 3; i++ {
		if shape, ok := targetStatues[i].GetShapeName().(api.Shape3d); ok {
			targetShapeNames[i] = shape
		} else {
			log.Default().Printf(fmt.Sprintf("targetStatues[%d] is not of type api.Shape3d\n", i))
		}
	}

	log.Println("Calculated target statue shapes for outside dissection")

	result := SimulateOutsideDissection(&initialStatues, &targetStatues)

	log.Println("Outside dissection simulation done")

	return result, targetShapeNames
}
