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
factory_abi = '[{"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"token0","type":"address"},{"indexed":true,"internalType":"address","name":"token1","type":"address"},{"indexed":false,"internalType":"address","name":"pair","type":"address"},{"indexed":false,"internalType":"uint256","name":"","type":"uint256"}],"name":"PairCreated","type":"event"},{"constant":true,"inputs":[],"name":"INIT_CODE_PAIR_HASH","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"allPairs","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"allPairsLength","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"tokenA","type":"address"},{"internalType":"address","name":"tokenB","type":"address"}],"name":"createPair","outputs":[{"internalType":"address","name":"pair","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"feeToSetter","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"address","name":"","type":"address"}],"name":"getPair","outputs":[{"internalType":"address","name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeTo","type":"address"}],"name":"setFeeTo","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"internalType":"address","name":"_feeToSetter","type":"address"}],"name":"setFeeToSetter","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"}]'
factory_address = '0xcA143Ce32Fe78f1f7019d7d551a6402fC5350c73'
random_address = '0x0000000000000000000000000000000000000000'
ca = w3.eth.contract(random_address,abi=sniperabi)
factory_ca = w3.eth.contract(factory_address,abi=factory_abi)
def get(start, end):
    encoded_data = ca.functions.Read(factory_address, start,end).buildTransaction({'gas' : 3000000, 'gasPrice' : 5})['data']
    params = {
        'to' : ca.address,
        'data' : encoded_data,
    }
    resp = w3.provider.make_request(
        'eth_call',
        [params, 'latest', {str(ca.address) : {"code" : "0x" + sniperbytecode}}]
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
        print(f"{pair=} {token0=} {token1=}")

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