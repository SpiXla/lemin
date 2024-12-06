package functions

import (
	"errors"
	"fmt"

	input "lemin/Input"
)

func FindUniquePaths(rooms map[string]input.Room, connections map[string][]string) ([][]string,[][][]string, error) {
	start, end := FindStartEnd(rooms)
	// var paths [][]string 
	visited := map[string]bool{}

	if start == "" || end == "" {
		return nil,nil, errors.New("start or end room not defined")
	}

	shortestPath := [][]string{}

	// DepthFirstSearch([]string{start}, start, end, &paths, visited, connections)
	ConnWithoutStart := RemoveStart(connections, start)
	// fmt.Println(connections[start])
	for _, StartConnection := range connections[start] {
		path := BFS(ConnWithoutStart, StartConnection, end)
		shortestPath = append(shortestPath, path)
	}
	SortPaths(&shortestPath)
	for _, path := range shortestPath {
		p := path[:len(path)-1]
		for i, room := range p {
			if !visited[room] {
				visited[room] = true
			} else {
				roomBefore := p[i-1]
				if len(connections[roomBefore]) > 2 {
					connections = RemoveConn(connections, roomBefore, room)
					path := BFS(ConnWithoutStart, p[0], end)
					shortestPath = append(shortestPath, path)
				}
			}
		}
	}

	SortPaths(&shortestPath)
	AllGroups := [][][]string{}
	for _, path := range shortestPath {
		startConn := []string{}
		for _, room := range connections[start] {
			if room != path[0] {
				startConn = append(startConn, room)
			}
		}
		Groups := GroupsPaths(path, startConn, ConnWithoutStart, end)
		AllGroups = append(AllGroups, Groups)
		// fmt.Println(Groups)
	}
    for _, group := range AllGroups {
        roomSet := make(map[string]bool)
        for _, path := range group {
            for _, room := range path {
                if room != end {
                    if roomSet[room] {
                        fmt.Printf("Room %s is duplicated!\n", room)
                    }
                    roomSet[room] = true
                }
            }
        }
    }

	return shortestPath,AllGroups, nil
}

func GroupsPaths(path, roomsLinkedWStart []string, connWithoutStart map[string][]string, end string) [][]string {
    globalVisited := make(map[string]bool) // Global visited map
    groups := [][]string{}

    // Add the first path to groups and mark rooms as visited
    for _, room := range path {
        if room != end {
            globalVisited[room] = true
        }
    }
    groups = append(groups, path)

    // Process other connections starting from rooms linked to the start
    for _, pathStart := range roomsLinkedWStart {
        if !globalVisited[pathStart] {
            newGroup := BfsGroups(connWithoutStart, pathStart, end, &globalVisited)
            if len(newGroup) > 0 {
                groups = append(groups, newGroup...)
            }
        }
    }

    return groups
}

func BfsGroups(graph map[string][]string, start, end string, globalVisited *map[string]bool) [][]string {
    localVisited := make(map[string]bool)
    for room, isVisited := range *globalVisited {
        localVisited[room] = isVisited
    }

    queue := [][]string{{start}}
    groups := [][]string{}

    for len(queue) > 0 {
        currentPath := queue[0]
        queue = queue[1:]

        currentRoom := currentPath[len(currentPath)-1]

        if currentRoom == end {
            // Check if this path introduces any new duplicate rooms
            isDuplicate := false
            roomSet := make(map[string]bool)
            
            for _, room := range currentPath {
                if room != end && roomSet[room] {
                    isDuplicate = true
                    break
                }
                roomSet[room] = true
            }

            if !isDuplicate {
                groups = append(groups, currentPath)
            }
            continue
        }

        for _, neighbor := range graph[currentRoom] {
            // Strict check to prevent duplicates
            if !localVisited[neighbor] {
                newPath := append([]string{}, currentPath...) // Clone path
                newPath = append(newPath, neighbor)
                
                // Additional check for duplicate rooms
                roomSet := make(map[string]bool)
                hasDuplicate := false
                for _, room := range newPath {
                    if room != end && roomSet[room] {
                        hasDuplicate = true
                        break
                    }
                    roomSet[room] = true
                }

                if !hasDuplicate {
                    queue = append(queue, newPath)
                    localVisited[neighbor] = true
                }
            }
        }
    }

    // Update global visited map
    for _, path := range groups {
        for _, room := range path {
            if room != end {
                (*globalVisited)[room] = true
            }
        }
    }

    return groups
}




func BFS(graph map[string][]string, start, end string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]

		node := path[len(path)-1]
		if node == end {
			return path
		}
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return nil
}

func RemoveConn(connections map[string][]string, roomBefore, link string) map[string][]string {
	// fmt.Println(link, roomBefore)
	links := []string{}
	for _, r := range connections[roomBefore] {
		if r != link {
			links = append(links, r)
		}
	}
	connections[roomBefore] = links
	rooms := []string{}
	for _, r := range connections[link] {
		if r != roomBefore {
			rooms = append(rooms, r)
		}
	}
	connections[link] = rooms
	return connections
}

func RemoveStart(connections map[string][]string, start string) map[string][]string {
	// for i, path := range *paths {
	//     (*paths)[i] = path[1:]
	// }

	for key, value := range connections {
		if key == start {
			delete(connections, key)
		}
		var newSlice []string

		for _, room := range value {
			if room != start {
				newSlice = append(newSlice, room)
			}
		}
		connections[key] = newSlice

	}
	return connections
}

func FindStartEnd(rooms map[string]input.Room) (string, string) {
	var start, end string
	for name, room := range rooms {
		if room.IsStart {
			start = name
		} else if room.IsEnd {
			end = name
		}
	}
	return start, end
}
