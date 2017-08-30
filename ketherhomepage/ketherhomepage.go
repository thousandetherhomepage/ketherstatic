// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package ketherhomepage

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// KetherHomepageABI is the input ABI used to generate the binding from.
const KetherHomepageABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ads\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"width\",\"type\":\"uint256\"},{\"name\":\"height\",\"type\":\"uint256\"},{\"name\":\"link\",\"type\":\"string\"},{\"name\":\"image\",\"type\":\"string\"},{\"name\":\"title\",\"type\":\"string\"},{\"name\":\"NSFW\",\"type\":\"bool\"},{\"name\":\"forceNSFW\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_width\",\"type\":\"uint256\"},{\"name\":\"_height\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[{\"name\":\"idx\",\"type\":\"uint256\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pixelsPerCell\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"grid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_idx\",\"type\":\"uint256\"},{\"name\":\"_link\",\"type\":\"string\"},{\"name\":\"_image\",\"type\":\"string\"},{\"name\":\"_title\",\"type\":\"string\"},{\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"publish\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_idx\",\"type\":\"uint256\"},{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setAdOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdsLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_idx\",\"type\":\"uint256\"},{\"name\":\"_NSFW\",\"type\":\"bool\"}],\"name\":\"forceNSFW\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiPixelPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_contractOwner\",\"type\":\"address\"},{\"name\":\"_withdrawWallet\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"width\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"link\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"image\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"NSFW\",\"type\":\"bool\"}],\"name\":\"Publish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"idx\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SetAdOwner\",\"type\":\"event\"}]"

// KetherHomepage is an auto generated Go binding around an Ethereum contract.
type KetherHomepage struct {
	KetherHomepageCaller     // Read-only binding to the contract
	KetherHomepageTransactor // Write-only binding to the contract
}

