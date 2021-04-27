// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package asset

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AssetABI is the input ABI used to generate the binding from.
const AssetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Address\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ProofOfTitle\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_getaddress\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ownership\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"OwnerSocialMedia\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_initalSupply\",\"type\":\"uint64\"},{\"name\":\"_initialName\",\"type\":\"string\"},{\"name\":\"_initialSymbol\",\"type\":\"string\"},{\"name\":\"_initialOwnership\",\"type\":\"string\"},{\"name\":\"_initialAddress\",\"type\":\"string\"},{\"name\":\"_initialOwnerSocialMedia\",\"type\":\"string\"},{\"name\":\"_initialProofOfTitle\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approve\",\"type\":\"event\"}]"

// AssetBin is the compiled bytecode used for deploying new contracts.
const AssetBin = `60806040523480156200001157600080fd5b50604051620010c5380380620010c5833981018060405260e08110156200003757600080fd5b810190808051906020019092919080516401000000008111156200005a57600080fd5b828101905060208101848111156200007157600080fd5b81518560018202830111640100000000821117156200008f57600080fd5b50509291906020018051640100000000811115620000ac57600080fd5b82810190506020810184811115620000c357600080fd5b8151856001820283011164010000000082111715620000e157600080fd5b50509291906020018051640100000000811115620000fe57600080fd5b828101905060208101848111156200011557600080fd5b81518560018202830111640100000000821117156200013357600080fd5b505092919060200180516401000000008111156200015057600080fd5b828101905060208101848111156200016757600080fd5b81518560018202830111640100000000821117156200018557600080fd5b50509291906020018051640100000000811115620001a257600080fd5b82810190506020810184811115620001b957600080fd5b8151856001820283011164010000000082111715620001d757600080fd5b50509291906020018051640100000000811115620001f457600080fd5b828101905060208101848111156200020b57600080fd5b81518560018202830111640100000000821117156200022957600080fd5b50509291905050508667ffffffffffffffff16600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550866000806101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055508560019080519060200190620002bf9291906200034a565b508460029080519060200190620002d89291906200034a565b508360039080519060200190620002f19291906200034a565b5082600490805190602001906200030a9291906200034a565b508160059080519060200190620003239291906200034a565b5080600690805190602001906200033c9291906200034a565b5050505050505050620003f9565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200038d57805160ff1916838001178555620003be565b82800160010185558215620003be579182015b82811115620003bd578251825591602001919060010190620003a0565b5b509050620003cd9190620003d1565b5090565b620003f691905b80821115620003f2576000816000905550600101620003d8565b5090565b90565b610cbc80620004096000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063a9059cbb11610071578063a9059cbb14610332578063dd62ed3e14610398578063ea5f338314610410578063f8b2cb4f14610493578063fbdd7852146104eb578063ffac9bd11461056e576100b4565b8063095ea7b3146100b95780633045aaf31461011f5780633aeaccf5146101a257806370a08231146102255780638052474d1461027d578063a44b47f714610300575b600080fd5b610105600480360360408110156100cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506105f1565b604051808215151515815260200191505060405180910390f35b6101276106e3565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561016757808201518184015260208101905061014c565b50505050905090810190601f1680156101945780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101aa610781565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101ea5780820151818401526020810190506101cf565b50505050905090810190601f1680156102175780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102676004803603602081101561023b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061081f565b6040518082815260200191505060405180910390f35b610285610837565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102c55780820151818401526020810190506102aa565b50505050905090810190601f1680156102f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103086108d5565b604051808267ffffffffffffffff1667ffffffffffffffff16815260200191505060405180910390f35b61037e6004803603604081101561034857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108ee565b604051808215151515815260200191505060405180910390f35b6103fa600480360360408110156103ae57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a45565b6040518082815260200191505060405180910390f35b610418610a6a565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561045857808201518184015260208101905061043d565b50505050905090810190601f1680156104855780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6104d5600480360360208110156104a957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610b08565b6040518082815260200191505060405180910390f35b6104f3610b54565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610533578082015181840152602081019050610518565b50505050905090810190601f1680156105605780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610576610bf2565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105b657808201518184015260208101905061059b565b50505050905090810190601f1680156105e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600081600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e846040518082815260200191505060405180910390a36001905092915050565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107795780601f1061074e57610100808354040283529160200191610779565b820191906000526020600020905b81548152906001019060200180831161075c57829003601f168201915b505050505081565b60048054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108175780601f106107ec57610100808354040283529160200191610817565b820191906000526020600020905b8154815290600101906020018083116107fa57829003601f168201915b505050505081565b60076020528060005260406000206000915090505481565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108cd5780601f106108a2576101008083540402835291602001916108cd565b820191906000526020600020905b8154815290600101906020018083116108b057829003601f168201915b505050505081565b6000809054906101000a900467ffffffffffffffff1681565b600081600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054101561093c57600080fd5b81600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254039250508190555081600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825401925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6008602052816000526040600020602052806000526040600020600091509150505481565b60068054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b005780601f10610ad557610100808354040283529160200191610b00565b820191906000526020600020905b815481529060010190602001808311610ae357829003601f168201915b505050505081565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050809050919050565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bea5780601f10610bbf57610100808354040283529160200191610bea565b820191906000526020600020905b815481529060010190602001808311610bcd57829003601f168201915b505050505081565b60058054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c885780601f10610c5d57610100808354040283529160200191610c88565b820191906000526020600020905b815481529060010190602001808311610c6b57829003601f168201915b50505050508156fea165627a7a72305820b4c5085943663fbd9b4b9e9b3504c000cde299ca3b214bf5f43d2fac3573974c0029`

