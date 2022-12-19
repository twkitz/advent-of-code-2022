package helper

type LavaSpector struct {
	droplets    map[[3]int]int
	minPosition [3]int
	maxPosition [3]int
}

func (ls *LavaSpector) UpdateDropletSide(x, y, z int) {
	if ls.droplets[[3]int{x, y, z}] == -1 {
		return
	}
	ls.droplets[[3]int{x, y, z}] += 1
}

func (ls *LavaSpector) AddDroplet(x, y, z int) {
	ls.droplets[[3]int{x, y, z}] = -1
	ls.UpdateDropletSide(x-1, y, z)
	ls.UpdateDropletSide(x+1, y, z)
	ls.UpdateDropletSide(x, y-1, z)
	ls.UpdateDropletSide(x, y+1, z)
	ls.UpdateDropletSide(x, y, z-1)
	ls.UpdateDropletSide(x, y, z+1)
}

func (ls LavaSpector) GetExposedDropletSide(shouldExcludeAirPocket bool) int {
	sum := 0
	for position, exposing := range ls.droplets {
		if exposing > 0 && (!shouldExcludeAirPocket || !ls.IsAirPocket(position)) {
			sum += exposing
		}
	}
	return sum
}

func (ls LavaSpector) IsAirPocket(position [3]int) bool {
	if ls.droplets[position] == 6 {
		return true
	}
	return ls.IsConnectedAirPocket(position)
}

func (ls *LavaSpector) IsConnectedAirPocket(position [3]int) bool {
	offsets := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	positions := [][3]int{position}
	checked := make(map[[3]int]bool)
	checked[position] = true
	for len(positions) > 0 {
		for _, offset := range offsets {
			adjacentX := positions[0][0] + offset[0]
			adjacentY := positions[0][1] + offset[1]
			adjacentZ := positions[0][2] + offset[2]
			adjacentPosition := [3]int{adjacentX, adjacentY, adjacentZ}

			if checked[adjacentPosition] {
				continue
			}

			exposing := ls.droplets[adjacentPosition]
			if adjacentX < ls.minPosition[0] || adjacentX > ls.maxPosition[0] ||
				adjacentY < ls.minPosition[1] || adjacentY > ls.maxPosition[1] ||
				adjacentZ < ls.minPosition[2] || adjacentZ > ls.maxPosition[2] {
				return false
			}

			if exposing > -1 {
				positions = append(positions, adjacentPosition)
				checked[adjacentPosition] = true
			}
		}
		positions = positions[1:]
	}
	return true
}
