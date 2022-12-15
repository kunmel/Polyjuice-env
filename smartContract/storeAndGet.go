// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontract

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

// StoreAndGetMetaData contains all meta data concerning the StoreAndGet contract.
var StoreAndGetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"}],\"name\":\"batchRetrieve\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_keys\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_values\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"_version\",\"type\":\"string[]\"}],\"name\":\"batchStore\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"}],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_key\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StoreAndGetABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreAndGetMetaData.ABI instead.
var StoreAndGetABI = StoreAndGetMetaData.ABI

// StoreAndGet is an auto generated Go binding around an Ethereum contract.
type StoreAndGet struct {
	StoreAndGetCaller     // Read-only binding to the contract
	StoreAndGetTransactor // Write-only binding to the contract
	StoreAndGetFilterer   // Log filterer for contract events
}

// StoreAndGetCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreAndGetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreAndGetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreAndGetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreAndGetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreAndGetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreAndGetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreAndGetSession struct {
	Contract     *StoreAndGet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreAndGetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreAndGetCallerSession struct {
	Contract *StoreAndGetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StoreAndGetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreAndGetTransactorSession struct {
	Contract     *StoreAndGetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StoreAndGetRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreAndGetRaw struct {
	Contract *StoreAndGet // Generic contract binding to access the raw methods on
}

// StoreAndGetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreAndGetCallerRaw struct {
	Contract *StoreAndGetCaller // Generic read-only contract binding to access the raw methods on
}

// StoreAndGetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreAndGetTransactorRaw struct {
	Contract *StoreAndGetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoreAndGet creates a new instance of StoreAndGet, bound to a specific deployed contract.
func NewStoreAndGet(address common.Address, backend bind.ContractBackend) (*StoreAndGet, error) {
	contract, err := bindStoreAndGet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StoreAndGet{StoreAndGetCaller: StoreAndGetCaller{contract: contract}, StoreAndGetTransactor: StoreAndGetTransactor{contract: contract}, StoreAndGetFilterer: StoreAndGetFilterer{contract: contract}}, nil
}

// NewStoreAndGetCaller creates a new read-only instance of StoreAndGet, bound to a specific deployed contract.
func NewStoreAndGetCaller(address common.Address, caller bind.ContractCaller) (*StoreAndGetCaller, error) {
	contract, err := bindStoreAndGet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreAndGetCaller{contract: contract}, nil
}

// NewStoreAndGetTransactor creates a new write-only instance of StoreAndGet, bound to a specific deployed contract.
func NewStoreAndGetTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreAndGetTransactor, error) {
	contract, err := bindStoreAndGet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreAndGetTransactor{contract: contract}, nil
}

// NewStoreAndGetFilterer creates a new log filterer instance of StoreAndGet, bound to a specific deployed contract.
func NewStoreAndGetFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreAndGetFilterer, error) {
	contract, err := bindStoreAndGet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreAndGetFilterer{contract: contract}, nil
}

// bindStoreAndGet binds a generic wrapper to an already deployed contract.
func bindStoreAndGet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreAndGetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoreAndGet *StoreAndGetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoreAndGet.Contract.StoreAndGetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoreAndGet *StoreAndGetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoreAndGet.Contract.StoreAndGetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoreAndGet *StoreAndGetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoreAndGet.Contract.StoreAndGetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StoreAndGet *StoreAndGetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StoreAndGet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StoreAndGet *StoreAndGetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StoreAndGet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StoreAndGet *StoreAndGetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StoreAndGet.Contract.contract.Transact(opts, method, params...)
}

// BatchRetrieve is a free data retrieval call binding the contract method 0x437d851a.
//
// Solidity: function batchRetrieve(string[] _keys) view returns(string[], string[])
func (_StoreAndGet *StoreAndGetCaller) BatchRetrieve(opts *bind.CallOpts, _keys []string) ([]string, []string, error) {
	var out []interface{}
	err := _StoreAndGet.contract.Call(opts, &out, "batchRetrieve", _keys)

	if err != nil {
		return *new([]string), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)

	return out0, out1, err

}

// BatchRetrieve is a free data retrieval call binding the contract method 0x437d851a.
//
// Solidity: function batchRetrieve(string[] _keys) view returns(string[], string[])
func (_StoreAndGet *StoreAndGetSession) BatchRetrieve(_keys []string) ([]string, []string, error) {
	return _StoreAndGet.Contract.BatchRetrieve(&_StoreAndGet.CallOpts, _keys)
}

