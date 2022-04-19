from copyreg import constructor
from web3 import Web3
import json



config = json.load(open('config.json'))
condata = json.load(open('sniperdata.json'))
networks = config['networks']
to_chose = n = 0 
print("Please Choose a network: ")
for ii in networks:
    name = ii['name']
    print(n+1,". ",name)
num = int(input("Enter the network number-> "))-1
choose_network = networks[num]
pks = config['private_keys']
if len(pks) == 0:
    print("Please add a private key in your config to deploy a config")
    exit()
url = choose_network['node_url']
if url.startswith('http'):
    w3 = Web3(Web3.HTTPProvider(url))
elif url.startswith('ws'):
    w3 = Web3(Web3.WebsocketProvider(url))
else:
    #  ipc
    w3 = Web3(Web3.IPCProvider(url))
wallet = w3.eth.account.privateKeyToAccount(pks[0])
contract = w3.eth.contract(abi=condata['abi'], bytecode=condata['bytecode'])
weth = input("Enter WETH/WBNB/WMATIC et... address of the chain: ")
chigas = input("Enter CHIGAS address of the chain if none press enter: ")

try:
    weth = w3.toChecksumAddress(weth)
except:
    print("Invalid WETH address")
    exit()

if chigas == "":
    # 0x0000 address
    chigas = "0x0000000000000000000000000000000000000000"
else:
    try:
        chigas = w3.toChecksumAddress(chigas)
    except:
        print("Invalid CHIGAS address")
        exit()
        
constructed_tx = contract.constructor(weth,chigas)
tx_params = ***REMOVED***
    "from": wallet.address,
    "nonce": w3.eth.getTransactionCount(wallet.address),
***REMOVED***
tx = constructed_tx.buildTransaction(tx_params)
estimategas = w3.eth.estimateGas(tx)
tx['gas'] = estimategas
signed_tx = w3.eth.account.signTransaction(tx, private_key=wallet.privateKey)
tx_hash = w3.eth.sendRawTransaction(signed_tx.rawTransaction)
tx_receipt = w3.eth.waitForTransactionReceipt(tx_hash)
if tx_receipt['status'] == 1:
        contract_address = tx_receipt['contractAddress']
        print("Deployed Contract Address: ",contract_address)
else:
    print("Deployment Failed!")