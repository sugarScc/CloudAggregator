// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CloudAggregator

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

// SimplePaymentChannelJobInformation is an auto generated low-level Go binding around an user-defined struct.
type SimplePaymentChannelJobInformation struct {
	DockerImage string
	Port        string
	FlagMessage [32]byte
	Url         string
	IpAddress   string
}

// CloudAggregatorMetaData contains all meta data concerning the CloudAggregator contract.
var CloudAggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"ChainlinkRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"LogAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"LogBool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"LogBytes\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"LogBytes32\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"name\":\"LogInt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"LogUint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"commitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"consumers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"customerToTransactions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"cloudProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"enumSimplePaymentChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"dockerImage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"flagMessage\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"}],\"internalType\":\"structSimplePaymentChannel.JobInformation\",\"name\":\"jobs\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"flagMessage\",\"type\":\"bytes32\"}],\"name\":\"fulfillLivenessCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dockerImage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"flagMessage\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"creationTimestamp\",\"type\":\"uint256\"}],\"name\":\"publishTask\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requestIdToTransactionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieveAllUnfinishedTask\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"dockerImages\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"transactionIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"consumerAddress\",\"type\":\"address\"}],\"name\":\"retrieveTasksInfoFromTargetUser\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"transactionIds\",\"type\":\"uint256[]\"},{\"internalType\":\"enumSimplePaymentChannel.State[]\",\"name\":\"states\",\"type\":\"uint8[]\"},{\"internalType\":\"string[]\",\"name\":\"dockerImages\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"ports\",\"type\":\"string[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"flagMessages\",\"type\":\"bytes32[]\"},{\"internalType\":\"string[]\",\"name\":\"urls\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"creationTimeStamps\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"returnMoneyBack\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tidToTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"customer\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"cloudProvider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"money\",\"type\":\"uint256\"},{\"internalType\":\"enumSimplePaymentChannel.State\",\"name\":\"state\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"dockerImage\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"port\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"flagMessage\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"ipAddress\",\"type\":\"string\"}],\"internalType\":\"structSimplePaymentChannel.JobInformation\",\"name\":\"jobs\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"creationTimeStamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CloudAggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use CloudAggregatorMetaData.ABI instead.
var CloudAggregatorABI = CloudAggregatorMetaData.ABI

// CloudAggregator is an auto generated Go binding around an Ethereum contract.
type CloudAggregator struct {
	CloudAggregatorCaller     // Read-only binding to the contract
	CloudAggregatorTransactor // Write-only binding to the contract
	CloudAggregatorFilterer   // Log filterer for contract events
}

// CloudAggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CloudAggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudAggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CloudAggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudAggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CloudAggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CloudAggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CloudAggregatorSession struct {
	Contract     *CloudAggregator  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CloudAggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CloudAggregatorCallerSession struct {
	Contract *CloudAggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CloudAggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CloudAggregatorTransactorSession struct {
	Contract     *CloudAggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CloudAggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CloudAggregatorRaw struct {
	Contract *CloudAggregator // Generic contract binding to access the raw methods on
}

// CloudAggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CloudAggregatorCallerRaw struct {
	Contract *CloudAggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// CloudAggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CloudAggregatorTransactorRaw struct {
	Contract *CloudAggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCloudAggregator creates a new instance of CloudAggregator, bound to a specific deployed contract.
func NewCloudAggregator(address common.Address, backend bind.ContractBackend) (*CloudAggregator, error) {
	contract, err := bindCloudAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CloudAggregator{CloudAggregatorCaller: CloudAggregatorCaller{contract: contract}, CloudAggregatorTransactor: CloudAggregatorTransactor{contract: contract}, CloudAggregatorFilterer: CloudAggregatorFilterer{contract: contract}}, nil
}

// NewCloudAggregatorCaller creates a new read-only instance of CloudAggregator, bound to a specific deployed contract.
func NewCloudAggregatorCaller(address common.Address, caller bind.ContractCaller) (*CloudAggregatorCaller, error) {
	contract, err := bindCloudAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorCaller{contract: contract}, nil
}

// NewCloudAggregatorTransactor creates a new write-only instance of CloudAggregator, bound to a specific deployed contract.
func NewCloudAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*CloudAggregatorTransactor, error) {
	contract, err := bindCloudAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorTransactor{contract: contract}, nil
}

// NewCloudAggregatorFilterer creates a new log filterer instance of CloudAggregator, bound to a specific deployed contract.
func NewCloudAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*CloudAggregatorFilterer, error) {
	contract, err := bindCloudAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorFilterer{contract: contract}, nil
}

