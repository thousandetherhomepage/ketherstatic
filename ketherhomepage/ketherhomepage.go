// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ketherhomepage

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

// IKetherHomepageMetaData contains all meta data concerning the IKetherHomepage contract.
var IKetherHomepageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"width\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"NSFW\",\"type\":\"bool\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SetAdOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"ads\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_width\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"forceNSFW\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setAdOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"11a7a4c0": "ads(uint256)",
		"1281311d": "buy(uint256,uint256,uint256,uint256)",
		"acff2fce": "forceNSFW(uint256,bool)",
		"7a6adab6": "getAdsLength()",
		"45ebc145": "publish(uint256,string,string,string,bool)",
		"759c7a58": "setAdOwner(uint256,address)",
		"3ccfd60b": "withdraw()",
	},
}

// IKetherHomepageABI is the input ABI used to generate the binding from.
// Deprecated: Use IKetherHomepageMetaData.ABI instead.
var IKetherHomepageABI = IKetherHomepageMetaData.ABI

// Deprecated: Use IKetherHomepageMetaData.Sigs instead.
// IKetherHomepageFuncSigs maps the 4-byte function signature to its string representation.
var IKetherHomepageFuncSigs = IKetherHomepageMetaData.Sigs

// IKetherHomepage is an auto generated Go binding around an Ethereum contract.
type IKetherHomepage struct {
	IKetherHomepageCaller     // Read-only binding to the contract
	IKetherHomepageTransactor // Write-only binding to the contract
	IKetherHomepageFilterer   // Log filterer for contract events
}

// IKetherHomepageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IKetherHomepageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKetherHomepageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IKetherHomepageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKetherHomepageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IKetherHomepageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IKetherHomepageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IKetherHomepageSession struct {
	Contract     *IKetherHomepage  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IKetherHomepageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IKetherHomepageCallerSession struct {
	Contract *IKetherHomepageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IKetherHomepageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IKetherHomepageTransactorSession struct {
	Contract     *IKetherHomepageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IKetherHomepageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IKetherHomepageRaw struct {
	Contract *IKetherHomepage // Generic contract binding to access the raw methods on
}

// IKetherHomepageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IKetherHomepageCallerRaw struct {
	Contract *IKetherHomepageCaller // Generic read-only contract binding to access the raw methods on
}

// IKetherHomepageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IKetherHomepageTransactorRaw struct {
	Contract *IKetherHomepageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIKetherHomepage creates a new instance of IKetherHomepage, bound to a specific deployed contract.
func NewIKetherHomepage(address common.Address, backend bind.ContractBackend) (*IKetherHomepage, error) {
	contract, err := bindIKetherHomepage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepage{IKetherHomepageCaller: IKetherHomepageCaller{contract: contract}, IKetherHomepageTransactor: IKetherHomepageTransactor{contract: contract}, IKetherHomepageFilterer: IKetherHomepageFilterer{contract: contract}}, nil
}

// NewIKetherHomepageCaller creates a new read-only instance of IKetherHomepage, bound to a specific deployed contract.
func NewIKetherHomepageCaller(address common.Address, caller bind.ContractCaller) (*IKetherHomepageCaller, error) {
	contract, err := bindIKetherHomepage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepageCaller{contract: contract}, nil
}

// NewIKetherHomepageTransactor creates a new write-only instance of IKetherHomepage, bound to a specific deployed contract.
func NewIKetherHomepageTransactor(address common.Address, transactor bind.ContractTransactor) (*IKetherHomepageTransactor, error) {
	contract, err := bindIKetherHomepage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepageTransactor{contract: contract}, nil
}

// NewIKetherHomepageFilterer creates a new log filterer instance of IKetherHomepage, bound to a specific deployed contract.
func NewIKetherHomepageFilterer(address common.Address, filterer bind.ContractFilterer) (*IKetherHomepageFilterer, error) {
	contract, err := bindIKetherHomepage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepageFilterer{contract: contract}, nil
}

// bindIKetherHomepage binds a generic wrapper to an already deployed contract.
func bindIKetherHomepage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IKetherHomepageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKetherHomepage *IKetherHomepageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKetherHomepage.Contract.IKetherHomepageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKetherHomepage *IKetherHomepageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.IKetherHomepageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKetherHomepage *IKetherHomepageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.IKetherHomepageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IKetherHomepage *IKetherHomepageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IKetherHomepage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IKetherHomepage *IKetherHomepageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IKetherHomepage *IKetherHomepageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.contract.Transact(opts, method, params...)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 _idx) view returns(address, uint256, uint256, uint256, uint256, string, string, string, bool, bool)
func (_IKetherHomepage *IKetherHomepageCaller) Ads(opts *bind.CallOpts, _idx *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, string, string, string, bool, bool, error) {
	var out []interface{}
	err := _IKetherHomepage.contract.Call(opts, &out, "ads", _idx)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(string), *new(string), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new(string)).(*string)
	out7 := *abi.ConvertType(out[7], new(string)).(*string)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)
	out9 := *abi.ConvertType(out[9], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 _idx) view returns(address, uint256, uint256, uint256, uint256, string, string, string, bool, bool)
