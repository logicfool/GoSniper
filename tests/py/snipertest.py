from web3 import Web3
from abis import *
# import solc compiler
import json
import solcx as solc
contract_file = "contract.sol"
solc.install_solc('0.8.6')
res = solc.compile_source(     open(contract_file,"r").read(),
                     output_values=["abi", "bin"],     solc_version="0.6.12")



sniperabi = res['<stdin>:GoSniper']['abi']
sniperbytecode = res['<stdin>:GoSniper']['bin']


from web3.middleware import geth_poa_middleware

url = "http://127.0.0.1:8545"
# url = "https://data-seed-prebsc-1-s1.binance.org:8545"
w3 = Web3(Web3.HTTPProvider(url))
w3.middleware_onion.inject(geth_poa_middleware, layer=0)
pk = w3.eth.account.privateKeyToAccount("***REMOVED***")

# bal = w3.fromWei(w3.eth.getBalance(pk.address),"ether")
# if (bal<10):
#     tx_params = ***REMOVED***"value":w3.toWei("15","ether"),"from":w3.eth.accounts[0],"to":pk.address***REMOVED***
#     tx_hash = w3.eth.sendTransaction(tx_params)
    

# deploy a contract from bytecode

	


 
# print(w3.eth.chainId)
# exit()
if w3.eth.chainId == 1:
    token_contract1 = w3.eth.contract(address = "0xdAC17F958D2ee523a2206206994597C13D831ec7",abi=erc20abi)
    token_contract2 = w3.eth.contract(address = "0xf4d2888d29D722226FafA5d9B24F9164c092421E",abi=erc20abi)
    weth = w3.eth.contract(address = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",abi=erc20abi)
    buy_args_1  = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2","0xf4d2888d29D722226FafA5d9B24F9164c092421E"],1000000000000000000 #eth
    buy_args_2 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xdAC17F958D2ee523a2206206994597C13D831ec7","0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"]
    buy_args_3 = "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D",["0xf4d2888d29D722226FafA5d9B24F9164c092421E","0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"]
elif (w3.eth.chainId == 56) or (w3.eth.chainId == 31337):
    contract = w3.eth.contract(abi=sniperabi, bytecode=sniperbytecode)
    data11 = ***REMOVED******REMOVED***
    data11['abi'] = sniperabi
    data11["bytecode"] = sniperbytecode
    strr = json.dump(data11,open("sniperdata.json","w"))
    tx = contract.constructor("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0x0000000000004946c0e9F43F4Dee607b0eF1fA1c").buildTransaction(***REMOVED***
		'from': pk.address,
		'nonce': w3.eth.getTransactionCount(pk.address),
		'gas': 5000000,
		'gasPrice': w3.toWei('40', 'gwei')
	***REMOVED***)
    signed = pk.signTransaction(tx)
    tx_hash = w3.eth.sendRawTransaction(signed.rawTransaction)
    tx_receipt = w3.eth.waitForTransactionReceipt(tx_hash)
    if tx_receipt['status'] == 1:
        contract_address = tx_receipt['contractAddress']
        contract = w3.eth.contract(address=contract_address, abi=sniperabi)
        print(contract_address)
    else:
        print("Failed to deploy contract")
        exit()
    
    token_contract1 = w3.eth.contract(address = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",abi=erc20abi)
    token_contract2 = w3.eth.contract(address = "0x55d398326f99059fF775485246999027B3197955",abi=erc20abi)
    weth = w3.eth.contract(address = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",abi=erc20abi)
    buy_args_1 = ["0x10ED43C718714eb63d5aA57B78B54704E256024E","0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73","0x949aFf84b31B1c1D38ca1A365Fcc27A560028b23"],["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"],pk.address
    buy_args_2 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39"]
    buy_args_3 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"]
    
elif (w3.eth.chainId == 97):
    contract = w3.eth.contract(abi=sniperabi, bytecode=sniperbytecode)
    data11 = ***REMOVED******REMOVED***
    data11['abi'] = sniperabi
    data11["bytecode"] = sniperbytecode
    strr = json.dump(data11,open("sniperdata.json","w"))
    tx = contract.constructor("0xae13d989daC2f0dEbFf460aC112a837C89BAa7cd","0x0000000000000000000000000000000000000000").buildTransaction(***REMOVED***
		'from': pk.address,
		'nonce': w3.eth.getTransactionCount(pk.address),
		'gas': 5000000,
		'gasPrice': w3.toWei('40', 'gwei')
	***REMOVED***)
    signed = pk.signTransaction(tx)
    tx_hash = w3.eth.sendRawTransaction(signed.rawTransaction)
    tx_receipt = w3.eth.waitForTransactionReceipt(tx_hash)
    if tx_receipt['status'] == 1:
        contract_address = tx_receipt['contractAddress']
        contract = w3.eth.contract(address=contract_address, abi=sniperabi)
        print(contract_address)
    else:
        print("Failed to deploy contract")
        exit()
    
    token_contract1 = w3.eth.contract(address = "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",abi=erc20abi)
    token_contract2 = w3.eth.contract(address = "0x55d398326f99059fF775485246999027B3197955",abi=erc20abi)
    weth = w3.eth.contract(address = "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",abi=erc20abi)
    buy_args_1 = ["0x10ED43C718714eb63d5aA57B78B54704E256024E","0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73","0x949aFf84b31B1c1D38ca1A365Fcc27A560028b23"],["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"],pk.address
    buy_args_2 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0x2b591e99afE9f32eAA6214f7B7629768c40Eeb39"]
    buy_args_3 = "0x10ED43C718714eb63d5aA57B78B54704E256024E",["0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82","0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"]
    
elif w3.eth.chainId == 137:
    token_contract1 = w3.eth.contract(address = "0x831753DD7087CaC61aB5644b308642cc1c33Dc13",abi=erc20abi)
    weth = w3.eth.contract(address = "0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",abi=erc20abi)
    buy_args_1 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270","0x831753DD7087CaC61aB5644b308642cc1c33Dc13"]
    buy_args_2 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x831753DD7087CaC61aB5644b308642cc1c33Dc13","0x55d398326f99059fF775485246999027B3197955"]
    buy_args_3 = "0xa5E0829CaCEd8fFDD4De3c43696c57F7D7A678ff",["0x831753DD7087CaC61aB5644b308642cc1c33Dc13","0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270"]
    
# try:
#     tx = contract.functions.TradeDirectlyByPairWithETH(*buy_args_1).buildTransaction(***REMOVED***"value":1000000000000000000,"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000***REMOVED***)
#     signed_tx = pk.sign_transaction(tx)
#     txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
#     tx_receipt = w3.eth.waitForTransactionReceipt(txx)
#     if tx_receipt['status'] == 1:
#     	print("Buy Tx 1 Succeeded: ", w3.toHex(txx))
# except Exception as e:
#     print("Buy Tx 1 failed!",e)
#     exit();
    
	# txx = weth.functions.approve(contract.address,115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction(***REMOVED***"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000***REMOVED***)
	# signed_tx = pk.sign_transaction(txx)
	# txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
	# tx_receipt = w3.eth.waitForTransactionReceipt(txx)
	# if tx_receipt['status'] == 1:
	# 	print("Approve Tx Succeeded: ", w3.toHex(txx))

	# print("Token 1 contract: ",token_contract1.address)
	# tx = token_contract1.functions.approve(contract.address ,115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction(***REMOVED***"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000***REMOVED***)
	# signed_tx = pk.sign_transaction(tx)
	# txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
	# receipt = w3.eth.waitForTransactionReceipt(txx)
	# print("Approval: ",w3.toHex(txx))

	# # TestReserves = contract.functions.HoneyPotCheck("0x10ED43C718714eb63d5aA57B78B54704E256024E",["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56"],token_contract1.functions.balanceOf(pk.address).call(),1).call(***REMOVED***"from":pk.address,"value":1000000000000000000,"gasPrice":5000000000***REMOVED***)
	# # print("TestReserves: ",TestReserves)


	# majoraddresses = ["0x10ED43C718714eb63d5aA57B78B54704E256024E","0x000000000000000000000000000000000000dEaD","0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16"]
	# path =["0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56","0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"]

	# # HoneyCheck = contract.functions.HoneyPotCheck(majoraddresses,path,int(token_contract1.functions.balanceOf(pk.address).call()/2),1).call(***REMOVED***"from":pk.address,"gasPrice":50000000000***REMOVED***)
	# # print("HoneyCheck Using Pair: ",HoneyCheck)


	# path =["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c","0xF700D4c708C2be1463E355F337603183D20E0808"]
	# HoneyCheck = contract.functions.HoneyPotCheck(majoraddresses,path,0,1).call(***REMOVED***"from":pk.address,"gasPrice":50000000000,"value":1000000000000000000***REMOVED***)
	# print("HoneyCheck Using ETH: ",HoneyCheck)
















# pairs_to_get = [
# 	["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",w3.toChecksumAddress("0x26193c7fa4354ae49ec53ea2cebc513dc39a10aa")],
# 	[w3.toChecksumAddress("0x69b14e8d3cebfdd8196bfe530954a0c226e5008e")],
# 	[w3.toChecksumAddress("0x965f527d9159dce6288a2219db51fc6eef120dd1")],
# 	[w3.toChecksumAddress("0x1d871ed702FA8c220F03Ddf8e20DCf512aF74Ade")],
# 	[w3.toChecksumAddress("0x3D5629a4174D6522BA030ebc7727A319E02E4913")],
#  	["0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",w3.toChecksumAddress("0xf700d4c708c2be1463e355f337603183d20e0808")],
#   ["0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82"]
# ]
# stablecoin = "0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE"
# router = "0x10ED43C718714eb63d5aA57B78B54704E256024E"

# getprices = contract.functions.GetPrices(router,stablecoin,pairs_to_get).call(***REMOVED***"from":pk.address,"gasPrice":50000000000,"value":int(0.1*10**18)***REMOVED***)
# print("GetPrices: ",getprices) 
# bnbvalue= getprices[0]/10**18
# print(f"BNB price: ***REMOVED***bnbvalue***REMOVED***")
# for ii in range(0,len(pairs_to_get)):
#     price,decimal =getprices[1][ii] 
#     price = price/10**decimal
#     price = bnbvalue/price
#     print(f"Price Of ***REMOVED***pairs_to_get[ii]***REMOVED*** is ***REMOVED***price***REMOVED***,Amount: ***REMOVED***getprices[1][ii][0]***REMOVED***")
















# TradingCheck = contract.functions.TradeDirectlyByPair(majoraddresses,path,int(token_contract1.functions.balanceOf(pk.address).call()/2),pk.address,0,1).call(***REMOVED***"from":pk.address,"gasPrice":50000000000***REMOVED***)
# print("TradingCheck: ",TradingCheck)


# Tradecheck = contract.functions.directSwapBuyAndSell()

# tx = token_contract1.functions.approve("0x10ED43C718714eb63d5aA57B78B54704E256024E",115792089237316195423570985008687907853269984665640564039457584007913129639935).buildTransaction(***REMOVED***"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000***REMOVED***)
# signed_tx = pk.sign_transaction(tx)
# txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
# receipt = w3.eth.waitForTransactionReceipt(txx)
# print("Approval WETH: ",w3.toHex(txx))

# try:
#     tx = contract.functions.RouterSwapBySplit(*buy_args_3,[0,token_contract1.functions.balanceOf(pk.address).call()],pk.address,1).buildTransaction(***REMOVED***"value":0,"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":6000000000***REMOVED***)
#     signed_tx = pk.sign_transaction(tx)
#     txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
#     tx_receipt = w3.eth.waitForTransactionReceipt(txx)
#     if tx_receipt['status'] == 1:
#     	print("Buy Tx 3 Succeeded: ", w3.toHex(txx))
# except Exception as e:
#     print("Buy Tx 3 failed!",e)
#     exit();
    

# hp_check = contract.functions.HoneyPotCheck(*buy_args_1[:2],1).call(***REMOVED***"from":pk.address,"value":1000000000000000000***REMOVED***)

# tx = contract.functions.GetBalance("0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c").call()

# tx = weth.functions.balanceOf(pk.address).call(),weth.functions.balanceOf(contract_address).call(),w3.eth.getBalance(pk.address)
# print("Before: ",tx)

# tx = weth.functions.withdraw(tx[0]).buildTransaction(***REMOVED***"from":pk.address,"nonce":w3.eth.getTransactionCount(pk.address),"gasPrice":5000000000***REMOVED***)
# signed_tx = pk.sign_transaction(tx)
# txx = w3.eth.sendRawTransaction(signed_tx.rawTransaction) 

# tx = weth.functions.balanceOf(pk.address).call(),weth.functions.balanceOf(contract_address).call(), w3.eth.getBalance(pk.address)
# print("After: ",tx)
