package utils

import "fmt"

func SimplyError(funcName string, err error) {
	fmt.Println("running " + funcName + " failed, errInfo: " + err.Error())
}

func ShowStepInfo(info string) {
	fmt.Println("========================= " + info + " =========================")
}