// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VMCABI is the input ABI used to generate the binding from.
const VMCABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_shardId\",\"type\":\"int256\"},{\"name\":\"_expectedPeriodNumber\",\"type\":\"uint256\"},{\"name\":\"_periodStartPrevHash\",\"type\":\"bytes32\"},{\"name\":\"_parentHash\",\"type\":\"bytes32\"},{\"name\":\"_transactionRoot\",\"type\":\"bytes32\"},{\"name\":\"_coinbase\",\"type\":\"address\"},{\"name\":\"_stateRoot\",\"type\":\"bytes32\"},{\"name\":\"_receiptRoot\",\"type\":\"bytes32\"},{\"name\":\"_number\",\"type\":\"int256\"}],\"name\":\"addHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"shardCount\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_receiptId\",\"type\":\"int256\"},{\"name\":\"_txGasprice\",\"type\":\"uint256\"}],\"name\":\"updateGasPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_shardId\",\"type\":\"int256\"},{\"name\":\"_txStartgas\",\"type\":\"uint256\"},{\"name\":\"_txGasprice\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes12\"}],\"name\":\"txToShard\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"periodHead\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"validators\",\"outputs\":[{\"name\":\"deposit\",\"type\":\"uint256\"},{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validatorIndex\",\"type\":\"int256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCollationGasLimit\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"collationHeaders\",\"outputs\":[{\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"name\":\"score\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isValidatorDeposited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"int256\"}],\"name\":\"receipts\",\"outputs\":[{\"name\":\"shardId\",\"type\":\"int256\"},{\"name\":\"txStartgas\",\"type\":\"uint256\"},{\"name\":\"txGasprice\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes32\"},{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_shardId\",\"type\":\"int256\"},{\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"getEligibleProposer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"shardId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"receiptId\",\"type\":\"int256\"}],\"name\":\"TxToShard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"shardId\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"expectedPeriodNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"periodStartPrevHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"parentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"transactionRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"stateRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"receiptRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"number\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"isNewHead\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"score\",\"type\":\"int256\"}],\"name\":\"CollationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"index\",\"type\":\"int256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// VMCBin is the compiled bytecode used for deploying new contracts.
const VMCBin = `0x6060604052341561000f57600080fd5b610bcd8061001e6000396000f3006060604052600436106100cf5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630341518d81146100d457806304e9c77a146101205780632213138914610145578063372a9e2a14610153578063584475db1461017d5780635d593f8d146101935780635ee9514c146101a65780637e62eab8146101dd578063934586ec146101f5578063b9d8ef9614610208578063c6cbe85c14610239578063d0e30db014610258578063d19ca56814610260578063e29de3ad146102bf575b600080fd5b34156100df57600080fd5b61010c600435602435604435606435608435600160a060020a0360a4351660c43560e435610104356102f4565b604051901515815260200160405180910390f35b341561012b57600080fd5b6101336105da565b60405190815260200160405180910390f35b61010c6004356024356105df565b610133600160a060020a0360043516602435604435606435600160a060020a031960843516610624565b341561018857600080fd5b610133600435610747565b341561019e57600080fd5b610133610759565b34156101b157600080fd5b6101bc60043561075f565b604051918252600160a060020a031660208201526040908101905180910390f35b34156101e857600080fd5b6101f3600435610781565b005b341561020057600080fd5b610133610884565b341561021357600080fd5b61022160043560243561088c565b60405191825260208201526040908101905180910390f35b341561024457600080fd5b61010c600160a060020a03600435166108ad565b6101336108c2565b341561026b57600080fd5b6102766004356109e7565b604051968752602087019590955260408087019490945260608601929092526080850152600160a060020a0390811660a085015290911660c083015260e0909101905180910390f35b34156102ca57600080fd5b6102d8600435602435610a31565b604051600160a060020a03909116815260200160405180910390f35b60006102fe610b7a565b60008b1215801561030f575060648b125b151561031a57600080fd5b600543101561032857600080fd5b600543048a1461033757600080fd5b60001960058b020140891461034b57600080fd5b8a8a8a8a8a600160a060020a038b168a8a8a60405198895260208901979097526040808901969096526060880194909452608087019290925260a086015260c085015260e08401526101008301919091526101209091019051908190039020815260008b81526001602052604081209082518152602081019190915260400160002060010154156103d857fe5b87156104005760008b81526001602081815260408084208c8552909152822001541361040057fe5b60008b8152600960205260409020548a901261041857fe5b6104258b60054304610a31565b600160a060020a03166040820190815251600160a060020a0316151561044a57600080fd5b8060400151600160a060020a031633600160a060020a031614151561046e57600080fd5b60008b81526001602081815260408084208c855282529092208101540190820190815251831461049d57600080fd5b60408051908101604052888152602080820190830151905260008c81526001602052604081209083518152602081019190915260400160002081518155602082015160019182015560008d81526009602090815260408083208f9055838252808320600383528184205484528252909120909101549150820151131561053657805160008c815260036020526040902055600160608201525b8a7f788a01fdbeb989d24e706cdd2993ca4213400e7b5fa631819b2acfe8ebad41358b8b8b8b8b8b8b8b8a606001518b60200151604051998a5260208a01989098526040808a01979097526060890195909552600160a060020a03909316608088015260a087019190915260c086015260e08501521515610100840152610120830191909152610140909101905180910390a25060019a9950505050505050505050565b606481565b60008281526002602052604081206005015433600160a060020a0390811691161461060957600080fd5b50600091825260026020819052604090922090910155600190565b60008060e060405190810160409081528782526020808301889052818301879052346060840152600160a060020a031986166080840152600160a060020a0333811660a08501528a1660c08401526005546000908152600290915220815181556020820151816001015560408201518160020155606082015181600301556080820151600482015560a0820151600582018054600160a060020a031916600160a060020a039290921691909117905560c08201516006919091018054600160a060020a031916600160a060020a039283161790556005805460018101909155925087915088167ffc322e0c42ee41e0d74b940ceeee9cd5971acdd6ace8ff8010ee7134c31d9ea58360405190815260200160405180910390a39695505050505050565b60096020526000908152604090205481565b60045481565b60006020819052908152604090208054600190910154600160a060020a031682565b60008181526020819052604090206001015433600160a060020a039081169116146107ab57600080fd5b6000818152602081905260409081902060018101549054600160a060020a039091169181156108fc02919051600060405180830381858888f1935050505015156107f457600080fd5b600081815260208181526040808320600181018054600160a060020a0316855260088452918420805460ff19169055848452918390529190558054600160a060020a031916905561084481610abf565b600480546000190190557fe13f360aa18d414ccdb598da6c447faa89d0477ffc7549dab5678fca76910b8c8160405190815260200160405180910390a150565b629896805b90565b60016020818152600093845260408085209091529183529120805491015482565b60086020526000908152604090205460ff1681565b600160a060020a033316600090815260086020526040812054819060ff16156108ea57600080fd5b3468056bc75e2d63100000146108ff57600080fd5b610907610ade565b151561091c57610915610ae5565b9050610921565b506004545b604080519081016040908152348252600160a060020a0333166020808401919091526000848152908190522081518155602082015160019182018054600160a060020a031916600160a060020a0392831617905560048054830190553390811660009081526008602052604090819020805460ff19169093179092557fd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478925090839051600160a060020a03909216825260208201526040908101905180910390a1919050565b60026020819052600091825260409091208054600182015492820154600383015460048401546005850154600690950154939594929391929091600160a060020a03918216911687565b60006004821015610a4157600080fd5b43600319830160050210610a5457600080fd5b60045460009013610a6457600080fd5b600080610a6f610b1c565b60031985016005024086604051918252602082015260409081019051908190039020811515610a9a57fe5b068152602081019190915260400160002060010154600160a060020a03169392505050565b6007805460009081526006602052604090209190915580546001019055565b6007541590565b6000610aef610ade565b15610afd5750600019610889565b5060078054600019019081905560009081526006602052604090205490565b600754600454600091829101815b610400811215610b6f57818112610b4057610b6f565b600081815260208190526040902060010154600160a060020a031615610b67576001830192505b600101610b2a565b505060075401919050565b608060405190810160409081526000808352602083018190529082018190526060820152905600a165627a7a72305820c2d8b6b6f94dc618137b8d964bb0bc718f94e39cccbd56de8f9bca19ec1fa0700029`

