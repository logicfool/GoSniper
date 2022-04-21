package main

import (
	"fmt"
	"os"

	"github.com/logicfool/GoSniper/sniper"
)

type Exit struct{ Code int }

func handleExit() {
	fmt.Print("Press enter to exit")
	if e := recover(); e != nil {
		if exit, ok := e.(Exit); ok == true {
			os.Exit(exit.Code)
		}
		panic(e) // not an Exit, bubble up
	}
}
func main() {
	Sniper := sniper.QuickMaintainConfig()
	Sniper.Start()
}
