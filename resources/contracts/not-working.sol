// SPDX-License-Identifier: Unlicensed 
pragma solidity ^0.8.6;

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

    // IWETH functions but not declaring another interface to probs lower gas consumption!
    function deposit() external payable;
    function transfer(address to, uint value) external returns (bool);
    function withdraw(uint) external;
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
    function token0() external view returns (address);
    function token1() external view returns (address);
    function getReserves() external view returns (uint112 reserve0, uint112 reserve1, uint32 blockTimestampLast);
    function price0CumulativeLast() external view returns (uint);
    function price1CumulativeLast() external view returns (uint);
    function kLast() external view returns (uint);
    function swap(uint amount0Out, uint amount1Out, address to, bytes calldata data) external;
    function skim(address to) external;
    function sync() external;
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

library SwapHelper{
    using SafeMath for uint;
    function sortTokens(address tokenA, address tokenB) internal pure returns (address token0, address token1) {
        require(tokenA != tokenB, ' Error: IDENTICAL_ADDRESSES');
        (token0, token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        require(token0 != address(0), 'Error: ZERO_ADDRESS');
    }

    function getAmountOut(uint amountIn, uint reserveIn, uint reserveOut) internal pure returns (uint amountOut) {
        require(amountIn > 0, 'Error: INSUFFICIENT_INPUT_AMOUNT');
        require(reserveIn > 0 && reserveOut > 0, 'Error: INSUFFICIENT_LIQUIDITY');
        uint amountInWithFee = amountIn.mul(9975);
        uint numerator = amountInWithFee.mul(reserveOut);
        uint denominator = reserveIn.mul(10000).add(amountInWithFee);
        amountOut = numerator / denominator;
    }

    function getReserves(ILiquidityPair _pair, address tokenA, address tokenB) internal view returns (uint reserveA, uint reserveB) {
        (address token0,) = sortTokens(tokenA, tokenB);
        (uint reserve0, uint reserve1,) = _pair.getReserves();
        (reserveA, reserveB) = tokenA == token0 ? (reserve0, reserve1) : (reserve1, reserve0);
    }

     function getAmountsOut( IFactory _factory ,uint amountIn, address[] memory path) internal view returns (uint[] memory amounts) {
        require(path.length >= 2, 'Error: INVALID_PATH');
        amounts = new uint[](path.length);
        amounts[0] = amountIn;
        ILiquidityPair _pair = ILiquidityPair(_factory.getPair(path[0], path[1]));
        for (uint i; i < path.length - 1; i++) {
            (uint reserveIn, uint reserveOut) = getReserves(_pair, path[i], path[i + 1]);
            amounts[i + 1] = getAmountOut(amounts[i], reserveIn, reserveOut);
        }
    }
}


contract GoSniper {
    using SafeMath for uint;
    address public owner;
    uint public chainid;
    ChiGasToken private chi;
    address public WETH ;

    struct ExchangeData{
        address factory;
        string init_code;
    }
    mapping (address => ExchangeData) exchanges;

    function setChainData(uint _chainid ) internal{
        if (_chainid == 1 ){
            // Setting in Default Required values
            WETH = address(0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2);
            chi = ChiGasToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);
        }
        else if(_chainid == 56){
            WETH = address(0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c);
            chi = ChiGasToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);
        }else if(_chainid == 137){
            WETH = address(0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270);
            chi = ChiGasToken(0x0000000000004946c0e9F43F4Dee607b0eF1fA1c);
        }else if(_chainid == 43114){
            WETH = address(0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7);
            chi = ChiGasToken(address(0x0));
        }
        else{
            WETH = address(0x0);
            chi = ChiGasToken(address(0x0));
            
        }

    }

    constructor () {
        owner = msg.sender;
        chainid = block.chainid;   
        setChainData(chainid);
        
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

    function transferOwnership(address newOwner) onlyOwner public {
        owner = newOwner;
    }

    function RouterSwapBySplit(address _router,address _tokenIn, address _tokenOut, uint _amountOutMin, address _to,uint _split)   external payable discountCHI{
        address[] memory path;
        path = new address[](2);
        path[0] = _tokenIn;
        path[1] = _tokenOut;
        uint splittedPayment = msg.value/_split;
        for(uint i =0;i<_split;i++){
            IRouter(_router).swapExactETHForTokens{value:splittedPayment}(_amountOutMin, path, _to, block.timestamp);
        }

    }

    function RouterSwapBySplitWithGasRefund(address _router,address _tokenIn, address _tokenOut, uint _amountOutMin, address _to,uint _split)   external payable discountCHI{
        address[] memory path;
        path = new address[](2);
        path[0] = _tokenIn;
        path[1] = _tokenOut;
        uint splittedPayment = msg.value/_split;
        for(uint i =0;i<_split;i++){
            IRouter(_router).swapExactETHForTokens{value:splittedPayment}(_amountOutMin, path, _to, block.timestamp);
        }

    }

    function TradeDirectlyByPair(address _router,address[] memory path,uint _amount,uint _split) public payable{
        // IFactory factory;
        ILiquidityPair pair;
        uint amountout;
        address token0;address token1;
        if(exchanges[_router].factory == address(0x0)){
            // factory = IFactory(IRouter(_router).factory());
            // exchanges[_router].factory = address(factory);
            exchanges[_router].factory = IRouter(_router).factory();
            pair = ILiquidityPair(IFactory(exchanges[_router].factory).getPair(path[0],path[1]));
            
        }else{
            pair = ILiquidityPair(IFactory(exchanges[_router].factory).getPair(path[0],path[1]));
        }
        // pair = ILiquidityPair(factory.getPair(path[0],path[1]));
        if (address(pair) == address(0x0)){
            require(false,string(abi.encodePacked('Error : Pair not found or Not available!')));
        }
        (token0,token1)= SwapHelper.sortTokens(path[0], path[1]);
        _amount = _amount/_split;
        if(token0 == WETH && msg.value >0){
            IERC20(WETH).deposit{value:msg.value}();
        }

        for(uint i;i<_split;i++){
                if(path[0] == WETH){
                    require(IERC20(WETH).balanceOf(msg.sender) >= _amount*_split ||msg.value >= _amount*_split ,string(abi.encodePacked('Error : Insufficient WETH Balance Or No Value Passed!')));
                    IERC20(WETH).transfer(address(pair),_amount);
                    (uint reserve0,uint reserve1) = SwapHelper.getReserves(pair,token0,token1);
                    amountout = IRouter(_router).getAmountOut(_amount,reserve0,reserve1);
            }else{
                require(checkContractAllowance(path[0],_amount) == true,'Error: Token0 does not have enough allowance');
                TransferHelper.safeTransferFrom(path[0],msg.sender,address(pair),_amount);
                (uint reserve0,uint reserve1) = SwapHelper.getReserves(pair,token0,token1);
                (uint reservein,uint reserveout) = path[0] == token0 ?(reserve0, reserve1) : (reserve1, reserve0);
                amountout = IRouter(_router).getAmountOut(IERC20(path[0]).balanceOf(address(pair)).sub(reservein), reservein, reserveout);
            }
            if(address(token0) == address(path[0])){
                if (path[1] == WETH){
                    pair.swap(0,amountout,address(this),new bytes(0));
                }else{
                    pair.swap(0,amountout,msg.sender,new bytes(0));
                }
            }else{
                if (path[1] == WETH){
                    pair.swap(amountout,0,address(this),new bytes(0));
                }else{
                    pair.swap(amountout,0,msg.sender,new bytes(0));
                }
            }
            pair.sync();
        }
        if(path[1] == WETH){
            amountout = IERC20(WETH).balanceOf(address(this));
            IERC20(WETH).withdraw(amountout);
            TransferHelper.safeTransferETH(msg.sender, amountout);
        }
    }

    function TradeDirectlyByPairWithGasRefund(address _router,address[] memory path,uint _amount,uint _split) external payable discountCHI{
        TradeDirectlyByPair(_router,path,_amount,_split);
    }
    
    function HoneyPotCheck(address _router, address[] memory path,uint _split) external payable returns (uint, uint, uint, uint){
        uint tokens_to_get = IRouter(_router).getAmountsOut(msg.value, path)[1];
        address token = path[1];
        address _weth = path[0];
        for(uint i;i<_split;i++){
            IRouter(_router).swapExactETHForTokensSupportingFeeOnTransferTokens{value : msg.value/_split}(0, path, address(this), block.timestamp + 15);
        }
        uint tokens_got = IERC20(token).balanceOf(address(this)); 
        IERC20(token).approve(address(_router), tokens_got);
        path[0] = token; 
        path[1] = _weth; 
        uint eth_to_get = IRouter(_router).getAmountsOut(tokens_got, path)[1];
        for(uint i;i<_split;i++){
            IRouter(_router).swapExactTokensForETHSupportingFeeOnTransferTokens(tokens_got/_split, 0, path, address(this), block.timestamp + 15);
        }
        uint eth_got = address(this).balance;
        return (tokens_to_get, tokens_got, eth_to_get, eth_got);
    }

    function checkContractAllowance(address _token,uint _amount) view public returns(bool){
        uint allowance = IERC20(_token).allowance(msg.sender, address(this));
        return allowance >= _amount;
    }


    receive() external payable {
        }

    fallback() external payable {
        
        }
        
    }