// DeployVMC deploys a new Ethereum contract, binding an instance of VMC to it.
func DeployVMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VMC, error) {
	parsed, err := abi.JSON(strings.NewReader(VMCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VMC{VMCCaller: VMCCaller{contract: contract}, VMCTransactor: VMCTransactor{contract: contract}, VMCFilterer: VMCFilterer{contract: contract}}, nil
}

// VMC is an auto generated Go binding around an Ethereum contract.
type VMC struct {
	VMCCaller     // Read-only binding to the contract
	VMCTransactor // Write-only binding to the contract
	VMCFilterer   // Log filterer for contract events
}

// VMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMCSession struct {
	Contract     *VMC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCCallerSession struct {
	Contract *VMCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMCTransactorSession struct {
	Contract     *VMCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMCRaw struct {
	Contract *VMC // Generic contract binding to access the raw methods on
}

// VMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCCallerRaw struct {
	Contract *VMCCaller // Generic read-only contract binding to access the raw methods on
}

// VMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMCTransactorRaw struct {
	Contract *VMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVMC creates a new instance of VMC, bound to a specific deployed contract.
func NewVMC(address common.Address, backend bind.ContractBackend) (*VMC, error) {
	contract, err := bindVMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VMC{VMCCaller: VMCCaller{contract: contract}, VMCTransactor: VMCTransactor{contract: contract}, VMCFilterer: VMCFilterer{contract: contract}}, nil
}

// NewVMCCaller creates a new read-only instance of VMC, bound to a specific deployed contract.
func NewVMCCaller(address common.Address, caller bind.ContractCaller) (*VMCCaller, error) {
	contract, err := bindVMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCCaller{contract: contract}, nil
}

// NewVMCTransactor creates a new write-only instance of VMC, bound to a specific deployed contract.
func NewVMCTransactor(address common.Address, transactor bind.ContractTransactor) (*VMCTransactor, error) {
	contract, err := bindVMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMCTransactor{contract: contract}, nil
}

// NewVMCFilterer creates a new log filterer instance of VMC, bound to a specific deployed contract.
func NewVMCFilterer(address common.Address, filterer bind.ContractFilterer) (*VMCFilterer, error) {
	contract, err := bindVMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMCFilterer{contract: contract}, nil
}

// bindVMC binds a generic wrapper to an already deployed contract.
func bindVMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMC *VMCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMC.Contract.VMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMC *VMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMC.Contract.VMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMC *VMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMC.Contract.VMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VMC *VMCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VMC *VMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VMC *VMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VMC.Contract.contract.Transact(opts, method, params...)
}

// CollationHeaders is a free data retrieval call binding the contract method 0xb9d8ef96.
//
// Solidity: function collationHeaders( int256,  bytes32) constant returns(parentHash bytes32, score int256)
func (_VMC *VMCCaller) CollationHeaders(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ParentHash [32]byte
	Score      *big.Int
}, error) {
	ret := new(struct {
		ParentHash [32]byte
		Score      *big.Int
	})
	out := ret
	err := _VMC.contract.Call(opts, out, "collationHeaders", arg0, arg1)
	return *ret, err
}

// CollationHeaders is a free data retrieval call binding the contract method 0xb9d8ef96.
//
// Solidity: function collationHeaders( int256,  bytes32) constant returns(parentHash bytes32, score int256)
func (_VMC *VMCSession) CollationHeaders(arg0 *big.Int, arg1 [32]byte) (struct {
	ParentHash [32]byte
	Score      *big.Int
}, error) {
	return _VMC.Contract.CollationHeaders(&_VMC.CallOpts, arg0, arg1)
}

// CollationHeaders is a free data retrieval call binding the contract method 0xb9d8ef96.
//
// Solidity: function collationHeaders( int256,  bytes32) constant returns(parentHash bytes32, score int256)
func (_VMC *VMCCallerSession) CollationHeaders(arg0 *big.Int, arg1 [32]byte) (struct {
	ParentHash [32]byte
	Score      *big.Int
}, error) {
	return _VMC.Contract.CollationHeaders(&_VMC.CallOpts, arg0, arg1)
}

// GetCollationGasLimit is a free data retrieval call binding the contract method 0x934586ec.
//
// Solidity: function getCollationGasLimit() constant returns(uint256)
func (_VMC *VMCCaller) GetCollationGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "getCollationGasLimit")
	return *ret0, err
}

