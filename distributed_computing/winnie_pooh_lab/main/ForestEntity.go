package main

type ForestRegion struct {
	regionId int
	isWinniePoohHere bool
}

type Forest struct {
	forestRegions []ForestRegion
	forestName string
}

func (forest *Forest) initializeForestRegions(regionsNumber int, forestName string, winniePosition int) {
	forest.forestRegions = make([]ForestRegion, regionsNumber)
	forest.forestName = forestName
	forest.forestRegions[winniePosition].isWinniePoohHere = true

	for index := range forest.forestRegions {
		forest.forestRegions[index].regionId = index
	}
}

func (forest Forest) String() string {
	return forest.forestName
}
