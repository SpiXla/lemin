package input

import (
	"errors"
	"strings"
)

func parseTunnel(line string, connections map[string][]string) error {

	parts := strings.Split(line, "-")
	if len(parts) != 2 {
		return errors.New("invalid links")
	}
	if parts[0] == parts[1] {
		return errors.New("invalid links")
	}
	connections[parts[0]] = append(connections[parts[0]], parts[1])
	connections[parts[1]] = append(connections[parts[1]], parts[0])
	return nil
}