// GetCollationGasLimit is a free data retrieval call binding the contract method 0x934586ec.
//
// Solidity: function getCollationGasLimit() constant returns(uint256)
func (_VMC *VMCSession) GetCollationGasLimit() (*big.Int, error) {
	return _VMC.Contract.GetCollationGasLimit(&_VMC.CallOpts)
}

// GetCollationGasLimit is a free data retrieval call binding the contract method 0x934586ec.
//
// Solidity: function getCollationGasLimit() constant returns(uint256)
func (_VMC *VMCCallerSession) GetCollationGasLimit() (*big.Int, error) {
	return _VMC.Contract.GetCollationGasLimit(&_VMC.CallOpts)
}

// GetEligibleProposer is a free data retrieval call binding the contract method 0xe29de3ad.
//
// Solidity: function getEligibleProposer(_shardId int256, _period uint256) constant returns(address)
func (_VMC *VMCCaller) GetEligibleProposer(opts *bind.CallOpts, _shardId *big.Int, _period *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "getEligibleProposer", _shardId, _period)
	return *ret0, err
}

// GetEligibleProposer is a free data retrieval call binding the contract method 0xe29de3ad.
//
// Solidity: function getEligibleProposer(_shardId int256, _period uint256) constant returns(address)
func (_VMC *VMCSession) GetEligibleProposer(_shardId *big.Int, _period *big.Int) (common.Address, error) {
	return _VMC.Contract.GetEligibleProposer(&_VMC.CallOpts, _shardId, _period)
}

