package movement

import (
	"strings"
)

func GenerateSteps(antPaths map[int][]string, validPaths [][]string) [][][]string {
	result := make([][][]string, len(validPaths))

	for pathIdx, path := range validPaths {
		steps := [][]string{}

		ants := antPaths[pathIdx]

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

func MergeSteps(steps [][][]string) [][]string {
	var merged [][]string

	maxSteps := 0
	for _, pathSteps := range steps {
		if len(pathSteps) > maxSteps {
			maxSteps = len(pathSteps)
		}
	}

	for stepIdx := 0; stepIdx < maxSteps; stepIdx++ {
		mergedStep := []string{}

		for _, pathSteps := range steps {
			if stepIdx < len(pathSteps) {
				mergedStep = append(mergedStep, pathSteps[stepIdx]...)
			}
		}

		merged = append(merged, mergedStep)
	}

	return merged
}

func RemoveStartRoom(paths [][]string) [][]string {
	for i, v := range paths {
		v = v[1:]
		paths[i] = v
	}
	return paths
}

func JoinStepsWithNewLine(mergedSteps [][]string) string {
	var result string

	for _, step := range mergedSteps {
		stepStr := strings.Join(step, " ")
		result += stepStr + "\n"
	}

	return strings.TrimSpace(result)
}

func GetBestResult(res1 [][]string, res2 [][]string) [][]string {

	if len(res1) < len(res2) {
		return res1
	}
	return res2
}