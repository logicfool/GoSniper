package consts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/logicfool/GoSniper/utils"
)

var (
	ZeroAddress         = common.Address{}
	MAX_APPROVAL_HEX    = "0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	MAX_APPROVAL_STRING = "11579208923731619542357098500868790785326998466564056403945758400"
	MAX_APPROVAL_INT    = utils.StringToBigInt(MAX_APPROVAL_STRING)
	CHIGASTOKEN         = common.HexToAddress("0x0000000000004946c0e9F43F4Dee607b0eF1fA1c")
)

/*
//Incase if ever lost all the functions lol

var h [][]byte
	h = append(h, []byte("addLiquidityETH(address,uint256,uint256,uint256,address,uint256)"))
	h = append(h, []byte("addLiquidity(address,address,uint256,uint256,uint256,uint256,address,uint256)"))
	h = append(h, []byte("removeLiquidity(address,address,uint256,uint256,uint256,address,uint256)"))
	h = append(h, []byte("removeLiquidityETH(address,uint256,uint256,uint256,address,uint256)"))
	h = append(h, []byte("removeLiquidityWithPermit(address,address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)"))
	h = append(h, []byte("removeLiquidityETHWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)"))
	h = append(h, []byte("removeLiquidityETHSupportingFeeOnTransferTokens(address,uint256,uint256,uint256,address,uint256)"))
	h = append(h, []byte("removeLiquidityETHSupportingFeeOnTransferTokensWithPermit(address,uint256,uint256,uint256,address,uint256,bool,uint8,bytes32,bytes32)"))

	for _, v := range h {
		fmt.Printf("%v: %v \n", string(v), utils.EncodeFunctionName(v))
	}


*/
