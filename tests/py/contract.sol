// SPDX-License-Identifier: Unlicensed 
pragma solidity ^0.6.12;
pragma experimental ABIEncoderV2;


interface IERC20 ***REMOVED***
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


***REMOVED***


interface IRouter ***REMOVED***
    function factory() external pure returns (address);
    function getAmountsOut(uint amountIn, address[] calldata path) external view returns (uint[] memory amounts);
    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) external pure returns (uint amountOut);
    function swapExactETHForTokens(uint amountOutMin, address[] calldata path, address to, uint deadline)external payable returns (uint[] memory amounts);
    function swapExactTokensForETH(uint amountIn, address[] calldata path, uint deadline)external payable returns (uint[] memory amounts);
    function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint amountIn,uint amountOutMin,address[] calldata path,address to,uint deadline) external payable returns (uint[] memory amounts);
    function swapExactETHForTokensSupportingFeeOnTransferTokens(uint amountOutMin,address[] calldata path,address to,uint deadline) external payable;
    function swapExactTokensForETHSupportingFeeOnTransferTokens(uint amountIn,uint amountOutMin,address[] calldata path,address to,uint deadline) external payable;
***REMOVED***

interface IFactory ***REMOVED***
    function getPair(address tokenA, address tokenB) external view returns (address pair);
    function allPairs(uint) external view returns (address pair);
    function allPairsLength() external view returns (uint);
***REMOVED***

interface ILiquidityPair ***REMOVED***
    function factory() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function sync() external;
    function token0() external pure returns(address);
***REMOVED***

interface ChiGasToken***REMOVED***
    function free(uint256 value) external returns (uint256);
    function freeUpTo(uint256 value) external returns (uint256);
    function freeFrom(address from, uint256 value) external returns (uint256);
    function freeFromUpTo(address from, uint256 value) external returns (uint256);
***REMOVED***



library TransferHelper ***REMOVED***
    function safeApprove(address token, address to, uint value) internal ***REMOVED***
        // bytes4(keccak256(bytes('approve(address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x095ea7b3, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: APPROVE_FAILED');
***REMOVED***

    function safeTransfer(address token, address to, uint value) internal ***REMOVED***
        // bytes4(keccak256(bytes('transfer(address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0xa9059cbb, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: TRANSFER_FAILED');
***REMOVED***

    function safeTransferFrom(address token, address from, address to, uint value) internal ***REMOVED***
        // bytes4(keccak256(bytes('transferFrom(address,address,uint256)')));
        (bool success, bytes memory data) = token.call(abi.encodeWithSelector(0x23b872dd, from, to, value));
        require(success && (data.length == 0 || abi.decode(data, (bool))), 'TransferHelper: TRANSFER_FROM_FAILED');
***REMOVED***

    function safeTransferETH(address to, uint value) internal ***REMOVED***
        (bool success,) = to.call***REMOVED***value:value***REMOVED***(new bytes(0));
        require(success, 'TransferHelper: ETH_TRANSFER_FAILED');
***REMOVED***

***REMOVED***


library SafeMath ***REMOVED***
    function add(uint x, uint y) internal pure returns (uint z) ***REMOVED***
        require((z = x + y) >= x, 'ds-math-add-overflow');
***REMOVED***

    function sub(uint x, uint y) internal pure returns (uint z) ***REMOVED***
        require((z = x - y) <= x, 'ds-math-sub-underflow');
***REMOVED***

    function mul(uint x, uint y) internal pure returns (uint z) ***REMOVED***
        require(y == 0 || (z = x * y) / y == x, 'ds-math-mul-overflow');
***REMOVED***
***REMOVED***


library Strings***REMOVED***
    function uint2str(uint _i) internal pure returns (string memory _uintAsString) ***REMOVED***
        if (_i == 0) ***REMOVED***
            return "0";
    ***REMOVED***
        uint j = _i;
        uint len;
        while (j != 0) ***REMOVED***
            len++;
            j /= 10;
    ***REMOVED***
        bytes memory bstr = new bytes(len);
        uint k = len;
        while (_i != 0) ***REMOVED***
            k = k-1;
            uint8 temp = (48 + uint8(_i - _i / 10 * 10));
            bytes1 b1 = bytes1(temp);
            bstr[k] = b1;
            _i /= 10;
    ***REMOVED***
        return string(bstr);
***REMOVED***
    function toAsciiString(address x) internal pure returns (string memory) ***REMOVED***
    bytes memory s = new bytes(40);
    for (uint i = 0; i < 20; i++) ***REMOVED***
        bytes1 b = bytes1(uint8(uint(uint160(x)) / (2**(8*(19 - i)))));
        bytes1 hi = bytes1(uint8(b) / 16);
        bytes1 lo = bytes1(uint8(b) - 16 * uint8(hi));
        s[2*i] = char(hi);
        s[2*i+1] = char(lo);            
