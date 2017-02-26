package main

import (
	"time"
	"math/rand"
	"fmt"
	"strconv"
)

type Officer struct {
	name string
}

func handleInnermostOfficer(goodsHolder GoodsHolder, officer Officer, middleChan chan int) {
	for _, value := range goodsHolder.goods {
		time.Sleep(time.Duration(rand.Intn(3e3)) * time.Millisecond)
		fmt.Printf("Innermost %v done with good #%v \n", officer.name, value)
		middleChan <- value
	}
	close(middleChan)
}

func handleMiddleOfficer(officer Officer, middleChan chan int, outerChan chan int) {
	for {
		v, ok := <- middleChan
		if ok {
			fmt.Printf("Middle officer %v is here \n", officer.name)
			outerChan <- v
		} else {
			close(outerChan)
			return
		}
	}
}

func handleOuterOfficer(officer Officer, outerChan chan int, reportChan chan string) {
	for {
		v, ok := <- outerChan
		if ok {
			fmt.Printf("Outer officer %v is here\n", officer.name)
			reportChan <- "Sequence " + strconv.Itoa(v) + " is done"
		} else {
			close(reportChan)
			return
		}
	}
}
