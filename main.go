package main

import (
	"fmt"
	"os"

	input "lemin/Input"
	"lemin/functions"
	"lemin/movement"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	fileContent, numAnts, rooms, connections, err := input.ParseInput(os.Args[1])
	_ = fileContent
	_ = numAnts
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	_, AllGroups, err := functions.FindUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	turnsOfgroups := [][][]string{}
	for _, group := range AllGroups {
		turnOfGroup := movement.MergeTurnsOfPaths(movement.GenerateStepsOfAnts(movement.BeforeMovingAntsInPaths(group, numAnts), group))
		turnsOfgroups = append(turnsOfgroups, turnOfGroup)
		// fmt.Println(group)
	}
	tol := len(turnsOfgroups[0])
	// fmt.Println(tol)
	for _,turns := range turnsOfgroups {
		if len(turns) < tol {
			tol = len(turns)
		}
		// fmt.Println("tri9:")
		// fmt.Println(len(r))
		// fmt.Println()
		// fmt.Println(r)
	}
	// fmt.Println(tol)
	ShortestTurn := [][]string{}
	for _,turns := range turnsOfgroups {
		if len(turns) == tol {
			ShortestTurn = append(ShortestTurn, turns...)
			break
		}
		// fmt.Println("tri9:")
		// fmt.Println(len(r))
		// fmt.Println()
		// fmt.Println(r)
	}
	// fmt.Println(turnsOfgroups)
	// start, _ := functions.FindStartEnd(rooms)
	// PathsWrooms := [][]string{}
	// for _, p := range paths {
	// 	path := []string{}
	// 	path = append(path, start)
	// 	path = append(path, p...)
	// 	PathsWrooms = append(PathsWrooms, path)
	// }
	// fmt.Println(PathsWrooms)
	// fmt.Println()
	fmt.Println(fileContent)

	fmt.Println(movement.JoinStepsWithNewLine(ShortestTurn))
}
