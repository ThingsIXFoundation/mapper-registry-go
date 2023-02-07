// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mapper_registry

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

// IMapperRegistryMapper is an auto generated low-level Go binding around an user-defined struct.
type IMapperRegistryMapper struct {
	Id            [32]byte
	Revision      uint16
	Owner         common.Address
	FrequencyPlan uint8
	Active        bool
}

// MapperRegistryMetaData contains all meta data concerning the MapperRegistry contract.
var MapperRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidNewOwnerErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MapperAlreadyOnboardedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperAlreadyRegisteredErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperNotAddedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperNotOnboardedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MapperNotOwnedErr\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"UnknownMapperErr\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"FrequencyPlanAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"FrequencyPlanRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperInactive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MapperOnboarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"}],\"name\":\"MapperRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mapperId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"MapperTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"plan\",\"type\":\"uint8\"}],\"name\":\"frequencyPlans\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"mappers\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"revision\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structIMapperRegistry.Mapper\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mappersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mappersOnboarded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mappersOwned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"mappersOwnedPaged\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"revision\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structIMapperRegistry.Mapper[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"mappersPaged\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"revision\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"frequencyPlan\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structIMapperRegistry.Mapper[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mappersRegistered\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MapperRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MapperRegistryMetaData.ABI instead.
var MapperRegistryABI = MapperRegistryMetaData.ABI

// MapperRegistry is an auto generated Go binding around an Ethereum contract.
type MapperRegistry struct {
	MapperRegistryCaller     // Read-only binding to the contract
	MapperRegistryTransactor // Write-only binding to the contract
	MapperRegistryFilterer   // Log filterer for contract events
}

// MapperRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MapperRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MapperRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MapperRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MapperRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MapperRegistrySession struct {
	Contract     *MapperRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MapperRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MapperRegistryCallerSession struct {
	Contract *MapperRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MapperRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MapperRegistryTransactorSession struct {
	Contract     *MapperRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MapperRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MapperRegistryRaw struct {
	Contract *MapperRegistry // Generic contract binding to access the raw methods on
}

// MapperRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MapperRegistryCallerRaw struct {
	Contract *MapperRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MapperRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MapperRegistryTransactorRaw struct {
	Contract *MapperRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMapperRegistry creates a new instance of MapperRegistry, bound to a specific deployed contract.
func NewMapperRegistry(address common.Address, backend bind.ContractBackend) (*MapperRegistry, error) {
	contract, err := bindMapperRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MapperRegistry{MapperRegistryCaller: MapperRegistryCaller{contract: contract}, MapperRegistryTransactor: MapperRegistryTransactor{contract: contract}, MapperRegistryFilterer: MapperRegistryFilterer{contract: contract}}, nil
}

// NewMapperRegistryCaller creates a new read-only instance of MapperRegistry, bound to a specific deployed contract.
func NewMapperRegistryCaller(address common.Address, caller bind.ContractCaller) (*MapperRegistryCaller, error) {
	contract, err := bindMapperRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryCaller{contract: contract}, nil
}

// NewMapperRegistryTransactor creates a new write-only instance of MapperRegistry, bound to a specific deployed contract.
func NewMapperRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MapperRegistryTransactor, error) {
	contract, err := bindMapperRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryTransactor{contract: contract}, nil
}

// NewMapperRegistryFilterer creates a new log filterer instance of MapperRegistry, bound to a specific deployed contract.
func NewMapperRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MapperRegistryFilterer, error) {
	contract, err := bindMapperRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryFilterer{contract: contract}, nil
}

// bindMapperRegistry binds a generic wrapper to an already deployed contract.
func bindMapperRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MapperRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapperRegistry *MapperRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapperRegistry.Contract.MapperRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapperRegistry *MapperRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapperRegistry.Contract.MapperRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapperRegistry *MapperRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapperRegistry.Contract.MapperRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MapperRegistry *MapperRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MapperRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MapperRegistry *MapperRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MapperRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MapperRegistry *MapperRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MapperRegistry.Contract.contract.Transact(opts, method, params...)
}

// Mappers is a free data retrieval call binding the contract method 0xa2e8380b.
//
// Solidity: function mappers(bytes32 id) view returns((bytes32,uint16,address,uint8,bool))
func (_MapperRegistry *MapperRegistryCaller) Mappers(opts *bind.CallOpts, id [32]byte) (IMapperRegistryMapper, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappers", id)

	if err != nil {
		return *new(IMapperRegistryMapper), err
	}

	out0 := *abi.ConvertType(out[0], new(IMapperRegistryMapper)).(*IMapperRegistryMapper)

	return out0, err

}

// Mappers is a free data retrieval call binding the contract method 0xa2e8380b.
//
// Solidity: function mappers(bytes32 id) view returns((bytes32,uint16,address,uint8,bool))
func (_MapperRegistry *MapperRegistrySession) Mappers(id [32]byte) (IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.Mappers(&_MapperRegistry.CallOpts, id)
}

// Mappers is a free data retrieval call binding the contract method 0xa2e8380b.
//
// Solidity: function mappers(bytes32 id) view returns((bytes32,uint16,address,uint8,bool))
func (_MapperRegistry *MapperRegistryCallerSession) Mappers(id [32]byte) (IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.Mappers(&_MapperRegistry.CallOpts, id)
}

// MappersCount is a free data retrieval call binding the contract method 0x2adc0ba0.
//
// Solidity: function mappersCount() view returns(uint256)
func (_MapperRegistry *MapperRegistryCaller) MappersCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappersCount is a free data retrieval call binding the contract method 0x2adc0ba0.
//
// Solidity: function mappersCount() view returns(uint256)
func (_MapperRegistry *MapperRegistrySession) MappersCount() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersCount(&_MapperRegistry.CallOpts)
}

