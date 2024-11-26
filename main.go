package main

import (
	"fmt"
	"os"

	input "lemin/Input"
	"lemin/functions"
)

var linkstart, linkend = false, false

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	numAnts, rooms, connections, err := input.ParseInput(os.Args[1])
	_ = numAnts
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
	paths, err := functions.FindUniquePaths(rooms, connections)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}

	for _, path := range paths {
		fmt.Println(path)
	}

	// movement.Antmovement(numAnts, paths)
}
