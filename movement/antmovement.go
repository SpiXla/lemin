package movement

import "fmt"



type Ant struct {
	Current string
	Num     int
	path    []string
}


func Antmovement(ants int, start, end string, paths [][]string) {
	Ants := make([]Ant, ants)

	for i := 1; i <= ants; i++ {
		Ants[i-1].Num = i
		Ants[i-1].path = paths[i%len(paths)]
		Ants[i-1].Current = Ants[i-1].path[0]
	}

	moves := []string{}

	for _, ant := range Ants {
		for _, p := range ant.path[1:] {
			num := ant.Num
			moves = append(moves, fmt.Sprintf("L%d-%v", num, p))
		}
	}
	fmt.Println(Ants)
	fmt.Println(moves)
}
