// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollupContract

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
)

// RollupContractMetaData contains all meta data concerning the RollupContract contract.
var RollupContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rollupNumber\",\"type\":\"uint256\"}],\"name\":\"getTransactions\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollupHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parentRollupHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rollupNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes[]\",\"name\":\"transactions\",\"type\":\"bytes[]\"}],\"name\":\"postRollup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rollups\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"rollupHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"parentRollupHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RollupContractABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupContractMetaData.ABI instead.
var RollupContractABI = RollupContractMetaData.ABI

// RollupContract is an auto generated Go binding around an Ethereum contract.
type RollupContract struct {
	RollupContractCaller     // Read-only binding to the contract
	RollupContractTransactor // Write-only binding to the contract
	RollupContractFilterer   // Log filterer for contract events
}

// RollupContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupContractSession struct {
	Contract     *RollupContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupContractCallerSession struct {
	Contract *RollupContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// RollupContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupContractTransactorSession struct {
	Contract     *RollupContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RollupContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupContractRaw struct {
	Contract *RollupContract // Generic contract binding to access the raw methods on
}

// RollupContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupContractCallerRaw struct {
	Contract *RollupContractCaller // Generic read-only contract binding to access the raw methods on
}

// RollupContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupContractTransactorRaw struct {
	Contract *RollupContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupContract creates a new instance of RollupContract, bound to a specific deployed contract.
func NewRollupContract(address common.Address, backend bind.ContractBackend) (*RollupContract, error) {
	contract, err := bindRollupContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupContract{RollupContractCaller: RollupContractCaller{contract: contract}, RollupContractTransactor: RollupContractTransactor{contract: contract}, RollupContractFilterer: RollupContractFilterer{contract: contract}}, nil
}

// NewRollupContractCaller creates a new read-only instance of RollupContract, bound to a specific deployed contract.
func NewRollupContractCaller(address common.Address, caller bind.ContractCaller) (*RollupContractCaller, error) {
	contract, err := bindRollupContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupContractCaller{contract: contract}, nil
}

// NewRollupContractTransactor creates a new write-only instance of RollupContract, bound to a specific deployed contract.
func NewRollupContractTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupContractTransactor, error) {
	contract, err := bindRollupContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupContractTransactor{contract: contract}, nil
}

// NewRollupContractFilterer creates a new log filterer instance of RollupContract, bound to a specific deployed contract.
func NewRollupContractFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupContractFilterer, error) {
	contract, err := bindRollupContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupContractFilterer{contract: contract}, nil
}

// bindRollupContract binds a generic wrapper to an already deployed contract.
func bindRollupContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupContract *RollupContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupContract.Contract.RollupContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupContract *RollupContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupContract.Contract.RollupContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupContract *RollupContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupContract.Contract.RollupContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupContract *RollupContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupContract *RollupContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupContract *RollupContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupContract.Contract.contract.Transact(opts, method, params...)
}

// GetTransactions is a free data retrieval call binding the contract method 0x5742177c.
//
// Solidity: function getTransactions(uint256 rollupNumber) view returns(bytes[])
func (_RollupContract *RollupContractCaller) GetTransactions(opts *bind.CallOpts, rollupNumber *big.Int) ([][]byte, error) {
	var out []interface{}
	err := _RollupContract.contract.Call(opts, &out, "getTransactions", rollupNumber)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetTransactions is a free data retrieval call binding the contract method 0x5742177c.
//
// Solidity: function getTransactions(uint256 rollupNumber) view returns(bytes[])
func (_RollupContract *RollupContractSession) GetTransactions(rollupNumber *big.Int) ([][]byte, error) {
	return _RollupContract.Contract.GetTransactions(&_RollupContract.CallOpts, rollupNumber)
}

// GetTransactions is a free data retrieval call binding the contract method 0x5742177c.
//
// Solidity: function getTransactions(uint256 rollupNumber) view returns(bytes[])
func (_RollupContract *RollupContractCallerSession) GetTransactions(rollupNumber *big.Int) ([][]byte, error) {
	return _RollupContract.Contract.GetTransactions(&_RollupContract.CallOpts, rollupNumber)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bytes32 rollupHash, bytes32 parentRollupHash)
func (_RollupContract *RollupContractCaller) Rollups(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RollupHash       [32]byte
	ParentRollupHash [32]byte
}, error) {
	var out []interface{}
	err := _RollupContract.contract.Call(opts, &out, "rollups", arg0)

	outstruct := new(struct {
		RollupHash       [32]byte
		ParentRollupHash [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RollupHash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.ParentRollupHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bytes32 rollupHash, bytes32 parentRollupHash)
func (_RollupContract *RollupContractSession) Rollups(arg0 *big.Int) (struct {
	RollupHash       [32]byte
	ParentRollupHash [32]byte
}, error) {
	return _RollupContract.Contract.Rollups(&_RollupContract.CallOpts, arg0)
}

// Rollups is a free data retrieval call binding the contract method 0xb794e5a3.
//
// Solidity: function rollups(uint256 ) view returns(bytes32 rollupHash, bytes32 parentRollupHash)
func (_RollupContract *RollupContractCallerSession) Rollups(arg0 *big.Int) (struct {
	RollupHash       [32]byte
	ParentRollupHash [32]byte
}, error) {
	return _RollupContract.Contract.Rollups(&_RollupContract.CallOpts, arg0)
}

// PostRollup is a paid mutator transaction binding the contract method 0xdc5c5917.
//
// Solidity: function postRollup(bytes32 rollupHash, bytes32 parentRollupHash, uint256 rollupNumber, bytes[] transactions) returns()
func (_RollupContract *RollupContractTransactor) PostRollup(opts *bind.TransactOpts, rollupHash [32]byte, parentRollupHash [32]byte, rollupNumber *big.Int, transactions [][]byte) (*types.Transaction, error) {
	return _RollupContract.contract.Transact(opts, "postRollup", rollupHash, parentRollupHash, rollupNumber, transactions)
}

// PostRollup is a paid mutator transaction binding the contract method 0xdc5c5917.
//
// Solidity: function postRollup(bytes32 rollupHash, bytes32 parentRollupHash, uint256 rollupNumber, bytes[] transactions) returns()
func (_RollupContract *RollupContractSession) PostRollup(rollupHash [32]byte, parentRollupHash [32]byte, rollupNumber *big.Int, transactions [][]byte) (*types.Transaction, error) {
	return _RollupContract.Contract.PostRollup(&_RollupContract.TransactOpts, rollupHash, parentRollupHash, rollupNumber, transactions)
}

// PostRollup is a paid mutator transaction binding the contract method 0xdc5c5917.
//
// Solidity: function postRollup(bytes32 rollupHash, bytes32 parentRollupHash, uint256 rollupNumber, bytes[] transactions) returns()
func (_RollupContract *RollupContractTransactorSession) PostRollup(rollupHash [32]byte, parentRollupHash [32]byte, rollupNumber *big.Int, transactions [][]byte) (*types.Transaction, error) {
	return _RollupContract.Contract.PostRollup(&_RollupContract.TransactOpts, rollupHash, parentRollupHash, rollupNumber, transactions)
}