// MappersCount is a free data retrieval call binding the contract method 0x2adc0ba0.
//
// Solidity: function mappersCount() view returns(uint256)
func (_MapperRegistry *MapperRegistryCallerSession) MappersCount() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersCount(&_MapperRegistry.CallOpts)
}

// MappersOnboarded is a free data retrieval call binding the contract method 0xe9a446f1.
//
// Solidity: function mappersOnboarded() view returns(uint256)
func (_MapperRegistry *MapperRegistryCaller) MappersOnboarded(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersOnboarded")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappersOnboarded is a free data retrieval call binding the contract method 0xe9a446f1.
//
// Solidity: function mappersOnboarded() view returns(uint256)
func (_MapperRegistry *MapperRegistrySession) MappersOnboarded() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersOnboarded(&_MapperRegistry.CallOpts)
}

// MappersOnboarded is a free data retrieval call binding the contract method 0xe9a446f1.
//
// Solidity: function mappersOnboarded() view returns(uint256)
func (_MapperRegistry *MapperRegistryCallerSession) MappersOnboarded() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersOnboarded(&_MapperRegistry.CallOpts)
}

// MappersOwned is a free data retrieval call binding the contract method 0x7f1ff48e.
//
// Solidity: function mappersOwned(address owner) view returns(uint256)
func (_MapperRegistry *MapperRegistryCaller) MappersOwned(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersOwned", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappersOwned is a free data retrieval call binding the contract method 0x7f1ff48e.
//
// Solidity: function mappersOwned(address owner) view returns(uint256)
func (_MapperRegistry *MapperRegistrySession) MappersOwned(owner common.Address) (*big.Int, error) {
	return _MapperRegistry.Contract.MappersOwned(&_MapperRegistry.CallOpts, owner)
}

// MappersOwned is a free data retrieval call binding the contract method 0x7f1ff48e.
//
// Solidity: function mappersOwned(address owner) view returns(uint256)
func (_MapperRegistry *MapperRegistryCallerSession) MappersOwned(owner common.Address) (*big.Int, error) {
	return _MapperRegistry.Contract.MappersOwned(&_MapperRegistry.CallOpts, owner)
}

// MappersOwnedPaged is a free data retrieval call binding the contract method 0x75da3e78.
//
// Solidity: function mappersOwnedPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistryCaller) MappersOwnedPaged(opts *bind.CallOpts, owner common.Address, start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersOwnedPaged", owner, start, end)

	if err != nil {
		return *new([]IMapperRegistryMapper), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMapperRegistryMapper)).(*[]IMapperRegistryMapper)

	return out0, err

}

