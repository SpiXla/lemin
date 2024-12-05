package functions

import (
    "errors"
    "fmt"

    input "lemin/Input"
)


func FindUniquePaths(rooms map[string]input.Room, connections map[string][]string) ([][]string, error) {
    start, end := FindStartEnd(rooms)
    // var paths [][]string
    visited := map[string]bool{}

    if start == "" || end == "" {
        return nil, errors.New("start or end room not defined")
    }


    shortestPath := [][]string{}

    // DepthFirstSearch([]string{start}, start, end, &paths, visited, connections)
    ConnWithoutStart := RemoveStart(connections, start)
    fmt.Println(connections[start])
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
        Groups := GroupsPaths(path, connections[start],ConnWithoutStart, end)
        AllGroups = append(AllGroups, Groups)
        // fmt.Println(Groups)
    }
    for _, group := range AllGroups {

        fmt.Println(group)
    }
    return shortestPath, nil
}

func GroupsPaths(path, RoomsLinkedWstart []string, connWithoutStart map[string][]string, end string) [][]string {
    localVisited := make(map[string]bool)
    Groups := [][]string{}

    for _, room := range path[1 : len(path)-1] {
        localVisited[room] = true
    }

    for _, pathStart := range RoomsLinkedWstart {
        if !localVisited[pathStart] {
            group := BfsGroups(connWithoutStart, pathStart, end, localVisited)
            if group != nil {
                Groups = append(Groups, group)
            }
        }
    }

    return Groups
}


func BfsGroups(graph map[string][]string, start, end string, visited map[string]bool) []string {
    queue := [][]string{{start}}
    localVisited := make(map[string]bool) 
    localVisited[start] = true

    for len(queue) > 0 {
        path := queue[0]
        queue = queue[1:]

        node := path[len(path)-1]
        if node == end {
            return path
        }
        for _, neighbor := range graph[node] {
            if !visited[neighbor] && !localVisited[neighbor] {
                localVisited[neighbor] = true
                newPath := append([]string{}, path...)
                newPath = append(newPath, neighbor)
                queue = append(queue, newPath)
            }
        }
    }
    return nil
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
    fmt.Println(link,roomBefore)
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