// GetEligibleProposer is a free data retrieval call binding the contract method 0xe29de3ad.
//
// Solidity: function getEligibleProposer(_shardId int256, _period uint256) constant returns(address)
func (_VMC *VMCCallerSession) GetEligibleProposer(_shardId *big.Int, _period *big.Int) (common.Address, error) {
	return _VMC.Contract.GetEligibleProposer(&_VMC.CallOpts, _shardId, _period)
}

// IsValidatorDeposited is a free data retrieval call binding the contract method 0xc6cbe85c.
//
// Solidity: function isValidatorDeposited( address) constant returns(bool)
func (_VMC *VMCCaller) IsValidatorDeposited(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "isValidatorDeposited", arg0)
	return *ret0, err
}

// IsValidatorDeposited is a free data retrieval call binding the contract method 0xc6cbe85c.
//
// Solidity: function isValidatorDeposited( address) constant returns(bool)
func (_VMC *VMCSession) IsValidatorDeposited(arg0 common.Address) (bool, error) {
	return _VMC.Contract.IsValidatorDeposited(&_VMC.CallOpts, arg0)
}

// IsValidatorDeposited is a free data retrieval call binding the contract method 0xc6cbe85c.
//
// Solidity: function isValidatorDeposited( address) constant returns(bool)
func (_VMC *VMCCallerSession) IsValidatorDeposited(arg0 common.Address) (bool, error) {
	return _VMC.Contract.IsValidatorDeposited(&_VMC.CallOpts, arg0)
}

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(int256)
func (_VMC *VMCCaller) NumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "numValidators")
	return *ret0, err
}

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(int256)
func (_VMC *VMCSession) NumValidators() (*big.Int, error) {
	return _VMC.Contract.NumValidators(&_VMC.CallOpts)
}

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(int256)
func (_VMC *VMCCallerSession) NumValidators() (*big.Int, error) {
	return _VMC.Contract.NumValidators(&_VMC.CallOpts)
}

// PeriodHead is a free data retrieval call binding the contract method 0x584475db.
//
// Solidity: function periodHead( int256) constant returns(int256)
func (_VMC *VMCCaller) PeriodHead(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "periodHead", arg0)
	return *ret0, err
}

// PeriodHead is a free data retrieval call binding the contract method 0x584475db.
//
// Solidity: function periodHead( int256) constant returns(int256)
func (_VMC *VMCSession) PeriodHead(arg0 *big.Int) (*big.Int, error) {
	return _VMC.Contract.PeriodHead(&_VMC.CallOpts, arg0)
}

// PeriodHead is a free data retrieval call binding the contract method 0x584475db.
//
// Solidity: function periodHead( int256) constant returns(int256)
func (_VMC *VMCCallerSession) PeriodHead(arg0 *big.Int) (*big.Int, error) {
	return _VMC.Contract.PeriodHead(&_VMC.CallOpts, arg0)
}