***REMOVED***
    return string(s);
***REMOVED***

function char(bytes1 b) internal pure returns (bytes1 c) ***REMOVED***
    if (uint8(b) < 10) return bytes1(uint8(b) + 0x30);
    else return bytes1(uint8(b) + 0x57);
***REMOVED***
***REMOVED***

contract GoSniper ***REMOVED***
    using SafeMath for uint;
    address public owner;
    ChiGasToken public chi;
    address public WETH ;
    mapping (address => bool) public Allowed;
    constructor (address _weth,address _chi) public ***REMOVED***
        owner = msg.sender;
        WETH = _weth;
        chi = ChiGasToken(_chi);
        
***REMOVED***

    modifier discountCHI ***REMOVED***
        if (address(chi)!=address(0x0))***REMOVED***
            uint256 gasStart = gasleft();

        _;

        uint256 initialGas = 21000 + 16 * msg.data.length;
        uint256 gasSpent = initialGas + gasStart - gasleft();
        uint256 freeUpValue = (gasSpent + 14154) / 41947;

        chi.freeFromUpTo(msg.sender, freeUpValue);
    ***REMOVED***
        
***REMOVED***

    modifier onlyOwner() ***REMOVED***
        require(owner == msg.sender, "Not Owner");
        _;
***REMOVED***

    modifier onlyOwnerOrAllowed(address _addr) ***REMOVED***
        require(owner == msg.sender || Allowed[_addr], "Not Owner or Allowed");
        _;
***REMOVED***

    function setWETH(address _weth) onlyOwner public ***REMOVED***
        WETH = _weth;
***REMOVED***
    
    function setCHI(address _chi) onlyOwner public ***REMOVED***
        chi = ChiGasToken(_chi);
***REMOVED***
    

    function transferOwnership(address newOwner) onlyOwner public ***REMOVED***
        owner = newOwner;
***REMOVED***
    
    function SetAllowed(address _addr,bool _res) onlyOwner public ***REMOVED***
        Allowed[_addr] = _res;
***REMOVED***

    function SetAllowedBulk(address[] memory _addrs,bool _res) onlyOwner public ***REMOVED***
        for (uint i = 0; i < _addrs.length; i++) ***REMOVED***
            Allowed[_addrs[i]] = _res;
    ***REMOVED***