// MappersOwnedPaged is a free data retrieval call binding the contract method 0x75da3e78.
//
// Solidity: function mappersOwnedPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistrySession) MappersOwnedPaged(owner common.Address, start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.MappersOwnedPaged(&_MapperRegistry.CallOpts, owner, start, end)
}

// MappersOwnedPaged is a free data retrieval call binding the contract method 0x75da3e78.
//
// Solidity: function mappersOwnedPaged(address owner, uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistryCallerSession) MappersOwnedPaged(owner common.Address, start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.MappersOwnedPaged(&_MapperRegistry.CallOpts, owner, start, end)
}

// MappersPaged is a free data retrieval call binding the contract method 0x86c41533.
//
// Solidity: function mappersPaged(uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistryCaller) MappersPaged(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersPaged", start, end)

	if err != nil {
		return *new([]IMapperRegistryMapper), err
	}

	out0 := *abi.ConvertType(out[0], new([]IMapperRegistryMapper)).(*[]IMapperRegistryMapper)

	return out0, err

}

// MappersPaged is a free data retrieval call binding the contract method 0x86c41533.
//
// Solidity: function mappersPaged(uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistrySession) MappersPaged(start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.MappersPaged(&_MapperRegistry.CallOpts, start, end)
}

// MappersPaged is a free data retrieval call binding the contract method 0x86c41533.
//
// Solidity: function mappersPaged(uint256 start, uint256 end) view returns((bytes32,uint16,address,uint8,bool)[])
func (_MapperRegistry *MapperRegistryCallerSession) MappersPaged(start *big.Int, end *big.Int) ([]IMapperRegistryMapper, error) {
	return _MapperRegistry.Contract.MappersPaged(&_MapperRegistry.CallOpts, start, end)
}

// MappersRegistered is a free data retrieval call binding the contract method 0x3213e9c4.
//
// Solidity: function mappersRegistered() view returns(uint256)
func (_MapperRegistry *MapperRegistryCaller) MappersRegistered(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MapperRegistry.contract.Call(opts, &out, "mappersRegistered")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MappersRegistered is a free data retrieval call binding the contract method 0x3213e9c4.
//
// Solidity: function mappersRegistered() view returns(uint256)
func (_MapperRegistry *MapperRegistrySession) MappersRegistered() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersRegistered(&_MapperRegistry.CallOpts)
}

// MappersRegistered is a free data retrieval call binding the contract method 0x3213e9c4.
//
// Solidity: function mappersRegistered() view returns(uint256)
func (_MapperRegistry *MapperRegistryCallerSession) MappersRegistered() (*big.Int, error) {
	return _MapperRegistry.Contract.MappersRegistered(&_MapperRegistry.CallOpts)
}

// FrequencyPlans is a paid mutator transaction binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 plan) returns(string)
func (_MapperRegistry *MapperRegistryTransactor) FrequencyPlans(opts *bind.TransactOpts, plan uint8) (*types.Transaction, error) {
	return _MapperRegistry.contract.Transact(opts, "frequencyPlans", plan)
}

// FrequencyPlans is a paid mutator transaction binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 plan) returns(string)
func (_MapperRegistry *MapperRegistrySession) FrequencyPlans(plan uint8) (*types.Transaction, error) {
	return _MapperRegistry.Contract.FrequencyPlans(&_MapperRegistry.TransactOpts, plan)
}

// FrequencyPlans is a paid mutator transaction binding the contract method 0x809ae202.
//
// Solidity: function frequencyPlans(uint8 plan) returns(string)
func (_MapperRegistry *MapperRegistryTransactorSession) FrequencyPlans(plan uint8) (*types.Transaction, error) {
	return _MapperRegistry.Contract.FrequencyPlans(&_MapperRegistry.TransactOpts, plan)
}

