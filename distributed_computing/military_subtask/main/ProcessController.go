package main

import (
	"fmt"
)

var goodsNumber int = getGoodsNumber()
var goodsHolder GoodsHolder = initGoodsHolder()
var reportChannel chan string = make(chan string)

func initGoodsHolder() GoodsHolder {
	var goodsHolder GoodsHolder
	goodsHolder.initGoods(goodsNumber)
	return goodsHolder
}

func getGoodsNumber() int {
	var goodsNumber int
	fmt.Print("Enter goods number: ")
	fmt.Scan(&goodsNumber)

	return goodsNumber
}

func main() {
	startOperation()

	for  {
		v, ok := <- reportChannel
		if ok {
			fmt.Printf("Conroller has got a report: %v \n\n", v)
		} else {
			return
		}
	}
}
func startOperation() {
	middleChan := make(chan int)
	outerChan := make(chan int)

	go handleInnermostOfficer(goodsHolder, Officer{"Innermost"}, middleChan)
	go handleMiddleOfficer(Officer{"Middle"}, middleChan, outerChan)
	go handleOuterOfficer(Officer{"Outer"}, outerChan, reportChannel)
}

