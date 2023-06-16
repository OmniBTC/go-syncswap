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

// IPoolTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type IPoolTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// StablePoolMetaData contains all meta data concerning the StablePool contract.
var StablePoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"_amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burnSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_protocolFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"permit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0PrecisionMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1PrecisionMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StablePoolABI is the input ABI used to generate the binding from.
// Deprecated: Use StablePoolMetaData.ABI instead.
var StablePoolABI = StablePoolMetaData.ABI

// StablePool is an auto generated Go binding around an Ethereum contract.
type StablePool struct {
	StablePoolCaller     // Read-only binding to the contract
	StablePoolTransactor // Write-only binding to the contract
	StablePoolFilterer   // Log filterer for contract events
}

// StablePoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type StablePoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StablePoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StablePoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablePoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StablePoolSession struct {
	Contract     *StablePool       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StablePoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StablePoolCallerSession struct {
	Contract *StablePoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StablePoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StablePoolTransactorSession struct {
	Contract     *StablePoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StablePoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type StablePoolRaw struct {
	Contract *StablePool // Generic contract binding to access the raw methods on
}

// StablePoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StablePoolCallerRaw struct {
	Contract *StablePoolCaller // Generic read-only contract binding to access the raw methods on
}

// StablePoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StablePoolTransactorRaw struct {
	Contract *StablePoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStablePool creates a new instance of StablePool, bound to a specific deployed contract.
func NewStablePool(address common.Address, backend bind.ContractBackend) (*StablePool, error) {
	contract, err := bindStablePool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StablePool{StablePoolCaller: StablePoolCaller{contract: contract}, StablePoolTransactor: StablePoolTransactor{contract: contract}, StablePoolFilterer: StablePoolFilterer{contract: contract}}, nil
}

// NewStablePoolCaller creates a new read-only instance of StablePool, bound to a specific deployed contract.
func NewStablePoolCaller(address common.Address, caller bind.ContractCaller) (*StablePoolCaller, error) {
	contract, err := bindStablePool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StablePoolCaller{contract: contract}, nil
}

// NewStablePoolTransactor creates a new write-only instance of StablePool, bound to a specific deployed contract.
func NewStablePoolTransactor(address common.Address, transactor bind.ContractTransactor) (*StablePoolTransactor, error) {
	contract, err := bindStablePool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StablePoolTransactor{contract: contract}, nil
}

// NewStablePoolFilterer creates a new log filterer instance of StablePool, bound to a specific deployed contract.
func NewStablePoolFilterer(address common.Address, filterer bind.ContractFilterer) (*StablePoolFilterer, error) {
	contract, err := bindStablePool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StablePoolFilterer{contract: contract}, nil
}

// bindStablePool binds a generic wrapper to an already deployed contract.
func bindStablePool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StablePoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablePool *StablePoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablePool.Contract.StablePoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablePool *StablePoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablePool.Contract.StablePoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablePool *StablePoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablePool.Contract.StablePoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablePool *StablePoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablePool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablePool *StablePoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablePool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablePool *StablePoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablePool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StablePool *StablePoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StablePool *StablePoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StablePool.Contract.DOMAINSEPARATOR(&_StablePool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_StablePool *StablePoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _StablePool.Contract.DOMAINSEPARATOR(&_StablePool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StablePool *StablePoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StablePool *StablePoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StablePool.Contract.Allowance(&_StablePool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_StablePool *StablePoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _StablePool.Contract.Allowance(&_StablePool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StablePool *StablePoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StablePool *StablePoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StablePool.Contract.BalanceOf(&_StablePool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_StablePool *StablePoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _StablePool.Contract.BalanceOf(&_StablePool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablePool *StablePoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablePool *StablePoolSession) Decimals() (uint8, error) {
	return _StablePool.Contract.Decimals(&_StablePool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablePool *StablePoolCallerSession) Decimals() (uint8, error) {
	return _StablePool.Contract.Decimals(&_StablePool.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_StablePool *StablePoolCaller) GetAmountIn(opts *bind.CallOpts, _tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getAmountIn", _tokenOut, _amountOut, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_StablePool *StablePoolSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _StablePool.Contract.GetAmountIn(&_StablePool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_StablePool *StablePoolCallerSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _StablePool.Contract.GetAmountIn(&_StablePool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_StablePool *StablePoolCaller) GetAmountOut(opts *bind.CallOpts, _tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getAmountOut", _tokenIn, _amountIn, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_StablePool *StablePoolSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _StablePool.Contract.GetAmountOut(&_StablePool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_StablePool *StablePoolCallerSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _StablePool.Contract.GetAmountOut(&_StablePool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_StablePool *StablePoolCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_StablePool *StablePoolSession) GetAssets() ([]common.Address, error) {
	return _StablePool.Contract.GetAssets(&_StablePool.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_StablePool *StablePoolCallerSession) GetAssets() ([]common.Address, error) {
	return _StablePool.Contract.GetAssets(&_StablePool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_StablePool *StablePoolCaller) GetProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_StablePool *StablePoolSession) GetProtocolFee() (*big.Int, error) {
	return _StablePool.Contract.GetProtocolFee(&_StablePool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_StablePool *StablePoolCallerSession) GetProtocolFee() (*big.Int, error) {
	return _StablePool.Contract.GetProtocolFee(&_StablePool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_StablePool *StablePoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_StablePool *StablePoolSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _StablePool.Contract.GetReserves(&_StablePool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_StablePool *StablePoolCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _StablePool.Contract.GetReserves(&_StablePool.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_StablePool *StablePoolCaller) GetSwapFee(opts *bind.CallOpts, _sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "getSwapFee", _sender, _tokenIn, _tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_StablePool *StablePoolSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _StablePool.Contract.GetSwapFee(&_StablePool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_StablePool *StablePoolCallerSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _StablePool.Contract.GetSwapFee(&_StablePool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_StablePool *StablePoolCaller) InvariantLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "invariantLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_StablePool *StablePoolSession) InvariantLast() (*big.Int, error) {
	return _StablePool.Contract.InvariantLast(&_StablePool.CallOpts)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_StablePool *StablePoolCallerSession) InvariantLast() (*big.Int, error) {
	return _StablePool.Contract.InvariantLast(&_StablePool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_StablePool *StablePoolCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_StablePool *StablePoolSession) Master() (common.Address, error) {
	return _StablePool.Contract.Master(&_StablePool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_StablePool *StablePoolCallerSession) Master() (common.Address, error) {
	return _StablePool.Contract.Master(&_StablePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablePool *StablePoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablePool *StablePoolSession) Name() (string, error) {
	return _StablePool.Contract.Name(&_StablePool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablePool *StablePoolCallerSession) Name() (string, error) {
	return _StablePool.Contract.Name(&_StablePool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StablePool *StablePoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StablePool *StablePoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StablePool.Contract.Nonces(&_StablePool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_StablePool *StablePoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _StablePool.Contract.Nonces(&_StablePool.CallOpts, arg0)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_StablePool *StablePoolCaller) PoolType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_StablePool *StablePoolSession) PoolType() (uint16, error) {
	return _StablePool.Contract.PoolType(&_StablePool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_StablePool *StablePoolCallerSession) PoolType() (uint16, error) {
	return _StablePool.Contract.PoolType(&_StablePool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_StablePool *StablePoolCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_StablePool *StablePoolSession) Reserve0() (*big.Int, error) {
	return _StablePool.Contract.Reserve0(&_StablePool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_StablePool *StablePoolCallerSession) Reserve0() (*big.Int, error) {
	return _StablePool.Contract.Reserve0(&_StablePool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_StablePool *StablePoolCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_StablePool *StablePoolSession) Reserve1() (*big.Int, error) {
	return _StablePool.Contract.Reserve1(&_StablePool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_StablePool *StablePoolCallerSession) Reserve1() (*big.Int, error) {
	return _StablePool.Contract.Reserve1(&_StablePool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_StablePool *StablePoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_StablePool *StablePoolSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _StablePool.Contract.SupportsInterface(&_StablePool.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_StablePool *StablePoolCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _StablePool.Contract.SupportsInterface(&_StablePool.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablePool *StablePoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablePool *StablePoolSession) Symbol() (string, error) {
	return _StablePool.Contract.Symbol(&_StablePool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablePool *StablePoolCallerSession) Symbol() (string, error) {
	return _StablePool.Contract.Symbol(&_StablePool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StablePool *StablePoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StablePool *StablePoolSession) Token0() (common.Address, error) {
	return _StablePool.Contract.Token0(&_StablePool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_StablePool *StablePoolCallerSession) Token0() (common.Address, error) {
	return _StablePool.Contract.Token0(&_StablePool.CallOpts)
}

// Token0PrecisionMultiplier is a free data retrieval call binding the contract method 0xbaa8c7cb.
//
// Solidity: function token0PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolCaller) Token0PrecisionMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "token0PrecisionMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Token0PrecisionMultiplier is a free data retrieval call binding the contract method 0xbaa8c7cb.
//
// Solidity: function token0PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolSession) Token0PrecisionMultiplier() (*big.Int, error) {
	return _StablePool.Contract.Token0PrecisionMultiplier(&_StablePool.CallOpts)
}

// Token0PrecisionMultiplier is a free data retrieval call binding the contract method 0xbaa8c7cb.
//
// Solidity: function token0PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolCallerSession) Token0PrecisionMultiplier() (*big.Int, error) {
	return _StablePool.Contract.Token0PrecisionMultiplier(&_StablePool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StablePool *StablePoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StablePool *StablePoolSession) Token1() (common.Address, error) {
	return _StablePool.Contract.Token1(&_StablePool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_StablePool *StablePoolCallerSession) Token1() (common.Address, error) {
	return _StablePool.Contract.Token1(&_StablePool.CallOpts)
}

// Token1PrecisionMultiplier is a free data retrieval call binding the contract method 0x4e25dc47.
//
// Solidity: function token1PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolCaller) Token1PrecisionMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "token1PrecisionMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Token1PrecisionMultiplier is a free data retrieval call binding the contract method 0x4e25dc47.
//
// Solidity: function token1PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolSession) Token1PrecisionMultiplier() (*big.Int, error) {
	return _StablePool.Contract.Token1PrecisionMultiplier(&_StablePool.CallOpts)
}

// Token1PrecisionMultiplier is a free data retrieval call binding the contract method 0x4e25dc47.
//
// Solidity: function token1PrecisionMultiplier() view returns(uint256)
func (_StablePool *StablePoolCallerSession) Token1PrecisionMultiplier() (*big.Int, error) {
	return _StablePool.Contract.Token1PrecisionMultiplier(&_StablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablePool *StablePoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablePool *StablePoolSession) TotalSupply() (*big.Int, error) {
	return _StablePool.Contract.TotalSupply(&_StablePool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablePool *StablePoolCallerSession) TotalSupply() (*big.Int, error) {
	return _StablePool.Contract.TotalSupply(&_StablePool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StablePool *StablePoolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablePool.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StablePool *StablePoolSession) Vault() (common.Address, error) {
	return _StablePool.Contract.Vault(&_StablePool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_StablePool *StablePoolCallerSession) Vault() (common.Address, error) {
	return _StablePool.Contract.Vault(&_StablePool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StablePool *StablePoolSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.Approve(&_StablePool.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.Approve(&_StablePool.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_StablePool *StablePoolTransactor) Burn(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "burn", _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_StablePool *StablePoolSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Burn(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_StablePool *StablePoolTransactorSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Burn(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolTransactor) BurnSingle(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "burnSingle", _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.BurnSingle(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolTransactorSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.BurnSingle(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_StablePool *StablePoolTransactor) Mint(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "mint", _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_StablePool *StablePoolSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Mint(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_StablePool *StablePoolTransactorSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Mint(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StablePool *StablePoolTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "permit", _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StablePool *StablePoolSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StablePool.Contract.Permit(&_StablePool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_StablePool *StablePoolTransactorSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _StablePool.Contract.Permit(&_StablePool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_StablePool *StablePoolTransactor) Permit2(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "permit2", _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_StablePool *StablePoolSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Permit2(&_StablePool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_StablePool *StablePoolTransactorSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Permit2(&_StablePool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolTransactor) Swap(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "swap", _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Swap(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_StablePool *StablePoolTransactorSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _StablePool.Contract.Swap(&_StablePool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.Transfer(&_StablePool.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.Transfer(&_StablePool.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.TransferFrom(&_StablePool.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_StablePool *StablePoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _StablePool.Contract.TransferFrom(&_StablePool.TransactOpts, _from, _to, _amount)
}

// StablePoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StablePool contract.
type StablePoolApprovalIterator struct {
	Event *StablePoolApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolApproval represents a Approval event raised by the StablePool contract.
type StablePoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StablePool *StablePoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StablePoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StablePoolApprovalIterator{contract: _StablePool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StablePool *StablePoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StablePoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolApproval)
				if err := _StablePool.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_StablePool *StablePoolFilterer) ParseApproval(log types.Log) (*StablePoolApproval, error) {
	event := new(StablePoolApproval)
	if err := _StablePool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the StablePool contract.
type StablePoolBurnIterator struct {
	Event *StablePoolBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolBurn represents a Burn event raised by the StablePool contract.
type StablePoolBurn struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*StablePoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StablePoolBurnIterator{contract: _StablePool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StablePoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolBurn)
				if err := _StablePool.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBurn is a log parse operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) ParseBurn(log types.Log) (*StablePoolBurn, error) {
	event := new(StablePoolBurn)
	if err := _StablePool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the StablePool contract.
type StablePoolMintIterator struct {
	Event *StablePoolMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolMint represents a Mint event raised by the StablePool contract.
type StablePoolMint struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*StablePoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StablePoolMintIterator{contract: _StablePool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *StablePoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolMint)
				if err := _StablePool.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMint is a log parse operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_StablePool *StablePoolFilterer) ParseMint(log types.Log) (*StablePoolMint, error) {
	event := new(StablePoolMint)
	if err := _StablePool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the StablePool contract.
type StablePoolSwapIterator struct {
	Event *StablePoolSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolSwap represents a Swap event raised by the StablePool contract.
type StablePoolSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_StablePool *StablePoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*StablePoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StablePoolSwapIterator{contract: _StablePool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_StablePool *StablePoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *StablePoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolSwap)
				if err := _StablePool.contract.UnpackLog(event, "Swap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_StablePool *StablePoolFilterer) ParseSwap(log types.Log) (*StablePoolSwap, error) {
	event := new(StablePoolSwap)
	if err := _StablePool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the StablePool contract.
type StablePoolSyncIterator struct {
	Event *StablePoolSync // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolSync)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolSync)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolSync represents a Sync event raised by the StablePool contract.
type StablePoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_StablePool *StablePoolFilterer) FilterSync(opts *bind.FilterOpts) (*StablePoolSyncIterator, error) {

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &StablePoolSyncIterator{contract: _StablePool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_StablePool *StablePoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *StablePoolSync) (event.Subscription, error) {

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolSync)
				if err := _StablePool.contract.UnpackLog(event, "Sync", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_StablePool *StablePoolFilterer) ParseSync(log types.Log) (*StablePoolSync, error) {
	event := new(StablePoolSync)
	if err := _StablePool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablePoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StablePool contract.
type StablePoolTransferIterator struct {
	Event *StablePoolTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StablePoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablePoolTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StablePoolTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StablePoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablePoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablePoolTransfer represents a Transfer event raised by the StablePool contract.
type StablePoolTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StablePool *StablePoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StablePoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StablePoolTransferIterator{contract: _StablePool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StablePool *StablePoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StablePoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablePool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablePoolTransfer)
				if err := _StablePool.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_StablePool *StablePoolFilterer) ParseTransfer(log types.Log) (*StablePoolTransfer, error) {
	event := new(StablePoolTransfer)
	if err := _StablePool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

