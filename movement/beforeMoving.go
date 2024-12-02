package movement

import (
	"strconv"
)

func BeforeMoving(paths [][]string, NumberOfAnts int) map[int][]string {

	indexOfCurentPath := 0
	antsSelection := make(map[int][]string, len(paths))

	for i := 1; i <= NumberOfAnts; i++ {

		if indexOfCurentPath == len(paths)-1 {
			indexOfCurentPath = 0
		} else if i != 1 {
			if len(antsSelection[indexOfCurentPath])+len(paths[indexOfCurentPath]) > len(antsSelection[indexOfCurentPath+1])+len(paths[indexOfCurentPath+1]) {
				indexOfCurentPath += 1
			} else {
				for j := 0; j < len(paths); j++ {
					if len(antsSelection[indexOfCurentPath])+len(paths[indexOfCurentPath]) > len(antsSelection[j])+len(paths[j]) {
						indexOfCurentPath = j
						break
					}
				}
			}
		}
		antsSelection[indexOfCurentPath] = append(antsSelection[indexOfCurentPath], "L"+strconv.Itoa(i))
	}

	return antsSelection
}

