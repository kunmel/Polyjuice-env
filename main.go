package main

import "C"
import (
	"Polyjuice-env/config"
	"Polyjuice-env/ethConnecter"
	"Polyjuice-env/transaction"
	"Polyjuice-env/utils"
	"fmt"
)


func main() {
	for i := 1; i <= config.THREAD_COUNT; i++ {
		transaction.MicroTxFileBuilder(100, i)
	}
	utils.ShowStepInfo("All TxFile built")
	ethConnecter.InitETH()
	fmt.Println(ethConnecter.QueryETH("4500"))
	fmt.Println(ethConnecter.QueryETH("32000"))
	fmt.Println(ethConnecter.QueryETH("2-250"))
	utils.ShowStepInfo("All KV initialized in ETH")
}

