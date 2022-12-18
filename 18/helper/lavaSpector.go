package helper

type LavaSpector struct {
	droplets map[[3]int]int
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

func (ls *LavaSpector) GetExposedDropletSide() int {
	sum := 0
	for _, exposing := range ls.droplets {
		if exposing > 0 {
			sum += exposing
		}
	}
	return sum
}
