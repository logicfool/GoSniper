from web3 import Web3
erc20abi = """[
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "owner",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Approval",
		"type": "event"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "from",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "to",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "value",
				"type": "uint256"
			}
		],
		"name": "Transfer",
		"type": "event"
	},
 {"constant":false,"inputs":[{"name":"wad","type":"uint256"}],"name":"withdraw","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "allowance",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "spender",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "approve",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "",
				"type": "address"
			}
		],
		"name": "balanceOf",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "burn",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "decimals",
		"outputs": [
			{
				"internalType": "uint8",
				"name": "",
				"type": "uint8"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "mint",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "name",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "symbol",
		"outputs": [
			{
				"internalType": "string",
				"name": "",
				"type": "string"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"name": "totalSupply",
		"outputs": [
			{
				"internalType": "uint256",
				"name": "",
				"type": "uint256"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "transfer",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "address",
				"name": "sender",
				"type": "address"
			},
			{
				"internalType": "address",
				"name": "recipient",
				"type": "address"
			},
			{
				"internalType": "uint256",
				"name": "amount",
				"type": "uint256"
			}
		],
		"name": "transferFrom",
		"outputs": [
			{
				"internalType": "bool",
				"name": "",
				"type": "bool"
			}
		],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]"""

# import solc compiler
import solcx as solc
contract_file = "contract.sol"
solc.install_solc('0.8.6')
res = solc.compile_source(     open(contract_file,"r").read(),
                     output_values=["abi", "bin"],     solc_version="0.8.6")



sniperabi = res['<stdin>:GoSniper']['abi']
# sniperbytecode = res['<stdin>:GoSniper']['bin']
from web3.middleware import geth_poa_middleware
rpc = "https://rpc-mainnet.maticvigil.com/v1/2e79e727831c8068242e336ef887ab50eaab44ee" #poly
rpc = "https://speedy-nodes-nyc.moralis.io/78fd8cc9150d7e2a80c27dc4/bsc/mainnet" #bsc
w3 = Web3(Web3.HTTPProvider(rpc))
w3.middleware_onion.inject(geth_poa_middleware, layer=0)
pk = w3.eth.account.privateKeyToAccount("")
 

if w3.eth.chainId == 1:
    token_contract1 = w3.eth.contract(address = "0xdAC17F958D2ee523a2206206994597C13D831ec7",abi=erc20abi)
    token_contract2 = w3.eth.contract(address = "0xf4d2888d29D722226FafA5d9B24F9164c092421E",abi=erc20abi)
    weth = w3.eth.contract(address = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",abi=erc20abi)
    buy_args_1  = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2","0xf4d2888d29D722226FafA5d9B24F9164c092421E"],1000000000000000000 #eth
    buy_args_2 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xdAC17F958D2ee523a2206206994597C13D831ec7","0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"]
    buy_args_3 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xf4d2888d29D722226FafA5d9B24F9164c092421E","0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"]
elif w3.eth.chainId == 56:
    contract_address = "0xEB835Ba1db2C46626E2A709b18F996c81B573CE9"
    token_contract1 = w3.eth.contract(address = "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82",abi=erc20abi)
    token_contract2 = w3.eth.contract(address = "0x55d398326f99059fF775485246999027B3197955",abi=erc20abi)
    weth = w3.eth.contract(address = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",abi=erc20abi)
    buy_args_1 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"],10000000000000000 #bsc
    buy_args_2 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39"]
    buy_args_3 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"]
elif w3.eth.chainId == 137:
    contract_address = "0x5DaE8719F7254bb7ec202FD164dce348ab40E20B"
    print("POLYGON DETECTED")
    token_contract1 = w3.eth.contract(address = "0x831753DD7087CaC61aB5644b308642cc1c33Dc13",abi=erc20abi)
    # chi_token = w3.eth.contract(address = "0x0000000000004946c0e9F43F4Dee607b0eF1fA1c",abi=erc20abi)
    weth = w3.eth.contract(address = "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",abi=erc20abi)
    buy_args_1 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270","0x831753DD7087CaC61aB5644b308642cc1c33Dc13"],1000000000000000 #eth
    # buy_args_2 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x831753DD7087CaC61aB5644b308642cc1c33Dc13","0x55d398326f99059fF775485246999027B3197955"]
    buy_args_3 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x831753DD7087CaC61aB5644b308642cc1c33Dc13","0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"]

chi_token = w3.eth.contract(address = "0x0000000000004946c0e9F43F4Dee607b0eF1fA1c",abi=erc20abi)
contract = w3.eth.contract(abi=sniperabi, address=contract_address)  #bytecode=sniperbytecode)
# check the allowance of chi token to the contract address

if chi_token.functions.allowance(pk.address,contract_address).call() < 1000000000000000000:
    print("CHI TOKEN ALLOWANCE IS NOT ENOUGH Approving now!")
    tx_hash = chi_token.functions.approve(contract_address,115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction({"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address)})
    signed_tx = pk.signTransaction(tx_hash)
    tx_hash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
    print("Chi allowance TX HASH:",tx_hash.hex())

exit();

try:
    tx = contract.functions.TradeDirectlyByPairWithGasRefund(*buy_args_1,3).buildTransaction({"value":1000000000000000,"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address)})
    signed_tx = pk.sign_transaction(tx)
    txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
    receipt = w3.eth.waitForTransactionReceipt(txx)
    print("Buy Tx 1 Succeeded: ", w3.toHex(txx))
except Exception as e:
    print("Buy Tx 1 failed!",e)
    exit();
        
tx = token_contract1.functions.approve(contract_address,115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction({"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address)})
signed_tx = pk.sign_transaction(tx)
txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
receipt = w3.eth.waitForTransactionReceipt(txx)
print("Approval: ",w3.toHex(txx))

tx = weth.functions.approve(contract_address,115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction({"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address)})
signed_tx = pk.sign_transaction(tx)
txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
receipt = w3.eth.waitForTransactionReceipt(txx)
print("Approval WETH: ",w3.toHex(txx))

try:
    tx = contract.functions.TradeDirectlyByPairWithGasRefund(*buy_args_3,token_contract1.functions.balanceOf(pk.address).call(),5).buildTransaction({"value":0,"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address)})
    # get tx data
    signed_tx = pk.sign_transaction(tx)
    txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
    receipt = w3.eth.waitForTransactionReceipt(txx)
    print("Buy Tx 3 Succeeded: ", w3.toHex(txx))
except Exception as e:
    print("Buy Tx 3 failed!",e)
    exit();
    

# hp_check = contract.functions.HoneyPotCheck(*buy_args_1[:2],5).call({"from":pk.address,"value":1000000000000000000})

# tx = contract.functions.GetBalance("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c").call()

# tx = weth.functions.balanceOf(pk.address).call(),weth.functions.balanceOf(contract_address).call(),w3.eth.getBalance(pk.address)
# print("Before: ",tx)

# tx = weth.functions.withdraw(tx[0]).buildTransaction({"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000})
# signed_tx = pk.sign_transaction(tx)
# txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction) 

# tx = weth.functions.balanceOf(pk.address).call(),weth.functions.balanceOf(contract_address).call(), w3.eth.getBalance(pk.address)
# print("After: ",tx)
