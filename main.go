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

	pathsOne, err := functions.FindUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	start ,_ := functions.FindStartEnd(rooms)
	PathsWrooms := [][]string{}
	for _,r := range pathsOne {
		path := []string{}
		path = append(path, start)
		path = append(path, r...)
		PathsWrooms = append(PathsWrooms, path)
	}
	fmt.Println(PathsWrooms)
	fmt.Println()
	fmt.Println(fileContent)

	ResultOneOfTurns := movement.MergeTurnsOfPaths(movement.GenerateStepsOfAnts(movement.BeforeMovingAntsInPaths(pathsOne, numAnts), movement.RemoveStartRoom(pathsOne)))

	fmt.Println(movement.JoinStepsWithNewLine(ResultOneOfTurns))
}
