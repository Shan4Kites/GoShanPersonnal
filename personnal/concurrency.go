package main

import (
	"fmt"
	"sync"
	"time"
)

type item struct {
	name int
}

func reachedUpstairs(item item) {
	fmt.Println("Item reached :" , item.name)
}

func walkAllFloorDirectly(item item) item {
	time.Sleep(time.Second*6)
	reachedUpstairs(item)
	return item
}

func walkAllFloorDirectly_wg(item item, wg *sync.WaitGroup) {
	time.Sleep(time.Second*6)
	reachedUpstairs(item)
	wg.Done()
}

func walkAllFloorDirectlyCh(item item, processedCh chan bool) {
	time.Sleep(time.Second*6)
	reachedUpstairs(item)
	processedCh <- true
}

func walkAllFloorDirectly_ch(itemCountLocal int) {

	items := []item{}
	for i:=0;i<itemCountLocal;i++ {
		items = append(items, item{name:i})
	}

	processedCh := make(chan bool, itemCountLocal)
	for _, item := range items {
		go walkAllFloorDirectlyCh(item, processedCh)
	}

	i := 0
	for _ = range processedCh {
		if i == itemCountLocal-1 {
			break
		}
		i++
	}
}

type internalfloorTask struct {
	name string
	myItemCh chan item
	nextFloorCh chan item
	timeD time.Duration
}

func (ft *internalfloorTask) init(capacity int, name string, timeD time.Duration, nextFloorCh chan item) {
	ft.name = name
	ft.myItemCh = make(chan item, capacity)
	ft.nextFloorCh = nextFloorCh
	ft.timeD = timeD
	go ft.processAndMoveToNextFloor()
}

func processItem(timeD time.Duration, nextFloorCh chan item, item item) {
	time.Sleep(timeD)
	nextFloorCh <- item
}

func (ft internalfloorTask) processAndMoveToNextFloor() {
	for item := range ft.myItemCh {
		//fmt.Println("Processing at " , ft.name, " Item: ", item.name)
		//time.Sleep(ft.timeD)
		//ft.nextFloorCh <- item
		go processItem(ft.timeD, ft.nextFloorCh, item)
	}
}


type finalfloorTask struct {
	myItemCh chan item
	timeD time.Duration
	taskCompletionCh chan bool
}

func (ft *finalfloorTask) init(capacity int, timeD time.Duration) {
	ft.myItemCh = make(chan item, capacity)
	ft.taskCompletionCh = make(chan bool, capacity)
	ft.timeD = timeD
	go ft.processIt()
}

func processItemInFinalFloor(timeD time.Duration, taskCompletionCh chan bool, item item) {
	time.Sleep(timeD)
	reachedUpstairs(item)
	taskCompletionCh <- true
}

func (ft finalfloorTask) processIt() {
	for item := range ft.myItemCh {
		go processItemInFinalFloor(ft.timeD, ft.taskCompletionCh, item)
	}
}

func walkStepByStep_ch(itemCountLocal int) {
	items := []item{}
	for i:=0;i<itemCountLocal;i++ {
		items = append(items, item{name:i})
	}

	tFloor := finalfloorTask{}
	tFloor.init(7000, time.Second*4)
	sFloor := internalfloorTask{}
	sFloor.init(5000, "secondF", time.Second, tFloor.myItemCh)
	fFloor := internalfloorTask{}
	fFloor.init(5000, "firstF", time.Second, sFloor.myItemCh)

	for _, item := range items {
		fFloor.myItemCh <- item
	}

	i := 0
	for _= range tFloor.taskCompletionCh {
		if i == itemCountLocal-1 {
			break
		}
		i++
	}
}


func main() {

	//wg := &sync.WaitGroup{}
	//start := time.Now()
	//
	//items := []item{}
	//for i:=0;i<itemCountLocal;i++ {
	//	items = append(items, item{name:i})
	//}
	//for _, item := range items {
	//	wg.Add(1)
	//	go walkAllFloorDirectly_wg(item, wg)
	//}
	//wg.Wait()
	//elapsed := time.Since(start)
	//fmt.Println("elapsed :", elapsed)
	//----------------

	itemCountLocal := 1050000
	//itemCountLocal := 5

	//items := []item{}
	//for i:=0;i<itemCountLocal;i++ {
	//	items = append(items, item{name:i})
	//}
	//start := time.Now()
	//for _, item := range items {
	//	walkAllFloorDirectly(item)
	//}
	//elapsed := time.Since(start)
	//fmt.Println("elapsed :", elapsed)




	//start := time.Now()
	//walkAllFloorDirectly_ch(itemCountLocal)
	//elapsed := time.Since(start)
	//fmt.Println("elapsed :", elapsed)

	start := time.Now()
	walkStepByStep_ch(itemCountLocal)
	elapsed := time.Since(start)
	fmt.Println("elapsed :", elapsed)

}
