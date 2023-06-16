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

// ClassicPoolFactoryMetaData contains all meta data concerning the ClassicPoolFactory contract.
var ClassicPoolFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_master\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidTokens\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"createPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"deployData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ClassicPoolFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ClassicPoolFactoryMetaData.ABI instead.
var ClassicPoolFactoryABI = ClassicPoolFactoryMetaData.ABI

// ClassicPoolFactory is an auto generated Go binding around an Ethereum contract.
type ClassicPoolFactory struct {
	ClassicPoolFactoryCaller     // Read-only binding to the contract
	ClassicPoolFactoryTransactor // Write-only binding to the contract
	ClassicPoolFactoryFilterer   // Log filterer for contract events
}

// ClassicPoolFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClassicPoolFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClassicPoolFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClassicPoolFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClassicPoolFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClassicPoolFactorySession struct {
	Contract     *ClassicPoolFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ClassicPoolFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClassicPoolFactoryCallerSession struct {
	Contract *ClassicPoolFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ClassicPoolFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClassicPoolFactoryTransactorSession struct {
	Contract     *ClassicPoolFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ClassicPoolFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClassicPoolFactoryRaw struct {
	Contract *ClassicPoolFactory // Generic contract binding to access the raw methods on
}

// ClassicPoolFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClassicPoolFactoryCallerRaw struct {
	Contract *ClassicPoolFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ClassicPoolFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClassicPoolFactoryTransactorRaw struct {
	Contract *ClassicPoolFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClassicPoolFactory creates a new instance of ClassicPoolFactory, bound to a specific deployed contract.
func NewClassicPoolFactory(address common.Address, backend bind.ContractBackend) (*ClassicPoolFactory, error) {
	contract, err := bindClassicPoolFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFactory{ClassicPoolFactoryCaller: ClassicPoolFactoryCaller{contract: contract}, ClassicPoolFactoryTransactor: ClassicPoolFactoryTransactor{contract: contract}, ClassicPoolFactoryFilterer: ClassicPoolFactoryFilterer{contract: contract}}, nil
}

// NewClassicPoolFactoryCaller creates a new read-only instance of ClassicPoolFactory, bound to a specific deployed contract.
func NewClassicPoolFactoryCaller(address common.Address, caller bind.ContractCaller) (*ClassicPoolFactoryCaller, error) {
	contract, err := bindClassicPoolFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFactoryCaller{contract: contract}, nil
}

// NewClassicPoolFactoryTransactor creates a new write-only instance of ClassicPoolFactory, bound to a specific deployed contract.
func NewClassicPoolFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ClassicPoolFactoryTransactor, error) {
	contract, err := bindClassicPoolFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFactoryTransactor{contract: contract}, nil
}

// NewClassicPoolFactoryFilterer creates a new log filterer instance of ClassicPoolFactory, bound to a specific deployed contract.
func NewClassicPoolFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ClassicPoolFactoryFilterer, error) {
	contract, err := bindClassicPoolFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFactoryFilterer{contract: contract}, nil
}

// bindClassicPoolFactory binds a generic wrapper to an already deployed contract.
func bindClassicPoolFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ClassicPoolFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClassicPoolFactory *ClassicPoolFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClassicPoolFactory.Contract.ClassicPoolFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClassicPoolFactory *ClassicPoolFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.ClassicPoolFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClassicPoolFactory *ClassicPoolFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.ClassicPoolFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClassicPoolFactory *ClassicPoolFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClassicPoolFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClassicPoolFactory *ClassicPoolFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClassicPoolFactory *ClassicPoolFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.contract.Transact(opts, method, params...)
}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes deployData)
func (_ClassicPoolFactory *ClassicPoolFactoryCaller) GetDeployData(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _ClassicPoolFactory.contract.Call(opts, &out, "getDeployData")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes deployData)
func (_ClassicPoolFactory *ClassicPoolFactorySession) GetDeployData() ([]byte, error) {
	return _ClassicPoolFactory.Contract.GetDeployData(&_ClassicPoolFactory.CallOpts)
}

// GetDeployData is a free data retrieval call binding the contract method 0xd039f622.
//
// Solidity: function getDeployData() view returns(bytes deployData)
func (_ClassicPoolFactory *ClassicPoolFactoryCallerSession) GetDeployData() ([]byte, error) {
	return _ClassicPoolFactory.Contract.GetDeployData(&_ClassicPoolFactory.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address , address ) view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactoryCaller) GetPool(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ClassicPoolFactory.contract.Call(opts, &out, "getPool", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address , address ) view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactorySession) GetPool(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _ClassicPoolFactory.Contract.GetPool(&_ClassicPoolFactory.CallOpts, arg0, arg1)
}

