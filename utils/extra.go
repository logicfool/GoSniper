package utils

import (
	"encoding/hex"
	"fmt"
	"log"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func ColorPrint(message interface{}, color ...interface{}) {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	var color_str string
	var tolog bool
	color_str = color[0].(string)
	if len(color) > 1 {
		tolog = color[1].(bool)
	} else {
		tolog = true
	}

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

	switch color_str {
	case "red":
		if tolog {
			log.Printf("%s%s%s", colorRed, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorRed, message, colorReset)
		}
	case "green":
		if tolog {
			log.Printf("%s%s%s", colorGreen, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorGreen, message, colorReset)
		}
	case "yellow":
		if tolog {
			log.Printf("%s%s%s", colorYellow, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorYellow, message, colorReset)
		}
	case "blue":
		if tolog {
			log.Printf("%s%s%s", colorBlue, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorBlue, message, colorReset)
		}
	case "purple":
		if tolog {
			log.Printf("%s%s%s", colorPurple, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorPurple, message, colorReset)
		}
	case "cyan":
		if tolog {
			log.Printf("%s%s%s", colorCyan, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorCyan, message, colorReset)
		}
	case "white":
		if tolog {
			log.Printf("%s%s%s", colorWhite, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorWhite, message, colorReset)
		}
	default:
		if tolog {
			log.Printf("%s%s%s", colorWhite, message, colorReset)
		} else {
			fmt.Printf("%s%s%s", colorWhite, message, colorReset)
		}
	}
}

func ArrayContainsString(thearray []string, tofind string) bool {
	sort.Strings(thearray)
	i := sort.SearchStrings(thearray, tofind)
	if i < len(thearray) && thearray[i] == tofind {
		// fmt.Printf("found \"%s\" at files[%d]\n", files[i], i)
		return true
	}
	return false
}

func DecodeTxParams(abi abi.ABI, data interface{}) (map[string]interface{}, error) {
	var data1 []byte
	var err error
	var v map[string]interface{} = make(map[string]interface{})
	switch data.(type) {
	case []byte:
		data1 = data.([]byte)
	case string:
		data1, err = hex.DecodeString(data.(string))
		if err != nil {
			return nil, err
		}
	}
	m, err := abi.MethodById(data1[:4])
	if err != nil {
		return map[string]interface{}{}, err
	}
	if err := m.Inputs.UnpackIntoMap(v, data1[4:]); err != nil {
		return map[string]interface{}{}, err
	}
	return v, nil
}

func DecodeBytecodeinput(data interface{}) []common.Address {
	var inputbyte []byte
	var resultarray []common.Address
	switch data.(type) {
	case []byte:
		// data = data.([]byte)
		inputbyte = data.([]byte)
	case string:
		// data, _ = hex.DecodeString(data.(string))
		inputbyte, _ = hex.DecodeString(data.(string))
	}
	// fmt.Println(inputbyte)
	var i int = 0
	lengthofslice := len(inputbyte)
	for {

		ress := common.HexToAddress(hex.EncodeToString(inputbyte[i : i+32]))
		// fmt.Println(ress)
		resultarray = append(resultarray, ress)
		i = i + 32
		if i >= lengthofslice {
			break
		}
	}
	return resultarray
}