// Receipts is a free data retrieval call binding the contract method 0xd19ca568.
//
// Solidity: function receipts( int256) constant returns(shardId int256, txStartgas uint256, txGasprice uint256, value uint256, data bytes32, sender address, to address)
func (_VMC *VMCCaller) Receipts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ShardId    *big.Int
	TxStartgas *big.Int
	TxGasprice *big.Int
	Value      *big.Int
	Data       [32]byte
	Sender     common.Address
	To         common.Address
}, error) {
	ret := new(struct {
		ShardId    *big.Int
		TxStartgas *big.Int
		TxGasprice *big.Int
		Value      *big.Int
		Data       [32]byte
		Sender     common.Address
		To         common.Address
	})
	out := ret
	err := _VMC.contract.Call(opts, out, "receipts", arg0)
	return *ret, err
}

// Receipts is a free data retrieval call binding the contract method 0xd19ca568.
//
// Solidity: function receipts( int256) constant returns(shardId int256, txStartgas uint256, txGasprice uint256, value uint256, data bytes32, sender address, to address)
func (_VMC *VMCSession) Receipts(arg0 *big.Int) (struct {
	ShardId    *big.Int
	TxStartgas *big.Int
	TxGasprice *big.Int
	Value      *big.Int
	Data       [32]byte
	Sender     common.Address
	To         common.Address
}, error) {
	return _VMC.Contract.Receipts(&_VMC.CallOpts, arg0)
}

