package main

import (
	"fmt"
)

const FOREST_NAME = "Fangorn Forest"

var communicationChan = make(chan bool)
var unitsNumber int = getUnitsNumber()
var forest Forest = initForest()
var hive Hive = initHive()

func initHive() Hive {
	var hive Hive
	hive.initHive(unitsNumber, communicationChan)
	return hive
}

func initForest() Forest {
	poohPosition := getPoohPosition()
	var forest Forest
	forest.initializeForestRegions(unitsNumber, FOREST_NAME, poohPosition)
	return forest
}

func getUnitsNumber() int {
	var unitsNumber int
	fmt.Print("Enter regions number: ")
	fmt.Scan(&unitsNumber)

	return unitsNumber
}

func getPoohPosition() int {
	var poohPosition int
	fmt.Print("Enter Winnie Pooh position: ")
	fmt.Scan(&poohPosition)

	return poohPosition
}

func main() {
	startScanning()
	for i := 0; i < unitsNumber; i++ {
		fmt.Println(<-communicationChan)
	}
}

func startScanning() {
	for i := 0; i < unitsNumber; i++ {
		go forest.forestRegions[i].scan(hive.beeGroups[i])
	}
}