// GetPool is a free data retrieval call binding the contract method 0x531aa03e.
//
// Solidity: function getPool(address , address ) view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactoryCallerSession) GetPool(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _ClassicPoolFactory.Contract.GetPool(&_ClassicPoolFactory.CallOpts, arg0, arg1)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ClassicPoolFactory *ClassicPoolFactoryCaller) GetSwapFee(opts *bind.CallOpts, pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _ClassicPoolFactory.contract.Call(opts, &out, "getSwapFee", pool, sender, tokenIn, tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ClassicPoolFactory *ClassicPoolFactorySession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ClassicPoolFactory.Contract.GetSwapFee(&_ClassicPoolFactory.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x4625a94d.
//
// Solidity: function getSwapFee(address pool, address sender, address tokenIn, address tokenOut, bytes data) view returns(uint24 swapFee)
func (_ClassicPoolFactory *ClassicPoolFactoryCallerSession) GetSwapFee(pool common.Address, sender common.Address, tokenIn common.Address, tokenOut common.Address, data []byte) (*big.Int, error) {
	return _ClassicPoolFactory.Contract.GetSwapFee(&_ClassicPoolFactory.CallOpts, pool, sender, tokenIn, tokenOut, data)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactoryCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClassicPoolFactory.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactorySession) Master() (common.Address, error) {
	return _ClassicPoolFactory.Contract.Master(&_ClassicPoolFactory.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_ClassicPoolFactory *ClassicPoolFactoryCallerSession) Master() (common.Address, error) {
	return _ClassicPoolFactory.Contract.Master(&_ClassicPoolFactory.CallOpts)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ClassicPoolFactory *ClassicPoolFactoryTransactor) CreatePool(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ClassicPoolFactory.contract.Transact(opts, "createPool", data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ClassicPoolFactory *ClassicPoolFactorySession) CreatePool(data []byte) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.CreatePool(&_ClassicPoolFactory.TransactOpts, data)
}

// CreatePool is a paid mutator transaction binding the contract method 0x13b8683f.
//
// Solidity: function createPool(bytes data) returns(address pool)
func (_ClassicPoolFactory *ClassicPoolFactoryTransactorSession) CreatePool(data []byte) (*types.Transaction, error) {
	return _ClassicPoolFactory.Contract.CreatePool(&_ClassicPoolFactory.TransactOpts, data)
}

// ClassicPoolFactoryPoolCreatedIterator is returned from FilterPoolCreated and is used to iterate over the raw logs and unpacked data for PoolCreated events raised by the ClassicPoolFactory contract.
type ClassicPoolFactoryPoolCreatedIterator struct {
	Event *ClassicPoolFactoryPoolCreated // Event containing the contract specifics and raw log

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
func (it *ClassicPoolFactoryPoolCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClassicPoolFactoryPoolCreated)
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
		it.Event = new(ClassicPoolFactoryPoolCreated)
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
func (it *ClassicPoolFactoryPoolCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClassicPoolFactoryPoolCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClassicPoolFactoryPoolCreated represents a PoolCreated event raised by the ClassicPoolFactory contract.
type ClassicPoolFactoryPoolCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pool   common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolCreated is a free log retrieval operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ClassicPoolFactory *ClassicPoolFactoryFilterer) FilterPoolCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*ClassicPoolFactoryPoolCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ClassicPoolFactory.contract.FilterLogs(opts, "PoolCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &ClassicPoolFactoryPoolCreatedIterator{contract: _ClassicPoolFactory.contract, event: "PoolCreated", logs: logs, sub: sub}, nil
}

// WatchPoolCreated is a free log subscription operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ClassicPoolFactory *ClassicPoolFactoryFilterer) WatchPoolCreated(opts *bind.WatchOpts, sink chan<- *ClassicPoolFactoryPoolCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _ClassicPoolFactory.contract.WatchLogs(opts, "PoolCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClassicPoolFactoryPoolCreated)
				if err := _ClassicPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
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

// ParsePoolCreated is a log parse operation binding the contract event 0x9c5d829b9b23efc461f9aeef91979ec04bb903feb3bee4f26d22114abfc7335b.
//
// Solidity: event PoolCreated(address indexed token0, address indexed token1, address pool)
func (_ClassicPoolFactory *ClassicPoolFactoryFilterer) ParsePoolCreated(log types.Log) (*ClassicPoolFactoryPoolCreated, error) {
	event := new(ClassicPoolFactoryPoolCreated)
	if err := _ClassicPoolFactory.contract.UnpackLog(event, "PoolCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