// KetherHomepageCaller is an auto generated read-only Go binding around an Ethereum contract.
type KetherHomepageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KetherHomepageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KetherHomepageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KetherHomepageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KetherHomepageSession struct {
	Contract     *KetherHomepage   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KetherHomepageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KetherHomepageCallerSession struct {
	Contract *KetherHomepageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KetherHomepageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KetherHomepageTransactorSession struct {
	Contract     *KetherHomepageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KetherHomepageRaw is an auto generated low-level Go binding around an Ethereum contract.
type KetherHomepageRaw struct {
	Contract *KetherHomepage // Generic contract binding to access the raw methods on
}

// KetherHomepageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KetherHomepageCallerRaw struct {
	Contract *KetherHomepageCaller // Generic read-only contract binding to access the raw methods on
}

// KetherHomepageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KetherHomepageTransactorRaw struct {
	Contract *KetherHomepageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKetherHomepage creates a new instance of KetherHomepage, bound to a specific deployed contract.
func NewKetherHomepage(address common.Address, backend bind.ContractBackend) (*KetherHomepage, error) {
	contract, err := bindKetherHomepage(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KetherHomepage{KetherHomepageCaller: KetherHomepageCaller{contract: contract}, KetherHomepageTransactor: KetherHomepageTransactor{contract: contract}}, nil
}

// NewKetherHomepageCaller creates a new read-only instance of KetherHomepage, bound to a specific deployed contract.
func NewKetherHomepageCaller(address common.Address, caller bind.ContractCaller) (*KetherHomepageCaller, error) {
	contract, err := bindKetherHomepage(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageCaller{contract: contract}, nil
}

// NewKetherHomepageTransactor creates a new write-only instance of KetherHomepage, bound to a specific deployed contract.
func NewKetherHomepageTransactor(address common.Address, transactor bind.ContractTransactor) (*KetherHomepageTransactor, error) {
	contract, err := bindKetherHomepage(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &KetherHomepageTransactor{contract: contract}, nil
}

// bindKetherHomepage binds a generic wrapper to an already deployed contract.
func bindKetherHomepage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KetherHomepageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KetherHomepage *KetherHomepageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KetherHomepage.Contract.KetherHomepageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KetherHomepage *KetherHomepageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepage.Contract.KetherHomepageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KetherHomepage *KetherHomepageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KetherHomepage.Contract.KetherHomepageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KetherHomepage *KetherHomepageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _KetherHomepage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KetherHomepage *KetherHomepageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KetherHomepage *KetherHomepageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KetherHomepage.Contract.contract.Transact(opts, method, params...)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads( uint256) constant returns(owner address, x uint256, y uint256, width uint256, height uint256, link string, image string, title string, NSFW bool, forceNSFW bool)
func (_KetherHomepage *KetherHomepageCaller) Ads(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	ret := new(struct {
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
	out := ret
	err := _KetherHomepage.contract.Call(opts, out, "ads", arg0)
	return *ret, err
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads( uint256) constant returns(owner address, x uint256, y uint256, width uint256, height uint256, link string, image string, title string, NSFW bool, forceNSFW bool)
func (_KetherHomepage *KetherHomepageSession) Ads(arg0 *big.Int) (struct {
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
	return _KetherHomepage.Contract.Ads(&_KetherHomepage.CallOpts, arg0)
}

// Ads is a free data retrieval call binding the contract method 0x11a7a4c0.
//
// Solidity: function ads( uint256) constant returns(owner address, x uint256, y uint256, width uint256, height uint256, link string, image string, title string, NSFW bool, forceNSFW bool)
func (_KetherHomepage *KetherHomepageCallerSession) Ads(arg0 *big.Int) (struct {
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
	return _KetherHomepage.Contract.Ads(&_KetherHomepage.CallOpts, arg0)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCaller) GetAdsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KetherHomepage.contract.Call(opts, out, "getAdsLength")
	return *ret0, err
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() constant returns(uint256)
func (_KetherHomepage *KetherHomepageSession) GetAdsLength() (*big.Int, error) {
	return _KetherHomepage.Contract.GetAdsLength(&_KetherHomepage.CallOpts)
}

// GetAdsLength is a free data retrieval call binding the contract method 0x7a6adab6.
//
// Solidity: function getAdsLength() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCallerSession) GetAdsLength() (*big.Int, error) {
	return _KetherHomepage.Contract.GetAdsLength(&_KetherHomepage.CallOpts)
}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid( uint256,  uint256) constant returns(bool)
func (_KetherHomepage *KetherHomepageCaller) Grid(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _KetherHomepage.contract.Call(opts, out, "grid", arg0, arg1)
	return *ret0, err
}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid( uint256,  uint256) constant returns(bool)
func (_KetherHomepage *KetherHomepageSession) Grid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _KetherHomepage.Contract.Grid(&_KetherHomepage.CallOpts, arg0, arg1)
}

// Grid is a free data retrieval call binding the contract method 0x146008e3.
//
// Solidity: function grid( uint256,  uint256) constant returns(bool)
func (_KetherHomepage *KetherHomepageCallerSession) Grid(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _KetherHomepage.Contract.Grid(&_KetherHomepage.CallOpts, arg0, arg1)
}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCaller) PixelsPerCell(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KetherHomepage.contract.Call(opts, out, "pixelsPerCell")
	return *ret0, err
}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() constant returns(uint256)
func (_KetherHomepage *KetherHomepageSession) PixelsPerCell() (*big.Int, error) {
	return _KetherHomepage.Contract.PixelsPerCell(&_KetherHomepage.CallOpts)
}

// PixelsPerCell is a free data retrieval call binding the contract method 0x13f4b42c.
//
// Solidity: function pixelsPerCell() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCallerSession) PixelsPerCell() (*big.Int, error) {
	return _KetherHomepage.Contract.PixelsPerCell(&_KetherHomepage.CallOpts)
}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCaller) WeiPixelPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _KetherHomepage.contract.Call(opts, out, "weiPixelPrice")
	return *ret0, err
}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() constant returns(uint256)
func (_KetherHomepage *KetherHomepageSession) WeiPixelPrice() (*big.Int, error) {
	return _KetherHomepage.Contract.WeiPixelPrice(&_KetherHomepage.CallOpts)
}

// WeiPixelPrice is a free data retrieval call binding the contract method 0xd5bec84e.
//
// Solidity: function weiPixelPrice() constant returns(uint256)
func (_KetherHomepage *KetherHomepageCallerSession) WeiPixelPrice() (*big.Int, error) {
	return _KetherHomepage.Contract.WeiPixelPrice(&_KetherHomepage.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(_x uint256, _y uint256, _width uint256, _height uint256) returns(idx uint256)
func (_KetherHomepage *KetherHomepageTransactor) Buy(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepage.contract.Transact(opts, "buy", _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(_x uint256, _y uint256, _width uint256, _height uint256) returns(idx uint256)
func (_KetherHomepage *KetherHomepageSession) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepage.Contract.Buy(&_KetherHomepage.TransactOpts, _x, _y, _width, _height)
}

// Buy is a paid mutator transaction binding the contract method 0x1281311d.
//
// Solidity: function buy(_x uint256, _y uint256, _width uint256, _height uint256) returns(idx uint256)
func (_KetherHomepage *KetherHomepageTransactorSession) Buy(_x *big.Int, _y *big.Int, _width *big.Int, _height *big.Int) (*types.Transaction, error) {
	return _KetherHomepage.Contract.Buy(&_KetherHomepage.TransactOpts, _x, _y, _width, _height)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(_idx uint256, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageTransactor) ForceNSFW(opts *bind.TransactOpts, _idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.contract.Transact(opts, "forceNSFW", _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(_idx uint256, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageSession) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.Contract.ForceNSFW(&_KetherHomepage.TransactOpts, _idx, _NSFW)
}

// ForceNSFW is a paid mutator transaction binding the contract method 0xacff2fce.
//
// Solidity: function forceNSFW(_idx uint256, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageTransactorSession) ForceNSFW(_idx *big.Int, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.Contract.ForceNSFW(&_KetherHomepage.TransactOpts, _idx, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(_idx uint256, _link string, _image string, _title string, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageTransactor) Publish(opts *bind.TransactOpts, _idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.contract.Transact(opts, "publish", _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(_idx uint256, _link string, _image string, _title string, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageSession) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.Contract.Publish(&_KetherHomepage.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// Publish is a paid mutator transaction binding the contract method 0x45ebc145.
//
// Solidity: function publish(_idx uint256, _link string, _image string, _title string, _NSFW bool) returns()
func (_KetherHomepage *KetherHomepageTransactorSession) Publish(_idx *big.Int, _link string, _image string, _title string, _NSFW bool) (*types.Transaction, error) {
	return _KetherHomepage.Contract.Publish(&_KetherHomepage.TransactOpts, _idx, _link, _image, _title, _NSFW)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(_idx uint256, _newOwner address) returns()
func (_KetherHomepage *KetherHomepageTransactor) SetAdOwner(opts *bind.TransactOpts, _idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepage.contract.Transact(opts, "setAdOwner", _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(_idx uint256, _newOwner address) returns()
func (_KetherHomepage *KetherHomepageSession) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepage.Contract.SetAdOwner(&_KetherHomepage.TransactOpts, _idx, _newOwner)
}

// SetAdOwner is a paid mutator transaction binding the contract method 0x759c7a58.
//
// Solidity: function setAdOwner(_idx uint256, _newOwner address) returns()
func (_KetherHomepage *KetherHomepageTransactorSession) SetAdOwner(_idx *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _KetherHomepage.Contract.SetAdOwner(&_KetherHomepage.TransactOpts, _idx, _newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepage *KetherHomepageTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KetherHomepage.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepage *KetherHomepageSession) Withdraw() (*types.Transaction, error) {
	return _KetherHomepage.Contract.Withdraw(&_KetherHomepage.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_KetherHomepage *KetherHomepageTransactorSession) Withdraw() (*types.Transaction, error) {
	return _KetherHomepage.Contract.Withdraw(&_KetherHomepage.TransactOpts)
}

