package main

import (
	"fmt"
	"os"

	input "lemin/Input"
	"lemin/functions"
	movement "lemin/movement"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	fileContent, numAnts, rooms, connections, err := input.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pathsOne, pathsTwo, err := functions.FindUniquePaths(rooms, connections)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Println(fileContent)

	ResultOneOfTurns := movement.MergeTurnsOfPaths(movement.GenerateStepsOfAnts(movement.BeforeMovingAntsInPaths(pathsOne, numAnts), movement.RemoveStartRoom(pathsOne)))
	ResultTwoOfTurns := movement.MergeTurnsOfPaths(movement.GenerateStepsOfAnts(movement.BeforeMovingAntsInPaths(pathsTwo, numAnts), movement.RemoveStartRoom(pathsTwo)))

	fmt.Println(movement.JoinStepsWithNewLine(movement.GetBestResult(ResultOneOfTurns, ResultTwoOfTurns)))
}
