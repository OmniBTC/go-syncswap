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

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wETH\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ETHTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashLoanFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashLoanFeePercentage\",\"type\":\"uint256\"}],\"name\":\"FlashLoanFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC3156_CALLBACK_SUCCESS\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"depositETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"flashFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC3156FlashBorrower\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFlashLoanRecipient\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"flashLoanMultiple\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"maxFlashLoan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFlashLoanFeePercentage\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_flashLoanFeeRecipient\",\"type\":\"address\"}],\"name\":\"setFlashLoanFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAndDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"mode\",\"type\":\"uint8\"}],\"name\":\"withdrawAlternative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Vault *VaultCaller) ERC3156CALLBACKSUCCESS(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "ERC3156_CALLBACK_SUCCESS")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Vault *VaultSession) ERC3156CALLBACKSUCCESS() ([32]byte, error) {
	return _Vault.Contract.ERC3156CALLBACKSUCCESS(&_Vault.CallOpts)
}

// ERC3156CALLBACKSUCCESS is a free data retrieval call binding the contract method 0xec85b12b.
//
// Solidity: function ERC3156_CALLBACK_SUCCESS() view returns(bytes32)
func (_Vault *VaultCallerSession) ERC3156CALLBACKSUCCESS() ([32]byte, error) {
	return _Vault.Contract.ERC3156CALLBACKSUCCESS(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Vault *VaultSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256 balance)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token, account)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Vault *VaultCaller) FlashFee(opts *bind.CallOpts, arg0 common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "flashFee", arg0, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Vault *VaultSession) FlashFee(arg0 common.Address, amount *big.Int) (*big.Int, error) {
	return _Vault.Contract.FlashFee(&_Vault.CallOpts, arg0, amount)
}

// FlashFee is a free data retrieval call binding the contract method 0xd9d98ce4.
//
// Solidity: function flashFee(address , uint256 amount) view returns(uint256)
func (_Vault *VaultCallerSession) FlashFee(arg0 common.Address, amount *big.Int) (*big.Int, error) {
	return _Vault.Contract.FlashFee(&_Vault.CallOpts, arg0, amount)
}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Vault *VaultCaller) FlashLoanFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "flashLoanFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Vault *VaultSession) FlashLoanFeePercentage() (*big.Int, error) {
	return _Vault.Contract.FlashLoanFeePercentage(&_Vault.CallOpts)
}

// FlashLoanFeePercentage is a free data retrieval call binding the contract method 0xc499f8ce.
//
// Solidity: function flashLoanFeePercentage() view returns(uint256)
func (_Vault *VaultCallerSession) FlashLoanFeePercentage() (*big.Int, error) {
	return _Vault.Contract.FlashLoanFeePercentage(&_Vault.CallOpts)
}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Vault *VaultCaller) FlashLoanFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "flashLoanFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Vault *VaultSession) FlashLoanFeeRecipient() (common.Address, error) {
	return _Vault.Contract.FlashLoanFeeRecipient(&_Vault.CallOpts)
}

// FlashLoanFeeRecipient is a free data retrieval call binding the contract method 0xa16e5112.
//
// Solidity: function flashLoanFeeRecipient() view returns(address)
func (_Vault *VaultCallerSession) FlashLoanFeeRecipient() (common.Address, error) {
	return _Vault.Contract.FlashLoanFeeRecipient(&_Vault.CallOpts)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Vault *VaultCaller) MaxFlashLoan(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "maxFlashLoan", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Vault *VaultSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxFlashLoan(&_Vault.CallOpts, token)
}

