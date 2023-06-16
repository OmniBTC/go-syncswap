// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package syncswap

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)


// IRouterArrayPermitParams is an auto generated low-level Go binding around an user-defined struct.
type IRouterArrayPermitParams struct {
	ApproveAmount *big.Int
	Deadline      *big.Int
	Signature     []byte
}

// IRouterSplitPermitParams is an auto generated low-level Go binding around an user-defined struct.
type IRouterSplitPermitParams struct {
	Token         common.Address
	ApproveAmount *big.Int
	Deadline      *big.Int
	V             uint8
	R             [32]byte
	S             [32]byte
}

// IRouterSwapPath is an auto generated low-level Go binding around an user-defined struct.
type IRouterSwapPath struct {
	Steps    []IRouterSwapStep
	TokenIn  common.Address
	AmountIn *big.Int
}

// IRouterSwapStep is an auto generated low-level Go binding around an user-defined struct.
type IRouterSwapStep struct {
	Pool         common.Address
	Data         []byte
	Callback     common.Address
	CallbackData []byte
}

// SyncSwapRouterTokenInput is an auto generated low-level Go binding around an user-defined struct.
type SyncSwapRouterTokenInput struct {
	Token  common.Address
	Amount *big.Int
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApproveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSyncSwapRouter.TokenInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSyncSwapRouter.TokenInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"addLiquidity2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSyncSwapRouter.TokenInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIRouter.SplitPermitParams[]\",\"name\":\"permits\",\"type\":\"tuple[]\"}],\"name\":\"addLiquidityWithPermit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structSyncSwapRouter.TokenInput[]\",\"name\":\"inputs\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIRouter.SplitPermitParams[]\",\"name\":\"permits\",\"type\":\"tuple[]\"}],\"name\":\"addLiquidityWithPermit2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"burnLiquidity\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"burnLiquiditySingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"amountOut\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRouter.ArrayPermitParams\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"burnLiquiditySingleWithPermit\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"amountOut\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structIRouter.ArrayPermitParams\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"burnLiquidityWithPermit\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"enteredPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"enteredPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isPoolEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"selfPermit2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"selfPermit2IfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowed\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitAllowedIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"selfPermitIfNecessary\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stakingPool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalf\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"internalType\":\"structIRouter.SwapStep[]\",\"name\":\"steps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.SwapPath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"amountOut\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"internalType\":\"structIRouter.SwapStep[]\",\"name\":\"steps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structIRouter.SwapPath[]\",\"name\":\"paths\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"approveAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIRouter.SplitPermitParams\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"swapWithPermit\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"amountOut\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// EnteredPools is a free data retrieval call binding the contract method 0x2b4abadb.
//
// Solidity: function enteredPools(address , uint256 ) view returns(address)
func (_Router *RouterCaller) EnteredPools(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "enteredPools", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EnteredPools is a free data retrieval call binding the contract method 0x2b4abadb.
//
// Solidity: function enteredPools(address , uint256 ) view returns(address)
func (_Router *RouterSession) EnteredPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Router.Contract.EnteredPools(&_Router.CallOpts, arg0, arg1)
}

// EnteredPools is a free data retrieval call binding the contract method 0x2b4abadb.
//
// Solidity: function enteredPools(address , uint256 ) view returns(address)
func (_Router *RouterCallerSession) EnteredPools(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Router.Contract.EnteredPools(&_Router.CallOpts, arg0, arg1)
}

// EnteredPoolsLength is a free data retrieval call binding the contract method 0xb956b3fb.
//
// Solidity: function enteredPoolsLength(address account) view returns(uint256)
func (_Router *RouterCaller) EnteredPoolsLength(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "enteredPoolsLength", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnteredPoolsLength is a free data retrieval call binding the contract method 0xb956b3fb.
//
// Solidity: function enteredPoolsLength(address account) view returns(uint256)
func (_Router *RouterSession) EnteredPoolsLength(account common.Address) (*big.Int, error) {
	return _Router.Contract.EnteredPoolsLength(&_Router.CallOpts, account)
}

// EnteredPoolsLength is a free data retrieval call binding the contract method 0xb956b3fb.
//
// Solidity: function enteredPoolsLength(address account) view returns(uint256)
func (_Router *RouterCallerSession) EnteredPoolsLength(account common.Address) (*big.Int, error) {
	return _Router.Contract.EnteredPoolsLength(&_Router.CallOpts, account)
}

// IsPoolEntered is a free data retrieval call binding the contract method 0x4f25b858.
//
// Solidity: function isPoolEntered(address , address ) view returns(bool)
func (_Router *RouterCaller) IsPoolEntered(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "isPoolEntered", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolEntered is a free data retrieval call binding the contract method 0x4f25b858.
//
// Solidity: function isPoolEntered(address , address ) view returns(bool)
func (_Router *RouterSession) IsPoolEntered(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Router.Contract.IsPoolEntered(&_Router.CallOpts, arg0, arg1)
}

// IsPoolEntered is a free data retrieval call binding the contract method 0x4f25b858.
//
// Solidity: function isPoolEntered(address , address ) view returns(bool)
func (_Router *RouterCallerSession) IsPoolEntered(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Router.Contract.IsPoolEntered(&_Router.CallOpts, arg0, arg1)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterSession) Vault() (common.Address, error) {
	return _Router.Contract.Vault(&_Router.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Router *RouterCallerSession) Vault() (common.Address, error) {
	return _Router.Contract.Vault(&_Router.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Router *RouterCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Router *RouterSession) WETH() (common.Address, error) {
	return _Router.Contract.WETH(&_Router.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Router *RouterCallerSession) WETH() (common.Address, error) {
	return _Router.Contract.WETH(&_Router.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6cbe96fa.
//
// Solidity: function addLiquidity(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidity(opts *bind.TransactOpts, pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity", pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6cbe96fa.
//
// Solidity: function addLiquidity(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterSession) AddLiquidity(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x6cbe96fa.
//
// Solidity: function addLiquidity(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidity(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x94ec6d78.
//
// Solidity: function addLiquidity2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidity2(opts *bind.TransactOpts, pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidity2", pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x94ec6d78.
//
// Solidity: function addLiquidity2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterSession) AddLiquidity2(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity2(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidity2 is a paid mutator transaction binding the contract method 0x94ec6d78.
//
// Solidity: function addLiquidity2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData) payable returns(uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidity2(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidity2(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xc4b3fc40.
//
// Solidity: function addLiquidityWithPermit(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidityWithPermit(opts *bind.TransactOpts, pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityWithPermit", pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xc4b3fc40.
//
// Solidity: function addLiquidityWithPermit(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterSession) AddLiquidityWithPermit(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityWithPermit(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// AddLiquidityWithPermit is a paid mutator transaction binding the contract method 0xc4b3fc40.
//
// Solidity: function addLiquidityWithPermit(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidityWithPermit(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityWithPermit(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// AddLiquidityWithPermit2 is a paid mutator transaction binding the contract method 0xced78795.
//
// Solidity: function addLiquidityWithPermit2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterTransactor) AddLiquidityWithPermit2(opts *bind.TransactOpts, pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addLiquidityWithPermit2", pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// AddLiquidityWithPermit2 is a paid mutator transaction binding the contract method 0xced78795.
//
// Solidity: function addLiquidityWithPermit2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterSession) AddLiquidityWithPermit2(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityWithPermit2(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// AddLiquidityWithPermit2 is a paid mutator transaction binding the contract method 0xced78795.
//
// Solidity: function addLiquidityWithPermit2(address pool, (address,uint256)[] inputs, bytes data, uint256 minLiquidity, address callback, bytes callbackData, (address,uint256,uint256,uint8,bytes32,bytes32)[] permits) payable returns(uint256 liquidity)
func (_Router *RouterTransactorSession) AddLiquidityWithPermit2(pool common.Address, inputs []SyncSwapRouterTokenInput, data []byte, minLiquidity *big.Int, callback common.Address, callbackData []byte, permits []IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.AddLiquidityWithPermit2(&_Router.TransactOpts, pool, inputs, data, minLiquidity, callback, callbackData, permits)
}

// BurnLiquidity is a paid mutator transaction binding the contract method 0xad271fa3.
//
// Solidity: function burnLiquidity(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData) returns((address,uint256)[] amounts)
func (_Router *RouterTransactor) BurnLiquidity(opts *bind.TransactOpts, pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "burnLiquidity", pool, liquidity, data, minAmounts, callback, callbackData)
}

// BurnLiquidity is a paid mutator transaction binding the contract method 0xad271fa3.
//
// Solidity: function burnLiquidity(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData) returns((address,uint256)[] amounts)
func (_Router *RouterSession) BurnLiquidity(pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquidity(&_Router.TransactOpts, pool, liquidity, data, minAmounts, callback, callbackData)
}

// BurnLiquidity is a paid mutator transaction binding the contract method 0xad271fa3.
//
// Solidity: function burnLiquidity(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData) returns((address,uint256)[] amounts)
func (_Router *RouterTransactorSession) BurnLiquidity(pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquidity(&_Router.TransactOpts, pool, liquidity, data, minAmounts, callback, callbackData)
}

// BurnLiquiditySingle is a paid mutator transaction binding the contract method 0x53c43f15.
//
// Solidity: function burnLiquiditySingle(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData) returns((address,uint256) amountOut)
func (_Router *RouterTransactor) BurnLiquiditySingle(opts *bind.TransactOpts, pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "burnLiquiditySingle", pool, liquidity, data, minAmount, callback, callbackData)
}

// BurnLiquiditySingle is a paid mutator transaction binding the contract method 0x53c43f15.
//
// Solidity: function burnLiquiditySingle(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData) returns((address,uint256) amountOut)
func (_Router *RouterSession) BurnLiquiditySingle(pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquiditySingle(&_Router.TransactOpts, pool, liquidity, data, minAmount, callback, callbackData)
}

// BurnLiquiditySingle is a paid mutator transaction binding the contract method 0x53c43f15.
//
// Solidity: function burnLiquiditySingle(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData) returns((address,uint256) amountOut)
func (_Router *RouterTransactorSession) BurnLiquiditySingle(pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquiditySingle(&_Router.TransactOpts, pool, liquidity, data, minAmount, callback, callbackData)
}

// BurnLiquiditySingleWithPermit is a paid mutator transaction binding the contract method 0x7d10c9d6.
//
// Solidity: function burnLiquiditySingleWithPermit(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256) amountOut)
func (_Router *RouterTransactor) BurnLiquiditySingleWithPermit(opts *bind.TransactOpts, pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "burnLiquiditySingleWithPermit", pool, liquidity, data, minAmount, callback, callbackData, permit)
}

// BurnLiquiditySingleWithPermit is a paid mutator transaction binding the contract method 0x7d10c9d6.
//
// Solidity: function burnLiquiditySingleWithPermit(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256) amountOut)
func (_Router *RouterSession) BurnLiquiditySingleWithPermit(pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquiditySingleWithPermit(&_Router.TransactOpts, pool, liquidity, data, minAmount, callback, callbackData, permit)
}

// BurnLiquiditySingleWithPermit is a paid mutator transaction binding the contract method 0x7d10c9d6.
//
// Solidity: function burnLiquiditySingleWithPermit(address pool, uint256 liquidity, bytes data, uint256 minAmount, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256) amountOut)
func (_Router *RouterTransactorSession) BurnLiquiditySingleWithPermit(pool common.Address, liquidity *big.Int, data []byte, minAmount *big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquiditySingleWithPermit(&_Router.TransactOpts, pool, liquidity, data, minAmount, callback, callbackData, permit)
}

// BurnLiquidityWithPermit is a paid mutator transaction binding the contract method 0x353766c6.
//
// Solidity: function burnLiquidityWithPermit(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256)[] amounts)
func (_Router *RouterTransactor) BurnLiquidityWithPermit(opts *bind.TransactOpts, pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "burnLiquidityWithPermit", pool, liquidity, data, minAmounts, callback, callbackData, permit)
}

// BurnLiquidityWithPermit is a paid mutator transaction binding the contract method 0x353766c6.
//
// Solidity: function burnLiquidityWithPermit(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256)[] amounts)
func (_Router *RouterSession) BurnLiquidityWithPermit(pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquidityWithPermit(&_Router.TransactOpts, pool, liquidity, data, minAmounts, callback, callbackData, permit)
}

// BurnLiquidityWithPermit is a paid mutator transaction binding the contract method 0x353766c6.
//
// Solidity: function burnLiquidityWithPermit(address pool, uint256 liquidity, bytes data, uint256[] minAmounts, address callback, bytes callbackData, (uint256,uint256,bytes) permit) returns((address,uint256)[] amounts)
func (_Router *RouterTransactorSession) BurnLiquidityWithPermit(pool common.Address, liquidity *big.Int, data []byte, minAmounts []*big.Int, callback common.Address, callbackData []byte, permit IRouterArrayPermitParams) (*types.Transaction, error) {
	return _Router.Contract.BurnLiquidityWithPermit(&_Router.TransactOpts, pool, liquidity, data, minAmounts, callback, callbackData, permit)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address _factory, bytes data) payable returns(address)
func (_Router *RouterTransactor) CreatePool(opts *bind.TransactOpts, _factory common.Address, data []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "createPool", _factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address _factory, bytes data) payable returns(address)
func (_Router *RouterSession) CreatePool(_factory common.Address, data []byte) (*types.Transaction, error) {
	return _Router.Contract.CreatePool(&_Router.TransactOpts, _factory, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x9dd41df2.
//
// Solidity: function createPool(address _factory, bytes data) payable returns(address)
func (_Router *RouterTransactorSession) CreatePool(_factory common.Address, data []byte) (*types.Transaction, error) {
	return _Router.Contract.CreatePool(&_Router.TransactOpts, _factory, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) payable returns(bytes[] results)
func (_Router *RouterTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Router.Contract.Multicall(&_Router.TransactOpts, data)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactor) SelfPermit(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermit", token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit(&_Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit is a paid mutator transaction binding the contract method 0xf3995c67.
//
// Solidity: function selfPermit(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactorSession) SelfPermit(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit(&_Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermit2 is a paid mutator transaction binding the contract method 0x6cc781cd.
//
// Solidity: function selfPermit2(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterTransactor) SelfPermit2(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermit2", token, value, deadline, signature)
}

// SelfPermit2 is a paid mutator transaction binding the contract method 0x6cc781cd.
//
// Solidity: function selfPermit2(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterSession) SelfPermit2(token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit2(&_Router.TransactOpts, token, value, deadline, signature)
}

// SelfPermit2 is a paid mutator transaction binding the contract method 0x6cc781cd.
//
// Solidity: function selfPermit2(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterTransactorSession) SelfPermit2(token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit2(&_Router.TransactOpts, token, value, deadline, signature)
}

// SelfPermit2IfNecessary is a paid mutator transaction binding the contract method 0x688ee44c.
//
// Solidity: function selfPermit2IfNecessary(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterTransactor) SelfPermit2IfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermit2IfNecessary", token, value, deadline, signature)
}

// SelfPermit2IfNecessary is a paid mutator transaction binding the contract method 0x688ee44c.
//
// Solidity: function selfPermit2IfNecessary(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterSession) SelfPermit2IfNecessary(token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit2IfNecessary(&_Router.TransactOpts, token, value, deadline, signature)
}

// SelfPermit2IfNecessary is a paid mutator transaction binding the contract method 0x688ee44c.
//
// Solidity: function selfPermit2IfNecessary(address token, uint256 value, uint256 deadline, bytes signature) payable returns()
func (_Router *RouterTransactorSession) SelfPermit2IfNecessary(token common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermit2IfNecessary(&_Router.TransactOpts, token, value, deadline, signature)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactor) SelfPermitAllowed(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermitAllowed", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitAllowed(&_Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowed is a paid mutator transaction binding the contract method 0x4659a494.
//
// Solidity: function selfPermitAllowed(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactorSession) SelfPermitAllowed(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitAllowed(&_Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactor) SelfPermitAllowedIfNecessary(opts *bind.TransactOpts, token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermitAllowedIfNecessary", token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitAllowedIfNecessary(&_Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitAllowedIfNecessary is a paid mutator transaction binding the contract method 0xa4a78f0c.
//
// Solidity: function selfPermitAllowedIfNecessary(address token, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactorSession) SelfPermitAllowedIfNecessary(token common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitAllowedIfNecessary(&_Router.TransactOpts, token, nonce, expiry, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactor) SelfPermitIfNecessary(opts *bind.TransactOpts, token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "selfPermitIfNecessary", token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitIfNecessary(&_Router.TransactOpts, token, value, deadline, v, r, s)
}

// SelfPermitIfNecessary is a paid mutator transaction binding the contract method 0xc2e3140a.
//
// Solidity: function selfPermitIfNecessary(address token, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) payable returns()
func (_Router *RouterTransactorSession) SelfPermitIfNecessary(token common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Router.Contract.SelfPermitIfNecessary(&_Router.TransactOpts, token, value, deadline, v, r, s)
}

// Stake is a paid mutator transaction binding the contract method 0x6291027c.
//
// Solidity: function stake(address stakingPool, address token, uint256 amount, address onBehalf) returns()
func (_Router *RouterTransactor) Stake(opts *bind.TransactOpts, stakingPool common.Address, token common.Address, amount *big.Int, onBehalf common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "stake", stakingPool, token, amount, onBehalf)
}

// Stake is a paid mutator transaction binding the contract method 0x6291027c.
//
// Solidity: function stake(address stakingPool, address token, uint256 amount, address onBehalf) returns()
func (_Router *RouterSession) Stake(stakingPool common.Address, token common.Address, amount *big.Int, onBehalf common.Address) (*types.Transaction, error) {
	return _Router.Contract.Stake(&_Router.TransactOpts, stakingPool, token, amount, onBehalf)
}

// Stake is a paid mutator transaction binding the contract method 0x6291027c.
//
// Solidity: function stake(address stakingPool, address token, uint256 amount, address onBehalf) returns()
func (_Router *RouterTransactorSession) Stake(stakingPool common.Address, token common.Address, amount *big.Int, onBehalf common.Address) (*types.Transaction, error) {
	return _Router.Contract.Stake(&_Router.TransactOpts, stakingPool, token, amount, onBehalf)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns((address,uint256) amountOut)
func (_Router *RouterTransactor) Swap(opts *bind.TransactOpts, paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swap", paths, amountOutMin, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns((address,uint256) amountOut)
func (_Router *RouterSession) Swap(paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, paths, amountOutMin, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x2cc4081e.
//
// Solidity: function swap(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline) payable returns((address,uint256) amountOut)
func (_Router *RouterTransactorSession) Swap(paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, paths, amountOutMin, deadline)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns((address,uint256) amountOut)
func (_Router *RouterTransactor) SwapWithPermit(opts *bind.TransactOpts, paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swapWithPermit", paths, amountOutMin, deadline, permit)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns((address,uint256) amountOut)
func (_Router *RouterSession) SwapWithPermit(paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.SwapWithPermit(&_Router.TransactOpts, paths, amountOutMin, deadline, permit)
}

// SwapWithPermit is a paid mutator transaction binding the contract method 0xe84d494b.
//
// Solidity: function swapWithPermit(((address,bytes,address,bytes)[],address,uint256)[] paths, uint256 amountOutMin, uint256 deadline, (address,uint256,uint256,uint8,bytes32,bytes32) permit) payable returns((address,uint256) amountOut)
func (_Router *RouterTransactorSession) SwapWithPermit(paths []IRouterSwapPath, amountOutMin *big.Int, deadline *big.Int, permit IRouterSplitPermitParams) (*types.Transaction, error) {
	return _Router.Contract.SwapWithPermit(&_Router.TransactOpts, paths, amountOutMin, deadline, permit)
}

