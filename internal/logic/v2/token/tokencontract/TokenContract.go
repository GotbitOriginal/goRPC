// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokencontract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TokencontractABI is the input ABI used to generate the binding from.
const TokencontractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_upgradedAddress\",\"type\":\"address\"}],\"name\":\"deprecate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"deprecated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"upgradedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maximumFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBasisPoints\",\"type\":\"uint256\"},{\"name\":\"newMaxFee\",\"type\":\"uint256\"}],\"name\":\"setParams\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"basisPointsRate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_UINT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"feeBasisPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"Params\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// Tokencontract is an auto generated Go binding around an Ethereum contract.
type Tokencontract struct {
	TokencontractCaller     // Read-only binding to the contract
	TokencontractTransactor // Write-only binding to the contract
	TokencontractFilterer   // Log filterer for contract events
}

// TokencontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokencontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokencontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokencontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokencontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokencontractSession struct {
	Contract     *Tokencontract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokencontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokencontractCallerSession struct {
	Contract *TokencontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokencontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokencontractTransactorSession struct {
	Contract     *TokencontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokencontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokencontractRaw struct {
	Contract *Tokencontract // Generic contract binding to access the raw methods on
}

// TokencontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokencontractCallerRaw struct {
	Contract *TokencontractCaller // Generic read-only contract binding to access the raw methods on
}

// TokencontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokencontractTransactorRaw struct {
	Contract *TokencontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokencontract creates a new instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontract(address common.Address, backend bind.ContractBackend) (*Tokencontract, error) {
	contract, err := bindTokencontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokencontract{TokencontractCaller: TokencontractCaller{contract: contract}, TokencontractTransactor: TokencontractTransactor{contract: contract}, TokencontractFilterer: TokencontractFilterer{contract: contract}}, nil
}

// NewTokencontractCaller creates a new read-only instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractCaller(address common.Address, caller bind.ContractCaller) (*TokencontractCaller, error) {
	contract, err := bindTokencontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokencontractCaller{contract: contract}, nil
}

// NewTokencontractTransactor creates a new write-only instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractTransactor(address common.Address, transactor bind.ContractTransactor) (*TokencontractTransactor, error) {
	contract, err := bindTokencontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokencontractTransactor{contract: contract}, nil
}

// NewTokencontractFilterer creates a new log filterer instance of Tokencontract, bound to a specific deployed contract.
func NewTokencontractFilterer(address common.Address, filterer bind.ContractFilterer) (*TokencontractFilterer, error) {
	contract, err := bindTokencontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokencontractFilterer{contract: contract}, nil
}

// bindTokencontract binds a generic wrapper to an already deployed contract.
func bindTokencontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokencontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencontract *TokencontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencontract.Contract.TokencontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencontract *TokencontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokencontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencontract *TokencontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencontract.Contract.TokencontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokencontract *TokencontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokencontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokencontract *TokencontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokencontract *TokencontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokencontract.Contract.contract.Transact(opts, method, params...)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Tokencontract *TokencontractCaller) MAXUINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "MAX_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Tokencontract *TokencontractSession) MAXUINT() (*big.Int, error) {
	return _Tokencontract.Contract.MAXUINT(&_Tokencontract.CallOpts)
}

// MAXUINT is a free data retrieval call binding the contract method 0xe5b5019a.
//
// Solidity: function MAX_UINT() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) MAXUINT() (*big.Int, error) {
	return _Tokencontract.Contract.MAXUINT(&_Tokencontract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokencontract *TokencontractCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokencontract *TokencontractSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Allowance(&_Tokencontract.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256 remaining)
func (_Tokencontract *TokencontractCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Allowance(&_Tokencontract.CallOpts, _owner, _spender)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Tokencontract *TokencontractCaller) Allowed(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "allowed", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Tokencontract *TokencontractSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Allowed(&_Tokencontract.CallOpts, arg0, arg1)
}

// Allowed is a free data retrieval call binding the contract method 0x5c658165.
//
// Solidity: function allowed(address , address ) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) Allowed(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Allowed(&_Tokencontract.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Tokencontract *TokencontractCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "balanceOf", who)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Tokencontract *TokencontractSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.BalanceOf(&_Tokencontract.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.BalanceOf(&_Tokencontract.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokencontract *TokencontractCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokencontract *TokencontractSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Balances(&_Tokencontract.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Tokencontract.Contract.Balances(&_Tokencontract.CallOpts, arg0)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Tokencontract *TokencontractCaller) BasisPointsRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "basisPointsRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Tokencontract *TokencontractSession) BasisPointsRate() (*big.Int, error) {
	return _Tokencontract.Contract.BasisPointsRate(&_Tokencontract.CallOpts)
}

