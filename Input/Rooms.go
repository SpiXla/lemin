package input

import (
	"errors"
	"strconv"
	"strings"
)

var (
	seencor = make(map[string]bool)
	seename = make(map[string]bool)
)

func RoomParams(line string) (string, int, int, error) {
	info := strings.Fields(line)
	if len(info) != 3 {
		return "", 0, 0, errors.New("invalid data format, wrong room data")
	}
	name := info[0]
	if seename[name] || strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") {
		return "", 0, 0, errors.New("invalid data format, invalid room name")
	} else if !seename[name] {
		seename[name] = true
	}

	coor := strings.Join(info[1:], " ")
	if seencor[coor] {
		return "", 0, 0, errors.New("invalid data format, wrong coordinates")
	} else if !seencor[coor] {
		seencor[coor] = true
	}
	x, erX := strconv.Atoi(info[1])
	y, erY := strconv.Atoi(info[2])
	if erX != nil || erY != nil {
		return "", 0, 0, errors.New("invalid data format, wrong coordinates")
	}
	return name, x, y, nil
}
