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


// ClassicPoolMetaData contains all meta data concerning the ClassicPool contract.
var ClassicPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"_amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burnSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_protocolFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"permit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ClassicPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ClassicPoolMetaData.ABI instead.
var ClassicPoolABI = ClassicPoolMetaData.ABI

// ClassicPool is an auto generated Go binding around an Ethereum contract.
type ClassicPool struct {
	ClassicPoolCaller     // Read-only binding to the contract
	ClassicPoolTransactor // Write-only binding to the contract
	ClassicPoolFilterer   // Log filterer for contract events
}

// ClassicPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClassicPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClassicPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClassicPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClassicPoolSession struct {
	Contract     *ClassicPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClassicPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClassicPoolCallerSession struct {
	Contract *ClassicPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ClassicPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClassicPoolTransactorSession struct {
	Contract     *ClassicPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ClassicPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClassicPoolRaw struct {
	Contract *ClassicPool // Generic contract binding to access the raw methods on
}

// ClassicPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClassicPoolCallerRaw struct {
	Contract *ClassicPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ClassicPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClassicPoolTransactorRaw struct {
	Contract *ClassicPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClassicPool creates a new instance of ClassicPool, bound to a specific deployed contract.
func NewClassicPool(address common.Address, backend bind.ContractBackend) (*ClassicPool, error) {
	contract, err := bindClassicPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClassicPool{ClassicPoolCaller: ClassicPoolCaller{contract: contract}, ClassicPoolTransactor: ClassicPoolTransactor{contract: contract}, ClassicPoolFilterer: ClassicPoolFilterer{contract: contract}}, nil
}

// NewClassicPoolCaller creates a new read-only instance of ClassicPool, bound to a specific deployed contract.
func NewClassicPoolCaller(address common.Address, caller bind.ContractCaller) (*ClassicPoolCaller, error) {
	contract, err := bindClassicPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolCaller{contract: contract}, nil
}

// NewClassicPoolTransactor creates a new write-only instance of ClassicPool, bound to a specific deployed contract.
func NewClassicPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ClassicPoolTransactor, error) {
	contract, err := bindClassicPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolTransactor{contract: contract}, nil
}

// NewClassicPoolFilterer creates a new log filterer instance of ClassicPool, bound to a specific deployed contract.
func NewClassicPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ClassicPoolFilterer, error) {
	contract, err := bindClassicPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFilterer{contract: contract}, nil
}

// bindClassicPool binds a generic wrapper to an already deployed contract.
func bindClassicPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClassicPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClassicPool *ClassicPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClassicPool.Contract.ClassicPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClassicPool *ClassicPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClassicPool.Contract.ClassicPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClassicPool *ClassicPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClassicPool.Contract.ClassicPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClassicPool *ClassicPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClassicPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClassicPool *ClassicPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClassicPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClassicPool *ClassicPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClassicPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ClassicPool *ClassicPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ClassicPool *ClassicPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ClassicPool.Contract.DOMAINSEPARATOR(&_ClassicPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_ClassicPool *ClassicPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _ClassicPool.Contract.DOMAINSEPARATOR(&_ClassicPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ClassicPool *ClassicPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.Allowance(&_ClassicPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.Allowance(&_ClassicPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.BalanceOf(&_ClassicPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.BalanceOf(&_ClassicPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClassicPool *ClassicPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClassicPool *ClassicPoolSession) Decimals() (uint8, error) {
	return _ClassicPool.Contract.Decimals(&_ClassicPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ClassicPool *ClassicPoolCallerSession) Decimals() (uint8, error) {
	return _ClassicPool.Contract.Decimals(&_ClassicPool.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_ClassicPool *ClassicPoolCaller) GetAmountIn(opts *bind.CallOpts, _tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getAmountIn", _tokenOut, _amountOut, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_ClassicPool *ClassicPoolSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.GetAmountIn(&_ClassicPool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_ClassicPool *ClassicPoolCallerSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.GetAmountIn(&_ClassicPool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_ClassicPool *ClassicPoolCaller) GetAmountOut(opts *bind.CallOpts, _tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getAmountOut", _tokenIn, _amountIn, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_ClassicPool *ClassicPoolSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.GetAmountOut(&_ClassicPool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_ClassicPool *ClassicPoolCallerSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.GetAmountOut(&_ClassicPool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ClassicPool *ClassicPoolCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ClassicPool *ClassicPoolSession) GetAssets() ([]common.Address, error) {
	return _ClassicPool.Contract.GetAssets(&_ClassicPool.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_ClassicPool *ClassicPoolCallerSession) GetAssets() ([]common.Address, error) {
	return _ClassicPool.Contract.GetAssets(&_ClassicPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_ClassicPool *ClassicPoolCaller) GetProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_ClassicPool *ClassicPoolSession) GetProtocolFee() (*big.Int, error) {
	return _ClassicPool.Contract.GetProtocolFee(&_ClassicPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_ClassicPool *ClassicPoolCallerSession) GetProtocolFee() (*big.Int, error) {
	return _ClassicPool.Contract.GetProtocolFee(&_ClassicPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_ClassicPool *ClassicPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getReserves")

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
func (_ClassicPool *ClassicPoolSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _ClassicPool.Contract.GetReserves(&_ClassicPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_ClassicPool *ClassicPoolCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _ClassicPool.Contract.GetReserves(&_ClassicPool.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_ClassicPool *ClassicPoolCaller) GetSwapFee(opts *bind.CallOpts, _sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "getSwapFee", _sender, _tokenIn, _tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_ClassicPool *ClassicPoolSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ClassicPool.Contract.GetSwapFee(&_ClassicPool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_ClassicPool *ClassicPoolCallerSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ClassicPool.Contract.GetSwapFee(&_ClassicPool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) InvariantLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "invariantLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ClassicPool *ClassicPoolSession) InvariantLast() (*big.Int, error) {
	return _ClassicPool.Contract.InvariantLast(&_ClassicPool.CallOpts)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) InvariantLast() (*big.Int, error) {
	return _ClassicPool.Contract.InvariantLast(&_ClassicPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPool *ClassicPoolCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPool *ClassicPoolSession) Master() (common.Address, error) {
	return _ClassicPool.Contract.Master(&_ClassicPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPool *ClassicPoolCallerSession) Master() (common.Address, error) {
	return _ClassicPool.Contract.Master(&_ClassicPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClassicPool *ClassicPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClassicPool *ClassicPoolSession) Name() (string, error) {
	return _ClassicPool.Contract.Name(&_ClassicPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ClassicPool *ClassicPoolCallerSession) Name() (string, error) {
	return _ClassicPool.Contract.Name(&_ClassicPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.Nonces(&_ClassicPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ClassicPool.Contract.Nonces(&_ClassicPool.CallOpts, arg0)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ClassicPool *ClassicPoolCaller) PoolType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ClassicPool *ClassicPoolSession) PoolType() (uint16, error) {
	return _ClassicPool.Contract.PoolType(&_ClassicPool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_ClassicPool *ClassicPoolCallerSession) PoolType() (uint16, error) {
	return _ClassicPool.Contract.PoolType(&_ClassicPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ClassicPool *ClassicPoolSession) Reserve0() (*big.Int, error) {
	return _ClassicPool.Contract.Reserve0(&_ClassicPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) Reserve0() (*big.Int, error) {
	return _ClassicPool.Contract.Reserve0(&_ClassicPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ClassicPool *ClassicPoolSession) Reserve1() (*big.Int, error) {
	return _ClassicPool.Contract.Reserve1(&_ClassicPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) Reserve1() (*big.Int, error) {
	return _ClassicPool.Contract.Reserve1(&_ClassicPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ClassicPool *ClassicPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ClassicPool *ClassicPoolSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ClassicPool.Contract.SupportsInterface(&_ClassicPool.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ClassicPool *ClassicPoolCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ClassicPool.Contract.SupportsInterface(&_ClassicPool.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClassicPool *ClassicPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClassicPool *ClassicPoolSession) Symbol() (string, error) {
	return _ClassicPool.Contract.Symbol(&_ClassicPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ClassicPool *ClassicPoolCallerSession) Symbol() (string, error) {
	return _ClassicPool.Contract.Symbol(&_ClassicPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ClassicPool *ClassicPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ClassicPool *ClassicPoolSession) Token0() (common.Address, error) {
	return _ClassicPool.Contract.Token0(&_ClassicPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_ClassicPool *ClassicPoolCallerSession) Token0() (common.Address, error) {
	return _ClassicPool.Contract.Token0(&_ClassicPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ClassicPool *ClassicPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ClassicPool *ClassicPoolSession) Token1() (common.Address, error) {
	return _ClassicPool.Contract.Token1(&_ClassicPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_ClassicPool *ClassicPoolCallerSession) Token1() (common.Address, error) {
	return _ClassicPool.Contract.Token1(&_ClassicPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClassicPool *ClassicPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClassicPool *ClassicPoolSession) TotalSupply() (*big.Int, error) {
	return _ClassicPool.Contract.TotalSupply(&_ClassicPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ClassicPool *ClassicPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _ClassicPool.Contract.TotalSupply(&_ClassicPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ClassicPool *ClassicPoolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClassicPool.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ClassicPool *ClassicPoolSession) Vault() (common.Address, error) {
	return _ClassicPool.Contract.Vault(&_ClassicPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ClassicPool *ClassicPoolCallerSession) Vault() (common.Address, error) {
	return _ClassicPool.Contract.Vault(&_ClassicPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.Approve(&_ClassicPool.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.Approve(&_ClassicPool.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_ClassicPool *ClassicPoolTransactor) Burn(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "burn", _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_ClassicPool *ClassicPoolSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Burn(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_ClassicPool *ClassicPoolTransactorSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Burn(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolTransactor) BurnSingle(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "burnSingle", _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.BurnSingle(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolTransactorSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.BurnSingle(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_ClassicPool *ClassicPoolTransactor) Mint(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "mint", _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_ClassicPool *ClassicPoolSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Mint(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_ClassicPool *ClassicPoolTransactorSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Mint(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_ClassicPool *ClassicPoolTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "permit", _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_ClassicPool *ClassicPoolSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Permit(&_ClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_ClassicPool *ClassicPoolTransactorSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Permit(&_ClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_ClassicPool *ClassicPoolTransactor) Permit2(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "permit2", _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_ClassicPool *ClassicPoolSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Permit2(&_ClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_ClassicPool *ClassicPoolTransactorSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Permit2(&_ClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolTransactor) Swap(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "swap", _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Swap(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_ClassicPool *ClassicPoolTransactorSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _ClassicPool.Contract.Swap(&_ClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.Transfer(&_ClassicPool.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.Transfer(&_ClassicPool.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.TransferFrom(&_ClassicPool.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_ClassicPool *ClassicPoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ClassicPool.Contract.TransferFrom(&_ClassicPool.TransactOpts, _from, _to, _amount)
}

// ClassicPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ClassicPool contract.
type ClassicPoolApprovalIterator struct {
	Event *ClassicPoolApproval // Event containing the contract specifics and raw log

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
func (it *ClassicPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolApproval)
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
		it.Event = new(ClassicPoolApproval)
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
func (it *ClassicPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolApproval represents a Approval event raised by the ClassicPool contract.
type ClassicPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ClassicPool *ClassicPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ClassicPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolApprovalIterator{contract: _ClassicPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_ClassicPool *ClassicPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ClassicPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolApproval)
				if err := _ClassicPool.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseApproval(log types.Log) (*ClassicPoolApproval, error) {
	event := new(ClassicPoolApproval)
	if err := _ClassicPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClassicPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the ClassicPool contract.
type ClassicPoolBurnIterator struct {
	Event *ClassicPoolBurn // Event containing the contract specifics and raw log

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
func (it *ClassicPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolBurn)
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
		it.Event = new(ClassicPoolBurn)
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
func (it *ClassicPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolBurn represents a Burn event raised by the ClassicPool contract.
type ClassicPoolBurn struct {
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
func (_ClassicPool *ClassicPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ClassicPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolBurnIterator{contract: _ClassicPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ClassicPool *ClassicPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *ClassicPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolBurn)
				if err := _ClassicPool.contract.UnpackLog(event, "Burn", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseBurn(log types.Log) (*ClassicPoolBurn, error) {
	event := new(ClassicPoolBurn)
	if err := _ClassicPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClassicPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the ClassicPool contract.
type ClassicPoolMintIterator struct {
	Event *ClassicPoolMint // Event containing the contract specifics and raw log

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
func (it *ClassicPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolMint)
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
		it.Event = new(ClassicPoolMint)
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
func (it *ClassicPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolMint represents a Mint event raised by the ClassicPool contract.
type ClassicPoolMint struct {
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
func (_ClassicPool *ClassicPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ClassicPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolMintIterator{contract: _ClassicPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_ClassicPool *ClassicPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *ClassicPoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolMint)
				if err := _ClassicPool.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseMint(log types.Log) (*ClassicPoolMint, error) {
	event := new(ClassicPoolMint)
	if err := _ClassicPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClassicPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ClassicPool contract.
type ClassicPoolSwapIterator struct {
	Event *ClassicPoolSwap // Event containing the contract specifics and raw log

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
func (it *ClassicPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolSwap)
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
		it.Event = new(ClassicPoolSwap)
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
func (it *ClassicPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolSwap represents a Swap event raised by the ClassicPool contract.
type ClassicPoolSwap struct {
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
func (_ClassicPool *ClassicPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ClassicPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolSwapIterator{contract: _ClassicPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_ClassicPool *ClassicPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ClassicPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolSwap)
				if err := _ClassicPool.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseSwap(log types.Log) (*ClassicPoolSwap, error) {
	event := new(ClassicPoolSwap)
	if err := _ClassicPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClassicPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the ClassicPool contract.
type ClassicPoolSyncIterator struct {
	Event *ClassicPoolSync // Event containing the contract specifics and raw log

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
func (it *ClassicPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolSync)
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
		it.Event = new(ClassicPoolSync)
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
func (it *ClassicPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolSync represents a Sync event raised by the ClassicPool contract.
type ClassicPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ClassicPool *ClassicPoolFilterer) FilterSync(opts *bind.FilterOpts) (*ClassicPoolSyncIterator, error) {

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &ClassicPoolSyncIterator{contract: _ClassicPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_ClassicPool *ClassicPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *ClassicPoolSync) (event.Subscription, error) {

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolSync)
				if err := _ClassicPool.contract.UnpackLog(event, "Sync", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseSync(log types.Log) (*ClassicPoolSync, error) {
	event := new(ClassicPoolSync)
	if err := _ClassicPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClassicPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ClassicPool contract.
type ClassicPoolTransferIterator struct {
	Event *ClassicPoolTransfer // Event containing the contract specifics and raw log

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
func (it *ClassicPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolTransfer)
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
		it.Event = new(ClassicPoolTransfer)
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
func (it *ClassicPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolTransfer represents a Transfer event raised by the ClassicPool contract.
type ClassicPoolTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ClassicPool *ClassicPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ClassicPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolTransferIterator{contract: _ClassicPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_ClassicPool *ClassicPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ClassicPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ClassicPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolTransfer)
				if err := _ClassicPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ClassicPool *ClassicPoolFilterer) ParseTransfer(log types.Log) (*ClassicPoolTransfer, error) {
	event := new(ClassicPoolTransfer)
	if err := _ClassicPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

