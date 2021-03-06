package paxosMain
//@author Remy Pannaud
import (
    "connector"
    "paxos/acceptor"
    "paxos/learner"
    "paxos/proposer"
	"paxos/slotsManager"
)
// global variables :
var (
    acceptToPropChan = make(chan int, 5)
)
// Initialization function
//@parameters :
func EntryPoint(debug chan int) (chan int, chan string, chan string, chan string, chan string, chan string) {
    pMap := connector.GetProcesses()
    var p = make([]int, len(pMap))
    i := 0
    for v, _ := range pMap {
        p[i] = v
        i++
    }
    //
	slotsManager.EntryPoint()
    trustChan, promChan, valueChan, slotChan := proposer.EntryPoint(p,debug)
    prepChan, acceptChan := acceptor.EntryPoint(p)
    learnChan := learner.EntryPoint(len(p), slotChan)
    return trustChan, prepChan, promChan, acceptChan, learnChan, valueChan
}
