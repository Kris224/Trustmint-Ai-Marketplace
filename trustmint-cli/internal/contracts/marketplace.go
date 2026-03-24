// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// TrustmintMarketplaceListing is an auto generated low-level Go binding around an user-defined struct.
type TrustmintMarketplaceListing struct {
	ListingId   *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Price       *big.Int
	Active      bool
}

// TrustmintMarketplaceMetaData contains all meta data concerning the TrustmintMarketplace contract.
var TrustmintMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"ListingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ListingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ModelPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"PlatformFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structTrustmintMarketplace.Listing\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"listModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"}],\"name\":\"purchaseModel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenToListing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TrustmintMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustmintMarketplaceMetaData.ABI instead.
var TrustmintMarketplaceABI = TrustmintMarketplaceMetaData.ABI

// TrustmintMarketplace is an auto generated Go binding around an Ethereum contract.
type TrustmintMarketplace struct {
	TrustmintMarketplaceCaller     // Read-only binding to the contract
	TrustmintMarketplaceTransactor // Write-only binding to the contract
	TrustmintMarketplaceFilterer   // Log filterer for contract events
}

// TrustmintMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustmintMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustmintMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustmintMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustmintMarketplaceSession struct {
	Contract     *TrustmintMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TrustmintMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustmintMarketplaceCallerSession struct {
	Contract *TrustmintMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// TrustmintMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustmintMarketplaceTransactorSession struct {
	Contract     *TrustmintMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// TrustmintMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustmintMarketplaceRaw struct {
	Contract *TrustmintMarketplace // Generic contract binding to access the raw methods on
}

// TrustmintMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustmintMarketplaceCallerRaw struct {
	Contract *TrustmintMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// TrustmintMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustmintMarketplaceTransactorRaw struct {
	Contract *TrustmintMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustmintMarketplace creates a new instance of TrustmintMarketplace, bound to a specific deployed contract.
func NewTrustmintMarketplace(address common.Address, backend bind.ContractBackend) (*TrustmintMarketplace, error) {
	contract, err := bindTrustmintMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplace{TrustmintMarketplaceCaller: TrustmintMarketplaceCaller{contract: contract}, TrustmintMarketplaceTransactor: TrustmintMarketplaceTransactor{contract: contract}, TrustmintMarketplaceFilterer: TrustmintMarketplaceFilterer{contract: contract}}, nil
}

// NewTrustmintMarketplaceCaller creates a new read-only instance of TrustmintMarketplace, bound to a specific deployed contract.
func NewTrustmintMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*TrustmintMarketplaceCaller, error) {
	contract, err := bindTrustmintMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceCaller{contract: contract}, nil
}

// NewTrustmintMarketplaceTransactor creates a new write-only instance of TrustmintMarketplace, bound to a specific deployed contract.
func NewTrustmintMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustmintMarketplaceTransactor, error) {
	contract, err := bindTrustmintMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceTransactor{contract: contract}, nil
}

// NewTrustmintMarketplaceFilterer creates a new log filterer instance of TrustmintMarketplace, bound to a specific deployed contract.
func NewTrustmintMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustmintMarketplaceFilterer, error) {
	contract, err := bindTrustmintMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceFilterer{contract: contract}, nil
}

// bindTrustmintMarketplace binds a generic wrapper to an already deployed contract.
func bindTrustmintMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustmintMarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustmintMarketplace *TrustmintMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustmintMarketplace.Contract.TrustmintMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustmintMarketplace *TrustmintMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.TrustmintMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustmintMarketplace *TrustmintMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.TrustmintMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustmintMarketplace *TrustmintMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustmintMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.contract.Transact(opts, method, params...)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,address,address,uint256,bool))
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (TrustmintMarketplaceListing, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "getListing", listingId)

	if err != nil {
		return *new(TrustmintMarketplaceListing), err
	}

	out0 := *abi.ConvertType(out[0], new(TrustmintMarketplaceListing)).(*TrustmintMarketplaceListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,address,address,uint256,bool))
