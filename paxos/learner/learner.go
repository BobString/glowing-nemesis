package learner

import (
	//"fmt"
	"strings"
)

var(
	learnChan = make(chan string,5)
	learnedValue int
	pairMap = make(map[Pair] int)
	nbProc int
)

type Pair struct {
	nv int
	val string
}

func EntryPoint (count int) (chan string) {
	nbProc = count
	go receivingMsgs()
	return learnChan
}

func receivingMsgs () {	
	for {
	mesg := <- learnChan
	res := strings.Split(mesg,"@")
	
	p := Pair {int(res[1]), int(res[2])}
	
	_,ok := pairMap [p]
	
	if ok {
		pairMap[p] = pairMap[p]+1
	} else {
		pairMap[p] = 1
	}
	
	if v,_ := pairMap[p]; v>(nbProc/2) {
		learnedValue = v 
	}
	}
}