// Receipts is a free data retrieval call binding the contract method 0xd19ca568.
//
// Solidity: function receipts( int256) constant returns(shardId int256, txStartgas uint256, txGasprice uint256, value uint256, data bytes32, sender address, to address)
func (_VMC *VMCCallerSession) Receipts(arg0 *big.Int) (struct {
	ShardId    *big.Int
	TxStartgas *big.Int
	TxGasprice *big.Int
	Value      *big.Int
	Data       [32]byte
	Sender     common.Address
	To         common.Address
}, error) {
	return _VMC.Contract.Receipts(&_VMC.CallOpts, arg0)
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(int256)
func (_VMC *VMCCaller) ShardCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VMC.contract.Call(opts, out, "shardCount")
	return *ret0, err
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(int256)
func (_VMC *VMCSession) ShardCount() (*big.Int, error) {
	return _VMC.Contract.ShardCount(&_VMC.CallOpts)
}

// ShardCount is a free data retrieval call binding the contract method 0x04e9c77a.
//
// Solidity: function shardCount() constant returns(int256)
func (_VMC *VMCCallerSession) ShardCount() (*big.Int, error) {
	return _VMC.Contract.ShardCount(&_VMC.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x5ee9514c.
//
// Solidity: function validators( int256) constant returns(deposit uint256, addr address)
func (_VMC *VMCCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Deposit *big.Int
	Addr    common.Address
}, error) {
	ret := new(struct {
		Deposit *big.Int
		Addr    common.Address
	})
	out := ret
	err := _VMC.contract.Call(opts, out, "validators", arg0)
	return *ret, err
}

// Validators is a free data retrieval call binding the contract method 0x5ee9514c.
//
// Solidity: function validators( int256) constant returns(deposit uint256, addr address)
func (_VMC *VMCSession) Validators(arg0 *big.Int) (struct {
	Deposit *big.Int
	Addr    common.Address
}, error) {
	return _VMC.Contract.Validators(&_VMC.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x5ee9514c.
//
// Solidity: function validators( int256) constant returns(deposit uint256, addr address)
func (_VMC *VMCCallerSession) Validators(arg0 *big.Int) (struct {
	Deposit *big.Int
	Addr    common.Address
}, error) {
	return _VMC.Contract.Validators(&_VMC.CallOpts, arg0)
}

// AddHeader is a paid mutator transaction binding the contract method 0x0341518d.
//
// Solidity: function addHeader(_shardId int256, _expectedPeriodNumber uint256, _periodStartPrevHash bytes32, _parentHash bytes32, _transactionRoot bytes32, _coinbase address, _stateRoot bytes32, _receiptRoot bytes32, _number int256) returns(bool)
func (_VMC *VMCTransactor) AddHeader(opts *bind.TransactOpts, _shardId *big.Int, _expectedPeriodNumber *big.Int, _periodStartPrevHash [32]byte, _parentHash [32]byte, _transactionRoot [32]byte, _coinbase common.Address, _stateRoot [32]byte, _receiptRoot [32]byte, _number *big.Int) (*types.Transaction, error) {
	return _VMC.contract.Transact(opts, "addHeader", _shardId, _expectedPeriodNumber, _periodStartPrevHash, _parentHash, _transactionRoot, _coinbase, _stateRoot, _receiptRoot, _number)
}

// AddHeader is a paid mutator transaction binding the contract method 0x0341518d.
//
// Solidity: function addHeader(_shardId int256, _expectedPeriodNumber uint256, _periodStartPrevHash bytes32, _parentHash bytes32, _transactionRoot bytes32, _coinbase address, _stateRoot bytes32, _receiptRoot bytes32, _number int256) returns(bool)
func (_VMC *VMCSession) AddHeader(_shardId *big.Int, _expectedPeriodNumber *big.Int, _periodStartPrevHash [32]byte, _parentHash [32]byte, _transactionRoot [32]byte, _coinbase common.Address, _stateRoot [32]byte, _receiptRoot [32]byte, _number *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.AddHeader(&_VMC.TransactOpts, _shardId, _expectedPeriodNumber, _periodStartPrevHash, _parentHash, _transactionRoot, _coinbase, _stateRoot, _receiptRoot, _number)
}

// AddHeader is a paid mutator transaction binding the contract method 0x0341518d.
//
// Solidity: function addHeader(_shardId int256, _expectedPeriodNumber uint256, _periodStartPrevHash bytes32, _parentHash bytes32, _transactionRoot bytes32, _coinbase address, _stateRoot bytes32, _receiptRoot bytes32, _number int256) returns(bool)
func (_VMC *VMCTransactorSession) AddHeader(_shardId *big.Int, _expectedPeriodNumber *big.Int, _periodStartPrevHash [32]byte, _parentHash [32]byte, _transactionRoot [32]byte, _coinbase common.Address, _stateRoot [32]byte, _receiptRoot [32]byte, _number *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.AddHeader(&_VMC.TransactOpts, _shardId, _expectedPeriodNumber, _periodStartPrevHash, _parentHash, _transactionRoot, _coinbase, _stateRoot, _receiptRoot, _number)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(int256)
func (_VMC *VMCTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VMC.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(int256)
func (_VMC *VMCSession) Deposit() (*types.Transaction, error) {
	return _VMC.Contract.Deposit(&_VMC.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(int256)
func (_VMC *VMCTransactorSession) Deposit() (*types.Transaction, error) {
	return _VMC.Contract.Deposit(&_VMC.TransactOpts)
}

// TxToShard is a paid mutator transaction binding the contract method 0x372a9e2a.
//
// Solidity: function txToShard(_to address, _shardId int256, _txStartgas uint256, _txGasprice uint256, _data bytes12) returns(int256)
func (_VMC *VMCTransactor) TxToShard(opts *bind.TransactOpts, _to common.Address, _shardId *big.Int, _txStartgas *big.Int, _txGasprice *big.Int, _data [12]byte) (*types.Transaction, error) {
	return _VMC.contract.Transact(opts, "txToShard", _to, _shardId, _txStartgas, _txGasprice, _data)
}

// TxToShard is a paid mutator transaction binding the contract method 0x372a9e2a.
//
// Solidity: function txToShard(_to address, _shardId int256, _txStartgas uint256, _txGasprice uint256, _data bytes12) returns(int256)
func (_VMC *VMCSession) TxToShard(_to common.Address, _shardId *big.Int, _txStartgas *big.Int, _txGasprice *big.Int, _data [12]byte) (*types.Transaction, error) {
	return _VMC.Contract.TxToShard(&_VMC.TransactOpts, _to, _shardId, _txStartgas, _txGasprice, _data)
}

// TxToShard is a paid mutator transaction binding the contract method 0x372a9e2a.
//
// Solidity: function txToShard(_to address, _shardId int256, _txStartgas uint256, _txGasprice uint256, _data bytes12) returns(int256)
func (_VMC *VMCTransactorSession) TxToShard(_to common.Address, _shardId *big.Int, _txStartgas *big.Int, _txGasprice *big.Int, _data [12]byte) (*types.Transaction, error) {
	return _VMC.Contract.TxToShard(&_VMC.TransactOpts, _to, _shardId, _txStartgas, _txGasprice, _data)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x22131389.
//
// Solidity: function updateGasPrice(_receiptId int256, _txGasprice uint256) returns(bool)
func (_VMC *VMCTransactor) UpdateGasPrice(opts *bind.TransactOpts, _receiptId *big.Int, _txGasprice *big.Int) (*types.Transaction, error) {
	return _VMC.contract.Transact(opts, "updateGasPrice", _receiptId, _txGasprice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x22131389.
//
// Solidity: function updateGasPrice(_receiptId int256, _txGasprice uint256) returns(bool)
func (_VMC *VMCSession) UpdateGasPrice(_receiptId *big.Int, _txGasprice *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.UpdateGasPrice(&_VMC.TransactOpts, _receiptId, _txGasprice)
}

// UpdateGasPrice is a paid mutator transaction binding the contract method 0x22131389.
//
// Solidity: function updateGasPrice(_receiptId int256, _txGasprice uint256) returns(bool)
func (_VMC *VMCTransactorSession) UpdateGasPrice(_receiptId *big.Int, _txGasprice *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.UpdateGasPrice(&_VMC.TransactOpts, _receiptId, _txGasprice)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(_validatorIndex int256) returns()
func (_VMC *VMCTransactor) Withdraw(opts *bind.TransactOpts, _validatorIndex *big.Int) (*types.Transaction, error) {
	return _VMC.contract.Transact(opts, "withdraw", _validatorIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(_validatorIndex int256) returns()
func (_VMC *VMCSession) Withdraw(_validatorIndex *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.Withdraw(&_VMC.TransactOpts, _validatorIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x7e62eab8.
//
// Solidity: function withdraw(_validatorIndex int256) returns()
func (_VMC *VMCTransactorSession) Withdraw(_validatorIndex *big.Int) (*types.Transaction, error) {
	return _VMC.Contract.Withdraw(&_VMC.TransactOpts, _validatorIndex)
}

// VMCCollationAddedIterator is returned from FilterCollationAdded and is used to iterate over the raw logs and unpacked data for CollationAdded events raised by the VMC contract.
type VMCCollationAddedIterator struct {
	Event *VMCCollationAdded // Event containing the contract specifics and raw log

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
func (it *VMCCollationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMCCollationAdded)
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
		it.Event = new(VMCCollationAdded)
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
func (it *VMCCollationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMCCollationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMCCollationAdded represents a CollationAdded event raised by the VMC contract.
type VMCCollationAdded struct {
	ShardId              *big.Int
	ExpectedPeriodNumber *big.Int
	PeriodStartPrevHash  [32]byte
	ParentHash           [32]byte
	TransactionRoot      [32]byte
	Coinbase             common.Address
	StateRoot            [32]byte
	ReceiptRoot          [32]byte
	Number               *big.Int
	IsNewHead            bool
	Score                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCollationAdded is a free log retrieval operation binding the contract event 0x788a01fdbeb989d24e706cdd2993ca4213400e7b5fa631819b2acfe8ebad4135.
//
// Solidity: event CollationAdded(shardId indexed int256, expectedPeriodNumber uint256, periodStartPrevHash bytes32, parentHash bytes32, transactionRoot bytes32, coinbase address, stateRoot bytes32, receiptRoot bytes32, number int256, isNewHead bool, score int256)
func (_VMC *VMCFilterer) FilterCollationAdded(opts *bind.FilterOpts, shardId []*big.Int) (*VMCCollationAddedIterator, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _VMC.contract.FilterLogs(opts, "CollationAdded", shardIdRule)
	if err != nil {
		return nil, err
	}
	return &VMCCollationAddedIterator{contract: _VMC.contract, event: "CollationAdded", logs: logs, sub: sub}, nil
}

// WatchCollationAdded is a free log subscription operation binding the contract event 0x788a01fdbeb989d24e706cdd2993ca4213400e7b5fa631819b2acfe8ebad4135.
//
// Solidity: event CollationAdded(shardId indexed int256, expectedPeriodNumber uint256, periodStartPrevHash bytes32, parentHash bytes32, transactionRoot bytes32, coinbase address, stateRoot bytes32, receiptRoot bytes32, number int256, isNewHead bool, score int256)
func (_VMC *VMCFilterer) WatchCollationAdded(opts *bind.WatchOpts, sink chan<- *VMCCollationAdded, shardId []*big.Int) (event.Subscription, error) {

	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _VMC.contract.WatchLogs(opts, "CollationAdded", shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMCCollationAdded)
				if err := _VMC.contract.UnpackLog(event, "CollationAdded", log); err != nil {
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

// VMCDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the VMC contract.
type VMCDepositIterator struct {
	Event *VMCDeposit // Event containing the contract specifics and raw log

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
func (it *VMCDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMCDeposit)
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
		it.Event = new(VMCDeposit)
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
func (it *VMCDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMCDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMCDeposit represents a Deposit event raised by the VMC contract.
type VMCDeposit struct {
	Validator common.Address
	Index     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(validator address, index int256)
func (_VMC *VMCFilterer) FilterDeposit(opts *bind.FilterOpts) (*VMCDepositIterator, error) {

	logs, sub, err := _VMC.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VMCDepositIterator{contract: _VMC.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd8a6d38df847dcba70dfdeb4948fb1457d61a81d132801f40dc9c00d52dfd478.
//
// Solidity: event Deposit(validator address, index int256)
func (_VMC *VMCFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VMCDeposit) (event.Subscription, error) {

	logs, sub, err := _VMC.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMCDeposit)
				if err := _VMC.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// VMCTxToShardIterator is returned from FilterTxToShard and is used to iterate over the raw logs and unpacked data for TxToShard events raised by the VMC contract.
type VMCTxToShardIterator struct {
	Event *VMCTxToShard // Event containing the contract specifics and raw log

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
func (it *VMCTxToShardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMCTxToShard)
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
		it.Event = new(VMCTxToShard)
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
func (it *VMCTxToShardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMCTxToShardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMCTxToShard represents a TxToShard event raised by the VMC contract.
type VMCTxToShard struct {
	To        common.Address
	ShardId   *big.Int
	ReceiptId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTxToShard is a free log retrieval operation binding the contract event 0xfc322e0c42ee41e0d74b940ceeee9cd5971acdd6ace8ff8010ee7134c31d9ea5.
//
// Solidity: event TxToShard(to indexed address, shardId indexed int256, receiptId int256)
func (_VMC *VMCFilterer) FilterTxToShard(opts *bind.FilterOpts, to []common.Address, shardId []*big.Int) (*VMCTxToShardIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _VMC.contract.FilterLogs(opts, "TxToShard", toRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &VMCTxToShardIterator{contract: _VMC.contract, event: "TxToShard", logs: logs, sub: sub}, nil
}

// WatchTxToShard is a free log subscription operation binding the contract event 0xfc322e0c42ee41e0d74b940ceeee9cd5971acdd6ace8ff8010ee7134c31d9ea5.
//
// Solidity: event TxToShard(to indexed address, shardId indexed int256, receiptId int256)
func (_VMC *VMCFilterer) WatchTxToShard(opts *bind.WatchOpts, sink chan<- *VMCTxToShard, to []common.Address, shardId []*big.Int) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _VMC.contract.WatchLogs(opts, "TxToShard", toRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMCTxToShard)
				if err := _VMC.contract.UnpackLog(event, "TxToShard", log); err != nil {
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

// VMCWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the VMC contract.
type VMCWithdrawIterator struct {
	Event *VMCWithdraw // Event containing the contract specifics and raw log

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
func (it *VMCWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VMCWithdraw)
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
		it.Event = new(VMCWithdraw)
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
func (it *VMCWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VMCWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VMCWithdraw represents a Withdraw event raised by the VMC contract.
type VMCWithdraw struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe13f360aa18d414ccdb598da6c447faa89d0477ffc7549dab5678fca76910b8c.
//
// Solidity: event Withdraw(index int256)
func (_VMC *VMCFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VMCWithdrawIterator, error) {

	logs, sub, err := _VMC.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VMCWithdrawIterator{contract: _VMC.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe13f360aa18d414ccdb598da6c447faa89d0477ffc7549dab5678fca76910b8c.
//
// Solidity: event Withdraw(index int256)
func (_VMC *VMCFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VMCWithdraw) (event.Subscription, error) {

	logs, sub, err := _VMC.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VMCWithdraw)
				if err := _VMC.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
