package main

import (
	"fmt"
	"os"

	input "lemin/Input"
	"lemin/functions"
	movement "lemin/movement"
)

var linkstart, linkend = false, false


func GetLargestLenOfResult(res1 [][]string, res2 [][]string) [][]string {

	if len(res1) < len(res2) {
		return res1
	}
	return res2
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	numAnts, rooms, connections, err := input.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	start, end := functions.FindStartEnd(rooms)
	for a := range connections {
		if a == start {
			linkstart = true
		} else if a == end {
			linkend = true
		}
	}

	if !linkend || !linkstart {
		fmt.Println("start or end isnt linked")
		return
	}
	paths1, paths2, err := functions.FindUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}

	result1 := movement.MergeSteps(movement.GenerateSteps(movement.BeforeMoving(paths1, numAnts), movement.RemoveStartRoom(paths1)))
	result2 := movement.MergeSteps(movement.GenerateSteps(movement.BeforeMoving(paths2, numAnts), movement.RemoveStartRoom(paths2)))
	fmt.Println(movement.JoinStepsWithNewLine(GetLargestLenOfResult(result1, result2)))
}
