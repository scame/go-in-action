package main

import (
	"time"
	"math/rand"
	"fmt"
)

type BeeGroup struct {
	groupId int
	groupChannel chan bool
}

type Hive struct {
	beeGroups []BeeGroup
}

func (hive *Hive) initHive(hiveSize int, groupChannel chan bool) {
	hive.beeGroups = make([]BeeGroup, hiveSize)
	for index := range hive.beeGroups {
		hive.beeGroups[index].groupId = index
		hive.beeGroups[index].groupChannel = groupChannel
	}
}

func (region * ForestRegion) scan(beeGroup BeeGroup) {
	fmt.Printf("Bee group %v has started\n", beeGroup.groupId)

	time.Sleep(time.Duration(rand.Intn(3e3)) * time.Millisecond)

	if region.isWinniePoohHere {
		beeGroup.groupChannel <- true
	} else {
		beeGroup.groupChannel <- false
	}
}
