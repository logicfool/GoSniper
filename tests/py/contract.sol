// SPDX-License-Identifier: Unlicensed 
pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;


interface IERC20 {
    event Approval(address indexed owner, address indexed spender, uint value);
    event Transfer(address indexed from, address indexed to, uint value);
    function name() external view returns (string memory);
    function symbol() external view returns (string memory);
    function decimals() external view returns (uint8);
    function totalSupply() external view returns (uint);
    function balanceOf(address owner) external view returns (uint);
    function allowance(address owner, address spender) external view returns (uint);
    function approve(address spender, uint value) external returns (bool);
    function transferFrom(address from, address to, uint value) external returns (bool);

    // WETH/IERC20 functions but not declaring another interface to probs lower gas consumption!
    function deposit() external payable;
    function transfer(address to, uint value) external returns (bool);
    function withdraw(uint) external;
    function latestRoundData() external view
    returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt,
        uint80 answeredInRound
    );


}


interface IRouter {
    function factory() external pure returns (address);
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)external payable returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, address[] calldata path, uint deadline)external payable returns (uint[] memory amounts);
    function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint amountIn,uint amountOutMin,address[] calldata path,address to,uint deadline) external payable returns (uint[] memory amounts);
    function swapExactETHForTokensSupportingFeeOnTransferTokens(uint amountOutMin,address[] calldata path,address to,uint deadline) external payable;
    function swapExactTokensForETHSupportingFeeOnTransferTokens(uint amountIn,uint amountOutMin,address[] calldata path,address to,uint deadline) external payable;
    }

interface IFactory {
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
}

interface ILiquidityPair {
    function factory() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function sync() external;
    function token0() external pure returns(address);
}

interface ChiGasToken{
    function free(uint256 value) external returns (uint256);
    function freeUpTo(uint256 value) external returns (uint256);
    function freeFrom(address from, uint256 value) external returns (uint256);
    function freeFromUpTo(address from, uint256 value) external returns (uint256);
}



library TransferHelper {
    function safeApprove(address token, address to, uint value) internal {
        // bytes4(keccak256(bytes('approve(address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x095ea7b3, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: APPROVE_FAILED');
    }

    function safeTransfer(address token, address to, uint value) internal {
        // bytes4(keccak256(bytes('transfer(address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0xa9059cbb, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: TRANSFER_FAILED');
    }

    function safeTransferFrom(address token, address from, address to, uint value) internal {
        // bytes4(keccak256(bytes('transferFrom(address,address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x23b872dd, from, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: TRANSFER_FROM_FAILED');
    }

    function safeTransferETH(address to, uint value) internal {
        (bool success,) = to.call{value:value}(new bytes(0));
        require(success, 'TransferHelper: ETH_TRANSFER_FAILED');
    }

}


library SafeMath {
    function add(uint x, uint y) internal pure returns (uint z) {
        require((z = x + y) >= x, 'ds-math-add-overflow');
    }

    function sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x, 'ds-math-sub-underflow');
    }

    function mul(uint x, uint y) internal pure returns (uint z) {
        require(y == 0 || (z = x * y) / y == x, 'ds-math-mul-overflow');
    }
}


library Strings{
    function uint2str(uint _i) internal pure returns (string memory _uintAsString) {
        if (_i == 0) {
            return "0";
        }
        uint j = _i;
        uint len;
        while (j != 0) {
            len++;
            j /= 10;
        }
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) {
            k = k-1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
        }
        return string(bstr);
    }
    function toAsciiString(address x) internal pure returns (string memory) {
    bytes memory s = new bytes(40);
    for (uint i = 0; i < 20; i++) {
        bytes1 b = bytes1(uint8(uint(uint160(x)) / (2**(8*(19 - i)))));
        bytes1 hi = bytes1(uint8(b) / 16);
        bytes1 lo = bytes1(uint8(b) - 16 * uint8(hi));
        s[2*i] = char(hi);
        s[2*i+1] = char(lo);            
    }
    return string(s);
}

function char(bytes1 b) internal pure returns (bytes1 c) {
    if (uint8(b) < 10) return bytes1(uint8(b) + 0x30);
    else return bytes1(uint8(b) + 0x57);
}
}

