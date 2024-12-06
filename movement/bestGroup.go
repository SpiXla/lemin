package movement

func BesGroup(AllGroups [][][]string,numAnts int) [][]string {
	turnsOfgroups := [][][]string{}
	for _, group := range AllGroups {
		turnOfGroup := MergeTurnsOfPaths(GenerateStepsOfAnts(BeforeMovingAntsInPaths(group, numAnts), group))
		turnsOfgroups = append(turnsOfgroups, turnOfGroup)
	}
	tol := len(turnsOfgroups[0])
	for _, turns := range turnsOfgroups {
		if len(turns) < tol {
			tol = len(turns)
		}
	}
	ShortestTurn := [][]string{}
	for _, turns := range turnsOfgroups {
		if len(turns) == tol {
			ShortestTurn = append(ShortestTurn, turns...)
			break
		}

	}
	return ShortestTurn
}