// bindCloudAggregator binds a generic wrapper to an already deployed contract.
func bindCloudAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CloudAggregatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloudAggregator *CloudAggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloudAggregator.Contract.CloudAggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloudAggregator *CloudAggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloudAggregator.Contract.CloudAggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloudAggregator *CloudAggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloudAggregator.Contract.CloudAggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CloudAggregator *CloudAggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CloudAggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CloudAggregator *CloudAggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CloudAggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CloudAggregator *CloudAggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CloudAggregator.Contract.contract.Transact(opts, method, params...)
}

// Consumers is a free data retrieval call binding the contract method 0x4651ed3d.
//
// Solidity: function consumers(uint256 ) view returns(address)
func (_CloudAggregator *CloudAggregatorCaller) Consumers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "consumers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Consumers is a free data retrieval call binding the contract method 0x4651ed3d.
//
// Solidity: function consumers(uint256 ) view returns(address)
func (_CloudAggregator *CloudAggregatorSession) Consumers(arg0 *big.Int) (common.Address, error) {
	return _CloudAggregator.Contract.Consumers(&_CloudAggregator.CallOpts, arg0)
}

// Consumers is a free data retrieval call binding the contract method 0x4651ed3d.
//
// Solidity: function consumers(uint256 ) view returns(address)
func (_CloudAggregator *CloudAggregatorCallerSession) Consumers(arg0 *big.Int) (common.Address, error) {
	return _CloudAggregator.Contract.Consumers(&_CloudAggregator.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CloudAggregator *CloudAggregatorCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CloudAggregator *CloudAggregatorSession) Count() (*big.Int, error) {
	return _CloudAggregator.Contract.Count(&_CloudAggregator.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_CloudAggregator *CloudAggregatorCallerSession) Count() (*big.Int, error) {
	return _CloudAggregator.Contract.Count(&_CloudAggregator.CallOpts)
}

// CustomerToTransactions is a free data retrieval call binding the contract method 0x886b7239.
//
// Solidity: function customerToTransactions(address , uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorCaller) CustomerToTransactions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "customerToTransactions", arg0, arg1)

	outstruct := new(struct {
		TransactionId     *big.Int
		Customer          common.Address
		CloudProvider     common.Address
		Money             *big.Int
		State             uint8
		Jobs              SimplePaymentChannelJobInformation
		CreationTimeStamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransactionId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Customer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CloudProvider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Money = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Jobs = *abi.ConvertType(out[5], new(SimplePaymentChannelJobInformation)).(*SimplePaymentChannelJobInformation)
	outstruct.CreationTimeStamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CustomerToTransactions is a free data retrieval call binding the contract method 0x886b7239.
//
// Solidity: function customerToTransactions(address , uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorSession) CustomerToTransactions(arg0 common.Address, arg1 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	return _CloudAggregator.Contract.CustomerToTransactions(&_CloudAggregator.CallOpts, arg0, arg1)
}

// CustomerToTransactions is a free data retrieval call binding the contract method 0x886b7239.
//
// Solidity: function customerToTransactions(address , uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorCallerSession) CustomerToTransactions(arg0 common.Address, arg1 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	return _CloudAggregator.Contract.CustomerToTransactions(&_CloudAggregator.CallOpts, arg0, arg1)
}

// RequestIdToTransactionId is a free data retrieval call binding the contract method 0x76992720.
//
// Solidity: function requestIdToTransactionId(bytes32 ) view returns(uint256)
func (_CloudAggregator *CloudAggregatorCaller) RequestIdToTransactionId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "requestIdToTransactionId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIdToTransactionId is a free data retrieval call binding the contract method 0x76992720.
//
// Solidity: function requestIdToTransactionId(bytes32 ) view returns(uint256)
func (_CloudAggregator *CloudAggregatorSession) RequestIdToTransactionId(arg0 [32]byte) (*big.Int, error) {
	return _CloudAggregator.Contract.RequestIdToTransactionId(&_CloudAggregator.CallOpts, arg0)
}

// RequestIdToTransactionId is a free data retrieval call binding the contract method 0x76992720.
//
// Solidity: function requestIdToTransactionId(bytes32 ) view returns(uint256)
func (_CloudAggregator *CloudAggregatorCallerSession) RequestIdToTransactionId(arg0 [32]byte) (*big.Int, error) {
	return _CloudAggregator.Contract.RequestIdToTransactionId(&_CloudAggregator.CallOpts, arg0)
}

// RetrieveAllUnfinishedTask is a free data retrieval call binding the contract method 0x9993179a.
//
// Solidity: function retrieveAllUnfinishedTask() view returns(string[] dockerImages, string[] ports, uint256[] transactionIds)
func (_CloudAggregator *CloudAggregatorCaller) RetrieveAllUnfinishedTask(opts *bind.CallOpts) (struct {
	DockerImages   []string
	Ports          []string
	TransactionIds []*big.Int
}, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "retrieveAllUnfinishedTask")

	outstruct := new(struct {
		DockerImages   []string
		Ports          []string
		TransactionIds []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DockerImages = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Ports = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.TransactionIds = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// RetrieveAllUnfinishedTask is a free data retrieval call binding the contract method 0x9993179a.
//
// Solidity: function retrieveAllUnfinishedTask() view returns(string[] dockerImages, string[] ports, uint256[] transactionIds)
func (_CloudAggregator *CloudAggregatorSession) RetrieveAllUnfinishedTask() (struct {
	DockerImages   []string
	Ports          []string
	TransactionIds []*big.Int
}, error) {
	return _CloudAggregator.Contract.RetrieveAllUnfinishedTask(&_CloudAggregator.CallOpts)
}

// RetrieveAllUnfinishedTask is a free data retrieval call binding the contract method 0x9993179a.
//
// Solidity: function retrieveAllUnfinishedTask() view returns(string[] dockerImages, string[] ports, uint256[] transactionIds)
func (_CloudAggregator *CloudAggregatorCallerSession) RetrieveAllUnfinishedTask() (struct {
	DockerImages   []string
	Ports          []string
	TransactionIds []*big.Int
}, error) {
	return _CloudAggregator.Contract.RetrieveAllUnfinishedTask(&_CloudAggregator.CallOpts)
}

// RetrieveTasksInfoFromTargetUser is a free data retrieval call binding the contract method 0xb9633b35.
//
// Solidity: function retrieveTasksInfoFromTargetUser(address consumerAddress) view returns(uint256[] transactionIds, uint8[] states, string[] dockerImages, string[] ports, bytes32[] flagMessages, string[] urls, uint256[] creationTimeStamps)
func (_CloudAggregator *CloudAggregatorCaller) RetrieveTasksInfoFromTargetUser(opts *bind.CallOpts, consumerAddress common.Address) (struct {
	TransactionIds     []*big.Int
	States             []uint8
	DockerImages       []string
	Ports              []string
	FlagMessages       [][32]byte
	Urls               []string
	CreationTimeStamps []*big.Int
}, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "retrieveTasksInfoFromTargetUser", consumerAddress)

	outstruct := new(struct {
		TransactionIds     []*big.Int
		States             []uint8
		DockerImages       []string
		Ports              []string
		FlagMessages       [][32]byte
		Urls               []string
		CreationTimeStamps []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransactionIds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.States = *abi.ConvertType(out[1], new([]uint8)).(*[]uint8)
	outstruct.DockerImages = *abi.ConvertType(out[2], new([]string)).(*[]string)
	outstruct.Ports = *abi.ConvertType(out[3], new([]string)).(*[]string)
	outstruct.FlagMessages = *abi.ConvertType(out[4], new([][32]byte)).(*[][32]byte)
	outstruct.Urls = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.CreationTimeStamps = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// RetrieveTasksInfoFromTargetUser is a free data retrieval call binding the contract method 0xb9633b35.
//
// Solidity: function retrieveTasksInfoFromTargetUser(address consumerAddress) view returns(uint256[] transactionIds, uint8[] states, string[] dockerImages, string[] ports, bytes32[] flagMessages, string[] urls, uint256[] creationTimeStamps)
func (_CloudAggregator *CloudAggregatorSession) RetrieveTasksInfoFromTargetUser(consumerAddress common.Address) (struct {
	TransactionIds     []*big.Int
	States             []uint8
	DockerImages       []string
	Ports              []string
	FlagMessages       [][32]byte
	Urls               []string
	CreationTimeStamps []*big.Int
}, error) {
	return _CloudAggregator.Contract.RetrieveTasksInfoFromTargetUser(&_CloudAggregator.CallOpts, consumerAddress)
}

// RetrieveTasksInfoFromTargetUser is a free data retrieval call binding the contract method 0xb9633b35.
//
// Solidity: function retrieveTasksInfoFromTargetUser(address consumerAddress) view returns(uint256[] transactionIds, uint8[] states, string[] dockerImages, string[] ports, bytes32[] flagMessages, string[] urls, uint256[] creationTimeStamps)
func (_CloudAggregator *CloudAggregatorCallerSession) RetrieveTasksInfoFromTargetUser(consumerAddress common.Address) (struct {
	TransactionIds     []*big.Int
	States             []uint8
	DockerImages       []string
	Ports              []string
	FlagMessages       [][32]byte
	Urls               []string
	CreationTimeStamps []*big.Int
}, error) {
	return _CloudAggregator.Contract.RetrieveTasksInfoFromTargetUser(&_CloudAggregator.CallOpts, consumerAddress)
}

// TidToTransaction is a free data retrieval call binding the contract method 0x743bb72d.
//
// Solidity: function tidToTransaction(uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorCaller) TidToTransaction(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	var out []interface{}
	err := _CloudAggregator.contract.Call(opts, &out, "tidToTransaction", arg0)

	outstruct := new(struct {
		TransactionId     *big.Int
		Customer          common.Address
		CloudProvider     common.Address
		Money             *big.Int
		State             uint8
		Jobs              SimplePaymentChannelJobInformation
		CreationTimeStamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransactionId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Customer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CloudProvider = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Money = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.State = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Jobs = *abi.ConvertType(out[5], new(SimplePaymentChannelJobInformation)).(*SimplePaymentChannelJobInformation)
	outstruct.CreationTimeStamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TidToTransaction is a free data retrieval call binding the contract method 0x743bb72d.
//
// Solidity: function tidToTransaction(uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorSession) TidToTransaction(arg0 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	return _CloudAggregator.Contract.TidToTransaction(&_CloudAggregator.CallOpts, arg0)
}

// TidToTransaction is a free data retrieval call binding the contract method 0x743bb72d.
//
// Solidity: function tidToTransaction(uint256 ) view returns(uint256 transactionId, address customer, address cloudProvider, uint256 money, uint8 state, (string,string,bytes32,string,string) jobs, uint256 creationTimeStamp)
func (_CloudAggregator *CloudAggregatorCallerSession) TidToTransaction(arg0 *big.Int) (struct {
	TransactionId     *big.Int
	Customer          common.Address
	CloudProvider     common.Address
	Money             *big.Int
	State             uint8
	Jobs              SimplePaymentChannelJobInformation
	CreationTimeStamp *big.Int
}, error) {
	return _CloudAggregator.Contract.TidToTransaction(&_CloudAggregator.CallOpts, arg0)
}

// CommitTask is a paid mutator transaction binding the contract method 0x8890f9ee.
//
// Solidity: function commitTask(string ipAddress, uint256 transactionId) returns()
func (_CloudAggregator *CloudAggregatorTransactor) CommitTask(opts *bind.TransactOpts, ipAddress string, transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.contract.Transact(opts, "commitTask", ipAddress, transactionId)
}

// CommitTask is a paid mutator transaction binding the contract method 0x8890f9ee.
//
// Solidity: function commitTask(string ipAddress, uint256 transactionId) returns()
func (_CloudAggregator *CloudAggregatorSession) CommitTask(ipAddress string, transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.CommitTask(&_CloudAggregator.TransactOpts, ipAddress, transactionId)
}

// CommitTask is a paid mutator transaction binding the contract method 0x8890f9ee.
//
// Solidity: function commitTask(string ipAddress, uint256 transactionId) returns()
func (_CloudAggregator *CloudAggregatorTransactorSession) CommitTask(ipAddress string, transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.CommitTask(&_CloudAggregator.TransactOpts, ipAddress, transactionId)
}

// FulfillLivenessCheck is a paid mutator transaction binding the contract method 0xeb932a89.
//
// Solidity: function fulfillLivenessCheck(bytes32 _requestId, bytes32 flagMessage) returns()
func (_CloudAggregator *CloudAggregatorTransactor) FulfillLivenessCheck(opts *bind.TransactOpts, _requestId [32]byte, flagMessage [32]byte) (*types.Transaction, error) {
	return _CloudAggregator.contract.Transact(opts, "fulfillLivenessCheck", _requestId, flagMessage)
}

// FulfillLivenessCheck is a paid mutator transaction binding the contract method 0xeb932a89.
//
// Solidity: function fulfillLivenessCheck(bytes32 _requestId, bytes32 flagMessage) returns()
func (_CloudAggregator *CloudAggregatorSession) FulfillLivenessCheck(_requestId [32]byte, flagMessage [32]byte) (*types.Transaction, error) {
	return _CloudAggregator.Contract.FulfillLivenessCheck(&_CloudAggregator.TransactOpts, _requestId, flagMessage)
}

// FulfillLivenessCheck is a paid mutator transaction binding the contract method 0xeb932a89.
//
// Solidity: function fulfillLivenessCheck(bytes32 _requestId, bytes32 flagMessage) returns()
func (_CloudAggregator *CloudAggregatorTransactorSession) FulfillLivenessCheck(_requestId [32]byte, flagMessage [32]byte) (*types.Transaction, error) {
	return _CloudAggregator.Contract.FulfillLivenessCheck(&_CloudAggregator.TransactOpts, _requestId, flagMessage)
}

// PublishTask is a paid mutator transaction binding the contract method 0x3d963065.
//
// Solidity: function publishTask(string dockerImage, string port, bytes32 flagMessage, string url, uint256 creationTimestamp) payable returns()
func (_CloudAggregator *CloudAggregatorTransactor) PublishTask(opts *bind.TransactOpts, dockerImage string, port string, flagMessage [32]byte, url string, creationTimestamp *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.contract.Transact(opts, "publishTask", dockerImage, port, flagMessage, url, creationTimestamp)
}

// PublishTask is a paid mutator transaction binding the contract method 0x3d963065.
//
// Solidity: function publishTask(string dockerImage, string port, bytes32 flagMessage, string url, uint256 creationTimestamp) payable returns()
func (_CloudAggregator *CloudAggregatorSession) PublishTask(dockerImage string, port string, flagMessage [32]byte, url string, creationTimestamp *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.PublishTask(&_CloudAggregator.TransactOpts, dockerImage, port, flagMessage, url, creationTimestamp)
}

// PublishTask is a paid mutator transaction binding the contract method 0x3d963065.
//
// Solidity: function publishTask(string dockerImage, string port, bytes32 flagMessage, string url, uint256 creationTimestamp) payable returns()
func (_CloudAggregator *CloudAggregatorTransactorSession) PublishTask(dockerImage string, port string, flagMessage [32]byte, url string, creationTimestamp *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.PublishTask(&_CloudAggregator.TransactOpts, dockerImage, port, flagMessage, url, creationTimestamp)
}

// ReturnMoneyBack is a paid mutator transaction binding the contract method 0xf18d0d0a.
//
// Solidity: function returnMoneyBack(uint256 transactionId) payable returns()
func (_CloudAggregator *CloudAggregatorTransactor) ReturnMoneyBack(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.contract.Transact(opts, "returnMoneyBack", transactionId)
}

// ReturnMoneyBack is a paid mutator transaction binding the contract method 0xf18d0d0a.
//
// Solidity: function returnMoneyBack(uint256 transactionId) payable returns()
func (_CloudAggregator *CloudAggregatorSession) ReturnMoneyBack(transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.ReturnMoneyBack(&_CloudAggregator.TransactOpts, transactionId)
}

// ReturnMoneyBack is a paid mutator transaction binding the contract method 0xf18d0d0a.
//
// Solidity: function returnMoneyBack(uint256 transactionId) payable returns()
func (_CloudAggregator *CloudAggregatorTransactorSession) ReturnMoneyBack(transactionId *big.Int) (*types.Transaction, error) {
	return _CloudAggregator.Contract.ReturnMoneyBack(&_CloudAggregator.TransactOpts, transactionId)
}

// CloudAggregatorChainlinkCancelledIterator is returned from FilterChainlinkCancelled and is used to iterate over the raw logs and unpacked data for ChainlinkCancelled events raised by the CloudAggregator contract.
type CloudAggregatorChainlinkCancelledIterator struct {
	Event *CloudAggregatorChainlinkCancelled // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorChainlinkCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorChainlinkCancelled)
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
		it.Event = new(CloudAggregatorChainlinkCancelled)
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
func (it *CloudAggregatorChainlinkCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorChainlinkCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorChainlinkCancelled represents a ChainlinkCancelled event raised by the CloudAggregator contract.
type CloudAggregatorChainlinkCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkCancelled is a free log retrieval operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) FilterChainlinkCancelled(opts *bind.FilterOpts, id [][32]byte) (*CloudAggregatorChainlinkCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorChainlinkCancelledIterator{contract: _CloudAggregator.contract, event: "ChainlinkCancelled", logs: logs, sub: sub}, nil
}

// WatchChainlinkCancelled is a free log subscription operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) WatchChainlinkCancelled(opts *bind.WatchOpts, sink chan<- *CloudAggregatorChainlinkCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "ChainlinkCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorChainlinkCancelled)
				if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
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

// ParseChainlinkCancelled is a log parse operation binding the contract event 0xe1fe3afa0f7f761ff0a8b89086790efd5140d2907ebd5b7ff6bfcb5e075fd4c5.
//
// Solidity: event ChainlinkCancelled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) ParseChainlinkCancelled(log types.Log) (*CloudAggregatorChainlinkCancelled, error) {
	event := new(CloudAggregatorChainlinkCancelled)
	if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorChainlinkFulfilledIterator is returned from FilterChainlinkFulfilled and is used to iterate over the raw logs and unpacked data for ChainlinkFulfilled events raised by the CloudAggregator contract.
type CloudAggregatorChainlinkFulfilledIterator struct {
	Event *CloudAggregatorChainlinkFulfilled // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorChainlinkFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorChainlinkFulfilled)
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
		it.Event = new(CloudAggregatorChainlinkFulfilled)
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
func (it *CloudAggregatorChainlinkFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorChainlinkFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorChainlinkFulfilled represents a ChainlinkFulfilled event raised by the CloudAggregator contract.
type CloudAggregatorChainlinkFulfilled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkFulfilled is a free log retrieval operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) FilterChainlinkFulfilled(opts *bind.FilterOpts, id [][32]byte) (*CloudAggregatorChainlinkFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorChainlinkFulfilledIterator{contract: _CloudAggregator.contract, event: "ChainlinkFulfilled", logs: logs, sub: sub}, nil
}

// WatchChainlinkFulfilled is a free log subscription operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) WatchChainlinkFulfilled(opts *bind.WatchOpts, sink chan<- *CloudAggregatorChainlinkFulfilled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "ChainlinkFulfilled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorChainlinkFulfilled)
				if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
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

// ParseChainlinkFulfilled is a log parse operation binding the contract event 0x7cc135e0cebb02c3480ae5d74d377283180a2601f8f644edf7987b009316c63a.
//
// Solidity: event ChainlinkFulfilled(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) ParseChainlinkFulfilled(log types.Log) (*CloudAggregatorChainlinkFulfilled, error) {
	event := new(CloudAggregatorChainlinkFulfilled)
	if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorChainlinkRequestedIterator is returned from FilterChainlinkRequested and is used to iterate over the raw logs and unpacked data for ChainlinkRequested events raised by the CloudAggregator contract.
type CloudAggregatorChainlinkRequestedIterator struct {
	Event *CloudAggregatorChainlinkRequested // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorChainlinkRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorChainlinkRequested)
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
		it.Event = new(CloudAggregatorChainlinkRequested)
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
func (it *CloudAggregatorChainlinkRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorChainlinkRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorChainlinkRequested represents a ChainlinkRequested event raised by the CloudAggregator contract.
type CloudAggregatorChainlinkRequested struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChainlinkRequested is a free log retrieval operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) FilterChainlinkRequested(opts *bind.FilterOpts, id [][32]byte) (*CloudAggregatorChainlinkRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorChainlinkRequestedIterator{contract: _CloudAggregator.contract, event: "ChainlinkRequested", logs: logs, sub: sub}, nil
}

// WatchChainlinkRequested is a free log subscription operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) WatchChainlinkRequested(opts *bind.WatchOpts, sink chan<- *CloudAggregatorChainlinkRequested, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "ChainlinkRequested", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorChainlinkRequested)
				if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
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

// ParseChainlinkRequested is a log parse operation binding the contract event 0xb5e6e01e79f91267dc17b4e6314d5d4d03593d2ceee0fbb452b750bd70ea5af9.
//
// Solidity: event ChainlinkRequested(bytes32 indexed id)
func (_CloudAggregator *CloudAggregatorFilterer) ParseChainlinkRequested(log types.Log) (*CloudAggregatorChainlinkRequested, error) {
	event := new(CloudAggregatorChainlinkRequested)
	if err := _CloudAggregator.contract.UnpackLog(event, "ChainlinkRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the CloudAggregator contract.
type CloudAggregatorLogAddressIterator struct {
	Event *CloudAggregatorLogAddress // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogAddress)
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
		it.Event = new(CloudAggregatorLogAddress)
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
func (it *CloudAggregatorLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogAddress represents a LogAddress event raised by the CloudAggregator contract.
type CloudAggregatorLogAddress struct {
	Arg0 string
	Arg1 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string arg0, address arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogAddress(opts *bind.FilterOpts) (*CloudAggregatorLogAddressIterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogAddressIterator{contract: _CloudAggregator.contract, event: "LogAddress", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string arg0, address arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogAddress) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogAddress)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogAddress", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x62ddffe5b5108385f7a590f100e1ee414ad9551a31f089e64e82998440785e1e.
//
// Solidity: event LogAddress(string arg0, address arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogAddress(log types.Log) (*CloudAggregatorLogAddress, error) {
	event := new(CloudAggregatorLogAddress)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogBoolIterator is returned from FilterLogBool and is used to iterate over the raw logs and unpacked data for LogBool events raised by the CloudAggregator contract.
type CloudAggregatorLogBoolIterator struct {
	Event *CloudAggregatorLogBool // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogBoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogBool)
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
		it.Event = new(CloudAggregatorLogBool)
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
func (it *CloudAggregatorLogBoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogBoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogBool represents a LogBool event raised by the CloudAggregator contract.
type CloudAggregatorLogBool struct {
	Arg0 string
	Arg1 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBool is a free log retrieval operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string arg0, bool arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogBool(opts *bind.FilterOpts) (*CloudAggregatorLogBoolIterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogBoolIterator{contract: _CloudAggregator.contract, event: "LogBool", logs: logs, sub: sub}, nil
}

// WatchLogBool is a free log subscription operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string arg0, bool arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogBool(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogBool) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogBool")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogBool)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogBool", log); err != nil {
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

// ParseLogBool is a log parse operation binding the contract event 0x4c34c2f9a78632f29fa59aaed5514cb742fd9fbcfd7ccc2c03c85f2bbc621c47.
//
// Solidity: event LogBool(string arg0, bool arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogBool(log types.Log) (*CloudAggregatorLogBool, error) {
	event := new(CloudAggregatorLogBool)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogBool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the CloudAggregator contract.
type CloudAggregatorLogBytesIterator struct {
	Event *CloudAggregatorLogBytes // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogBytes)
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
		it.Event = new(CloudAggregatorLogBytes)
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
func (it *CloudAggregatorLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogBytes represents a LogBytes event raised by the CloudAggregator contract.
type CloudAggregatorLogBytes struct {
	Arg0 string
	Arg1 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string arg0, bytes arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogBytes(opts *bind.FilterOpts) (*CloudAggregatorLogBytesIterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogBytesIterator{contract: _CloudAggregator.contract, event: "LogBytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string arg0, bytes arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogBytes) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogBytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogBytes)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogBytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0xe8407a0209fa99ec3a7228aff140c3d3e68bd279391739c7e0b65cd406cc93b5.
//
// Solidity: event LogBytes(string arg0, bytes arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogBytes(log types.Log) (*CloudAggregatorLogBytes, error) {
	event := new(CloudAggregatorLogBytes)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogBytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the CloudAggregator contract.
type CloudAggregatorLogBytes32Iterator struct {
	Event *CloudAggregatorLogBytes32 // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogBytes32)
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
		it.Event = new(CloudAggregatorLogBytes32)
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
func (it *CloudAggregatorLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogBytes32 represents a LogBytes32 event raised by the CloudAggregator contract.
type CloudAggregatorLogBytes32 struct {
	Arg0 string
	Arg1 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string arg0, bytes32 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*CloudAggregatorLogBytes32Iterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogBytes32Iterator{contract: _CloudAggregator.contract, event: "LogBytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string arg0, bytes32 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogBytes32) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogBytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogBytes32)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogBytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0x02d93529bba9d141e5e06733c52c7e6fbcb1149586adb5c24064b522ab26f1d7.
//
// Solidity: event LogBytes32(string arg0, bytes32 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogBytes32(log types.Log) (*CloudAggregatorLogBytes32, error) {
	event := new(CloudAggregatorLogBytes32)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogBytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the CloudAggregator contract.
type CloudAggregatorLogIntIterator struct {
	Event *CloudAggregatorLogInt // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogInt)
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
		it.Event = new(CloudAggregatorLogInt)
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
func (it *CloudAggregatorLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogInt represents a LogInt event raised by the CloudAggregator contract.
type CloudAggregatorLogInt struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string arg0, int256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogInt(opts *bind.FilterOpts) (*CloudAggregatorLogIntIterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogIntIterator{contract: _CloudAggregator.contract, event: "LogInt", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string arg0, int256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogInt) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogInt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogInt)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogInt", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x6a837ff3973aa4181e7b9f07756f8b7ee366dd85a36655d2cb42cd47f10b9638.
//
// Solidity: event LogInt(string arg0, int256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogInt(log types.Log) (*CloudAggregatorLogInt, error) {
	event := new(CloudAggregatorLogInt)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogInt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CloudAggregatorLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the CloudAggregator contract.
type CloudAggregatorLogUintIterator struct {
	Event *CloudAggregatorLogUint // Event containing the contract specifics and raw log

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
func (it *CloudAggregatorLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CloudAggregatorLogUint)
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
		it.Event = new(CloudAggregatorLogUint)
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
func (it *CloudAggregatorLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CloudAggregatorLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CloudAggregatorLogUint represents a LogUint event raised by the CloudAggregator contract.
type CloudAggregatorLogUint struct {
	Arg0 string
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string arg0, uint256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) FilterLogUint(opts *bind.FilterOpts) (*CloudAggregatorLogUintIterator, error) {

	logs, sub, err := _CloudAggregator.contract.FilterLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return &CloudAggregatorLogUintIterator{contract: _CloudAggregator.contract, event: "LogUint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string arg0, uint256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *CloudAggregatorLogUint) (event.Subscription, error) {

	logs, sub, err := _CloudAggregator.contract.WatchLogs(opts, "LogUint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CloudAggregatorLogUint)
				if err := _CloudAggregator.contract.UnpackLog(event, "LogUint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x941296a39ea107bde685522318a4b6c2b544904a5dd82a512748ca2cf839bef7.
//
// Solidity: event LogUint(string arg0, uint256 arg1)
func (_CloudAggregator *CloudAggregatorFilterer) ParseLogUint(log types.Log) (*CloudAggregatorLogUint, error) {
	event := new(CloudAggregatorLogUint)
	if err := _CloudAggregator.contract.UnpackLog(event, "LogUint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
