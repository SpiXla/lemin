package input

import (
	"errors"
	"strings"
)

func parseTunnel(line string, connections map[string][]string) error {
	str := strings.Trim(line," ")
	if strings.Contains(str," ") {
		return errors.New("invalid data format, invalid links")
	}

	parts := strings.Split(str, "-")
	if len(parts) != 2 || parts[0] == parts[1] {
		return errors.New("invalid data format, invalid links")
	}
	if !seename[parts[0]] || !seename[parts[1]] {
		return errors.New("invalid data format, invalid links")

	}

	connections[parts[0]] = append(connections[parts[0]], parts[1])
	connections[parts[1]] = append(connections[parts[1]], parts[0])
	return nil
}
