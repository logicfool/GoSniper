from web3.middleware import geth_poa_middleware
from web3 import Web3
from abis import *
w3 = Web3(Web3.HTTPProvider("https://bscrpc.com/"))
w3.middleware_onion.inject(geth_poa_middleware, layer=0)
pk = w3.eth.account.privateKeyToAccount("***REMOVED***")

# bal = w3.fromWei(w3.eth.getBalance(pk.address),"ether")
# if (bal<10):
#     tx_params = {"value":w3.toWei("15","ether"),"from":w3.eth.accounts[0],"to":pk.address}
#     tx_hash = w3.eth.sendTransaction(tx_params)

erc20 = w3.eth.contract(abi=erc20abi, address="0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56")
pair_contract = w3.eth.contract(abi=pair_abi, address="0x804678fa97d91B974ec2af3c843270886528a9E6")

# In pair storage slot 6 and 7 are token0 and token1
storage = w3.soliditySha3(['address','uint256'], [pk.address, 0])

w3.eth.getBalance

# tx = erc20.functions.balanceOf("0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56").buildTransaction({"from":pk.address})



first_data = {
    # "from":pk.address,
    "to":"0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56",
    "data":"0x0be5b6ba"
}


params = [
    first_data,
    "latest",
    {
        "0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56":{
            "code":"0x6080604052348015600f57600080fd5b506004361060285760003560e01c80630be5b6ba14602d575b600080fd5b60336045565b60408051918252519081900360200190f35b6007549056fea265627a7a723058206f26bd0433456354d8d1228d8fe524678a8aeeb0594851395bdbd35efc2a65f164736f6c634300050a0032"
        },
    }
]



import requests

# method = "eth_call"
# payload= {"jsonrpc":"2.0",
#            "method":method,
#            "params":[{"data":"9610cbcb0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000300000000000000000000000010ed43c718714eb63d5aa57b78b54704e256024e000000000000000000000000ca143ce32fe78f1f7019d7d551a6402fc5350c73000000000000000000000000737046ef759691fbc778cd59fdec77b1534c39a50000000000000000000000000000000000000000000000000000000000000002000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c00000000000000000000000026193c7fa4354ae49ec53ea2cebc513dc39a10aa","from":"0xabd3e1274a93a09bea79d6e8be9259eaaf57e7dd","to":"0xfbe69211B4E6E234c8b2cC3F3b322B4c3c310dCD","value":100000000000000000},"latest",{"0xABd3E1274A93a09BEa79d6e8Be9259eAAF57E7dd":{"balance":"0x56BC75E2D63100000"}}],
#            "id":1}
# headers = {'Content-type': 'application/json'}
# response = requests.post('http://127.0.0.1:8545', json=payload, headers=headers)
# print(response.json())

web3 = Web3(Web3.HTTPProvider('http://127.0.0.1:8545'))
abi = '[{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"description","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint80","name":"_roundId","type":"uint80"}],"name":"getRoundData","outputs":[{"internalType":"uint80","name":"roundId","type":"uint80"},{"internalType":"int256","name":"answer","type":"int256"},{"internalType":"uint256","name":"startedAt","type":"uint256"},{"internalType":"uint256","name":"updatedAt","type":"uint256"},{"internalType":"uint80","name":"answeredInRound","type":"uint80"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"latestRoundData","outputs":[{"internalType":"uint80","name":"roundId","type":"uint80"},{"internalType":"int256","name":"answer","type":"int256"},{"internalType":"uint256","name":"startedAt","type":"uint256"},{"internalType":"uint256","name":"updatedAt","type":"uint256"},{"internalType":"uint80","name":"answeredInRound","type":"uint80"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"version","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]'
addr = '0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE'
contract = web3.eth.contract(address=addr, abi=abi)
latestData = contract.functions.latestRoundData().call()
latestD = contract.functions.latestRoundData().buildTransaction({'from':'0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE'})
print(latestD)

payload = {
    "jsonrpc":"2.0",
    "method":"eth_call",
    "params":[{"data":"0xfeaf968c","from":"0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE","to":"0x0567F2323251f0Aab15c8dFb1967E4e8A7D42aeE"}],
    "id":1
}
headers = {'Content-type': 'application/json'}
response = requests.post('http://127.0.0.1:8545', json=payload, headers=headers)
print(response.json())