// MapperRegistryFrequencyPlanAddedIterator is returned from FilterFrequencyPlanAdded and is used to iterate over the raw logs and unpacked data for FrequencyPlanAdded events raised by the MapperRegistry contract.
type MapperRegistryFrequencyPlanAddedIterator struct {
	Event *MapperRegistryFrequencyPlanAdded // Event containing the contract specifics and raw log

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
func (it *MapperRegistryFrequencyPlanAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryFrequencyPlanAdded)
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
		it.Event = new(MapperRegistryFrequencyPlanAdded)
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
func (it *MapperRegistryFrequencyPlanAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryFrequencyPlanAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryFrequencyPlanAdded represents a FrequencyPlanAdded event raised by the MapperRegistry contract.
type MapperRegistryFrequencyPlanAdded struct {
	Arg0 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrequencyPlanAdded is a free log retrieval operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) FilterFrequencyPlanAdded(opts *bind.FilterOpts, arg0 []uint8) (*MapperRegistryFrequencyPlanAddedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "FrequencyPlanAdded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryFrequencyPlanAddedIterator{contract: _MapperRegistry.contract, event: "FrequencyPlanAdded", logs: logs, sub: sub}, nil
}

// WatchFrequencyPlanAdded is a free log subscription operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) WatchFrequencyPlanAdded(opts *bind.WatchOpts, sink chan<- *MapperRegistryFrequencyPlanAdded, arg0 []uint8) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "FrequencyPlanAdded", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryFrequencyPlanAdded)
				if err := _MapperRegistry.contract.UnpackLog(event, "FrequencyPlanAdded", log); err != nil {
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

// ParseFrequencyPlanAdded is a log parse operation binding the contract event 0x5eb1afacf56fce1c91289feeed93783eefc45e3111990d01a42bde4078ec6198.
//
// Solidity: event FrequencyPlanAdded(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) ParseFrequencyPlanAdded(log types.Log) (*MapperRegistryFrequencyPlanAdded, error) {
	event := new(MapperRegistryFrequencyPlanAdded)
	if err := _MapperRegistry.contract.UnpackLog(event, "FrequencyPlanAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryFrequencyPlanRemovedIterator is returned from FilterFrequencyPlanRemoved and is used to iterate over the raw logs and unpacked data for FrequencyPlanRemoved events raised by the MapperRegistry contract.
type MapperRegistryFrequencyPlanRemovedIterator struct {
	Event *MapperRegistryFrequencyPlanRemoved // Event containing the contract specifics and raw log

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
func (it *MapperRegistryFrequencyPlanRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryFrequencyPlanRemoved)
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
		it.Event = new(MapperRegistryFrequencyPlanRemoved)
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
func (it *MapperRegistryFrequencyPlanRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryFrequencyPlanRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryFrequencyPlanRemoved represents a FrequencyPlanRemoved event raised by the MapperRegistry contract.
type MapperRegistryFrequencyPlanRemoved struct {
	Arg0 uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterFrequencyPlanRemoved is a free log retrieval operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) FilterFrequencyPlanRemoved(opts *bind.FilterOpts, arg0 []uint8) (*MapperRegistryFrequencyPlanRemovedIterator, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "FrequencyPlanRemoved", arg0Rule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryFrequencyPlanRemovedIterator{contract: _MapperRegistry.contract, event: "FrequencyPlanRemoved", logs: logs, sub: sub}, nil
}

// WatchFrequencyPlanRemoved is a free log subscription operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) WatchFrequencyPlanRemoved(opts *bind.WatchOpts, sink chan<- *MapperRegistryFrequencyPlanRemoved, arg0 []uint8) (event.Subscription, error) {

	var arg0Rule []interface{}
	for _, arg0Item := range arg0 {
		arg0Rule = append(arg0Rule, arg0Item)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "FrequencyPlanRemoved", arg0Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryFrequencyPlanRemoved)
				if err := _MapperRegistry.contract.UnpackLog(event, "FrequencyPlanRemoved", log); err != nil {
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

// ParseFrequencyPlanRemoved is a log parse operation binding the contract event 0x898a81f51a8979704f8c50be34010e7d8aac481090dbff9badfd49079b58f06f.
//
// Solidity: event FrequencyPlanRemoved(uint8 indexed arg0)
func (_MapperRegistry *MapperRegistryFilterer) ParseFrequencyPlanRemoved(log types.Log) (*MapperRegistryFrequencyPlanRemoved, error) {
	event := new(MapperRegistryFrequencyPlanRemoved)
	if err := _MapperRegistry.contract.UnpackLog(event, "FrequencyPlanRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperActiveIterator is returned from FilterMapperActive and is used to iterate over the raw logs and unpacked data for MapperActive events raised by the MapperRegistry contract.
type MapperRegistryMapperActiveIterator struct {
	Event *MapperRegistryMapperActive // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperActive)
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
		it.Event = new(MapperRegistryMapperActive)
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
func (it *MapperRegistryMapperActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperActive represents a MapperActive event raised by the MapperRegistry contract.
type MapperRegistryMapperActive struct {
	MapperId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperActive is a free log retrieval operation binding the contract event 0x57434896cf673678905b6f2de982d17661e7f7f971e7b4677592aaea2276b890.
//
// Solidity: event MapperActive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperActive(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperActiveIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperActive", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperActiveIterator{contract: _MapperRegistry.contract, event: "MapperActive", logs: logs, sub: sub}, nil
}

// WatchMapperActive is a free log subscription operation binding the contract event 0x57434896cf673678905b6f2de982d17661e7f7f971e7b4677592aaea2276b890.
//
// Solidity: event MapperActive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperActive(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperActive, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperActive", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperActive)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperActive", log); err != nil {
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

// ParseMapperActive is a log parse operation binding the contract event 0x57434896cf673678905b6f2de982d17661e7f7f971e7b4677592aaea2276b890.
//
// Solidity: event MapperActive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperActive(log types.Log) (*MapperRegistryMapperActive, error) {
	event := new(MapperRegistryMapperActive)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperClaimedIterator is returned from FilterMapperClaimed and is used to iterate over the raw logs and unpacked data for MapperClaimed events raised by the MapperRegistry contract.
type MapperRegistryMapperClaimedIterator struct {
	Event *MapperRegistryMapperClaimed // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperClaimed)
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
		it.Event = new(MapperRegistryMapperClaimed)
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
func (it *MapperRegistryMapperClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperClaimed represents a MapperClaimed event raised by the MapperRegistry contract.
type MapperRegistryMapperClaimed struct {
	MapperId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperClaimed is a free log retrieval operation binding the contract event 0x16a6b348b479b8e7371e9850580976f3c4cb75a00047a126ca3cd769c8c95dc0.
//
// Solidity: event MapperClaimed(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperClaimed(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperClaimedIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperClaimed", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperClaimedIterator{contract: _MapperRegistry.contract, event: "MapperClaimed", logs: logs, sub: sub}, nil
}

// WatchMapperClaimed is a free log subscription operation binding the contract event 0x16a6b348b479b8e7371e9850580976f3c4cb75a00047a126ca3cd769c8c95dc0.
//
// Solidity: event MapperClaimed(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperClaimed(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperClaimed, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperClaimed", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperClaimed)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperClaimed", log); err != nil {
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

// ParseMapperClaimed is a log parse operation binding the contract event 0x16a6b348b479b8e7371e9850580976f3c4cb75a00047a126ca3cd769c8c95dc0.
//
// Solidity: event MapperClaimed(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperClaimed(log types.Log) (*MapperRegistryMapperClaimed, error) {
	event := new(MapperRegistryMapperClaimed)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperInactiveIterator is returned from FilterMapperInactive and is used to iterate over the raw logs and unpacked data for MapperInactive events raised by the MapperRegistry contract.
type MapperRegistryMapperInactiveIterator struct {
	Event *MapperRegistryMapperInactive // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperInactiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperInactive)
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
		it.Event = new(MapperRegistryMapperInactive)
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
func (it *MapperRegistryMapperInactiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperInactiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperInactive represents a MapperInactive event raised by the MapperRegistry contract.
type MapperRegistryMapperInactive struct {
	MapperId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperInactive is a free log retrieval operation binding the contract event 0x1d984a9841b38cf9731c72a565c5a2bc16a460add72b929e8f27380c14c2650b.
//
// Solidity: event MapperInactive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperInactive(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperInactiveIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperInactive", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperInactiveIterator{contract: _MapperRegistry.contract, event: "MapperInactive", logs: logs, sub: sub}, nil
}

// WatchMapperInactive is a free log subscription operation binding the contract event 0x1d984a9841b38cf9731c72a565c5a2bc16a460add72b929e8f27380c14c2650b.
//
// Solidity: event MapperInactive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperInactive(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperInactive, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperInactive", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperInactive)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperInactive", log); err != nil {
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

// ParseMapperInactive is a log parse operation binding the contract event 0x1d984a9841b38cf9731c72a565c5a2bc16a460add72b929e8f27380c14c2650b.
//
// Solidity: event MapperInactive(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperInactive(log types.Log) (*MapperRegistryMapperInactive, error) {
	event := new(MapperRegistryMapperInactive)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperInactive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperOnboardedIterator is returned from FilterMapperOnboarded and is used to iterate over the raw logs and unpacked data for MapperOnboarded events raised by the MapperRegistry contract.
type MapperRegistryMapperOnboardedIterator struct {
	Event *MapperRegistryMapperOnboarded // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperOnboardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperOnboarded)
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
		it.Event = new(MapperRegistryMapperOnboarded)
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
func (it *MapperRegistryMapperOnboardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperOnboardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperOnboarded represents a MapperOnboarded event raised by the MapperRegistry contract.
type MapperRegistryMapperOnboarded struct {
	MapperId [32]byte
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperOnboarded is a free log retrieval operation binding the contract event 0xd8d793d85957c382d4d0863f3e1af7298792551b545c4d39b571de8e779465a8.
//
// Solidity: event MapperOnboarded(bytes32 indexed mapperId, address indexed owner)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperOnboarded(opts *bind.FilterOpts, mapperId [][32]byte, owner []common.Address) (*MapperRegistryMapperOnboardedIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperOnboarded", mapperIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperOnboardedIterator{contract: _MapperRegistry.contract, event: "MapperOnboarded", logs: logs, sub: sub}, nil
}

// WatchMapperOnboarded is a free log subscription operation binding the contract event 0xd8d793d85957c382d4d0863f3e1af7298792551b545c4d39b571de8e779465a8.
//
// Solidity: event MapperOnboarded(bytes32 indexed mapperId, address indexed owner)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperOnboarded(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperOnboarded, mapperId [][32]byte, owner []common.Address) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperOnboarded", mapperIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperOnboarded)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperOnboarded", log); err != nil {
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

// ParseMapperOnboarded is a log parse operation binding the contract event 0xd8d793d85957c382d4d0863f3e1af7298792551b545c4d39b571de8e779465a8.
//
// Solidity: event MapperOnboarded(bytes32 indexed mapperId, address indexed owner)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperOnboarded(log types.Log) (*MapperRegistryMapperOnboarded, error) {
	event := new(MapperRegistryMapperOnboarded)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperOnboarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperRegisteredIterator is returned from FilterMapperRegistered and is used to iterate over the raw logs and unpacked data for MapperRegistered events raised by the MapperRegistry contract.
type MapperRegistryMapperRegisteredIterator struct {
	Event *MapperRegistryMapperRegistered // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperRegistered)
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
		it.Event = new(MapperRegistryMapperRegistered)
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
func (it *MapperRegistryMapperRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperRegistered represents a MapperRegistered event raised by the MapperRegistry contract.
type MapperRegistryMapperRegistered struct {
	MapperId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperRegistered is a free log retrieval operation binding the contract event 0x380500bbcd3e249b3a7b3cb6159f5cde82b6eaa11770ed36aa1cab58977c01e8.
//
// Solidity: event MapperRegistered(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperRegistered(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperRegisteredIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperRegistered", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperRegisteredIterator{contract: _MapperRegistry.contract, event: "MapperRegistered", logs: logs, sub: sub}, nil
}

// WatchMapperRegistered is a free log subscription operation binding the contract event 0x380500bbcd3e249b3a7b3cb6159f5cde82b6eaa11770ed36aa1cab58977c01e8.
//
// Solidity: event MapperRegistered(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperRegistered(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperRegistered, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperRegistered", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperRegistered)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperRegistered", log); err != nil {
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

// ParseMapperRegistered is a log parse operation binding the contract event 0x380500bbcd3e249b3a7b3cb6159f5cde82b6eaa11770ed36aa1cab58977c01e8.
//
// Solidity: event MapperRegistered(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperRegistered(log types.Log) (*MapperRegistryMapperRegistered, error) {
	event := new(MapperRegistryMapperRegistered)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperRemovedIterator is returned from FilterMapperRemoved and is used to iterate over the raw logs and unpacked data for MapperRemoved events raised by the MapperRegistry contract.
type MapperRegistryMapperRemovedIterator struct {
	Event *MapperRegistryMapperRemoved // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperRemoved)
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
		it.Event = new(MapperRegistryMapperRemoved)
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
func (it *MapperRegistryMapperRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperRemoved represents a MapperRemoved event raised by the MapperRegistry contract.
type MapperRegistryMapperRemoved struct {
	MapperId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperRemoved is a free log retrieval operation binding the contract event 0xb340bd3517a963c62d2a8d793b8fd2a23616a284d368d1a6e6755ace7613e77d.
//
// Solidity: event MapperRemoved(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperRemoved(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperRemovedIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperRemoved", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperRemovedIterator{contract: _MapperRegistry.contract, event: "MapperRemoved", logs: logs, sub: sub}, nil
}

// WatchMapperRemoved is a free log subscription operation binding the contract event 0xb340bd3517a963c62d2a8d793b8fd2a23616a284d368d1a6e6755ace7613e77d.
//
// Solidity: event MapperRemoved(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperRemoved(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperRemoved, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperRemoved", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperRemoved)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
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

// ParseMapperRemoved is a log parse operation binding the contract event 0xb340bd3517a963c62d2a8d793b8fd2a23616a284d368d1a6e6755ace7613e77d.
//
// Solidity: event MapperRemoved(bytes32 indexed mapperId)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperRemoved(log types.Log) (*MapperRegistryMapperRemoved, error) {
	event := new(MapperRegistryMapperRemoved)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MapperRegistryMapperTransferredIterator is returned from FilterMapperTransferred and is used to iterate over the raw logs and unpacked data for MapperTransferred events raised by the MapperRegistry contract.
type MapperRegistryMapperTransferredIterator struct {
	Event *MapperRegistryMapperTransferred // Event containing the contract specifics and raw log

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
func (it *MapperRegistryMapperTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MapperRegistryMapperTransferred)
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
		it.Event = new(MapperRegistryMapperTransferred)
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
func (it *MapperRegistryMapperTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MapperRegistryMapperTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MapperRegistryMapperTransferred represents a MapperTransferred event raised by the MapperRegistry contract.
type MapperRegistryMapperTransferred struct {
	MapperId [32]byte
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMapperTransferred is a free log retrieval operation binding the contract event 0xf6c771f0c1bcdd23d119476f4f877c71dced0dbb6751af687eed8da523035367.
//
// Solidity: event MapperTransferred(bytes32 indexed mapperId, address oldOwner, address newOwner)
func (_MapperRegistry *MapperRegistryFilterer) FilterMapperTransferred(opts *bind.FilterOpts, mapperId [][32]byte) (*MapperRegistryMapperTransferredIterator, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.FilterLogs(opts, "MapperTransferred", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return &MapperRegistryMapperTransferredIterator{contract: _MapperRegistry.contract, event: "MapperTransferred", logs: logs, sub: sub}, nil
}

// WatchMapperTransferred is a free log subscription operation binding the contract event 0xf6c771f0c1bcdd23d119476f4f877c71dced0dbb6751af687eed8da523035367.
//
// Solidity: event MapperTransferred(bytes32 indexed mapperId, address oldOwner, address newOwner)
func (_MapperRegistry *MapperRegistryFilterer) WatchMapperTransferred(opts *bind.WatchOpts, sink chan<- *MapperRegistryMapperTransferred, mapperId [][32]byte) (event.Subscription, error) {

	var mapperIdRule []interface{}
	for _, mapperIdItem := range mapperId {
		mapperIdRule = append(mapperIdRule, mapperIdItem)
	}

	logs, sub, err := _MapperRegistry.contract.WatchLogs(opts, "MapperTransferred", mapperIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MapperRegistryMapperTransferred)
				if err := _MapperRegistry.contract.UnpackLog(event, "MapperTransferred", log); err != nil {
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

// ParseMapperTransferred is a log parse operation binding the contract event 0xf6c771f0c1bcdd23d119476f4f877c71dced0dbb6751af687eed8da523035367.
//
// Solidity: event MapperTransferred(bytes32 indexed mapperId, address oldOwner, address newOwner)
func (_MapperRegistry *MapperRegistryFilterer) ParseMapperTransferred(log types.Log) (*MapperRegistryMapperTransferred, error) {
	event := new(MapperRegistryMapperTransferred)
	if err := _MapperRegistry.contract.UnpackLog(event, "MapperTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
