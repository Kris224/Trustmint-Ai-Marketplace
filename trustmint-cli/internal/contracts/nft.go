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

// TrustmintNFTMetaData contains all meta data concerning the TrustmintNFT contract.
var TrustmintNFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidDefaultRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidDefaultRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominator\",\"type\":\"uint256\"}],\"name\":\"ERC2981InvalidTokenRoyalty\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC2981InvalidTokenRoyaltyReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"marketplace\",\"type\":\"address\"}],\"name\":\"MarketplaceUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"modelHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ipfsCid\",\"type\":\"string\"}],\"name\":\"ModelMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"ModelVerified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getModelMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"modelHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"datasetHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsCid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplace\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"modelHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"datasetHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"metadataUri\",\"type\":\"string\"}],\"name\":\"mintModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"modelMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"modelHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"datasetHash\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"ipfsCid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"registeredModels\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplace\",\"type\":\"address\"}],\"name\":\"setMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verified\",\"type\":\"bool\"}],\"name\":\"setVerificationStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"modelHash\",\"type\":\"string\"}],\"name\":\"verifyModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TrustmintNFTABI is the input ABI used to generate the binding from.
// Deprecated: Use TrustmintNFTMetaData.ABI instead.
var TrustmintNFTABI = TrustmintNFTMetaData.ABI

// TrustmintNFT is an auto generated Go binding around an Ethereum contract.
type TrustmintNFT struct {
	TrustmintNFTCaller     // Read-only binding to the contract
	TrustmintNFTTransactor // Write-only binding to the contract
	TrustmintNFTFilterer   // Log filterer for contract events
}

// TrustmintNFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrustmintNFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintNFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrustmintNFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintNFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrustmintNFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrustmintNFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrustmintNFTSession struct {
	Contract     *TrustmintNFT     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrustmintNFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrustmintNFTCallerSession struct {
	Contract *TrustmintNFTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TrustmintNFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrustmintNFTTransactorSession struct {
	Contract     *TrustmintNFTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TrustmintNFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrustmintNFTRaw struct {
	Contract *TrustmintNFT // Generic contract binding to access the raw methods on
}

// TrustmintNFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrustmintNFTCallerRaw struct {
	Contract *TrustmintNFTCaller // Generic read-only contract binding to access the raw methods on
}

// TrustmintNFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrustmintNFTTransactorRaw struct {
	Contract *TrustmintNFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrustmintNFT creates a new instance of TrustmintNFT, bound to a specific deployed contract.
func NewTrustmintNFT(address common.Address, backend bind.ContractBackend) (*TrustmintNFT, error) {
	contract, err := bindTrustmintNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFT{TrustmintNFTCaller: TrustmintNFTCaller{contract: contract}, TrustmintNFTTransactor: TrustmintNFTTransactor{contract: contract}, TrustmintNFTFilterer: TrustmintNFTFilterer{contract: contract}}, nil
}

// NewTrustmintNFTCaller creates a new read-only instance of TrustmintNFT, bound to a specific deployed contract.
func NewTrustmintNFTCaller(address common.Address, caller bind.ContractCaller) (*TrustmintNFTCaller, error) {
	contract, err := bindTrustmintNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTCaller{contract: contract}, nil
}

// NewTrustmintNFTTransactor creates a new write-only instance of TrustmintNFT, bound to a specific deployed contract.
func NewTrustmintNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*TrustmintNFTTransactor, error) {
	contract, err := bindTrustmintNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTTransactor{contract: contract}, nil
}

// NewTrustmintNFTFilterer creates a new log filterer instance of TrustmintNFT, bound to a specific deployed contract.
func NewTrustmintNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*TrustmintNFTFilterer, error) {
	contract, err := bindTrustmintNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTFilterer{contract: contract}, nil
}