func (_IKetherHomepage *IKetherHomepageSession) Ads(_idx *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, string, string, string, bool, bool, error) {
	return _IKetherHomepage.Contract.Ads(&_IKetherHomepage.CallOpts, _idx)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 _idx) view returns(address, uint256, uint256, uint256, uint256, string, string, string, bool, bool)
func (_IKetherHomepage *IKetherHomepageCallerSession) Ads(_idx *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, *big.Int, string, string, string, bool, bool, error) {
	return _IKetherHomepage.Contract.Ads(&_IKetherHomepage.CallOpts, _idx)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_IKetherHomepage *IKetherHomepageCaller) GetAdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IKetherHomepage.contract.Call(opts, &out, "getAdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_IKetherHomepage *IKetherHomepageSession) GetAdsLength() (*big.Int, error) {
	return _IKetherHomepage.Contract.GetAdsLength(&_IKetherHomepage.CallOpts)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_IKetherHomepage *IKetherHomepageCallerSession) GetAdsLength() (*big.Int, error) {
	return _IKetherHomepage.Contract.GetAdsLength(&_IKetherHomepage.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_IKetherHomepage *IKetherHomepageTransactor) Buy(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _IKetherHomepage.contract.Transact(opts, "buy", _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_IKetherHomepage *IKetherHomepageSession) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Buy(&_IKetherHomepage.TransactOpts, _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_IKetherHomepage *IKetherHomepageTransactorSession) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Buy(&_IKetherHomepage.TransactOpts, _x, _y, _width, _height)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageTransactor) ForceNSFW(opts *bind.TransactOpts, _idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.contract.Transact(opts, "forceNSFW", _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageSession) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.ForceNSFW(&_IKetherHomepage.TransactOpts, _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageTransactorSession) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.ForceNSFW(&_IKetherHomepage.TransactOpts, _idx, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageTransactor) Publish(opts *bind.TransactOpts, _idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.contract.Transact(opts, "publish", _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageSession) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Publish(&_IKetherHomepage.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_IKetherHomepage *IKetherHomepageTransactorSession) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Publish(&_IKetherHomepage.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_IKetherHomepage *IKetherHomepageTransactor) SetAdOwner(opts *bind.TransactOpts, _idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _IKetherHomepage.contract.Transact(opts, "setAdOwner", _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_IKetherHomepage *IKetherHomepageSession) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.SetAdOwner(&_IKetherHomepage.TransactOpts, _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_IKetherHomepage *IKetherHomepageTransactorSession) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _IKetherHomepage.Contract.SetAdOwner(&_IKetherHomepage.TransactOpts, _idx, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IKetherHomepage *IKetherHomepageTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IKetherHomepage.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IKetherHomepage *IKetherHomepageSession) Withdraw() (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Withdraw(&_IKetherHomepage.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_IKetherHomepage *IKetherHomepageTransactorSession) Withdraw() (*types.Transaction, error) {
	return _IKetherHomepage.Contract.Withdraw(&_IKetherHomepage.TransactOpts)
}

// IKetherHomepageBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the IKetherHomepage contract.
type IKetherHomepageBuyIterator struct {
	Event *IKetherHomepageBuy // Event containing the contract specifics and raw log

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
func (it *IKetherHomepageBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKetherHomepageBuy)
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
		it.Event = new(IKetherHomepageBuy)
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
func (it *IKetherHomepageBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKetherHomepageBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKetherHomepageBuy represents a Buy event raised by the IKetherHomepage contract.
type IKetherHomepageBuy struct {
	Idx    *big.Int
	Owner  common.Address
	X      *big.Int
	Y      *big.Int
	Width  *big.Int
	Height *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_IKetherHomepage *IKetherHomepageFilterer) FilterBuy(opts *bind.FilterOpts, idx []*big.Int) (*IKetherHomepageBuyIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.FilterLogs(opts, "Buy", idxRule)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepageBuyIterator{contract: _IKetherHomepage.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_IKetherHomepage *IKetherHomepageFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *IKetherHomepageBuy, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.WatchLogs(opts, "Buy", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKetherHomepageBuy)
				if err := _IKetherHomepage.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_IKetherHomepage *IKetherHomepageFilterer) ParseBuy(log types.Log) (*IKetherHomepageBuy, error) {
	event := new(IKetherHomepageBuy)
	if err := _IKetherHomepage.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKetherHomepagePublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the IKetherHomepage contract.
type IKetherHomepagePublishIterator struct {
	Event *IKetherHomepagePublish // Event containing the contract specifics and raw log

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
func (it *IKetherHomepagePublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKetherHomepagePublish)
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
		it.Event = new(IKetherHomepagePublish)
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
func (it *IKetherHomepagePublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKetherHomepagePublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKetherHomepagePublish represents a Publish event raised by the IKetherHomepage contract.
type IKetherHomepagePublish struct {
	Idx   *big.Int
	Link  string
	Image string
	Title string
	NSFW  bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_IKetherHomepage *IKetherHomepageFilterer) FilterPublish(opts *bind.FilterOpts, idx []*big.Int) (*IKetherHomepagePublishIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.FilterLogs(opts, "Publish", idxRule)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepagePublishIterator{contract: _IKetherHomepage.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_IKetherHomepage *IKetherHomepageFilterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *IKetherHomepagePublish, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.WatchLogs(opts, "Publish", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKetherHomepagePublish)
				if err := _IKetherHomepage.contract.UnpackLog(event, "Publish", log); err != nil {
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

// ParsePublish is a log parse operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_IKetherHomepage *IKetherHomepageFilterer) ParsePublish(log types.Log) (*IKetherHomepagePublish, error) {
	event := new(IKetherHomepagePublish)
	if err := _IKetherHomepage.contract.UnpackLog(event, "Publish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IKetherHomepageSetAdOwnerIterator is returned from FilterSetAdOwner and is used to iterate over the raw logs and unpacked data for SetAdOwner events raised by the IKetherHomepage contract.
type IKetherHomepageSetAdOwnerIterator struct {
	Event *IKetherHomepageSetAdOwner // Event containing the contract specifics and raw log

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
func (it *IKetherHomepageSetAdOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IKetherHomepageSetAdOwner)
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
		it.Event = new(IKetherHomepageSetAdOwner)
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
func (it *IKetherHomepageSetAdOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IKetherHomepageSetAdOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IKetherHomepageSetAdOwner represents a SetAdOwner event raised by the IKetherHomepage contract.
type IKetherHomepageSetAdOwner struct {
	Idx  *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetAdOwner is a free log retrieval operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_IKetherHomepage *IKetherHomepageFilterer) FilterSetAdOwner(opts *bind.FilterOpts, idx []*big.Int) (*IKetherHomepageSetAdOwnerIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.FilterLogs(opts, "SetAdOwner", idxRule)
	if err != nil {
		return nil, err
	}
	return &IKetherHomepageSetAdOwnerIterator{contract: _IKetherHomepage.contract, event: "SetAdOwner", logs: logs, sub: sub}, nil
}

// WatchSetAdOwner is a free log subscription operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_IKetherHomepage *IKetherHomepageFilterer) WatchSetAdOwner(opts *bind.WatchOpts, sink chan<- *IKetherHomepageSetAdOwner, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _IKetherHomepage.contract.WatchLogs(opts, "SetAdOwner", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IKetherHomepageSetAdOwner)
				if err := _IKetherHomepage.contract.UnpackLog(event, "SetAdOwner", log); err != nil {
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

// ParseSetAdOwner is a log parse operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_IKetherHomepage *IKetherHomepageFilterer) ParseSetAdOwner(log types.Log) (*IKetherHomepageSetAdOwner, error) {
	event := new(IKetherHomepageSetAdOwner)
	if err := _IKetherHomepage.contract.UnpackLog(event, "SetAdOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KetherHomepageV2MetaData contains all meta data concerning the KetherHomepageV2 contract.
var KetherHomepageV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractOwner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_withdrawWallet\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"width\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"NSFW\",\"type\":\"bool\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SetAdOwner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ads\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"width\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"NSFW\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forceNSFW\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_width\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"forceNSFW\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"grid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pixelsPerCell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_link\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"publish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setAdOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiPixelPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"11a7a4c0": "ads(uint256)",
		"1281311d": "buy(uint256,uint256,uint256,uint256)",
		"acff2fce": "forceNSFW(uint256,bool)",
		"7a6adab6": "getAdsLength()",
		"146008e3": "grid(uint256,uint256)",
		"13f4b42c": "pixelsPerCell()",
		"45ebc145": "publish(uint256,string,string,string,bool)",
		"759c7a58": "setAdOwner(uint256,address)",
		"d5bec84e": "weiPixelPrice()",
		"3ccfd60b": "withdraw()",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516111ee3803806111ee83398101604081905261002f91610088565b6001600160a01b03821661004257600080fd5b6001600160a01b03811661005557600080fd5b61019080546001600160a01b039384166001600160a01b03199182161790915561019180549290931691161790556100da565b6000806040838503121561009b57600080fd5b82516100a6816100c2565b60208401519092506100b7816100c2565b809150509250929050565b6001600160a01b03811681146100d757600080fd5b50565b611105806100e96000396000f3fe6080604052600436106100915760003560e01c806345ebc1451161005957806345ebc14514610152578063759c7a58146101725780637a6adab614610192578063acff2fce146101a8578063d5bec84e146101c857600080fd5b806311a7a4c0146100965780631281311d146100d557806313f4b42c146100f6578063146008e31461010b5780633ccfd60b1461013b575b600080fd5b3480156100a257600080fd5b506100b66100b1366004610cb4565b6101e3565b6040516100cc9a99989796959493929190610f19565b60405180910390f35b6100e86100e3366004610dfa565b6103f0565b6040519081526020016100cc565b34801561010257600080fd5b506100e8606481565b34801561011757600080fd5b5061012b610126366004610dd8565b610843565b60405190151581526020016100cc565b34801561014757600080fd5b50610150610882565b005b34801561015e57600080fd5b5061015061016d366004610d35565b6108d7565b34801561017e57600080fd5b5061015061018d366004610ccd565b6109d5565b34801561019e57600080fd5b50610192546100e8565b3480156101b457600080fd5b506101506101c3366004610d09565b610ac0565b3480156101d457600080fd5b506100e866038d7ea4c6800081565b61019281815481106101f457600080fd5b60009182526020909120600990910201805460018201546002830154600384015460048501546005860180546001600160a01b03909616975093959294919390929161023f90611037565b80601f016020809104026020016040519081016040528092919081815260200182805461026b90611037565b80156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505050908060060180546102cd90611037565b80601f01602080910402602001604051908101604052809291908181526020018280546102f990611037565b80156103465780601f1061031b57610100808354040283529160200191610346565b820191906000526020600020905b81548152906001019060200180831161032957829003601f168201915b50505050509080600701805461035b90611037565b80601f016020809104026020016040519081016040528092919081815260200182805461038790611037565b80156103d45780601f106103a9576101008083540402835291602001916103d4565b820191906000526020600020905b8154815290600101906020018083116103b757829003601f168201915b5050506008909301549192505060ff808216916101009004168a565b60008066038d7ea4c6800060646104078587611001565b6104119190611001565b61041b9190611001565b90506000811161042a57600080fd5b8034101561048e5760405162461bcd60e51b815260206004820152602660248201527f4b6574686572486f6d65706167653a20696e73756666696369656e74206275796044820152652076616c756560d01b60648201526084015b60405180910390fd5b60005b848110156105d45760005b848110156105c15760006104b0838a610fe9565b606481106104c0576104c06110a3565b600402016104ce8289610fe9565b606481106104de576104de6110a3565b602081049091015460ff601f9092166101000a900416156105535760405162461bcd60e51b815260206004820152602960248201527f4b6574686572486f6d65706167653a2062757920616420736c6f7420616c726560448201526830b23c903a30b5b2b760b91b6064820152608401610485565b60016000610561848b610fe9565b60648110610571576105716110a3565b6004020161057f838a610fe9565b6064811061058f5761058f6110a3565b602091828204019190066101000a81548160ff02191690831515021790555080806105b990611072565b91505061049c565b50806105cc81611072565b915050610491565b5060408051610140810182523381526020808201898152828401898152606084018981526080850189815286518086018852600080825260a0880191825288518088018a5281815260c0890152885180880190995280895260e08801989098526101008701889052610120870188905261019280546001810182559852865160099098027ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67aba810180546001600160a01b039a909a166001600160a01b0319909a1699909917895594517ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67abb86015592517ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67abc85015590517ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67abd840155517ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67abe830155518051939485949093610769937ffcfced99f9d921eebdc59aa6f7a664084bd564a3d2d54ebc1a5c057c99c67abf01920190610b79565b5060c08201518051610785916006840191602090910190610b79565b5060e082015180516107a1916007840191602090910190610b79565b50610100828101516008909201805461012090940151151590910261ff00199215159290921661ffff1990931692909217179055610192546107e590600190611020565b60408051338152602081018a9052908101889052606081018790526080810186905290935083907fc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c19060a00160405180910390a25050949350505050565b6000826064811061085357600080fd5b60040201816064811061086557600080fd5b602081049091015460ff601f9092166101000a9004169150829050565b610190546001600160a01b0316331461089a57600080fd5b610191546040516001600160a01b03909116904780156108fc02916000818181858888f193505050501580156108d4573d6000803e3d6000fd5b50565b600061019286815481106108ed576108ed6110a3565b6000918252602090912060099091020180549091506001600160a01b0316331461091657600080fd5b845161092b9060058301906020880190610b79565b5083516109419060068301906020870190610b79565b5082516109579060078301906020860190610b79565b5060088101805460ff191683151590811790915586907f7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a719060058401906006850190600786019060ff16806109b557506008860154610100900460ff165b6040516109c59493929190610f9c565b60405180910390a2505050505050565b600061019283815481106109eb576109eb6110a3565b6000918252602090912060099091020180549091506001600160a01b03163314610a635760405162461bcd60e51b815260206004820152602360248201527f4b6574686572486f6d65706167653a2073656e646572206973206e6f74206f776044820152623732b960e91b6064820152608401610485565b80546001600160a01b0319166001600160a01b038316908117825560408051338152602081019290925284917f5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b91015b60405180910390a2505050565b610190546001600160a01b03163314610ad857600080fd5b60006101928381548110610aee57610aee6110a3565b60009182526020909120600860099092020190810180548415156101000261ff00198216811790925591925084917f7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a719160058501916006860191600787019160ff91821691161780610b6957506008860154610100900460ff165b604051610ab39493929190610f9c565b828054610b8590611037565b90600052602060002090601f016020900481019282610ba75760008555610bed565b82601f10610bc057805160ff1916838001178555610bed565b82800160010185558215610bed579182015b82811115610bed578251825591602001919060010190610bd2565b50610bf9929150610bfd565b5090565b5b80821115610bf95760008155600101610bfe565b80358015158114610c2257600080fd5b919050565b600082601f830112610c3857600080fd5b813567ffffffffffffffff80821115610c5357610c536110b9565b604051601f8301601f19908116603f01168101908282118183101715610c7b57610c7b6110b9565b81604052838152866020858801011115610c9457600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610cc657600080fd5b5035919050565b60008060408385031215610ce057600080fd5b8235915060208301356001600160a01b0381168114610cfe57600080fd5b809150509250929050565b60008060408385031215610d1c57600080fd5b82359150610d2c60208401610c12565b90509250929050565b600080600080600060a08688031215610d4d57600080fd5b85359450602086013567ffffffffffffffff80821115610d6c57600080fd5b610d7889838a01610c27565b95506040880135915080821115610d8e57600080fd5b610d9a89838a01610c27565b94506060880135915080821115610db057600080fd5b50610dbd88828901610c27565b925050610dcc60808701610c12565b90509295509295909350565b60008060408385031215610deb57600080fd5b50508035926020909101359150565b60008060008060808587031215610e1057600080fd5b5050823594602084013594506040840135936060013592509050565b6000815180845260005b81811015610e5257602081850181015186830182015201610e36565b81811115610e64576000602083870101525b50601f01601f19169290920160200192915050565b8054600090600181811c9080831680610e9357607f831692505b6020808410821415610eb557634e487b7160e01b600052602260045260246000fd5b83885260208801828015610ed05760018114610ee157610f0c565b60ff19871682528282019750610f0c565b60008981526020902060005b87811015610f0657815484820152908601908401610eed565b83019850505b5050505050505092915050565b600061014060018060a01b038d1683528b60208401528a60408401528960608401528860808401528060a0840152610f5381840189610e2c565b905082810360c0840152610f678188610e2c565b905082810360e0840152610f7b8187610e2c565b94151561010084015250509015156101209091015298975050505050505050565b608081526000610faf6080830187610e79565b8281036020840152610fc18187610e79565b90508281036040840152610fd58186610e79565b915050821515606083015295945050505050565b60008219821115610ffc57610ffc61108d565b500190565b600081600019048311821515161561101b5761101b61108d565b500290565b6000828210156110325761103261108d565b500390565b600181811c9082168061104b57607f821691505b6020821081141561106c57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156110865761108661108d565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212204dbefd0ee664b1d31d190f0d86715f2243a87bd4d512fe022a133f728710dcc664736f6c63430008070033",
}

// KetherHomepageV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use KetherHomepageV2MetaData.ABI instead.
var KetherHomepageV2ABI = KetherHomepageV2MetaData.ABI

// Deprecated: Use KetherHomepageV2MetaData.Sigs instead.
// KetherHomepageV2FuncSigs maps the 4-byte function signature to its string representation.
var KetherHomepageV2FuncSigs = KetherHomepageV2MetaData.Sigs

// KetherHomepageV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KetherHomepageV2MetaData.Bin instead.
var KetherHomepageV2Bin = KetherHomepageV2MetaData.Bin

// DeployKetherHomepageV2 deploys a new Ethereum contract, binding an instance of KetherHomepageV2 to it.
func DeployKetherHomepageV2(auth *bind.TransactOpts, backend bind.ContractBackend, _contractOwner common.Address, _withdrawWallet common.Address) (common.Address, *types.Transaction, *KetherHomepageV2, error) {
	parsed, err := KetherHomepageV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KetherHomepageV2Bin), backend, _contractOwner, _withdrawWallet)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KetherHomepageV2{KetherHomepageV2Caller: KetherHomepageV2Caller{contract: contract}, KetherHomepageV2Transactor: KetherHomepageV2Transactor{contract: contract}, KetherHomepageV2Filterer: KetherHomepageV2Filterer{contract: contract}}, nil
}

// KetherHomepageV2 is an auto generated Go binding around an Ethereum contract.
type KetherHomepageV2 struct {
	KetherHomepageV2Caller     // Read-only binding to the contract
	KetherHomepageV2Transactor // Write-only binding to the contract
	KetherHomepageV2Filterer   // Log filterer for contract events
}

// KetherHomepageV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type KetherHomepageV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KetherHomepageV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type KetherHomepageV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KetherHomepageV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KetherHomepageV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KetherHomepageV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KetherHomepageV2Session struct {
	Contract     *KetherHomepageV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KetherHomepageV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KetherHomepageV2CallerSession struct {
	Contract *KetherHomepageV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// KetherHomepageV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KetherHomepageV2TransactorSession struct {
	Contract     *KetherHomepageV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// KetherHomepageV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type KetherHomepageV2Raw struct {
	Contract *KetherHomepageV2 // Generic contract binding to access the raw methods on
}

// KetherHomepageV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KetherHomepageV2CallerRaw struct {
	Contract *KetherHomepageV2Caller // Generic read-only contract binding to access the raw methods on
}

// KetherHomepageV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KetherHomepageV2TransactorRaw struct {
	Contract *KetherHomepageV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewKetherHomepageV2 creates a new instance of KetherHomepageV2, bound to a specific deployed contract.
func NewKetherHomepageV2(address common.Address, backend bind.ContractBackend) (*KetherHomepageV2, error) {
	contract, err := bindKetherHomepageV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2{KetherHomepageV2Caller: KetherHomepageV2Caller{contract: contract}, KetherHomepageV2Transactor: KetherHomepageV2Transactor{contract: contract}, KetherHomepageV2Filterer: KetherHomepageV2Filterer{contract: contract}}, nil
}

// NewKetherHomepageV2Caller creates a new read-only instance of KetherHomepageV2, bound to a specific deployed contract.
func NewKetherHomepageV2Caller(address common.Address, caller bind.ContractCaller) (*KetherHomepageV2Caller, error) {
	contract, err := bindKetherHomepageV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2Caller{contract: contract}, nil
}

// NewKetherHomepageV2Transactor creates a new write-only instance of KetherHomepageV2, bound to a specific deployed contract.
func NewKetherHomepageV2Transactor(address common.Address, transactor bind.ContractTransactor) (*KetherHomepageV2Transactor, error) {
	contract, err := bindKetherHomepageV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2Transactor{contract: contract}, nil
}

// NewKetherHomepageV2Filterer creates a new log filterer instance of KetherHomepageV2, bound to a specific deployed contract.
func NewKetherHomepageV2Filterer(address common.Address, filterer bind.ContractFilterer) (*KetherHomepageV2Filterer, error) {
	contract, err := bindKetherHomepageV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2Filterer{contract: contract}, nil
}

// bindKetherHomepageV2 binds a generic wrapper to an already deployed contract.
func bindKetherHomepageV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KetherHomepageV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KetherHomepageV2 *KetherHomepageV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KetherHomepageV2.Contract.KetherHomepageV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KetherHomepageV2 *KetherHomepageV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.KetherHomepageV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KetherHomepageV2 *KetherHomepageV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.KetherHomepageV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KetherHomepageV2 *KetherHomepageV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KetherHomepageV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KetherHomepageV2 *KetherHomepageV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KetherHomepageV2 *KetherHomepageV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.contract.Transact(opts, method, params...)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 ) view returns(address owner, uint256 x, uint256 y, uint256 width, uint256 height, string link, string image, string title, bool NSFW, bool forceNSFW)
func (_KetherHomepageV2 *KetherHomepageV2Caller) Ads(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner     common.Address
	X         *big.Int
	Y         *big.Int
	Width     *big.Int
	Height    *big.Int
	Link      string
	Image     string
	Title     string
	NSFW      bool
	ForceNSFW bool
}, error) {
	var out []interface{}
	err := _KetherHomepageV2.contract.Call(opts, &out, "ads", arg0)

	outstruct := new(struct {
		Owner     common.Address
		X         *big.Int
		Y         *big.Int
		Width     *big.Int
		Height    *big.Int
		Link      string
		Image     string
		Title     string
		NSFW      bool
		ForceNSFW bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.X = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Y = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Width = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Height = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Link = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Image = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.NSFW = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.ForceNSFW = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 ) view returns(address owner, uint256 x, uint256 y, uint256 width, uint256 height, string link, string image, string title, bool NSFW, bool forceNSFW)
func (_KetherHomepageV2 *KetherHomepageV2Session) Ads(arg0 *big.Int) (struct {
	Owner     common.Address
	X         *big.Int
	Y         *big.Int
	Width     *big.Int
	Height    *big.Int
	Link      string
	Image     string
	Title     string
	NSFW      bool
	ForceNSFW bool
}, error) {
	return _KetherHomepageV2.Contract.Ads(&_KetherHomepageV2.CallOpts, arg0)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads(uint256 ) view returns(address owner, uint256 x, uint256 y, uint256 width, uint256 height, string link, string image, string title, bool NSFW, bool forceNSFW)
func (_KetherHomepageV2 *KetherHomepageV2CallerSession) Ads(arg0 *big.Int) (struct {
	Owner     common.Address
	X         *big.Int
	Y         *big.Int
	Width     *big.Int
	Height    *big.Int
	Link      string
	Image     string
	Title     string
	NSFW      bool
	ForceNSFW bool
}, error) {
	return _KetherHomepageV2.Contract.Ads(&_KetherHomepageV2.CallOpts, arg0)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Caller) GetAdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KetherHomepageV2.contract.Call(opts, &out, "getAdsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Session) GetAdsLength() (*big.Int, error) {
	return _KetherHomepageV2.Contract.GetAdsLength(&_KetherHomepageV2.CallOpts)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2CallerSession) GetAdsLength() (*big.Int, error) {
	return _KetherHomepageV2.Contract.GetAdsLength(&_KetherHomepageV2.CallOpts)
}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid(uint256 , uint256 ) view returns(bool)
func (_KetherHomepageV2 *KetherHomepageV2Caller) Grid(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _KetherHomepageV2.contract.Call(opts, &out, "grid", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid(uint256 , uint256 ) view returns(bool)
func (_KetherHomepageV2 *KetherHomepageV2Session) Grid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _KetherHomepageV2.Contract.Grid(&_KetherHomepageV2.CallOpts, arg0, arg1)
}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid(uint256 , uint256 ) view returns(bool)
func (_KetherHomepageV2 *KetherHomepageV2CallerSession) Grid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _KetherHomepageV2.Contract.Grid(&_KetherHomepageV2.CallOpts, arg0, arg1)
}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Caller) PixelsPerCell(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KetherHomepageV2.contract.Call(opts, &out, "pixelsPerCell")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Session) PixelsPerCell() (*big.Int, error) {
	return _KetherHomepageV2.Contract.PixelsPerCell(&_KetherHomepageV2.CallOpts)
}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2CallerSession) PixelsPerCell() (*big.Int, error) {
	return _KetherHomepageV2.Contract.PixelsPerCell(&_KetherHomepageV2.CallOpts)
}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Caller) WeiPixelPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KetherHomepageV2.contract.Call(opts, &out, "weiPixelPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2Session) WeiPixelPrice() (*big.Int, error) {
	return _KetherHomepageV2.Contract.WeiPixelPrice(&_KetherHomepageV2.CallOpts)
}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() view returns(uint256)
func (_KetherHomepageV2 *KetherHomepageV2CallerSession) WeiPixelPrice() (*big.Int, error) {
	return _KetherHomepageV2.Contract.WeiPixelPrice(&_KetherHomepageV2.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_KetherHomepageV2 *KetherHomepageV2Transactor) Buy(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepageV2.contract.Transact(opts, "buy", _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_KetherHomepageV2 *KetherHomepageV2Session) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Buy(&_KetherHomepageV2.TransactOpts, _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(uint256 _x, uint256 _y, uint256 _width, uint256 _height) payable returns(uint256 idx)
func (_KetherHomepageV2 *KetherHomepageV2TransactorSession) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Buy(&_KetherHomepageV2.TransactOpts, _x, _y, _width, _height)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2Transactor) ForceNSFW(opts *bind.TransactOpts, _idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.contract.Transact(opts, "forceNSFW", _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2Session) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.ForceNSFW(&_KetherHomepageV2.TransactOpts, _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(uint256 _idx, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2TransactorSession) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.ForceNSFW(&_KetherHomepageV2.TransactOpts, _idx, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2Transactor) Publish(opts *bind.TransactOpts, _idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.contract.Transact(opts, "publish", _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2Session) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Publish(&_KetherHomepageV2.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(uint256 _idx, string _link, string _image, string _title, bool _NSFW) returns()
func (_KetherHomepageV2 *KetherHomepageV2TransactorSession) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Publish(&_KetherHomepageV2.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_KetherHomepageV2 *KetherHomepageV2Transactor) SetAdOwner(opts *bind.TransactOpts, _idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepageV2.contract.Transact(opts, "setAdOwner", _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_KetherHomepageV2 *KetherHomepageV2Session) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.SetAdOwner(&_KetherHomepageV2.TransactOpts, _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(uint256 _idx, address _newOwner) returns()
func (_KetherHomepageV2 *KetherHomepageV2TransactorSession) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.SetAdOwner(&_KetherHomepageV2.TransactOpts, _idx, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepageV2 *KetherHomepageV2Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepageV2.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepageV2 *KetherHomepageV2Session) Withdraw() (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Withdraw(&_KetherHomepageV2.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepageV2 *KetherHomepageV2TransactorSession) Withdraw() (*types.Transaction, error) {
	return _KetherHomepageV2.Contract.Withdraw(&_KetherHomepageV2.TransactOpts)
}

// KetherHomepageV2BuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the KetherHomepageV2 contract.
type KetherHomepageV2BuyIterator struct {
	Event *KetherHomepageV2Buy // Event containing the contract specifics and raw log

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
func (it *KetherHomepageV2BuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KetherHomepageV2Buy)
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
		it.Event = new(KetherHomepageV2Buy)
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
func (it *KetherHomepageV2BuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KetherHomepageV2BuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KetherHomepageV2Buy represents a Buy event raised by the KetherHomepageV2 contract.
type KetherHomepageV2Buy struct {
	Idx    *big.Int
	Owner  common.Address
	X      *big.Int
	Y      *big.Int
	Width  *big.Int
	Height *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) FilterBuy(opts *bind.FilterOpts, idx []*big.Int) (*KetherHomepageV2BuyIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.FilterLogs(opts, "Buy", idxRule)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2BuyIterator{contract: _KetherHomepageV2.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *KetherHomepageV2Buy, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.WatchLogs(opts, "Buy", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KetherHomepageV2Buy)
				if err := _KetherHomepageV2.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0xc743092e1c1087d90fef606d97a56863ec6bbfd9e9cfbdddfe075a31094d70c1.
//
// Solidity: event Buy(uint256 indexed idx, address owner, uint256 x, uint256 y, uint256 width, uint256 height)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) ParseBuy(log types.Log) (*KetherHomepageV2Buy, error) {
	event := new(KetherHomepageV2Buy)
	if err := _KetherHomepageV2.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KetherHomepageV2PublishIterator is returned from FilterPublish and is used to iterate over the raw logs and unpacked data for Publish events raised by the KetherHomepageV2 contract.
type KetherHomepageV2PublishIterator struct {
	Event *KetherHomepageV2Publish // Event containing the contract specifics and raw log

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
func (it *KetherHomepageV2PublishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KetherHomepageV2Publish)
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
		it.Event = new(KetherHomepageV2Publish)
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
func (it *KetherHomepageV2PublishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KetherHomepageV2PublishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KetherHomepageV2Publish represents a Publish event raised by the KetherHomepageV2 contract.
type KetherHomepageV2Publish struct {
	Idx   *big.Int
	Link  string
	Image string
	Title string
	NSFW  bool
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPublish is a free log retrieval operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) FilterPublish(opts *bind.FilterOpts, idx []*big.Int) (*KetherHomepageV2PublishIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.FilterLogs(opts, "Publish", idxRule)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2PublishIterator{contract: _KetherHomepageV2.contract, event: "Publish", logs: logs, sub: sub}, nil
}

// WatchPublish is a free log subscription operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) WatchPublish(opts *bind.WatchOpts, sink chan<- *KetherHomepageV2Publish, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.WatchLogs(opts, "Publish", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KetherHomepageV2Publish)
				if err := _KetherHomepageV2.contract.UnpackLog(event, "Publish", log); err != nil {
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

// ParsePublish is a log parse operation binding the contract event 0x7c51272765f77e811c6ee6178cc8fa0d14abd934d7864e02d6261ffae52e2a71.
//
// Solidity: event Publish(uint256 indexed idx, string link, string image, string title, bool NSFW)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) ParsePublish(log types.Log) (*KetherHomepageV2Publish, error) {
	event := new(KetherHomepageV2Publish)
	if err := _KetherHomepageV2.contract.UnpackLog(event, "Publish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KetherHomepageV2SetAdOwnerIterator is returned from FilterSetAdOwner and is used to iterate over the raw logs and unpacked data for SetAdOwner events raised by the KetherHomepageV2 contract.
type KetherHomepageV2SetAdOwnerIterator struct {
	Event *KetherHomepageV2SetAdOwner // Event containing the contract specifics and raw log

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
func (it *KetherHomepageV2SetAdOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KetherHomepageV2SetAdOwner)
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
		it.Event = new(KetherHomepageV2SetAdOwner)
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
func (it *KetherHomepageV2SetAdOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KetherHomepageV2SetAdOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KetherHomepageV2SetAdOwner represents a SetAdOwner event raised by the KetherHomepageV2 contract.
type KetherHomepageV2SetAdOwner struct {
	Idx  *big.Int
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetAdOwner is a free log retrieval operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) FilterSetAdOwner(opts *bind.FilterOpts, idx []*big.Int) (*KetherHomepageV2SetAdOwnerIterator, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.FilterLogs(opts, "SetAdOwner", idxRule)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageV2SetAdOwnerIterator{contract: _KetherHomepageV2.contract, event: "SetAdOwner", logs: logs, sub: sub}, nil
}

// WatchSetAdOwner is a free log subscription operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) WatchSetAdOwner(opts *bind.WatchOpts, sink chan<- *KetherHomepageV2SetAdOwner, idx []*big.Int) (event.Subscription, error) {

	var idxRule []interface{}
	for _, idxItem := range idx {
		idxRule = append(idxRule, idxItem)
	}

	logs, sub, err := _KetherHomepageV2.contract.WatchLogs(opts, "SetAdOwner", idxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KetherHomepageV2SetAdOwner)
				if err := _KetherHomepageV2.contract.UnpackLog(event, "SetAdOwner", log); err != nil {
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

// ParseSetAdOwner is a log parse operation binding the contract event 0x5edf8d3cbd06178d149c2e008761e8055de4a427c9cb7a038d34617ee286d58b.
//
// Solidity: event SetAdOwner(uint256 indexed idx, address from, address to)
func (_KetherHomepageV2 *KetherHomepageV2Filterer) ParseSetAdOwner(log types.Log) (*KetherHomepageV2SetAdOwner, error) {
	event := new(KetherHomepageV2SetAdOwner)
	if err := _KetherHomepageV2.contract.UnpackLog(event, "SetAdOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
