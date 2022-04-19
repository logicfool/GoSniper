package main

import (
	"fmt"
	"os"

	"github.com/logicfool/GoSniper/sniper"
)

type Exit struct***REMOVED*** Code int ***REMOVED***

func handleExit() ***REMOVED***
	fmt.Print("Press enter to exit")
	if e := recover(); e != nil ***REMOVED***
		if exit, ok := e.(Exit); ok == true ***REMOVED***
			os.Exit(exit.Code)
		***REMOVED***
		panic(e) // not an Exit, bubble up
	***REMOVED***
***REMOVED***
func main() ***REMOVED***
	Sniper := sniper.QuickMaintainConfig()
	Sniper.Start()
***REMOVED***
