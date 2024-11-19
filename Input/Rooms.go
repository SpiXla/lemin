package input

import (
	"errors"
	"strings"
)

func RoomParams(line string) (string, string, string, error) {
	info := strings.Fields(line)
	if len(info) != 3 {
		return "", "", "", errors.New("wrong room data")
	}
	name := info[0]
	if seename[name] {
		return "", "", "", errors.New("wrong name")
	} else if !seename[name] {
		seename[name] = true
	}
	coor := strings.Join(info[1:], " ")
	if seencor[coor] {
		return "", "", "", errors.New("wrong coordinates")
	} else if !seencor[coor] {
		seencor[coor] = true
	}
	x := info[1]
	y := info[2]
	return name, x, y, nil
}