// MaxFlashLoan is a free data retrieval call binding the contract method 0x613255ab.
//
// Solidity: function maxFlashLoan(address token) view returns(uint256)
func (_Vault *VaultCallerSession) MaxFlashLoan(token common.Address) (*big.Int, error) {
	return _Vault.Contract.MaxFlashLoan(&_Vault.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vault *VaultCallerSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Vault *VaultCallerSession) Paused() (bool, error) {
	return _Vault.Contract.Paused(&_Vault.CallOpts)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Vault *VaultCaller) Reserves(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "reserves", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Vault *VaultSession) Reserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.Reserves(&_Vault.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256)
func (_Vault *VaultCallerSession) Reserves(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.Reserves(&_Vault.CallOpts, arg0)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Vault *VaultCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "wETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Vault *VaultSession) WETH() (common.Address, error) {
	return _Vault.Contract.WETH(&_Vault.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xf2428621.
//
// Solidity: function wETH() view returns(address)
func (_Vault *VaultCallerSession) WETH() (common.Address, error) {
	return _Vault.Contract.WETH(&_Vault.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, token common.Address, to common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", token, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Vault *VaultSession) Deposit(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, token, to)
}

// Deposit is a paid mutator transaction binding the contract method 0xf9609f08.
//
// Solidity: function deposit(address token, address to) payable returns(uint256 amount)
func (_Vault *VaultTransactorSession) Deposit(token common.Address, to common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, token, to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Vault *VaultTransactor) DepositETH(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositETH", to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Vault *VaultSession) DepositETH(to common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DepositETH(&_Vault.TransactOpts, to)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2d2da806.
//
// Solidity: function depositETH(address to) payable returns(uint256 amount)
func (_Vault *VaultTransactorSession) DepositETH(to common.Address) (*types.Transaction, error) {
	return _Vault.Contract.DepositETH(&_Vault.TransactOpts, to)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Vault *VaultTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "flashLoan", receiver, token, amount, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Vault *VaultSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoan(&_Vault.TransactOpts, receiver, token, amount, userData)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes userData) returns(bool)
func (_Vault *VaultTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoan(&_Vault.TransactOpts, receiver, token, amount, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultTransactor) FlashLoanMultiple(opts *bind.TransactOpts, recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "flashLoanMultiple", recipient, tokens, amounts, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultSession) FlashLoanMultiple(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoanMultiple(&_Vault.TransactOpts, recipient, tokens, amounts, userData)
}

// FlashLoanMultiple is a paid mutator transaction binding the contract method 0xcfaa541e.
//
// Solidity: function flashLoanMultiple(address recipient, address[] tokens, uint256[] amounts, bytes userData) returns()
func (_Vault *VaultTransactorSession) FlashLoanMultiple(recipient common.Address, tokens []common.Address, amounts []*big.Int, userData []byte) (*types.Transaction, error) {
	return _Vault.Contract.FlashLoanMultiple(&_Vault.TransactOpts, recipient, tokens, amounts, userData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vault.Contract.RenounceOwnership(&_Vault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vault *VaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vault.Contract.RenounceOwnership(&_Vault.TransactOpts)
}

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Vault *VaultTransactor) SetFlashLoanFeePercentage(opts *bind.TransactOpts, newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFlashLoanFeePercentage", newFlashLoanFeePercentage)
}

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Vault *VaultSession) SetFlashLoanFeePercentage(newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFlashLoanFeePercentage(&_Vault.TransactOpts, newFlashLoanFeePercentage)
}

// SetFlashLoanFeePercentage is a paid mutator transaction binding the contract method 0x6b6b9f69.
//
// Solidity: function setFlashLoanFeePercentage(uint256 newFlashLoanFeePercentage) returns()
func (_Vault *VaultTransactorSession) SetFlashLoanFeePercentage(newFlashLoanFeePercentage *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SetFlashLoanFeePercentage(&_Vault.TransactOpts, newFlashLoanFeePercentage)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Vault *VaultTransactor) SetFlashLoanFeeRecipient(opts *bind.TransactOpts, _flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setFlashLoanFeeRecipient", _flashLoanFeeRecipient)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Vault *VaultSession) SetFlashLoanFeeRecipient(_flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetFlashLoanFeeRecipient(&_Vault.TransactOpts, _flashLoanFeeRecipient)
}

// SetFlashLoanFeeRecipient is a paid mutator transaction binding the contract method 0xb914cc64.
//
// Solidity: function setFlashLoanFeeRecipient(address _flashLoanFeeRecipient) returns()
func (_Vault *VaultTransactorSession) SetFlashLoanFeeRecipient(_flashLoanFeeRecipient common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetFlashLoanFeeRecipient(&_Vault.TransactOpts, _flashLoanFeeRecipient)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Vault *VaultTransactor) SetPaused(opts *bind.TransactOpts, _status bool) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setPaused", _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Vault *VaultSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _Vault.Contract.SetPaused(&_Vault.TransactOpts, _status)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool _status) returns()
func (_Vault *VaultTransactorSession) SetPaused(_status bool) (*types.Transaction, error) {
	return _Vault.Contract.SetPaused(&_Vault.TransactOpts, _status)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) Transfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Vault *VaultTransactor) TransferAndDeposit(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferAndDeposit", token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Vault *VaultSession) TransferAndDeposit(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferAndDeposit(&_Vault.TransactOpts, token, to, amount)
}

// TransferAndDeposit is a paid mutator transaction binding the contract method 0x511de15b.
//
// Solidity: function transferAndDeposit(address token, address to, uint256 amount) payable returns(uint256)
func (_Vault *VaultTransactorSession) TransferAndDeposit(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferAndDeposit(&_Vault.TransactOpts, token, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vault.Contract.TransferOwnership(&_Vault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vault *VaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vault.Contract.TransferOwnership(&_Vault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, token, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) Withdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, token, to, amount)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Vault *VaultTransactor) WithdrawAlternative(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdrawAlternative", token, to, amount, mode)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Vault *VaultSession) WithdrawAlternative(token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawAlternative(&_Vault.TransactOpts, token, to, amount, mode)
}

// WithdrawAlternative is a paid mutator transaction binding the contract method 0x6cb568c1.
//
// Solidity: function withdrawAlternative(address token, address to, uint256 amount, uint8 mode) returns()
func (_Vault *VaultTransactorSession) WithdrawAlternative(token common.Address, to common.Address, amount *big.Int, mode uint8) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawAlternative(&_Vault.TransactOpts, token, to, amount, mode)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Vault *VaultTransactor) WithdrawETH(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdrawETH", to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Vault *VaultSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawETH(&_Vault.TransactOpts, to, amount)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x4782f779.
//
// Solidity: function withdrawETH(address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) WithdrawETH(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.WithdrawETH(&_Vault.TransactOpts, to, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// VaultFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the Vault contract.
type VaultFlashLoanIterator struct {
	Event *VaultFlashLoan // Event containing the contract specifics and raw log

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
func (it *VaultFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFlashLoan)
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
		it.Event = new(VaultFlashLoan)
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
func (it *VaultFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFlashLoan represents a FlashLoan event raised by the Vault contract.
type VaultFlashLoan struct {
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	FeeAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) FilterFlashLoan(opts *bind.FilterOpts, recipient []common.Address, token []common.Address) (*VaultFlashLoanIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultFlashLoanIterator{contract: _Vault.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *VaultFlashLoan, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "FlashLoan", recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFlashLoan)
				if err := _Vault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x0d7d75e01ab95780d3cd1c8ec0dd6c2ce19e3a20427eec8bf53283b6fb8e95f0.
//
// Solidity: event FlashLoan(address indexed recipient, address indexed token, uint256 amount, uint256 feeAmount)
func (_Vault *VaultFilterer) ParseFlashLoan(log types.Log) (*VaultFlashLoan, error) {
	event := new(VaultFlashLoan)
	if err := _Vault.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultFlashLoanFeePercentageChangedIterator is returned from FilterFlashLoanFeePercentageChanged and is used to iterate over the raw logs and unpacked data for FlashLoanFeePercentageChanged events raised by the Vault contract.
type VaultFlashLoanFeePercentageChangedIterator struct {
	Event *VaultFlashLoanFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultFlashLoanFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultFlashLoanFeePercentageChanged)
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
		it.Event = new(VaultFlashLoanFeePercentageChanged)
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
func (it *VaultFlashLoanFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultFlashLoanFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultFlashLoanFeePercentageChanged represents a FlashLoanFeePercentageChanged event raised by the Vault contract.
type VaultFlashLoanFeePercentageChanged struct {
	OldFlashLoanFeePercentage *big.Int
	NewFlashLoanFeePercentage *big.Int
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFlashLoanFeePercentageChanged is a free log retrieval operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Vault *VaultFilterer) FilterFlashLoanFeePercentageChanged(opts *bind.FilterOpts) (*VaultFlashLoanFeePercentageChangedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "FlashLoanFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return &VaultFlashLoanFeePercentageChangedIterator{contract: _Vault.contract, event: "FlashLoanFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchFlashLoanFeePercentageChanged is a free log subscription operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Vault *VaultFilterer) WatchFlashLoanFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultFlashLoanFeePercentageChanged) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "FlashLoanFeePercentageChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultFlashLoanFeePercentageChanged)
				if err := _Vault.contract.UnpackLog(event, "FlashLoanFeePercentageChanged", log); err != nil {
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

// ParseFlashLoanFeePercentageChanged is a log parse operation binding the contract event 0x36e8f57c180167765b2da71700ae4d0d3237d63cd1552cefa8bafca7c1d3fc3d.
//
// Solidity: event FlashLoanFeePercentageChanged(uint256 oldFlashLoanFeePercentage, uint256 newFlashLoanFeePercentage)
func (_Vault *VaultFilterer) ParseFlashLoanFeePercentageChanged(log types.Log) (*VaultFlashLoanFeePercentageChanged, error) {
	event := new(VaultFlashLoanFeePercentageChanged)
	if err := _Vault.contract.UnpackLog(event, "FlashLoanFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vault contract.
type VaultOwnershipTransferredIterator struct {
	Event *VaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultOwnershipTransferred)
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
		it.Event = new(VaultOwnershipTransferred)
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
func (it *VaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultOwnershipTransferred represents a OwnershipTransferred event raised by the Vault contract.
type VaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VaultOwnershipTransferredIterator{contract: _Vault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultOwnershipTransferred)
				if err := _Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vault *VaultFilterer) ParseOwnershipTransferred(log types.Log) (*VaultOwnershipTransferred, error) {
	event := new(VaultOwnershipTransferred)
	if err := _Vault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vault contract.
type VaultPausedIterator struct {
	Event *VaultPaused // Event containing the contract specifics and raw log

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
func (it *VaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPaused)
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
		it.Event = new(VaultPaused)
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
func (it *VaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPaused represents a Paused event raised by the Vault contract.
type VaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vault *VaultFilterer) FilterPaused(opts *bind.FilterOpts) (*VaultPausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VaultPausedIterator{contract: _Vault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vault *VaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VaultPaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPaused)
				if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Vault *VaultFilterer) ParsePaused(log types.Log) (*VaultPaused, error) {
	event := new(VaultPaused)
	if err := _Vault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Vault contract.
type VaultUnpausedIterator struct {
	Event *VaultUnpaused // Event containing the contract specifics and raw log

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
func (it *VaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnpaused)
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
		it.Event = new(VaultUnpaused)
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
func (it *VaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnpaused represents a Unpaused event raised by the Vault contract.
type VaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vault *VaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*VaultUnpausedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &VaultUnpausedIterator{contract: _Vault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vault *VaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *VaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnpaused)
				if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Vault *VaultFilterer) ParseUnpaused(log types.Log) (*VaultUnpaused, error) {
	event := new(VaultUnpaused)
	if err := _Vault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

