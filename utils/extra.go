package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func ColorPrint(message interface***REMOVED******REMOVED***, color ...interface***REMOVED******REMOVED***) ***REMOVED***
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	var color_str string
	var tolog bool
	color_str = color[0].(string)
	if len(color) > 1 ***REMOVED***
		tolog = color[1].(bool)
	***REMOVED*** else ***REMOVED***
		tolog = true
	***REMOVED***

	colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"
	colorPurple := "\033[35m"
	colorCyan := "\033[36m"
	colorWhite := "\033[37m"

	// Add time logs in message

	// variable to prevent printing on the new line and overwrite same line

	switch color_str ***REMOVED***
	case "red":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorRed, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorRed, message, colorReset)
		***REMOVED***
	case "green":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorGreen, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorGreen, message, colorReset)
		***REMOVED***
	case "yellow":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorYellow, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorYellow, message, colorReset)
		***REMOVED***
	case "blue":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorBlue, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorBlue, message, colorReset)
		***REMOVED***
	case "purple":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorPurple, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorPurple, message, colorReset)
		***REMOVED***
	case "cyan":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorCyan, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorCyan, message, colorReset)
		***REMOVED***
	case "white":
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorWhite, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorWhite, message, colorReset)
		***REMOVED***
	default:
		if tolog ***REMOVED***
			log.Printf("%s%s%s", colorWhite, message, colorReset)
		***REMOVED*** else ***REMOVED***
			fmt.Printf("%s%s%s", colorWhite, message, colorReset)
		***REMOVED***
	***REMOVED***
***REMOVED***

func ArrayContainsString(thearray []string, tofind string) bool ***REMOVED***
	sort.Strings(thearray)
	i := sort.SearchStrings(thearray, tofind)
	if i < len(thearray) && thearray[i] == tofind ***REMOVED***
		// fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
		return true
	***REMOVED***
	return false
***REMOVED***

func DecodeTxParams(abi abi.ABI, data interface***REMOVED******REMOVED***) (map[string]interface***REMOVED******REMOVED***, error) ***REMOVED***
	var data1 []byte
	var err error
	var v map[string]interface***REMOVED******REMOVED*** = make(map[string]interface***REMOVED******REMOVED***)
	switch data.(type) ***REMOVED***
	case []byte:
		data1 = data.([]byte)
	case string:
		data1, err = hex.DecodeString(data.(string))
		if err != nil ***REMOVED***
			return nil, err
		***REMOVED***
	***REMOVED***
	m, err := abi.MethodById(data1[:4])
	if err != nil ***REMOVED***
		return map[string]interface***REMOVED******REMOVED******REMOVED******REMOVED***, err
	***REMOVED***
	if err := m.Inputs.UnpackIntoMap(v, data1[4:]); err != nil ***REMOVED***
		return map[string]interface***REMOVED******REMOVED******REMOVED******REMOVED***, err
	***REMOVED***
	return v, nil
***REMOVED***

func DecodeBytecodeinput(data interface***REMOVED******REMOVED***) []common.Address ***REMOVED***
	var inputbyte []byte
	var resultarray []common.Address
	switch data.(type) ***REMOVED***
	case []byte:
		// data = data.([]byte)
		inputbyte = data.([]byte)
	case string:
		// data, _ = hex.DecodeString(data.(string))
		inputbyte, _ = hex.DecodeString(data.(string))
	***REMOVED***
	// fmt.Println(inputbyte)
	var i int = 0
	lengthofslice := len(inputbyte)
	for ***REMOVED***

		ress := common.HexToAddress(hex.EncodeToString(inputbyte[i : i+32]))
		// fmt.Println(ress)
		resultarray = append(resultarray, ress)
		i = i + 32
		if i >= lengthofslice ***REMOVED***
			break
		***REMOVED***
	***REMOVED***
	return resultarray
***REMOVED***