// BatchRetrieve is a free data retrieval call binding the contract method 0x437d851a.
//
// Solidity: function batchRetrieve(string[] _keys) view returns(string[], string[])
func (_StoreAndGet *StoreAndGetCallerSession) BatchRetrieve(_keys []string) ([]string, []string, error) {
	return _StoreAndGet.Contract.BatchRetrieve(&_StoreAndGet.CallOpts, _keys)
}

// Retrieve is a free data retrieval call binding the contract method 0x64cc7327.
//
// Solidity: function retrieve(string _key) view returns(string, string, string)
func (_StoreAndGet *StoreAndGetCaller) Retrieve(opts *bind.CallOpts, _key string) (string, string, string, error) {
	var out []interface{}
	err := _StoreAndGet.contract.Call(opts, &out, "retrieve", _key)

	if err != nil {
		return *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// Retrieve is a free data retrieval call binding the contract method 0x64cc7327.
//
// Solidity: function retrieve(string _key) view returns(string, string, string)
func (_StoreAndGet *StoreAndGetSession) Retrieve(_key string) (string, string, string, error) {
	return _StoreAndGet.Contract.Retrieve(&_StoreAndGet.CallOpts, _key)
}

// Retrieve is a free data retrieval call binding the contract method 0x64cc7327.
//
// Solidity: function retrieve(string _key) view returns(string, string, string)
func (_StoreAndGet *StoreAndGetCallerSession) Retrieve(_key string) (string, string, string, error) {
	return _StoreAndGet.Contract.Retrieve(&_StoreAndGet.CallOpts, _key)
}

// BatchStore is a paid mutator transaction binding the contract method 0x6fd56dce.
//
// Solidity: function batchStore(string[] _keys, string[] _values, string[] _version) payable returns()
func (_StoreAndGet *StoreAndGetTransactor) BatchStore(opts *bind.TransactOpts, _keys []string, _values []string, _version []string) (*types.Transaction, error) {
	return _StoreAndGet.contract.Transact(opts, "batchStore", _keys, _values, _version)
}

// BatchStore is a paid mutator transaction binding the contract method 0x6fd56dce.
//
// Solidity: function batchStore(string[] _keys, string[] _values, string[] _version) payable returns()
func (_StoreAndGet *StoreAndGetSession) BatchStore(_keys []string, _values []string, _version []string) (*types.Transaction, error) {
	return _StoreAndGet.Contract.BatchStore(&_StoreAndGet.TransactOpts, _keys, _values, _version)
}

// BatchStore is a paid mutator transaction binding the contract method 0x6fd56dce.
//
// Solidity: function batchStore(string[] _keys, string[] _values, string[] _version) payable returns()
func (_StoreAndGet *StoreAndGetTransactorSession) BatchStore(_keys []string, _values []string, _version []string) (*types.Transaction, error) {
	return _StoreAndGet.Contract.BatchStore(&_StoreAndGet.TransactOpts, _keys, _values, _version)
}

// Store is a paid mutator transaction binding the contract method 0x4ae69713.
//
// Solidity: function store(string _key, string _value, string _version) payable returns()
func (_StoreAndGet *StoreAndGetTransactor) Store(opts *bind.TransactOpts, _key string, _value string, _version string) (*types.Transaction, error) {
	return _StoreAndGet.contract.Transact(opts, "store", _key, _value, _version)
}

// Store is a paid mutator transaction binding the contract method 0x4ae69713.
//
// Solidity: function store(string _key, string _value, string _version) payable returns()
func (_StoreAndGet *StoreAndGetSession) Store(_key string, _value string, _version string) (*types.Transaction, error) {
	return _StoreAndGet.Contract.Store(&_StoreAndGet.TransactOpts, _key, _value, _version)
}

// Store is a paid mutator transaction binding the contract method 0x4ae69713.
//
// Solidity: function store(string _key, string _value, string _version) payable returns()
func (_StoreAndGet *StoreAndGetTransactorSession) Store(_key string, _value string, _version string) (*types.Transaction, error) {
	return _StoreAndGet.Contract.Store(&_StoreAndGet.TransactOpts, _key, _value, _version)
}
