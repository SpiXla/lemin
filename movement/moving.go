package movement

import (
	"strings"
)

func GenerateStepsOfAnts(AllPathsWithAnts map[int][]string, AllPathsWithoutRoomStart [][]string) [][][]string {

	result := make([][][]string, len(AllPathsWithoutRoomStart))

	for pathIdx, path := range AllPathsWithoutRoomStart {

		steps := [][]string{}

		ants := AllPathsWithAnts[pathIdx]

		for stepIdx := 0; stepIdx < len(path)+len(ants)-1; stepIdx++ {
			step := []string{}

			for antIdx, ant := range ants {
				position := stepIdx - antIdx
				if position >= 0 && position < len(path) {
					move := ant + "-" + path[position]
					step = append(step, move)
				}
			}

			if len(step) > 0 {
				steps = append(steps, step)
			}
		}

		result[pathIdx] = steps
	}

	return result
}

func MergeTurnsOfPaths(StepsOfAnts [][][]string) [][]string {

	var result [][]string

	maxSteps := 0
	for _, pathSteps := range StepsOfAnts {
		if len(pathSteps) > maxSteps {
			maxSteps = len(pathSteps)
		}
	}

	for stepIdx := 0; stepIdx < maxSteps; stepIdx++ {
		mergedStep := []string{}

		for _, pathSteps := range StepsOfAnts {
			if stepIdx < len(pathSteps) {
				mergedStep = append(mergedStep, pathSteps[stepIdx]...)
			}
		}

		result = append(result, mergedStep)
	}

	return result
}

func RemoveStartRoom(Paths [][]string) [][]string {

	for index, value := range Paths {
		value = value[1:]
		Paths[index] = value
	}
	return Paths
}

func JoinStepsWithNewLine(mergedStepsOfAnts [][]string) string {

	var result string
	for _, step := range mergedStepsOfAnts {

		stepStr := strings.Join(step, " ")
		result += stepStr + "\n"
	}
	return strings.TrimSpace(result)
}

func GetBestResult(ResultOneOfTurns [][]string, ResultTwoOfTurns [][]string) [][]string {

	if len(ResultOneOfTurns) < len(ResultTwoOfTurns) {
		return ResultOneOfTurns
	}
	return ResultTwoOfTurns
}