// BasisPointsRate is a free data retrieval call binding the contract method 0xdd644f72.
//
// Solidity: function basisPointsRate() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) BasisPointsRate() (*big.Int, error) {
	return _Tokencontract.Contract.BasisPointsRate(&_Tokencontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokencontract *TokencontractCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokencontract *TokencontractSession) Decimals() (*big.Int, error) {
	return _Tokencontract.Contract.Decimals(&_Tokencontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) Decimals() (*big.Int, error) {
	return _Tokencontract.Contract.Decimals(&_Tokencontract.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Tokencontract *TokencontractCaller) Deprecated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "deprecated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Tokencontract *TokencontractSession) Deprecated() (bool, error) {
	return _Tokencontract.Contract.Deprecated(&_Tokencontract.CallOpts)
}

// Deprecated is a free data retrieval call binding the contract method 0x0e136b19.
//
// Solidity: function deprecated() view returns(bool)
func (_Tokencontract *TokencontractCallerSession) Deprecated() (bool, error) {
	return _Tokencontract.Contract.Deprecated(&_Tokencontract.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Tokencontract *TokencontractCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Tokencontract *TokencontractSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _Tokencontract.Contract.GetBlackListStatus(&_Tokencontract.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_Tokencontract *TokencontractCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _Tokencontract.Contract.GetBlackListStatus(&_Tokencontract.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokencontract *TokencontractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokencontract *TokencontractSession) GetOwner() (common.Address, error) {
	return _Tokencontract.Contract.GetOwner(&_Tokencontract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Tokencontract *TokencontractCallerSession) GetOwner() (common.Address, error) {
	return _Tokencontract.Contract.GetOwner(&_Tokencontract.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Tokencontract *TokencontractCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Tokencontract *TokencontractSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _Tokencontract.Contract.IsBlackListed(&_Tokencontract.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_Tokencontract *TokencontractCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _Tokencontract.Contract.IsBlackListed(&_Tokencontract.CallOpts, arg0)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Tokencontract *TokencontractCaller) MaximumFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "maximumFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Tokencontract *TokencontractSession) MaximumFee() (*big.Int, error) {
	return _Tokencontract.Contract.MaximumFee(&_Tokencontract.CallOpts)
}

// MaximumFee is a free data retrieval call binding the contract method 0x35390714.
//
// Solidity: function maximumFee() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) MaximumFee() (*big.Int, error) {
	return _Tokencontract.Contract.MaximumFee(&_Tokencontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractSession) Name() (string, error) {
	return _Tokencontract.Contract.Name(&_Tokencontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Tokencontract *TokencontractCallerSession) Name() (string, error) {
	return _Tokencontract.Contract.Name(&_Tokencontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractSession) Owner() (common.Address, error) {
	return _Tokencontract.Contract.Owner(&_Tokencontract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokencontract *TokencontractCallerSession) Owner() (common.Address, error) {
	return _Tokencontract.Contract.Owner(&_Tokencontract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractSession) Paused() (bool, error) {
	return _Tokencontract.Contract.Paused(&_Tokencontract.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tokencontract *TokencontractCallerSession) Paused() (bool, error) {
	return _Tokencontract.Contract.Paused(&_Tokencontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractSession) Symbol() (string, error) {
	return _Tokencontract.Contract.Symbol(&_Tokencontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Tokencontract *TokencontractCallerSession) Symbol() (string, error) {
	return _Tokencontract.Contract.Symbol(&_Tokencontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractSession) TotalSupply() (*big.Int, error) {
	return _Tokencontract.Contract.TotalSupply(&_Tokencontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Tokencontract *TokencontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Tokencontract.Contract.TotalSupply(&_Tokencontract.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Tokencontract *TokencontractCaller) UpgradedAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokencontract.contract.Call(opts, &out, "upgradedAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Tokencontract *TokencontractSession) UpgradedAddress() (common.Address, error) {
	return _Tokencontract.Contract.UpgradedAddress(&_Tokencontract.CallOpts)
}

// UpgradedAddress is a free data retrieval call binding the contract method 0x26976e3f.
//
// Solidity: function upgradedAddress() view returns(address)
func (_Tokencontract *TokencontractCallerSession) UpgradedAddress() (common.Address, error) {
	return _Tokencontract.Contract.UpgradedAddress(&_Tokencontract.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Tokencontract *TokencontractTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Tokencontract *TokencontractSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.AddBlackList(&_Tokencontract.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_Tokencontract *TokencontractTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.AddBlackList(&_Tokencontract.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Tokencontract *TokencontractSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Approve(&_Tokencontract.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Approve(&_Tokencontract.TransactOpts, _spender, _value)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Tokencontract *TokencontractTransactor) Deprecate(opts *bind.TransactOpts, _upgradedAddress common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "deprecate", _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Tokencontract *TokencontractSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.Deprecate(&_Tokencontract.TransactOpts, _upgradedAddress)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0753c30c.
//
// Solidity: function deprecate(address _upgradedAddress) returns()
func (_Tokencontract *TokencontractTransactorSession) Deprecate(_upgradedAddress common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.Deprecate(&_Tokencontract.TransactOpts, _upgradedAddress)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Tokencontract *TokencontractTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Tokencontract *TokencontractSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.DestroyBlackFunds(&_Tokencontract.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_Tokencontract *TokencontractTransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.DestroyBlackFunds(&_Tokencontract.TransactOpts, _blackListedUser)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Tokencontract *TokencontractTransactor) Issue(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "issue", amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Tokencontract *TokencontractSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Issue(&_Tokencontract.TransactOpts, amount)
}

// Issue is a paid mutator transaction binding the contract method 0xcc872b66.
//
// Solidity: function issue(uint256 amount) returns()
func (_Tokencontract *TokencontractTransactorSession) Issue(amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Issue(&_Tokencontract.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractSession) Pause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Pause(&_Tokencontract.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tokencontract *TokencontractTransactorSession) Pause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Pause(&_Tokencontract.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Tokencontract *TokencontractTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Tokencontract *TokencontractSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Redeem(&_Tokencontract.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Tokencontract *TokencontractTransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Redeem(&_Tokencontract.TransactOpts, amount)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Tokencontract *TokencontractTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Tokencontract *TokencontractSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.RemoveBlackList(&_Tokencontract.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_Tokencontract *TokencontractTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.RemoveBlackList(&_Tokencontract.TransactOpts, _clearedUser)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Tokencontract *TokencontractTransactor) SetParams(opts *bind.TransactOpts, newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "setParams", newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Tokencontract *TokencontractSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.SetParams(&_Tokencontract.TransactOpts, newBasisPoints, newMaxFee)
}

// SetParams is a paid mutator transaction binding the contract method 0xc0324c77.
//
// Solidity: function setParams(uint256 newBasisPoints, uint256 newMaxFee) returns()
func (_Tokencontract *TokencontractTransactorSession) SetParams(newBasisPoints *big.Int, newMaxFee *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.SetParams(&_Tokencontract.TransactOpts, newBasisPoints, newMaxFee)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Transfer(&_Tokencontract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.Transfer(&_Tokencontract.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferFrom(&_Tokencontract.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns()
func (_Tokencontract *TokencontractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferFrom(&_Tokencontract.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokencontract *TokencontractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokencontract *TokencontractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferOwnership(&_Tokencontract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokencontract *TokencontractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokencontract.Contract.TransferOwnership(&_Tokencontract.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokencontract.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractSession) Unpause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Unpause(&_Tokencontract.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tokencontract *TokencontractTransactorSession) Unpause() (*types.Transaction, error) {
	return _Tokencontract.Contract.Unpause(&_Tokencontract.TransactOpts)
}

// TokencontractAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the Tokencontract contract.
type TokencontractAddedBlackListIterator struct {
	Event *TokencontractAddedBlackList // Event containing the contract specifics and raw log

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
func (it *TokencontractAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractAddedBlackList)
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
		it.Event = new(TokencontractAddedBlackList)
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
func (it *TokencontractAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractAddedBlackList represents a AddedBlackList event raised by the Tokencontract contract.
type TokencontractAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*TokencontractAddedBlackListIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &TokencontractAddedBlackListIterator{contract: _Tokencontract.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *TokencontractAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractAddedBlackList)
				if err := _Tokencontract.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) ParseAddedBlackList(log types.Log) (*TokencontractAddedBlackList, error) {
	event := new(TokencontractAddedBlackList)
	if err := _Tokencontract.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Tokencontract contract.
type TokencontractApprovalIterator struct {
	Event *TokencontractApproval // Event containing the contract specifics and raw log

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
func (it *TokencontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractApproval)
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
		it.Event = new(TokencontractApproval)
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
func (it *TokencontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractApproval represents a Approval event raised by the Tokencontract contract.
type TokencontractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokencontract *TokencontractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokencontractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractApprovalIterator{contract: _Tokencontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokencontract *TokencontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokencontractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractApproval)
				if err := _Tokencontract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Tokencontract *TokencontractFilterer) ParseApproval(log types.Log) (*TokencontractApproval, error) {
	event := new(TokencontractApproval)
	if err := _Tokencontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the Tokencontract contract.
type TokencontractDeprecateIterator struct {
	Event *TokencontractDeprecate // Event containing the contract specifics and raw log

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
func (it *TokencontractDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractDeprecate)
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
		it.Event = new(TokencontractDeprecate)
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
func (it *TokencontractDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractDeprecate represents a Deprecate event raised by the Tokencontract contract.
type TokencontractDeprecate struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Tokencontract *TokencontractFilterer) FilterDeprecate(opts *bind.FilterOpts) (*TokencontractDeprecateIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return &TokencontractDeprecateIterator{contract: _Tokencontract.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Tokencontract *TokencontractFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *TokencontractDeprecate) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Deprecate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractDeprecate)
				if err := _Tokencontract.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address newAddress)
func (_Tokencontract *TokencontractFilterer) ParseDeprecate(log types.Log) (*TokencontractDeprecate, error) {
	event := new(TokencontractDeprecate)
	if err := _Tokencontract.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the Tokencontract contract.
type TokencontractDestroyedBlackFundsIterator struct {
	Event *TokencontractDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *TokencontractDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractDestroyedBlackFunds)
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
		it.Event = new(TokencontractDestroyedBlackFunds)
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
func (it *TokencontractDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the Tokencontract contract.
type TokencontractDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Tokencontract *TokencontractFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*TokencontractDestroyedBlackFundsIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &TokencontractDestroyedBlackFundsIterator{contract: _Tokencontract.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Tokencontract *TokencontractFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *TokencontractDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractDestroyedBlackFunds)
				if err := _Tokencontract.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_Tokencontract *TokencontractFilterer) ParseDestroyedBlackFunds(log types.Log) (*TokencontractDestroyedBlackFunds, error) {
	event := new(TokencontractDestroyedBlackFunds)
	if err := _Tokencontract.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the Tokencontract contract.
type TokencontractIssueIterator struct {
	Event *TokencontractIssue // Event containing the contract specifics and raw log

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
func (it *TokencontractIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractIssue)
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
		it.Event = new(TokencontractIssue)
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
func (it *TokencontractIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractIssue represents a Issue event raised by the Tokencontract contract.
type TokencontractIssue struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Tokencontract *TokencontractFilterer) FilterIssue(opts *bind.FilterOpts) (*TokencontractIssueIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return &TokencontractIssueIterator{contract: _Tokencontract.contract, event: "Issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Tokencontract *TokencontractFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *TokencontractIssue) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractIssue)
				if err := _Tokencontract.contract.UnpackLog(event, "Issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0xcb8241adb0c3fdb35b70c24ce35c5eb0c17af7431c99f827d44a445ca624176a.
//
// Solidity: event Issue(uint256 amount)
func (_Tokencontract *TokencontractFilterer) ParseIssue(log types.Log) (*TokencontractIssue, error) {
	event := new(TokencontractIssue)
	if err := _Tokencontract.contract.UnpackLog(event, "Issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractParamsIterator is returned from FilterParams and is used to iterate over the raw logs and unpacked data for Params events raised by the Tokencontract contract.
type TokencontractParamsIterator struct {
	Event *TokencontractParams // Event containing the contract specifics and raw log

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
func (it *TokencontractParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractParams)
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
		it.Event = new(TokencontractParams)
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
func (it *TokencontractParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractParams represents a Params event raised by the Tokencontract contract.
type TokencontractParams struct {
	FeeBasisPoints *big.Int
	MaxFee         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterParams is a free log retrieval operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Tokencontract *TokencontractFilterer) FilterParams(opts *bind.FilterOpts) (*TokencontractParamsIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return &TokencontractParamsIterator{contract: _Tokencontract.contract, event: "Params", logs: logs, sub: sub}, nil
}

// WatchParams is a free log subscription operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Tokencontract *TokencontractFilterer) WatchParams(opts *bind.WatchOpts, sink chan<- *TokencontractParams) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Params")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractParams)
				if err := _Tokencontract.contract.UnpackLog(event, "Params", log); err != nil {
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

// ParseParams is a log parse operation binding the contract event 0xb044a1e409eac5c48e5af22d4af52670dd1a99059537a78b31b48c6500a6354e.
//
// Solidity: event Params(uint256 feeBasisPoints, uint256 maxFee)
func (_Tokencontract *TokencontractFilterer) ParseParams(log types.Log) (*TokencontractParams, error) {
	event := new(TokencontractParams)
	if err := _Tokencontract.contract.UnpackLog(event, "Params", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Tokencontract contract.
type TokencontractPauseIterator struct {
	Event *TokencontractPause // Event containing the contract specifics and raw log

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
func (it *TokencontractPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractPause)
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
		it.Event = new(TokencontractPause)
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
func (it *TokencontractPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractPause represents a Pause event raised by the Tokencontract contract.
type TokencontractPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Tokencontract *TokencontractFilterer) FilterPause(opts *bind.FilterOpts) (*TokencontractPauseIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &TokencontractPauseIterator{contract: _Tokencontract.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Tokencontract *TokencontractFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *TokencontractPause) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractPause)
				if err := _Tokencontract.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Tokencontract *TokencontractFilterer) ParsePause(log types.Log) (*TokencontractPause, error) {
	event := new(TokencontractPause)
	if err := _Tokencontract.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Tokencontract contract.
type TokencontractRedeemIterator struct {
	Event *TokencontractRedeem // Event containing the contract specifics and raw log

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
func (it *TokencontractRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractRedeem)
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
		it.Event = new(TokencontractRedeem)
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
func (it *TokencontractRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractRedeem represents a Redeem event raised by the Tokencontract contract.
type TokencontractRedeem struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Tokencontract *TokencontractFilterer) FilterRedeem(opts *bind.FilterOpts) (*TokencontractRedeemIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &TokencontractRedeemIterator{contract: _Tokencontract.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Tokencontract *TokencontractFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *TokencontractRedeem) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractRedeem)
				if err := _Tokencontract.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x702d5967f45f6513a38ffc42d6ba9bf230bd40e8f53b16363c7eb4fd2deb9a44.
//
// Solidity: event Redeem(uint256 amount)
func (_Tokencontract *TokencontractFilterer) ParseRedeem(log types.Log) (*TokencontractRedeem, error) {
	event := new(TokencontractRedeem)
	if err := _Tokencontract.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the Tokencontract contract.
type TokencontractRemovedBlackListIterator struct {
	Event *TokencontractRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *TokencontractRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractRemovedBlackList)
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
		it.Event = new(TokencontractRemovedBlackList)
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
func (it *TokencontractRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractRemovedBlackList represents a RemovedBlackList event raised by the Tokencontract contract.
type TokencontractRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*TokencontractRemovedBlackListIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &TokencontractRemovedBlackListIterator{contract: _Tokencontract.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *TokencontractRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractRemovedBlackList)
				if err := _Tokencontract.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_Tokencontract *TokencontractFilterer) ParseRemovedBlackList(log types.Log) (*TokencontractRemovedBlackList, error) {
	event := new(TokencontractRemovedBlackList)
	if err := _Tokencontract.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Tokencontract contract.
type TokencontractTransferIterator struct {
	Event *TokencontractTransfer // Event containing the contract specifics and raw log

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
func (it *TokencontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractTransfer)
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
		it.Event = new(TokencontractTransfer)
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
func (it *TokencontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractTransfer represents a Transfer event raised by the Tokencontract contract.
type TokencontractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokencontract *TokencontractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokencontractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokencontractTransferIterator{contract: _Tokencontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokencontract *TokencontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokencontractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractTransfer)
				if err := _Tokencontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Tokencontract *TokencontractFilterer) ParseTransfer(log types.Log) (*TokencontractTransfer, error) {
	event := new(TokencontractTransfer)
	if err := _Tokencontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokencontractUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Tokencontract contract.
type TokencontractUnpauseIterator struct {
	Event *TokencontractUnpause // Event containing the contract specifics and raw log

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
func (it *TokencontractUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokencontractUnpause)
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
		it.Event = new(TokencontractUnpause)
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
func (it *TokencontractUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokencontractUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokencontractUnpause represents a Unpause event raised by the Tokencontract contract.
type TokencontractUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Tokencontract *TokencontractFilterer) FilterUnpause(opts *bind.FilterOpts) (*TokencontractUnpauseIterator, error) {

	logs, sub, err := _Tokencontract.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &TokencontractUnpauseIterator{contract: _Tokencontract.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Tokencontract *TokencontractFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *TokencontractUnpause) (event.Subscription, error) {

	logs, sub, err := _Tokencontract.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokencontractUnpause)
				if err := _Tokencontract.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Tokencontract *TokencontractFilterer) ParseUnpause(log types.Log) (*TokencontractUnpause, error) {
	event := new(TokencontractUnpause)
	if err := _Tokencontract.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