contract GoSniper {
    using SafeMath for uint;
    address public owner;
    ChiGasToken public chi;
    address public WETH ;
    mapping (address => bool) public Allowed;
    constructor (address _weth,address _chi) public {
        owner = msg.sender;
        WETH = _weth;
        chi = ChiGasToken(_chi);
        
    }

    modifier discountCHI {
        if (address(chi)!=address(0x0)){
            uint256 gasStart = gasleft();

        _;

        uint256 initialGas = 21000 + 16 * msg.data.length;
        uint256 gasSpent = initialGas + gasStart - gasleft();
        uint256 freeUpValue = (gasSpent + 14154) / 41947;

        chi.freeFromUpTo(msg.sender, freeUpValue);
        }
        
    }

    modifier onlyOwner() {
        require(owner == msg.sender, "Not Owner");
        _;
    }

    modifier onlyOwnerOrAllowed(address _addr) {
        require(owner == msg.sender || Allowed[_addr], "Not Owner or Allowed");
        _;
    }

    function setWETH(address _weth) onlyOwner public {
        WETH = _weth;
    }
    
    function setCHI(address _chi) onlyOwner public {
        chi = ChiGasToken(_chi);
    }
    

    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }
    
    function SetAllowed(address _addr,bool _res) onlyOwner public {
        Allowed[_addr] = _res;
    }

    function SetAllowedBulk(address[] memory _addrs,bool _res) onlyOwner public {
        for (uint i = 0; i < _addrs.length; i++) {
            Allowed[_addrs[i]] = _res;
        }
    }

    function sortTokens(address tokenA, address tokenB) public pure returns (address token0, address token1) {
        require(tokenA != tokenB, ' Error: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'Error: ZERO_ADDRESS');
    }

    function _swap(address[] memory path, address _pair,address _router,address _to,uint _amountoutmin) internal virtual {
            (address input, address output) = (path[0], path[1]);
            (address token0,) = sortTokens(input, output);
            ILiquidityPair pair = ILiquidityPair(_pair) ;
            uint amountInput;
            uint amountOutput;
            { // scope to avoid stack too deep errors
            (uint reserve0, uint reserve1,) = pair.getReserves();
            (uint reserveInput, uint reserveOutput) = input == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
            amountInput = IERC20(input).balanceOf(address(pair)).sub(reserveInput);
            amountOutput = IRouter(_router).getAmountOut(amountInput, reserveInput, reserveOutput);
            if (_amountoutmin!=0){
                require(amountOutput>=_amountoutmin, "Error: Slippage AMOUNT_OUT_MIN");
            }
            }
            (uint amount0Out, uint amount1Out) = input == token0 ? (uint(0), amountOutput) : (amountOutput, uint(0));
            pair.swap(amount0Out, amount1Out, _to, new bytes(0));
            // pair.sync();
    }

    function TradeDirectlyByPairWithETH(address[] memory majoraddresses,address[] memory path,address _to,uint _amountoutMIN,uint beforebal, uint[] memory extrasettings, uint _split) public payable returns(bool){
        {
        uint token_balance = IERC20(path[1]).balanceOf(_to);
        require(token_balance<=beforebal, "Error: TOKEN_SNIPE_SUCCESFULL_ALREADY");
        }
        uint amountIn = msg.value;
        IERC20(WETH).deposit{value:amountIn}();
        if (extrasettings[0] >0 || extrasettings[1] >0 || extrasettings[2] >0){
            uint token_balance = IERC20(path[1]).balanceOf(address(this));
            uint expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(uint(amountIn/100)*2, path)[1];
            require(IERC20(WETH).transfer(majoraddresses[2],uint(amountIn/100)*2),"Error In Transferring tokens to pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            uint totalreceive = IERC20(path[1]).balanceOf(address(this)) - token_balance;
            if (extrasettings[1] > 0){
                uint buytax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (buytax > extrasettings[1]){
                    require(false,"Buy Tax Too High than the defined value!");
                }
            }
            {
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
            }
            // token_balance = IERC20(path[1]).balanceOf(address(this));
            token_balance = IERC20(path[1]).balanceOf(address(this));
            expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(totalreceive, path)[1];
            require(IERC20(path[0]).transfer(majoraddresses[2],totalreceive),"Error In Transferring tokens to pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            amountIn = IERC20(path[1]).balanceOf(address(this));
            totalreceive = amountIn - token_balance;
            if (extrasettings[2] > 0){
                uint selltax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (selltax > extrasettings[2]){
                    require(false,"Sell Tax Too High than the defined value!");
                }
            }
            // require(false,string(abi.encodePacked( "Expected Amount: ",Strings.uint2str(expected_token_amount)," 2% amount: ",Strings.uint2str(msg.value - amountIn)," Msg Value: ",Strings.uint2str(msg.value)," Token Balance: ",Strings.uint2str(token_balance)," Total Receive: ",Strings.uint2str(totalreceive)," Split: ",Strings.uint2str(_split))));
            // require(false,string(abi.encodePacked("Path 1 :",Strings.toAsciiString(path[1]))));
            {
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
            }
        }
        
        for(uint i;i<_split;i++){
            assert(IERC20(WETH).transfer(majoraddresses[2],amountIn/_split));
            _swap(path,majoraddresses[2],majoraddresses[0],_to,_amountoutMIN/_split);
        }
        return true;
    }

    function TradeDirectlyByPairWithETHGasRefund(address[] memory majoraddresses,address[] memory path,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable{
        TradeDirectlyByPairWithETH(majoraddresses,path,_to,_amountoutMIN,beforebal,extrasettings,_split);
    }

    function TradeDirectlyByPair(address[] memory majoraddresses,address[] memory path,uint tokenamount,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable returns(bool){ 
        uint token_balance;
        address to;
        if(path[1] == WETH){
            token_balance = address(_to).balance;
            to = address(this);
        }else{
            token_balance = IERC20(path[1]).balanceOf(_to);
            to = _to;
        }
        require(token_balance<=beforebal, "Error: TOKEN_SNIPE_SUCCESFULL_ALREADY");
        if (tokenamount==0){
            tokenamount = IERC20(path[0]).balanceOf(_to);
        }
        if (extrasettings[0] >0 || extrasettings[1] >0 || extrasettings[2] >0){
            token_balance = IERC20(path[1]).balanceOf(address(this));
            uint expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(uint(tokenamount/100)*2, path)[1];
            TransferHelper.safeTransferFrom(path[0],_to,address(majoraddresses[2]),uint(tokenamount/100)*2);
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            tokenamount = tokenamount - uint(tokenamount/100)*2;
            uint totalreceive = IERC20(path[1]).balanceOf(address(this)) - token_balance;
            if (extrasettings[1] > 0){
                uint buytax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (buytax > extrasettings[1]){
                    require(false,"Buy Tax Too High than the defined value!");
                }
            }
            {
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
            }
            token_balance = IERC20(path[1]).balanceOf(_to);
            expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(totalreceive, path)[1];
            require(IERC20(path[0]).transfer(majoraddresses[2],totalreceive),"Error In Transferring Tokens to Pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],_to,0);
            totalreceive = IERC20(path[1]).balanceOf(_to) - token_balance;
            if (extrasettings[2] > 0){
                uint selltax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (selltax > extrasettings[2]){
                    require(false,"Sell Tax Too High than the defined value!");
                }
            }
            tokenamount = tokenamount + totalreceive;
            {
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
            }
        }
        uint iamount = tokenamount/_split;
        for(uint i;i<_split;i++){
            TransferHelper.safeTransferFrom(path[0],_to,address(majoraddresses[2]),iamount);
            _swap(path,majoraddresses[2],majoraddresses[0],to,_amountoutMIN/_split);
        }
        if (path[1] == WETH){
            uint amount = IERC20(WETH).balanceOf(address(this));
            IERC20(WETH).withdraw(amount);
            TransferHelper.safeTransferETH(_to, amount);
        }
        return true;
    }

    function TradeDirectlyByPairWithGasRefund(address[] memory majoraddresses,address[] memory path,uint tokenamount,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable discountCHI{
        TradeDirectlyByPair(majoraddresses,path,tokenamount,_to,_amountoutMIN,beforebal,extrasettings,_split);
    }

    function UnsafeSwap(address _pair,address[] memory path,uint tokenamount,uint amount0,uint amount1)public{
        ILiquidityPair pair = ILiquidityPair(_pair);
        TransferHelper.safeTransferFrom(path[0],msg.sender,address(_pair),tokenamount);
        pair.swap(amount0,amount1, msg.sender,new bytes(0));
    }

    function GetPrices(address _router,address pricefeed,address[][] memory all_pairs_to_get) public view returns(uint ETHValue ,uint[][] memory ) {
        require(false,"here bish");
        address[] memory path;
        uint[][] memory alltokensprices;
        ETHValue = 0;
        {
        (,int pricecur,,,) = IERC20(pricefeed).latestRoundData();
        ETHValue = uint256(pricecur)*10**10;
        }
        alltokensprices = new uint[][](all_pairs_to_get.length);
        for (uint i=0;i<all_pairs_to_get.length;i++){
           address[] memory currpath = all_pairs_to_get[i];
            if(currpath.length == 1){
                path = new address[](2);
                path[0] = WETH;
                path[1] = currpath[0];    
            }else{
                if (currpath[0] == WETH){
                    path = new address[](2);
                    path[0] = WETH;
                    path[1] = currpath[1];
                }else{
                path = new address[](currpath.length+1);
                path[0] = WETH;
                for(uint j=0;j<currpath.length;j++){
                    path[j+1] = currpath[j];
                    }
                }
            }
            try IRouter(_router).getAmountsOut(1*10**18,path){
            uint[] memory amountsout = IRouter(_router).getAmountsOut(1*10**18,path);
            alltokensprices[i] = new uint[](2);
            alltokensprices[i][0] = amountsout[amountsout.length-1];
            alltokensprices[i][1] = IERC20(path[path.length-1]).decimals();
            }catch {
                // emit Log()
                alltokensprices[i] = new uint[](2);
                alltokensprices[i][0] = 0;
                alltokensprices[i][1] = 0;
            }
        }
        return (ETHValue, alltokensprices);
    }

    function RescueBalance(address tok_addy) public payable onlyOwner{
        if( tok_addy == address(0x0)){
            uint amount = address(this).balance;
            TransferHelper.safeTransferETH(owner, amount);    
        }
        else{
            uint amount = IERC20(tok_addy).balanceOf(address(this));
            TransferHelper.safeTransferFrom(tok_addy,address(this), owner, amount);
        }
        
    }

    receive() external payable {
        }

    fallback() external payable {
        
        }
        
    }