***REMOVED***

    function sortTokens(address tokenA, address tokenB) public pure returns (address token0, address token1) ***REMOVED***
        require(tokenA != tokenB, ' Error: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'Error: ZERO_ADDRESS');
***REMOVED***

    function _swap(address[] memory path, address _pair,address _router,address _to,uint _amountoutmin) internal virtual ***REMOVED***
            (address input, address output) = (path[0], path[1]);
            (address token0,) = sortTokens(input, output);
            ILiquidityPair pair = ILiquidityPair(_pair) ;
            uint amountInput;
            uint amountOutput;
            ***REMOVED*** // scope to avoid stack too deep errors
            (uint reserve0, uint reserve1,) = pair.getReserves();
            (uint reserveInput, uint reserveOutput) = input == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
            amountInput = IERC20(input).balanceOf(address(pair)).sub(reserveInput);
            amountOutput = IRouter(_router).getAmountOut(amountInput, reserveInput, reserveOutput);
            if (_amountoutmin!=0)***REMOVED***
                require(amountOutput>=_amountoutmin, "Error: Slippage AMOUNT_OUT_MIN");
        ***REMOVED***
        ***REMOVED***
            (uint amount0Out, uint amount1Out) = input == token0 ? (uint(0), amountOutput) : (amountOutput, uint(0));
            pair.swap(amount0Out, amount1Out, _to, new bytes(0));
            // pair.sync();
***REMOVED***

    function TradeDirectlyByPairWithETH(address[] memory majoraddresses,address[] memory path,address _to,uint _amountoutMIN,uint beforebal, uint[] memory extrasettings, uint _split) public payable returns(bool)***REMOVED***
        ***REMOVED***
        uint token_balance = IERC20(path[1]).balanceOf(_to);
        require(token_balance<=beforebal, "Error: TOKEN_SNIPE_SUCCESFULL_ALREADY");
    ***REMOVED***
        uint amountIn = msg.value;
        IERC20(WETH).deposit***REMOVED***value:amountIn***REMOVED***();
        if (extrasettings[0] >0 || extrasettings[1] >0 || extrasettings[2] >0)***REMOVED***
            uint token_balance = IERC20(path[1]).balanceOf(address(this));
            uint expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(uint(amountIn/100)*2, path)[1];
            require(IERC20(WETH).transfer(majoraddresses[2],uint(amountIn/100)*2),"Error In Transferring tokens to pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            uint totalreceive = IERC20(path[1]).balanceOf(address(this)) - token_balance;
            if (extrasettings[1] > 0)***REMOVED***
                uint buytax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (buytax > extrasettings[1])***REMOVED***
                    require(false,"Buy Tax Too High than the defined value!");
            ***REMOVED***
        ***REMOVED***
            ***REMOVED***
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
        ***REMOVED***
            // token_balance = IERC20(path[1]).balanceOf(address(this));
            token_balance = IERC20(path[1]).balanceOf(address(this));
            expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(totalreceive, path)[1];
            require(IERC20(path[0]).transfer(majoraddresses[2],totalreceive),"Error In Transferring tokens to pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            amountIn = IERC20(path[1]).balanceOf(address(this));
            totalreceive = amountIn - token_balance;
            if (extrasettings[2] > 0)***REMOVED***
                uint selltax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (selltax > extrasettings[2])***REMOVED***
                    require(false,"Sell Tax Too High than the defined value!");
            ***REMOVED***
        ***REMOVED***
            // require(false,string(abi.encodePacked( "Expected Amount: ",Strings.uint2str(expected_token_amount)," 2% amount: ",Strings.uint2str(msg.value - amountIn)," Msg Value: ",Strings.uint2str(msg.value)," Token Balance: ",Strings.uint2str(token_balance)," Total Receive: ",Strings.uint2str(totalreceive)," Split: ",Strings.uint2str(_split))));
            // require(false,string(abi.encodePacked("Path 1 :",Strings.toAsciiString(path[1]))));
            ***REMOVED***
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
        ***REMOVED***
    ***REMOVED***
        
        for(uint i;i<_split;i++)***REMOVED***
            assert(IERC20(WETH).transfer(majoraddresses[2],amountIn/_split));
            _swap(path,majoraddresses[2],majoraddresses[0],_to,_amountoutMIN/_split);
    ***REMOVED***
        return true;
***REMOVED***

    function TradeDirectlyByPairWithETHGasRefund(address[] memory majoraddresses,address[] memory path,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable***REMOVED***
        TradeDirectlyByPairWithETH(majoraddresses,path,_to,_amountoutMIN,beforebal,extrasettings,_split);
***REMOVED***

    function TradeDirectlyByPair(address[] memory majoraddresses,address[] memory path,uint tokenamount,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable returns(bool)***REMOVED*** 
        uint token_balance;
        address to;
        if(path[1] == WETH)***REMOVED***
            token_balance = address(_to).balance;
            to = address(this);
    ***REMOVED***else***REMOVED***
            token_balance = IERC20(path[1]).balanceOf(_to);
            to = _to;
    ***REMOVED***
        require(token_balance<=beforebal, "Error: TOKEN_SNIPE_SUCCESFULL_ALREADY");
        if (tokenamount==0)***REMOVED***
            tokenamount = IERC20(path[0]).balanceOf(_to);
    ***REMOVED***
        if (extrasettings[0] >0 || extrasettings[1] >0 || extrasettings[2] >0)***REMOVED***
            token_balance = IERC20(path[1]).balanceOf(address(this));
            uint expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(uint(tokenamount/100)*2, path)[1];
            TransferHelper.safeTransferFrom(path[0],_to,address(majoraddresses[2]),uint(tokenamount/100)*2);
            _swap(path,majoraddresses[2],majoraddresses[0],address(this),0);
            tokenamount = tokenamount - uint(tokenamount/100)*2;
            uint totalreceive = IERC20(path[1]).balanceOf(address(this)) - token_balance;
            if (extrasettings[1] > 0)***REMOVED***
                uint buytax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (buytax > extrasettings[1])***REMOVED***
                    require(false,"Buy Tax Too High than the defined value!");
            ***REMOVED***
        ***REMOVED***
            ***REMOVED***
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
        ***REMOVED***
            token_balance = IERC20(path[1]).balanceOf(_to);
            expected_token_amount = IRouter(majoraddresses[0]).getAmountsOut(totalreceive, path)[1];
            require(IERC20(path[0]).transfer(majoraddresses[2],totalreceive),"Error In Transferring Tokens to Pair during HP check");
            _swap(path,majoraddresses[2],majoraddresses[0],_to,0);
            totalreceive = IERC20(path[1]).balanceOf(_to) - token_balance;
            if (extrasettings[2] > 0)***REMOVED***
                uint selltax = uint(100.0 - ((totalreceive) / (expected_token_amount) * 100.0));
                if (selltax > extrasettings[2])***REMOVED***
                    require(false,"Sell Tax Too High than the defined value!");
            ***REMOVED***
        ***REMOVED***
            tokenamount = tokenamount + totalreceive;
            ***REMOVED***
            address bak = path[0];
            path[0] = path[1];
            path[1] = bak;
        ***REMOVED***
    ***REMOVED***
        uint iamount = tokenamount/_split;
        for(uint i;i<_split;i++)***REMOVED***
            TransferHelper.safeTransferFrom(path[0],_to,address(majoraddresses[2]),iamount);
            _swap(path,majoraddresses[2],majoraddresses[0],to,_amountoutMIN/_split);
    ***REMOVED***
        if (path[1] == WETH)***REMOVED***
            uint amount = IERC20(WETH).balanceOf(address(this));
            IERC20(WETH).withdraw(amount);
            TransferHelper.safeTransferETH(_to, amount);
    ***REMOVED***
        return true;
***REMOVED***

    function TradeDirectlyByPairWithGasRefund(address[] memory majoraddresses,address[] memory path,uint tokenamount,address _to,uint _amountoutMIN,uint beforebal,uint[] memory extrasettings,uint _split) public payable discountCHI***REMOVED***
        TradeDirectlyByPair(majoraddresses,path,tokenamount,_to,_amountoutMIN,beforebal,extrasettings,_split);
***REMOVED***

    function UnsafeSwap(address _pair,address[] memory path,uint tokenamount,uint amount0,uint amount1)public***REMOVED***
        ILiquidityPair pair = ILiquidityPair(_pair);
        TransferHelper.safeTransferFrom(path[0],msg.sender,address(_pair),tokenamount);
        pair.swap(amount0,amount1, msg.sender,new bytes(0));
***REMOVED***

    function GetPrices(address _router,address pricefeed,address[][] memory all_pairs_to_get) public view returns(uint ETHValue ,uint[][] memory ) ***REMOVED***
        require(false,"here bish");
        address[] memory path;
        uint[][] memory alltokensprices;
        ETHValue = 0;
        ***REMOVED***
        (,int pricecur,,,) = IERC20(pricefeed).latestRoundData();
        ETHValue = uint256(pricecur)*10**10;
    ***REMOVED***
        alltokensprices = new uint[][](all_pairs_to_get.length);
        for (uint i=0;i<all_pairs_to_get.length;i++)***REMOVED***
           address[] memory currpath = all_pairs_to_get[i];
            if(currpath.length == 1)***REMOVED***
                path = new address[](2);
                path[0] = WETH;
                path[1] = currpath[0];    
        ***REMOVED***else***REMOVED***
                if (currpath[0] == WETH)***REMOVED***
                    path = new address[](2);
                    path[0] = WETH;
                    path[1] = currpath[1];
            ***REMOVED***else***REMOVED***
                path = new address[](currpath.length+1);
                path[0] = WETH;
                for(uint j=0;j<currpath.length;j++)***REMOVED***
                    path[j+1] = currpath[j];
                ***REMOVED***
            ***REMOVED***
        ***REMOVED***
            try IRouter(_router).getAmountsOut(1*10**18,path)***REMOVED***
            uint[] memory amountsout = IRouter(_router).getAmountsOut(1*10**18,path);
            alltokensprices[i] = new uint[](2);
            alltokensprices[i][0] = amountsout[amountsout.length-1];
            alltokensprices[i][1] = IERC20(path[path.length-1]).decimals();
        ***REMOVED***catch ***REMOVED***
                // emit Log()
                alltokensprices[i] = new uint[](2);
                alltokensprices[i][0] = 0;
                alltokensprices[i][1] = 0;
        ***REMOVED***
    ***REMOVED***
        return (ETHValue, alltokensprices);
***REMOVED***

    function RescueBalance(address tok_addy) public payable onlyOwner***REMOVED***
        if( tok_addy == address(0x0))***REMOVED***
            uint amount = address(this).balance;
            TransferHelper.safeTransferETH(owner, amount);    
    ***REMOVED***
        else***REMOVED***
            uint amount = IERC20(tok_addy).balanceOf(address(this));
            TransferHelper.safeTransferFrom(tok_addy,address(this), owner, amount);
    ***REMOVED***
        
***REMOVED***

    receive() external payable ***REMOVED***
    ***REMOVED***

    fallback() external payable ***REMOVED***
        
    ***REMOVED***
        
***REMOVED***