// bindTrustmintNFT binds a generic wrapper to an already deployed contract.
func bindTrustmintNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TrustmintNFTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustmintNFT *TrustmintNFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustmintNFT.Contract.TrustmintNFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustmintNFT *TrustmintNFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TrustmintNFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustmintNFT *TrustmintNFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TrustmintNFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrustmintNFT *TrustmintNFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TrustmintNFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrustmintNFT *TrustmintNFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrustmintNFT *TrustmintNFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TrustmintNFT *TrustmintNFTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TrustmintNFT *TrustmintNFTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TrustmintNFT.Contract.BalanceOf(&_TrustmintNFT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_TrustmintNFT *TrustmintNFTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _TrustmintNFT.Contract.BalanceOf(&_TrustmintNFT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TrustmintNFT.Contract.GetApproved(&_TrustmintNFT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _TrustmintNFT.Contract.GetApproved(&_TrustmintNFT.CallOpts, tokenId)
}

// GetModelMetadata is a free data retrieval call binding the contract method 0xfe55168d.
//
// Solidity: function getModelMetadata(uint256 tokenId) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTCaller) GetModelMetadata(opts *bind.CallOpts, tokenId *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "getModelMetadata", tokenId)

	outstruct := new(struct {
		ModelHash   string
		DatasetHash string
		MerkleRoot  [32]byte
		IpfsCid     string
		Creator     common.Address
		Timestamp   *big.Int
		Verified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DatasetHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MerkleRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IpfsCid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Verified = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// GetModelMetadata is a free data retrieval call binding the contract method 0xfe55168d.
//
// Solidity: function getModelMetadata(uint256 tokenId) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTSession) GetModelMetadata(tokenId *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	return _TrustmintNFT.Contract.GetModelMetadata(&_TrustmintNFT.CallOpts, tokenId)
}

// GetModelMetadata is a free data retrieval call binding the contract method 0xfe55168d.
//
// Solidity: function getModelMetadata(uint256 tokenId) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTCallerSession) GetModelMetadata(tokenId *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	return _TrustmintNFT.Contract.GetModelMetadata(&_TrustmintNFT.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TrustmintNFT *TrustmintNFTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TrustmintNFT.Contract.IsApprovedForAll(&_TrustmintNFT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _TrustmintNFT.Contract.IsApprovedForAll(&_TrustmintNFT.CallOpts, owner, operator)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_TrustmintNFT *TrustmintNFTCaller) Marketplace(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "marketplace")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_TrustmintNFT *TrustmintNFTSession) Marketplace() (common.Address, error) {
	return _TrustmintNFT.Contract.Marketplace(&_TrustmintNFT.CallOpts)
}

// Marketplace is a free data retrieval call binding the contract method 0xabc8c7af.
//
// Solidity: function marketplace() view returns(address)
func (_TrustmintNFT *TrustmintNFTCallerSession) Marketplace() (common.Address, error) {
	return _TrustmintNFT.Contract.Marketplace(&_TrustmintNFT.CallOpts)
}

// ModelMetadata is a free data retrieval call binding the contract method 0x821ffb50.
//
// Solidity: function modelMetadata(uint256 ) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTCaller) ModelMetadata(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "modelMetadata", arg0)

	outstruct := new(struct {
		ModelHash   string
		DatasetHash string
		MerkleRoot  [32]byte
		IpfsCid     string
		Creator     common.Address
		Timestamp   *big.Int
		Verified    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelHash = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DatasetHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MerkleRoot = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.IpfsCid = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Verified = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// ModelMetadata is a free data retrieval call binding the contract method 0x821ffb50.
//
// Solidity: function modelMetadata(uint256 ) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTSession) ModelMetadata(arg0 *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	return _TrustmintNFT.Contract.ModelMetadata(&_TrustmintNFT.CallOpts, arg0)
}

// ModelMetadata is a free data retrieval call binding the contract method 0x821ffb50.
//
// Solidity: function modelMetadata(uint256 ) view returns(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, address creator, uint256 timestamp, bool verified)
func (_TrustmintNFT *TrustmintNFTCallerSession) ModelMetadata(arg0 *big.Int) (struct {
	ModelHash   string
	DatasetHash string
	MerkleRoot  [32]byte
	IpfsCid     string
	Creator     common.Address
	Timestamp   *big.Int
	Verified    bool
}, error) {
	return _TrustmintNFT.Contract.ModelMetadata(&_TrustmintNFT.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TrustmintNFT *TrustmintNFTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TrustmintNFT *TrustmintNFTSession) Name() (string, error) {
	return _TrustmintNFT.Contract.Name(&_TrustmintNFT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TrustmintNFT *TrustmintNFTCallerSession) Name() (string, error) {
	return _TrustmintNFT.Contract.Name(&_TrustmintNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintNFT *TrustmintNFTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintNFT *TrustmintNFTSession) Owner() (common.Address, error) {
	return _TrustmintNFT.Contract.Owner(&_TrustmintNFT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TrustmintNFT *TrustmintNFTCallerSession) Owner() (common.Address, error) {
	return _TrustmintNFT.Contract.Owner(&_TrustmintNFT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TrustmintNFT.Contract.OwnerOf(&_TrustmintNFT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_TrustmintNFT *TrustmintNFTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _TrustmintNFT.Contract.OwnerOf(&_TrustmintNFT.CallOpts, tokenId)
}

// RegisteredModels is a free data retrieval call binding the contract method 0xaa03b039.
//
// Solidity: function registeredModels(string ) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCaller) RegisteredModels(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "registeredModels", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredModels is a free data retrieval call binding the contract method 0xaa03b039.
//
// Solidity: function registeredModels(string ) view returns(bool)
func (_TrustmintNFT *TrustmintNFTSession) RegisteredModels(arg0 string) (bool, error) {
	return _TrustmintNFT.Contract.RegisteredModels(&_TrustmintNFT.CallOpts, arg0)
}

// RegisteredModels is a free data retrieval call binding the contract method 0xaa03b039.
//
// Solidity: function registeredModels(string ) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCallerSession) RegisteredModels(arg0 string) (bool, error) {
	return _TrustmintNFT.Contract.RegisteredModels(&_TrustmintNFT.CallOpts, arg0)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_TrustmintNFT *TrustmintNFTCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver common.Address
		Amount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_TrustmintNFT *TrustmintNFTSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _TrustmintNFT.Contract.RoyaltyInfo(&_TrustmintNFT.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 amount)
func (_TrustmintNFT *TrustmintNFTCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver common.Address
	Amount   *big.Int
}, error) {
	return _TrustmintNFT.Contract.RoyaltyInfo(&_TrustmintNFT.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustmintNFT *TrustmintNFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustmintNFT.Contract.SupportsInterface(&_TrustmintNFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TrustmintNFT.Contract.SupportsInterface(&_TrustmintNFT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TrustmintNFT *TrustmintNFTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TrustmintNFT *TrustmintNFTSession) Symbol() (string, error) {
	return _TrustmintNFT.Contract.Symbol(&_TrustmintNFT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TrustmintNFT *TrustmintNFTCallerSession) Symbol() (string, error) {
	return _TrustmintNFT.Contract.Symbol(&_TrustmintNFT.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TrustmintNFT *TrustmintNFTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TrustmintNFT *TrustmintNFTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TrustmintNFT.Contract.TokenURI(&_TrustmintNFT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_TrustmintNFT *TrustmintNFTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _TrustmintNFT.Contract.TokenURI(&_TrustmintNFT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TrustmintNFT *TrustmintNFTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TrustmintNFT *TrustmintNFTSession) TotalSupply() (*big.Int, error) {
	return _TrustmintNFT.Contract.TotalSupply(&_TrustmintNFT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TrustmintNFT *TrustmintNFTCallerSession) TotalSupply() (*big.Int, error) {
	return _TrustmintNFT.Contract.TotalSupply(&_TrustmintNFT.CallOpts)
}

// VerifyModel is a free data retrieval call binding the contract method 0x94cb282b.
//
// Solidity: function verifyModel(uint256 tokenId, string modelHash) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCaller) VerifyModel(opts *bind.CallOpts, tokenId *big.Int, modelHash string) (bool, error) {
	var out []interface{}
	err := _TrustmintNFT.contract.Call(opts, &out, "verifyModel", tokenId, modelHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyModel is a free data retrieval call binding the contract method 0x94cb282b.
//
// Solidity: function verifyModel(uint256 tokenId, string modelHash) view returns(bool)
func (_TrustmintNFT *TrustmintNFTSession) VerifyModel(tokenId *big.Int, modelHash string) (bool, error) {
	return _TrustmintNFT.Contract.VerifyModel(&_TrustmintNFT.CallOpts, tokenId, modelHash)
}

// VerifyModel is a free data retrieval call binding the contract method 0x94cb282b.
//
// Solidity: function verifyModel(uint256 tokenId, string modelHash) view returns(bool)
func (_TrustmintNFT *TrustmintNFTCallerSession) VerifyModel(tokenId *big.Int, modelHash string) (bool, error) {
	return _TrustmintNFT.Contract.VerifyModel(&_TrustmintNFT.CallOpts, tokenId, modelHash)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.Approve(&_TrustmintNFT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.Approve(&_TrustmintNFT.TransactOpts, to, tokenId)
}

// MintModel is a paid mutator transaction binding the contract method 0xe69d5d79.
//
// Solidity: function mintModel(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, string metadataUri) returns(uint256)
func (_TrustmintNFT *TrustmintNFTTransactor) MintModel(opts *bind.TransactOpts, modelHash string, datasetHash string, merkleRoot [32]byte, ipfsCid string, metadataUri string) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "mintModel", modelHash, datasetHash, merkleRoot, ipfsCid, metadataUri)
}

// MintModel is a paid mutator transaction binding the contract method 0xe69d5d79.
//
// Solidity: function mintModel(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, string metadataUri) returns(uint256)
func (_TrustmintNFT *TrustmintNFTSession) MintModel(modelHash string, datasetHash string, merkleRoot [32]byte, ipfsCid string, metadataUri string) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.MintModel(&_TrustmintNFT.TransactOpts, modelHash, datasetHash, merkleRoot, ipfsCid, metadataUri)
}

// MintModel is a paid mutator transaction binding the contract method 0xe69d5d79.
//
// Solidity: function mintModel(string modelHash, string datasetHash, bytes32 merkleRoot, string ipfsCid, string metadataUri) returns(uint256)
func (_TrustmintNFT *TrustmintNFTTransactorSession) MintModel(modelHash string, datasetHash string, merkleRoot [32]byte, ipfsCid string, metadataUri string) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.MintModel(&_TrustmintNFT.TransactOpts, modelHash, datasetHash, merkleRoot, ipfsCid, metadataUri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintNFT *TrustmintNFTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintNFT *TrustmintNFTSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrustmintNFT.Contract.RenounceOwnership(&_TrustmintNFT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TrustmintNFT.Contract.RenounceOwnership(&_TrustmintNFT.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SafeTransferFrom(&_TrustmintNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SafeTransferFrom(&_TrustmintNFT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TrustmintNFT *TrustmintNFTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SafeTransferFrom0(&_TrustmintNFT.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SafeTransferFrom0(&_TrustmintNFT.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TrustmintNFT *TrustmintNFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetApprovalForAll(&_TrustmintNFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetApprovalForAll(&_TrustmintNFT.TransactOpts, operator, approved)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) SetMarketplace(opts *bind.TransactOpts, _marketplace common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "setMarketplace", _marketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_TrustmintNFT *TrustmintNFTSession) SetMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetMarketplace(&_TrustmintNFT.TransactOpts, _marketplace)
}

// SetMarketplace is a paid mutator transaction binding the contract method 0x73ad6c2d.
//
// Solidity: function setMarketplace(address _marketplace) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) SetMarketplace(_marketplace common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetMarketplace(&_TrustmintNFT.TransactOpts, _marketplace)
}

// SetVerificationStatus is a paid mutator transaction binding the contract method 0x0ef644a5.
//
// Solidity: function setVerificationStatus(uint256 tokenId, bool verified) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) SetVerificationStatus(opts *bind.TransactOpts, tokenId *big.Int, verified bool) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "setVerificationStatus", tokenId, verified)
}

// SetVerificationStatus is a paid mutator transaction binding the contract method 0x0ef644a5.
//
// Solidity: function setVerificationStatus(uint256 tokenId, bool verified) returns()
func (_TrustmintNFT *TrustmintNFTSession) SetVerificationStatus(tokenId *big.Int, verified bool) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetVerificationStatus(&_TrustmintNFT.TransactOpts, tokenId, verified)
}

// SetVerificationStatus is a paid mutator transaction binding the contract method 0x0ef644a5.
//
// Solidity: function setVerificationStatus(uint256 tokenId, bool verified) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) SetVerificationStatus(tokenId *big.Int, verified bool) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.SetVerificationStatus(&_TrustmintNFT.TransactOpts, tokenId, verified)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TransferFrom(&_TrustmintNFT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TransferFrom(&_TrustmintNFT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintNFT *TrustmintNFTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintNFT *TrustmintNFTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TransferOwnership(&_TrustmintNFT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TrustmintNFT *TrustmintNFTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrustmintNFT.Contract.TransferOwnership(&_TrustmintNFT.TransactOpts, newOwner)
}

// TrustmintNFTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TrustmintNFT contract.
type TrustmintNFTApprovalIterator struct {
	Event *TrustmintNFTApproval // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTApproval)
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
		it.Event = new(TrustmintNFTApproval)
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
func (it *TrustmintNFTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTApproval represents a Approval event raised by the TrustmintNFT contract.
type TrustmintNFTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*TrustmintNFTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTApprovalIterator{contract: _TrustmintNFT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TrustmintNFTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTApproval)
				if err := _TrustmintNFT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseApproval(log types.Log) (*TrustmintNFTApproval, error) {
	event := new(TrustmintNFTApproval)
	if err := _TrustmintNFT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the TrustmintNFT contract.
type TrustmintNFTApprovalForAllIterator struct {
	Event *TrustmintNFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTApprovalForAll)
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
		it.Event = new(TrustmintNFTApprovalForAll)
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
func (it *TrustmintNFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTApprovalForAll represents a ApprovalForAll event raised by the TrustmintNFT contract.
type TrustmintNFTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*TrustmintNFTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTApprovalForAllIterator{contract: _TrustmintNFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *TrustmintNFTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTApprovalForAll)
				if err := _TrustmintNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseApprovalForAll(log types.Log) (*TrustmintNFTApprovalForAll, error) {
	event := new(TrustmintNFTApprovalForAll)
	if err := _TrustmintNFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the TrustmintNFT contract.
type TrustmintNFTBatchMetadataUpdateIterator struct {
	Event *TrustmintNFTBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTBatchMetadataUpdate)
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
		it.Event = new(TrustmintNFTBatchMetadataUpdate)
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
func (it *TrustmintNFTBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the TrustmintNFT contract.
type TrustmintNFTBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*TrustmintNFTBatchMetadataUpdateIterator, error) {

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTBatchMetadataUpdateIterator{contract: _TrustmintNFT.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TrustmintNFTBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTBatchMetadataUpdate)
				if err := _TrustmintNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseBatchMetadataUpdate(log types.Log) (*TrustmintNFTBatchMetadataUpdate, error) {
	event := new(TrustmintNFTBatchMetadataUpdate)
	if err := _TrustmintNFT.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTMarketplaceUpdatedIterator is returned from FilterMarketplaceUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceUpdated events raised by the TrustmintNFT contract.
type TrustmintNFTMarketplaceUpdatedIterator struct {
	Event *TrustmintNFTMarketplaceUpdated // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTMarketplaceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTMarketplaceUpdated)
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
		it.Event = new(TrustmintNFTMarketplaceUpdated)
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
func (it *TrustmintNFTMarketplaceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTMarketplaceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTMarketplaceUpdated represents a MarketplaceUpdated event raised by the TrustmintNFT contract.
type TrustmintNFTMarketplaceUpdated struct {
	Marketplace common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceUpdated is a free log retrieval operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address indexed marketplace)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterMarketplaceUpdated(opts *bind.FilterOpts, marketplace []common.Address) (*TrustmintNFTMarketplaceUpdatedIterator, error) {

	var marketplaceRule []interface{}
	for _, marketplaceItem := range marketplace {
		marketplaceRule = append(marketplaceRule, marketplaceItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "MarketplaceUpdated", marketplaceRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTMarketplaceUpdatedIterator{contract: _TrustmintNFT.contract, event: "MarketplaceUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceUpdated is a free log subscription operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address indexed marketplace)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchMarketplaceUpdated(opts *bind.WatchOpts, sink chan<- *TrustmintNFTMarketplaceUpdated, marketplace []common.Address) (event.Subscription, error) {

	var marketplaceRule []interface{}
	for _, marketplaceItem := range marketplace {
		marketplaceRule = append(marketplaceRule, marketplaceItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "MarketplaceUpdated", marketplaceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTMarketplaceUpdated)
				if err := _TrustmintNFT.contract.UnpackLog(event, "MarketplaceUpdated", log); err != nil {
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

// ParseMarketplaceUpdated is a log parse operation binding the contract event 0x210690abd7fd6cdbb8f2beb202b2a253d58d7a0813b2175c4172c14c0c1af6dc.
//
// Solidity: event MarketplaceUpdated(address indexed marketplace)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseMarketplaceUpdated(log types.Log) (*TrustmintNFTMarketplaceUpdated, error) {
	event := new(TrustmintNFTMarketplaceUpdated)
	if err := _TrustmintNFT.contract.UnpackLog(event, "MarketplaceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the TrustmintNFT contract.
type TrustmintNFTMetadataUpdateIterator struct {
	Event *TrustmintNFTMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTMetadataUpdate)
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
		it.Event = new(TrustmintNFTMetadataUpdate)
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
func (it *TrustmintNFTMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTMetadataUpdate represents a MetadataUpdate event raised by the TrustmintNFT contract.
type TrustmintNFTMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*TrustmintNFTMetadataUpdateIterator, error) {

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTMetadataUpdateIterator{contract: _TrustmintNFT.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *TrustmintNFTMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTMetadataUpdate)
				if err := _TrustmintNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseMetadataUpdate(log types.Log) (*TrustmintNFTMetadataUpdate, error) {
	event := new(TrustmintNFTMetadataUpdate)
	if err := _TrustmintNFT.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTModelMintedIterator is returned from FilterModelMinted and is used to iterate over the raw logs and unpacked data for ModelMinted events raised by the TrustmintNFT contract.
type TrustmintNFTModelMintedIterator struct {
	Event *TrustmintNFTModelMinted // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTModelMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTModelMinted)
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
		it.Event = new(TrustmintNFTModelMinted)
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
func (it *TrustmintNFTModelMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTModelMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTModelMinted represents a ModelMinted event raised by the TrustmintNFT contract.
type TrustmintNFTModelMinted struct {
	TokenId   *big.Int
	Creator   common.Address
	ModelHash string
	IpfsCid   string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterModelMinted is a free log retrieval operation binding the contract event 0x6cc399f284901c0712abd5514a48a1fb445a4b3264444ba12dd49271d50769fc.
//
// Solidity: event ModelMinted(uint256 indexed tokenId, address indexed creator, string modelHash, string ipfsCid)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterModelMinted(opts *bind.FilterOpts, tokenId []*big.Int, creator []common.Address) (*TrustmintNFTModelMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "ModelMinted", tokenIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTModelMintedIterator{contract: _TrustmintNFT.contract, event: "ModelMinted", logs: logs, sub: sub}, nil
}

// WatchModelMinted is a free log subscription operation binding the contract event 0x6cc399f284901c0712abd5514a48a1fb445a4b3264444ba12dd49271d50769fc.
//
// Solidity: event ModelMinted(uint256 indexed tokenId, address indexed creator, string modelHash, string ipfsCid)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchModelMinted(opts *bind.WatchOpts, sink chan<- *TrustmintNFTModelMinted, tokenId []*big.Int, creator []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "ModelMinted", tokenIdRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTModelMinted)
				if err := _TrustmintNFT.contract.UnpackLog(event, "ModelMinted", log); err != nil {
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

// ParseModelMinted is a log parse operation binding the contract event 0x6cc399f284901c0712abd5514a48a1fb445a4b3264444ba12dd49271d50769fc.
//
// Solidity: event ModelMinted(uint256 indexed tokenId, address indexed creator, string modelHash, string ipfsCid)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseModelMinted(log types.Log) (*TrustmintNFTModelMinted, error) {
	event := new(TrustmintNFTModelMinted)
	if err := _TrustmintNFT.contract.UnpackLog(event, "ModelMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTModelVerifiedIterator is returned from FilterModelVerified and is used to iterate over the raw logs and unpacked data for ModelVerified events raised by the TrustmintNFT contract.
type TrustmintNFTModelVerifiedIterator struct {
	Event *TrustmintNFTModelVerified // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTModelVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTModelVerified)
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
		it.Event = new(TrustmintNFTModelVerified)
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
func (it *TrustmintNFTModelVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTModelVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTModelVerified represents a ModelVerified event raised by the TrustmintNFT contract.
type TrustmintNFTModelVerified struct {
	TokenId  *big.Int
	Verified bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModelVerified is a free log retrieval operation binding the contract event 0xa1b0333125bdfed2c89a1b809f00f0d77b1db6c89940efab5f7c2108b8e6b57c.
//
// Solidity: event ModelVerified(uint256 indexed tokenId, bool verified)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterModelVerified(opts *bind.FilterOpts, tokenId []*big.Int) (*TrustmintNFTModelVerifiedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "ModelVerified", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTModelVerifiedIterator{contract: _TrustmintNFT.contract, event: "ModelVerified", logs: logs, sub: sub}, nil
}

// WatchModelVerified is a free log subscription operation binding the contract event 0xa1b0333125bdfed2c89a1b809f00f0d77b1db6c89940efab5f7c2108b8e6b57c.
//
// Solidity: event ModelVerified(uint256 indexed tokenId, bool verified)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchModelVerified(opts *bind.WatchOpts, sink chan<- *TrustmintNFTModelVerified, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "ModelVerified", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTModelVerified)
				if err := _TrustmintNFT.contract.UnpackLog(event, "ModelVerified", log); err != nil {
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

// ParseModelVerified is a log parse operation binding the contract event 0xa1b0333125bdfed2c89a1b809f00f0d77b1db6c89940efab5f7c2108b8e6b57c.
//
// Solidity: event ModelVerified(uint256 indexed tokenId, bool verified)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseModelVerified(log types.Log) (*TrustmintNFTModelVerified, error) {
	event := new(TrustmintNFTModelVerified)
	if err := _TrustmintNFT.contract.UnpackLog(event, "ModelVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TrustmintNFT contract.
type TrustmintNFTOwnershipTransferredIterator struct {
	Event *TrustmintNFTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTOwnershipTransferred)
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
		it.Event = new(TrustmintNFTOwnershipTransferred)
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
func (it *TrustmintNFTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTOwnershipTransferred represents a OwnershipTransferred event raised by the TrustmintNFT contract.
type TrustmintNFTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrustmintNFTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTOwnershipTransferredIterator{contract: _TrustmintNFT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrustmintNFTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTOwnershipTransferred)
				if err := _TrustmintNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TrustmintNFT *TrustmintNFTFilterer) ParseOwnershipTransferred(log types.Log) (*TrustmintNFTOwnershipTransferred, error) {
	event := new(TrustmintNFTOwnershipTransferred)
	if err := _TrustmintNFT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TrustmintNFTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TrustmintNFT contract.
type TrustmintNFTTransferIterator struct {
	Event *TrustmintNFTTransfer // Event containing the contract specifics and raw log

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
func (it *TrustmintNFTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrustmintNFTTransfer)
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
		it.Event = new(TrustmintNFTTransfer)
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
func (it *TrustmintNFTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrustmintNFTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrustmintNFTTransfer represents a Transfer event raised by the TrustmintNFT contract.
type TrustmintNFTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*TrustmintNFTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &TrustmintNFTTransferIterator{contract: _TrustmintNFT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TrustmintNFTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _TrustmintNFT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrustmintNFTTransfer)
				if err := _TrustmintNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_TrustmintNFT *TrustmintNFTFilterer) ParseTransfer(log types.Log) (*TrustmintNFTTransfer, error) {
	event := new(TrustmintNFTTransfer)
	if err := _TrustmintNFT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
