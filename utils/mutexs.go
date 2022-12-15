package utils

import "sync"

var GlobalTxid int
var GlobalSuccCount int
var GlobalAbortCount int
var GlobalExecCount int
var TxidLock sync.Mutex
var TxNumLock sync.Mutex

func init() {
	GlobalTxid = 0
}

func GetTXID() int {
	id := 0
	TxidLock.Lock()
	GlobalTxid++
	id = GlobalTxid
	TxidLock.Unlock()
	return id
}

func UpdateFinishCount() {
	TxNumLock.Lock()
	GlobalSuccCount++
	GlobalExecCount++
	TxNumLock.Unlock()
}

func UpdateAbortCount() {
	TxNumLock.Lock()
	GlobalAbortCount++
	GlobalExecCount++
	TxNumLock.Unlock()
}


func GetFinishTxCount() (int, int, int) {
	var doneNum, abortNum, execNum int
	TxNumLock.Lock()
	doneNum = GlobalSuccCount
	abortNum = GlobalAbortCount
	execNum = GlobalExecCount
	TxNumLock.Unlock()
	return doneNum, abortNum, execNum
}