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
	paths1, paths2, err := functions.FindUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(fileContent)
	result1 := movement.MergeSteps(movement.GenerateSteps(movement.BeforeMoving(paths1, numAnts), movement.RemoveStartRoom(paths1)))
	result2 := movement.MergeSteps(movement.GenerateSteps(movement.BeforeMoving(paths2, numAnts), movement.RemoveStartRoom(paths2)))
	fmt.Println(movement.JoinStepsWithNewLine(movement.GetBestResult(result1, result2)))
}
