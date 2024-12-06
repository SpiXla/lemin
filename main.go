package main

import (
	"fmt"
	"os"

	input "lemin/Input"
	"lemin/functions"
	"lemin/movement"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	fileContent, numAnts, rooms, connections, err := input.ParseInput(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	AllGroups, err := functions.PathFindingLogic(rooms, connections)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	ShortestTurn := movement.BesGroup(AllGroups,numAnts)

	fmt.Println(fileContent)
	fmt.Println(movement.JoinStepsWithNewLine(ShortestTurn))
}
