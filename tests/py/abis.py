erc20abi = """[
	***REMOVED***
		"anonymous": false,
		"inputs": [
			***REMOVED***
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			***REMOVED***
		],
		"name": "Approval",
		"type": "event"
	***REMOVED***,
	***REMOVED***
		"anonymous": false,
		"inputs": [
			***REMOVED***
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			***REMOVED***
		],
		"name": "Transfer",
		"type": "event"
	***REMOVED***,
 ***REMOVED***"constant":false,"inputs":[***REMOVED***"name":"wad","type":"uint256"***REMOVED***],"name":"withdraw","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "address",
				"name": "",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"internalType": "address",
				"name": "",
				"type": "address"
			***REMOVED***
		],
		"name": "allowance",
		"outputs": [
			***REMOVED***
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "address",
				"name": "spender",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			***REMOVED***
		],
		"name": "approve",
		"outputs": [
			***REMOVED***
				"internalType": "bool",
				"name": "",
				"type": "bool"
			***REMOVED***
		],
		"stateMutability": "nonpayable",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "address",
				"name": "",
				"type": "address"
			***REMOVED***
		],
		"name": "balanceOf",
		"outputs": [
			***REMOVED***
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			***REMOVED***
		],
		"name": "burn",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [],
		"name": "decimals",
		"outputs": [
			***REMOVED***
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			***REMOVED***
		],
		"name": "mint",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [],
		"name": "name",
		"outputs": [
			***REMOVED***
				"internalType": "string",
				"name": "",
				"type": "string"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [],
		"name": "symbol",
		"outputs": [
			***REMOVED***
				"internalType": "string",
				"name": "",
				"type": "string"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			***REMOVED***
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			***REMOVED***
		],
		"stateMutability": "view",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			***REMOVED***
		],
		"name": "transfer",
		"outputs": [
			***REMOVED***
				"internalType": "bool",
				"name": "",
				"type": "bool"
			***REMOVED***
		],
		"stateMutability": "nonpayable",
		"type": "function"
	***REMOVED***,
	***REMOVED***
		"inputs": [
			***REMOVED***
				"internalType": "address",
				"name": "sender",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			***REMOVED***,
			***REMOVED***
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			***REMOVED***
		],
		"name": "transferFrom",
		"outputs": [
			***REMOVED***
				"internalType": "bool",
				"name": "",
				"type": "bool"
			***REMOVED***
		],
		"stateMutability": "nonpayable",
		"type": "function"
	***REMOVED***
]"""


pair_abi = """[***REMOVED***"inputs":[],"payable":false,"stateMutability":"nonpayable","type":"constructor"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"owner","type":"address"***REMOVED***,***REMOVED***"indexed":true,"internalType":"address","name":"spender","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"value","type":"uint256"***REMOVED***],"name":"Approval","type":"event"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"sender","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"***REMOVED***,***REMOVED***"indexed":true,"internalType":"address","name":"to","type":"address"***REMOVED***],"name":"Burn","type":"event"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"sender","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount0","type":"uint256"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount1","type":"uint256"***REMOVED***],"name":"Mint","type":"event"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"sender","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount0In","type":"uint256"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount1In","type":"uint256"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount0Out","type":"uint256"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"amount1Out","type":"uint256"***REMOVED***,***REMOVED***"indexed":true,"internalType":"address","name":"to","type":"address"***REMOVED***],"name":"Swap","type":"event"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":false,"internalType":"uint112","name":"reserve0","type":"uint112"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint112","name":"reserve1","type":"uint112"***REMOVED***],"name":"Sync","type":"event"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"from","type":"address"***REMOVED***,***REMOVED***"indexed":true,"internalType":"address","name":"to","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"value","type":"uint256"***REMOVED***],"name":"Transfer","type":"event"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"DOMAIN_SEPARATOR","outputs":[***REMOVED***"internalType":"bytes32","name":"","type":"bytes32"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"MINIMUM_LIQUIDITY","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"PERMIT_TYPEHASH","outputs":[***REMOVED***"internalType":"bytes32","name":"","type":"bytes32"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"name":"allowance","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"spender","type":"address"***REMOVED***,***REMOVED***"internalType":"uint256","name":"value","type":"uint256"***REMOVED***],"name":"approve","outputs":[***REMOVED***"internalType":"bool","name":"","type":"bool"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"name":"balanceOf","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***],"name":"burn","outputs":[***REMOVED***"internalType":"uint256","name":"amount0","type":"uint256"***REMOVED***,***REMOVED***"internalType":"uint256","name":"amount1","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"decimals","outputs":[***REMOVED***"internalType":"uint8","name":"","type":"uint8"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"factory","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"getReserves","outputs":[***REMOVED***"internalType":"uint112","name":"_reserve0","type":"uint112"***REMOVED***,***REMOVED***"internalType":"uint112","name":"_reserve1","type":"uint112"***REMOVED***,***REMOVED***"internalType":"uint32","name":"_blockTimestampLast","type":"uint32"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"_token0","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"_token1","type":"address"***REMOVED***],"name":"initialize","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"kLast","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***],"name":"mint","outputs":[***REMOVED***"internalType":"uint256","name":"liquidity","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"name","outputs":[***REMOVED***"internalType":"string","name":"","type":"string"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"name":"nonces","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"owner","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"spender","type":"address"***REMOVED***,***REMOVED***"internalType":"uint256","name":"value","type":"uint256"***REMOVED***,***REMOVED***"internalType":"uint256","name":"deadline","type":"uint256"***REMOVED***,***REMOVED***"internalType":"uint8","name":"v","type":"uint8"***REMOVED***,***REMOVED***"internalType":"bytes32","name":"r","type":"bytes32"***REMOVED***,***REMOVED***"internalType":"bytes32","name":"s","type":"bytes32"***REMOVED***],"name":"permit","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"price0CumulativeLast","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"price1CumulativeLast","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***],"name":"skim","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"uint256","name":"amount0Out","type":"uint256"***REMOVED***,***REMOVED***"internalType":"uint256","name":"amount1Out","type":"uint256"***REMOVED***,***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***,***REMOVED***"internalType":"bytes","name":"data","type":"bytes"***REMOVED***],"name":"swap","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"symbol","outputs":[***REMOVED***"internalType":"string","name":"","type":"string"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[],"name":"sync","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"token0","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"token1","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"totalSupply","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***,***REMOVED***"internalType":"uint256","name":"value","type":"uint256"***REMOVED***],"name":"transfer","outputs":[***REMOVED***"internalType":"bool","name":"","type":"bool"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"from","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"to","type":"address"***REMOVED***,***REMOVED***"internalType":"uint256","name":"value","type":"uint256"***REMOVED***],"name":"transferFrom","outputs":[***REMOVED***"internalType":"bool","name":"","type":"bool"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***]"""