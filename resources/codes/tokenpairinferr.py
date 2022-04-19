from web3 import Web3
import solcx as solc
from eth_abi import decode_single
contract_file = "reader.sol"
from rich import print
from threading import Thread
solc.install_solc('0.8.12')
res = solc.compile_source(     open(contract_file,"r").read(),
                     output_values=["abi", "bin-runtime"],     solc_version="0.8.12")

# print(res.keys())
w3 = Web3(Web3.HTTPProvider("https://bscrpc.com"))
sniperabi = res['<stdin>:Reader']['abi']
sniperbytecode = res['<stdin>:Reader']['bin-runtime']
factory_abi = '[***REMOVED***"inputs":[***REMOVED***"internalType":"address","name":"_feeToSetter","type":"address"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"constructor"***REMOVED***,***REMOVED***"anonymous":false,"inputs":[***REMOVED***"indexed":true,"internalType":"address","name":"token0","type":"address"***REMOVED***,***REMOVED***"indexed":true,"internalType":"address","name":"token1","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"address","name":"pair","type":"address"***REMOVED***,***REMOVED***"indexed":false,"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"name":"PairCreated","type":"event"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"INIT_CODE_PAIR_HASH","outputs":[***REMOVED***"internalType":"bytes32","name":"","type":"bytes32"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"name":"allPairs","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"allPairsLength","outputs":[***REMOVED***"internalType":"uint256","name":"","type":"uint256"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"tokenA","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"tokenB","type":"address"***REMOVED***],"name":"createPair","outputs":[***REMOVED***"internalType":"address","name":"pair","type":"address"***REMOVED***],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"feeTo","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[],"name":"feeToSetter","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":true,"inputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***,***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"name":"getPair","outputs":[***REMOVED***"internalType":"address","name":"","type":"address"***REMOVED***],"payable":false,"stateMutability":"view","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"_feeTo","type":"address"***REMOVED***],"name":"setFeeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***,***REMOVED***"constant":false,"inputs":[***REMOVED***"internalType":"address","name":"_feeToSetter","type":"address"***REMOVED***],"name":"setFeeToSetter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"***REMOVED***]'
factory_address = '0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73'
random_address = '0x0000000000000000000000000000000000000000'
ca = w3.eth.contract(random_address,abi=sniperabi)
factory_ca = w3.eth.contract(factory_address,abi=factory_abi)
def get(start, end):
    encoded_data = ca.functions.Read(factory_address, start,end).buildTransaction(***REMOVED***'gas' : 3000000, 'gasPrice' : 5***REMOVED***)['data']
    params = ***REMOVED***
        'to' : ca.address,
        'data' : encoded_data,
***REMOVED***
    resp = w3.provider.make_request(
        'eth_call',
        [params, 'latest', ***REMOVED***str(ca.address) : ***REMOVED***"code" : "0x" + sniperbytecode***REMOVED******REMOVED***]
    )
    if 'error' in resp:
        print(resp['error'])
        return
    returned_data = resp.get('result')
    if returned_data is None:
        return

    data = (w3.codec.decode_abi(['address[][]'], bytes.fromhex(returned_data[2:])))
    for i in data[0]:
        pair, token0, token1 = i
        print(f"***REMOVED***pair=***REMOVED*** ***REMOVED***token0=***REMOVED*** ***REMOVED***token1=***REMOVED***")

# get(0, 2272)
sweet_spot = 2272
total = factory_ca.functions.allPairsLength().call()
total_threads = 50
for i in range(0, total, sweet_spot):
    t = Thread(target=get, args=(i, i+sweet_spot))
    t.start()
    total_threads -= 1
    if total_threads == 0:
        break