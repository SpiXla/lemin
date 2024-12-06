package functions

import input "lemin/Input"

func SortPaths(paths *[][]string) {
	for i := 0; i < len(*paths); i++ {
		for j := i + 1; j < len(*paths); j++ {
			if len((*paths)[i]) > len((*paths)[j]) {
				(*paths)[i], (*paths)[j] = (*paths)[j], (*paths)[i]
			}
		}
	}
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