// DeployAsset deploys a new Ethereum contract, binding an instance of Asset to it.
func DeployAsset(auth *bind.TransactOpts, backend bind.ContractBackend, _initalSupply uint64, _initialName string, _initialSymbol string, _initialOwnership string, _initialAddress string, _initialOwnerSocialMedia string, _initialProofOfTitle string) (common.Address, *types.Transaction, *Asset, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetBin), backend, _initalSupply, _initialName, _initialSymbol, _initialOwnership, _initialAddress, _initialOwnerSocialMedia, _initialProofOfTitle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// Asset is an auto generated Go binding around an Ethereum contract.
type Asset struct {
	AssetCaller     // Read-only binding to the contract
	AssetTransactor // Write-only binding to the contract
	AssetFilterer   // Log filterer for contract events
}

// AssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetSession struct {
	Contract     *Asset            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetCallerSession struct {
	Contract *AssetCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetTransactorSession struct {
	Contract     *AssetTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetRaw struct {
	Contract *Asset // Generic contract binding to access the raw methods on
}

// AssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetCallerRaw struct {
	Contract *AssetCaller // Generic read-only contract binding to access the raw methods on
}

// AssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetTransactorRaw struct {
	Contract *AssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsset creates a new instance of Asset, bound to a specific deployed contract.
func NewAsset(address common.Address, backend bind.ContractBackend) (*Asset, error) {
	contract, err := bindAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Asset{AssetCaller: AssetCaller{contract: contract}, AssetTransactor: AssetTransactor{contract: contract}, AssetFilterer: AssetFilterer{contract: contract}}, nil
}

// NewAssetCaller creates a new read-only instance of Asset, bound to a specific deployed contract.
func NewAssetCaller(address common.Address, caller bind.ContractCaller) (*AssetCaller, error) {
	contract, err := bindAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AssetCaller{contract: contract}, nil
}

// NewAssetTransactor creates a new write-only instance of Asset, bound to a specific deployed contract.
func NewAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetTransactor, error) {
	contract, err := bindAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AssetTransactor{contract: contract}, nil
}

// NewAssetFilterer creates a new log filterer instance of Asset, bound to a specific deployed contract.
func NewAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*AssetFilterer, error) {
	contract, err := bindAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AssetFilterer{contract: contract}, nil
}

// bindAsset binds a generic wrapper to an already deployed contract.
func bindAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.AssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.AssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Asset *AssetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Asset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Asset *AssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Asset *AssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Asset.Contract.contract.Transact(opts, method, params...)
}