func (_TrustmintMarketplace *TrustmintMarketplaceSession) GetListing(listingId *big.Int) (TrustmintMarketplaceListing, error) {
	return _TrustmintMarketplace.Contract.GetListing(&_TrustmintMarketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns((uint256,uint256,address,address,uint256,bool))
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) GetListing(listingId *big.Int) (TrustmintMarketplaceListing, error) {
	return _TrustmintMarketplace.Contract.GetListing(&_TrustmintMarketplace.CallOpts, listingId)
}

// IsListed is a free data retrieval call binding the contract method 0xcdb3cd25.
//
// Solidity: function isListed(address nftContract, uint256 tokenId) view returns(bool)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) IsListed(opts *bind.CallOpts, nftContract common.Address, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "isListed", nftContract, tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsListed is a free data retrieval call binding the contract method 0xcdb3cd25.
//
// Solidity: function isListed(address nftContract, uint256 tokenId) view returns(bool)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) IsListed(nftContract common.Address, tokenId *big.Int) (bool, error) {
	return _TrustmintMarketplace.Contract.IsListed(&_TrustmintMarketplace.CallOpts, nftContract, tokenId)
}

// IsListed is a free data retrieval call binding the contract method 0xcdb3cd25.
//
// Solidity: function isListed(address nftContract, uint256 tokenId) view returns(bool)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) IsListed(nftContract common.Address, tokenId *big.Int) (bool, error) {
	return _TrustmintMarketplace.Contract.IsListed(&_TrustmintMarketplace.CallOpts, nftContract, tokenId)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address nftContract, address seller, uint256 price, bool active)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId   *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Price       *big.Int
	Active      bool
}, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId   *big.Int
		TokenId     *big.Int
		NftContract common.Address
		Seller      common.Address
		Price       *big.Int
		Active      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NftContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Seller = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address nftContract, address seller, uint256 price, bool active)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) Listings(arg0 *big.Int) (struct {
	ListingId   *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Price       *big.Int
	Active      bool
}, error) {
	return _TrustmintMarketplace.Contract.Listings(&_TrustmintMarketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, uint256 tokenId, address nftContract, address seller, uint256 price, bool active)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId   *big.Int
	TokenId     *big.Int
	NftContract common.Address
	Seller      common.Address
	Price       *big.Int
	Active      bool
}, error) {
	return _TrustmintMarketplace.Contract.Listings(&_TrustmintMarketplace.CallOpts, arg0)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TrustmintMarketplace.Contract.OnERC721Received(&_TrustmintMarketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _TrustmintMarketplace.Contract.OnERC721Received(&_TrustmintMarketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) Owner() (common.Address, error) {
	return _TrustmintMarketplace.Contract.Owner(&_TrustmintMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) Owner() (common.Address, error) {
	return _TrustmintMarketplace.Contract.Owner(&_TrustmintMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) PlatformFee() (*big.Int, error) {
	return _TrustmintMarketplace.Contract.PlatformFee(&_TrustmintMarketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) PlatformFee() (*big.Int, error) {
	return _TrustmintMarketplace.Contract.PlatformFee(&_TrustmintMarketplace.CallOpts)
}

// TokenToListing is a free data retrieval call binding the contract method 0x6f7b4352.
//
// Solidity: function tokenToListing(address , uint256 ) view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) TokenToListing(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "tokenToListing", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToListing is a free data retrieval call binding the contract method 0x6f7b4352.
//
// Solidity: function tokenToListing(address , uint256 ) view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) TokenToListing(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TrustmintMarketplace.Contract.TokenToListing(&_TrustmintMarketplace.CallOpts, arg0, arg1)
}

// TokenToListing is a free data retrieval call binding the contract method 0x6f7b4352.
//
// Solidity: function tokenToListing(address , uint256 ) view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) TokenToListing(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TrustmintMarketplace.Contract.TokenToListing(&_TrustmintMarketplace.CallOpts, arg0, arg1)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCaller) TotalListings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustmintMarketplace.contract.Call(opts, &out, "totalListings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) TotalListings() (*big.Int, error) {
	return _TrustmintMarketplace.Contract.TotalListings(&_TrustmintMarketplace.CallOpts)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceCallerSession) TotalListings() (*big.Int, error) {
	return _TrustmintMarketplace.Contract.TotalListings(&_TrustmintMarketplace.CallOpts)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) CancelListing(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "cancelListing", listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.CancelListing(&_TrustmintMarketplace.TransactOpts, listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 listingId) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) CancelListing(listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.CancelListing(&_TrustmintMarketplace.TransactOpts, listingId)
}

// ListModel is a paid mutator transaction binding the contract method 0xd2870779.
//
// Solidity: function listModel(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) ListModel(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "listModel", nftContract, tokenId, price)
}

// ListModel is a paid mutator transaction binding the contract method 0xd2870779.
//
// Solidity: function listModel(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceSession) ListModel(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.ListModel(&_TrustmintMarketplace.TransactOpts, nftContract, tokenId, price)
}

// ListModel is a paid mutator transaction binding the contract method 0xd2870779.
//
// Solidity: function listModel(address nftContract, uint256 tokenId, uint256 price) returns(uint256)
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) ListModel(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.ListModel(&_TrustmintMarketplace.TransactOpts, nftContract, tokenId, price)
}

// PurchaseModel is a paid mutator transaction binding the contract method 0xdc746862.
//
// Solidity: function purchaseModel(uint256 listingId) payable returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) PurchaseModel(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "purchaseModel", listingId)
}

// PurchaseModel is a paid mutator transaction binding the contract method 0xdc746862.
//
// Solidity: function purchaseModel(uint256 listingId) payable returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) PurchaseModel(listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.PurchaseModel(&_TrustmintMarketplace.TransactOpts, listingId)
}

// PurchaseModel is a paid mutator transaction binding the contract method 0xdc746862.
//
// Solidity: function purchaseModel(uint256 listingId) payable returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) PurchaseModel(listingId *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.PurchaseModel(&_TrustmintMarketplace.TransactOpts, listingId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.RenounceOwnership(&_TrustmintMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.RenounceOwnership(&_TrustmintMarketplace.TransactOpts)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) SetPlatformFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "setPlatformFee", newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) SetPlatformFee(newFee *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.SetPlatformFee(&_TrustmintMarketplace.TransactOpts, newFee)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 newFee) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) SetPlatformFee(newFee *big.Int) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.SetPlatformFee(&_TrustmintMarketplace.TransactOpts, newFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.TransferOwnership(&_TrustmintMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.TransferOwnership(&_TrustmintMarketplace.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactor) WithdrawFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintMarketplace.contract.Transact(opts, "withdrawFees")
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceSession) WithdrawFees() (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.WithdrawFees(&_TrustmintMarketplace.TransactOpts)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x476343ee.
//
// Solidity: function withdrawFees() returns()
func (_TrustmintMarketplace *TrustmintMarketplaceTransactorSession) WithdrawFees() (*types.Transaction, error) {
	return _TrustmintMarketplace.Contract.WithdrawFees(&_TrustmintMarketplace.TransactOpts)
}

// TrustmintMarketplaceListingCancelledIterator is returned from FilterListingCancelled and is used to iterate over the raw logs and unpacked data for ListingCancelled events raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceListingCancelledIterator struct {
	Event *TrustmintMarketplaceListingCancelled // Event containing the contract specifics and raw log

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
func (it *TrustmintMarketplaceListingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintMarketplaceListingCancelled)
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
		it.Event = new(TrustmintMarketplaceListingCancelled)
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
func (it *TrustmintMarketplaceListingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintMarketplaceListingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintMarketplaceListingCancelled represents a ListingCancelled event raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceListingCancelled struct {
	ListingId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingCancelled is a free log retrieval operation binding the contract event 0x411aee90354c51b1b04cd563fcab2617142a9d50da19232d888547c8a1b7fd8a.
//
// Solidity: event ListingCancelled(uint256 indexed listingId)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) FilterListingCancelled(opts *bind.FilterOpts, listingId []*big.Int) (*TrustmintMarketplaceListingCancelledIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.FilterLogs(opts, "ListingCancelled", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceListingCancelledIterator{contract: _TrustmintMarketplace.contract, event: "ListingCancelled", logs: logs, sub: sub}, nil
}

// WatchListingCancelled is a free log subscription operation binding the contract event 0x411aee90354c51b1b04cd563fcab2617142a9d50da19232d888547c8a1b7fd8a.
//
// Solidity: event ListingCancelled(uint256 indexed listingId)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) WatchListingCancelled(opts *bind.WatchOpts, sink chan<- *TrustmintMarketplaceListingCancelled, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.WatchLogs(opts, "ListingCancelled", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintMarketplaceListingCancelled)
				if err := _TrustmintMarketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
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

// ParseListingCancelled is a log parse operation binding the contract event 0x411aee90354c51b1b04cd563fcab2617142a9d50da19232d888547c8a1b7fd8a.
//
// Solidity: event ListingCancelled(uint256 indexed listingId)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) ParseListingCancelled(log types.Log) (*TrustmintMarketplaceListingCancelled, error) {
	event := new(TrustmintMarketplaceListingCancelled)
	if err := _TrustmintMarketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintMarketplaceListingCreatedIterator is returned from FilterListingCreated and is used to iterate over the raw logs and unpacked data for ListingCreated events raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceListingCreatedIterator struct {
	Event *TrustmintMarketplaceListingCreated // Event containing the contract specifics and raw log

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
func (it *TrustmintMarketplaceListingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintMarketplaceListingCreated)
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
		it.Event = new(TrustmintMarketplaceListingCreated)
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
func (it *TrustmintMarketplaceListingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintMarketplaceListingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintMarketplaceListingCreated represents a ListingCreated event raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceListingCreated struct {
	ListingId *big.Int
	TokenId   *big.Int
	Seller    common.Address
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingCreated is a free log retrieval operation binding the contract event 0x8093adc67e6531e4dcb9f96fe66947096b91ad10573e8136f8ef1a1cebe2fad6.
//
// Solidity: event ListingCreated(uint256 indexed listingId, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) FilterListingCreated(opts *bind.FilterOpts, listingId []*big.Int, tokenId []*big.Int, seller []common.Address) (*TrustmintMarketplaceListingCreatedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.FilterLogs(opts, "ListingCreated", listingIdRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceListingCreatedIterator{contract: _TrustmintMarketplace.contract, event: "ListingCreated", logs: logs, sub: sub}, nil
}

// WatchListingCreated is a free log subscription operation binding the contract event 0x8093adc67e6531e4dcb9f96fe66947096b91ad10573e8136f8ef1a1cebe2fad6.
//
// Solidity: event ListingCreated(uint256 indexed listingId, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) WatchListingCreated(opts *bind.WatchOpts, sink chan<- *TrustmintMarketplaceListingCreated, listingId []*big.Int, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.WatchLogs(opts, "ListingCreated", listingIdRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintMarketplaceListingCreated)
				if err := _TrustmintMarketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
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

// ParseListingCreated is a log parse operation binding the contract event 0x8093adc67e6531e4dcb9f96fe66947096b91ad10573e8136f8ef1a1cebe2fad6.
//
// Solidity: event ListingCreated(uint256 indexed listingId, uint256 indexed tokenId, address indexed seller, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) ParseListingCreated(log types.Log) (*TrustmintMarketplaceListingCreated, error) {
	event := new(TrustmintMarketplaceListingCreated)
	if err := _TrustmintMarketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintMarketplaceModelPurchasedIterator is returned from FilterModelPurchased and is used to iterate over the raw logs and unpacked data for ModelPurchased events raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceModelPurchasedIterator struct {
	Event *TrustmintMarketplaceModelPurchased // Event containing the contract specifics and raw log

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
func (it *TrustmintMarketplaceModelPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintMarketplaceModelPurchased)
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
		it.Event = new(TrustmintMarketplaceModelPurchased)
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
func (it *TrustmintMarketplaceModelPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintMarketplaceModelPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintMarketplaceModelPurchased represents a ModelPurchased event raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceModelPurchased struct {
	ListingId *big.Int
	TokenId   *big.Int
	Seller    common.Address
	Buyer     common.Address
	Price     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModelPurchased is a free log retrieval operation binding the contract event 0xbd877f632d807754a1c87b928fd4a301d95b1ec1e699a282b5bc14afe8d6acef.
//
// Solidity: event ModelPurchased(uint256 indexed listingId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) FilterModelPurchased(opts *bind.FilterOpts, listingId []*big.Int, tokenId []*big.Int) (*TrustmintMarketplaceModelPurchasedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.FilterLogs(opts, "ModelPurchased", listingIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceModelPurchasedIterator{contract: _TrustmintMarketplace.contract, event: "ModelPurchased", logs: logs, sub: sub}, nil
}

// WatchModelPurchased is a free log subscription operation binding the contract event 0xbd877f632d807754a1c87b928fd4a301d95b1ec1e699a282b5bc14afe8d6acef.
//
// Solidity: event ModelPurchased(uint256 indexed listingId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) WatchModelPurchased(opts *bind.WatchOpts, sink chan<- *TrustmintMarketplaceModelPurchased, listingId []*big.Int, tokenId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.WatchLogs(opts, "ModelPurchased", listingIdRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintMarketplaceModelPurchased)
				if err := _TrustmintMarketplace.contract.UnpackLog(event, "ModelPurchased", log); err != nil {
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

// ParseModelPurchased is a log parse operation binding the contract event 0xbd877f632d807754a1c87b928fd4a301d95b1ec1e699a282b5bc14afe8d6acef.
//
// Solidity: event ModelPurchased(uint256 indexed listingId, uint256 indexed tokenId, address seller, address buyer, uint256 price)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) ParseModelPurchased(log types.Log) (*TrustmintMarketplaceModelPurchased, error) {
	event := new(TrustmintMarketplaceModelPurchased)
	if err := _TrustmintMarketplace.contract.UnpackLog(event, "ModelPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceOwnershipTransferredIterator struct {
	Event *TrustmintMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrustmintMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintMarketplaceOwnershipTransferred)
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
		it.Event = new(TrustmintMarketplaceOwnershipTransferred)
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
func (it *TrustmintMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the TrustmintMarketplace contract.
type TrustmintMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrustmintMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplaceOwnershipTransferredIterator{contract: _TrustmintMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrustmintMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrustmintMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintMarketplaceOwnershipTransferred)
				if err := _TrustmintMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*TrustmintMarketplaceOwnershipTransferred, error) {
	event := new(TrustmintMarketplaceOwnershipTransferred)
	if err := _TrustmintMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintMarketplacePlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the TrustmintMarketplace contract.
type TrustmintMarketplacePlatformFeeUpdatedIterator struct {
	Event *TrustmintMarketplacePlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *TrustmintMarketplacePlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintMarketplacePlatformFeeUpdated)
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
		it.Event = new(TrustmintMarketplacePlatformFeeUpdated)
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
func (it *TrustmintMarketplacePlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintMarketplacePlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintMarketplacePlatformFeeUpdated represents a PlatformFeeUpdated event raised by the TrustmintMarketplace contract.
type TrustmintMarketplacePlatformFeeUpdated struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*TrustmintMarketplacePlatformFeeUpdatedIterator, error) {

	logs, sub, err := _TrustmintMarketplace.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &TrustmintMarketplacePlatformFeeUpdatedIterator{contract: _TrustmintMarketplace.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *TrustmintMarketplacePlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _TrustmintMarketplace.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintMarketplacePlatformFeeUpdated)
				if err := _TrustmintMarketplace.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
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

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0x45610d581145924dd7090a5017e5f2b1d6f42213bb2e95707ff86846bbfcb1ca.
//
// Solidity: event PlatformFeeUpdated(uint256 newFee)
func (_TrustmintMarketplace *TrustmintMarketplaceFilterer) ParsePlatformFeeUpdated(log types.Log) (*TrustmintMarketplacePlatformFeeUpdated, error) {
	event := new(TrustmintMarketplacePlatformFeeUpdated)
	if err := _TrustmintMarketplace.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