// Address is a free data retrieval call binding the contract method 0x3aeaccf5.
//
// Solidity: function Address() constant returns(string)
func (_Asset *AssetCaller) Address(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "Address")
	return *ret0, err
}

// Address is a free data retrieval call binding the contract method 0x3aeaccf5.
//
// Solidity: function Address() constant returns(string)
func (_Asset *AssetSession) Address() (string, error) {
	return _Asset.Contract.Address(&_Asset.CallOpts)
}

// Address is a free data retrieval call binding the contract method 0x3aeaccf5.
//
// Solidity: function Address() constant returns(string)
func (_Asset *AssetCallerSession) Address() (string, error) {
	return _Asset.Contract.Address(&_Asset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() constant returns(string)
func (_Asset *AssetCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "Name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() constant returns(string)
func (_Asset *AssetSession) Name() (string, error) {
	return _Asset.Contract.Name(&_Asset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() constant returns(string)
func (_Asset *AssetCallerSession) Name() (string, error) {
	return _Asset.Contract.Name(&_Asset.CallOpts)
}

// OwnerSocialMedia is a free data retrieval call binding the contract method 0xffac9bd1.
//
// Solidity: function OwnerSocialMedia() constant returns(string)
func (_Asset *AssetCaller) OwnerSocialMedia(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "OwnerSocialMedia")
	return *ret0, err
}

// OwnerSocialMedia is a free data retrieval call binding the contract method 0xffac9bd1.
//
// Solidity: function OwnerSocialMedia() constant returns(string)
func (_Asset *AssetSession) OwnerSocialMedia() (string, error) {
	return _Asset.Contract.OwnerSocialMedia(&_Asset.CallOpts)
}

// OwnerSocialMedia is a free data retrieval call binding the contract method 0xffac9bd1.
//
// Solidity: function OwnerSocialMedia() constant returns(string)
func (_Asset *AssetCallerSession) OwnerSocialMedia() (string, error) {
	return _Asset.Contract.OwnerSocialMedia(&_Asset.CallOpts)
}

// Ownership is a free data retrieval call binding the contract method 0xfbdd7852.
//
// Solidity: function Ownership() constant returns(string)
func (_Asset *AssetCaller) Ownership(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "Ownership")
	return *ret0, err
}

// Ownership is a free data retrieval call binding the contract method 0xfbdd7852.
//
// Solidity: function Ownership() constant returns(string)
func (_Asset *AssetSession) Ownership() (string, error) {
	return _Asset.Contract.Ownership(&_Asset.CallOpts)
}

// Ownership is a free data retrieval call binding the contract method 0xfbdd7852.
//
// Solidity: function Ownership() constant returns(string)
func (_Asset *AssetCallerSession) Ownership() (string, error) {
	return _Asset.Contract.Ownership(&_Asset.CallOpts)
}

// ProofOfTitle is a free data retrieval call binding the contract method 0xea5f3383.
//
// Solidity: function ProofOfTitle() constant returns(string)
func (_Asset *AssetCaller) ProofOfTitle(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "ProofOfTitle")
	return *ret0, err
}

// ProofOfTitle is a free data retrieval call binding the contract method 0xea5f3383.
//
// Solidity: function ProofOfTitle() constant returns(string)
func (_Asset *AssetSession) ProofOfTitle() (string, error) {
	return _Asset.Contract.ProofOfTitle(&_Asset.CallOpts)
}

// ProofOfTitle is a free data retrieval call binding the contract method 0xea5f3383.
//
// Solidity: function ProofOfTitle() constant returns(string)
func (_Asset *AssetCallerSession) ProofOfTitle() (string, error) {
	return _Asset.Contract.ProofOfTitle(&_Asset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x3045aaf3.
//
// Solidity: function Symbol() constant returns(string)
func (_Asset *AssetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "Symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x3045aaf3.
//
// Solidity: function Symbol() constant returns(string)
func (_Asset *AssetSession) Symbol() (string, error) {
	return _Asset.Contract.Symbol(&_Asset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x3045aaf3.
//
// Solidity: function Symbol() constant returns(string)
func (_Asset *AssetCallerSession) Symbol() (string, error) {
	return _Asset.Contract.Symbol(&_Asset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xa44b47f7.
//
// Solidity: function TotalSupply() constant returns(uint64)
func (_Asset *AssetCaller) TotalSupply(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "TotalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0xa44b47f7.
//
// Solidity: function TotalSupply() constant returns(uint64)
func (_Asset *AssetSession) TotalSupply() (uint64, error) {
	return _Asset.Contract.TotalSupply(&_Asset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xa44b47f7.
//
// Solidity: function TotalSupply() constant returns(uint64)
func (_Asset *AssetCallerSession) TotalSupply() (uint64, error) {
	return _Asset.Contract.TotalSupply(&_Asset.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Asset *AssetCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Asset *AssetSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Asset.Contract.Allowance(&_Asset.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) constant returns(uint256)
func (_Asset *AssetCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Asset.Contract.Allowance(&_Asset.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Asset *AssetCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Asset *AssetSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Asset.Contract.BalanceOf(&_Asset.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) constant returns(uint256)
func (_Asset *AssetCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Asset.Contract.BalanceOf(&_Asset.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _getaddress) constant returns(uint256 balance)
func (_Asset *AssetCaller) GetBalance(opts *bind.CallOpts, _getaddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Asset.contract.Call(opts, out, "getBalance", _getaddress)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _getaddress) constant returns(uint256 balance)
func (_Asset *AssetSession) GetBalance(_getaddress common.Address) (*big.Int, error) {
	return _Asset.Contract.GetBalance(&_Asset.CallOpts, _getaddress)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _getaddress) constant returns(uint256 balance)
func (_Asset *AssetCallerSession) GetBalance(_getaddress common.Address) (*big.Int, error) {
	return _Asset.Contract.GetBalance(&_Asset.CallOpts, _getaddress)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Asset *AssetTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Asset *AssetSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Approve(&_Asset.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Asset *AssetTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Approve(&_Asset.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Asset *AssetTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Asset *AssetSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_Asset *AssetTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Asset.Contract.Transfer(&_Asset.TransactOpts, _to, _value)
}

// AssetApproveIterator is returned from FilterApprove and is used to iterate over the raw logs and unpacked data for Approve events raised by the Asset contract.
type AssetApproveIterator struct {
	Event *AssetApprove // Event containing the contract specifics and raw log

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
func (it *AssetApproveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetApprove)
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
		it.Event = new(AssetApprove)
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
func (it *AssetApproveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetApproveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetApprove represents a Approve event raised by the Asset contract.
type AssetApprove struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApprove is a free log retrieval operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed _owner, address indexed _spender, uint256 _value)
func (_Asset *AssetFilterer) FilterApprove(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*AssetApproveIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "Approve", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &AssetApproveIterator{contract: _Asset.contract, event: "Approve", logs: logs, sub: sub}, nil
}

// WatchApprove is a free log subscription operation binding the contract event 0x6e11fb1b7f119e3f2fa29896ef5fdf8b8a2d0d4df6fe90ba8668e7d8b2ffa25e.
//
// Solidity: event Approve(address indexed _owner, address indexed _spender, uint256 _value)
func (_Asset *AssetFilterer) WatchApprove(opts *bind.WatchOpts, sink chan<- *AssetApprove, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "Approve", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetApprove)
				if err := _Asset.contract.UnpackLog(event, "Approve", log); err != nil {
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

// AssetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Asset contract.
type AssetTransferIterator struct {
	Event *AssetTransfer // Event containing the contract specifics and raw log

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
func (it *AssetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AssetTransfer)
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
		it.Event = new(AssetTransfer)
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
func (it *AssetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AssetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AssetTransfer represents a Transfer event raised by the Asset contract.
type AssetTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Asset *AssetFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AssetTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Asset.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AssetTransferIterator{contract: _Asset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Asset *AssetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AssetTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Asset.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AssetTransfer)
				if err := _Asset.contract.UnpackLog(event, "Transfer", log); err != nil {
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
