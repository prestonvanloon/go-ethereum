pragma solidity ^0.4.19;

contract VMC {
<<<<<<< HEAD
  event TxToShard(address indexed to, int indexed shardId, int receiptId);
  event CollationAdded(int indexed shardId, uint expectedPeriodNumber, 
                     bytes32 periodStartPrevHash, bytes32 parentHash,
                     bytes32 transactionRoot, address coinbase,
                     bytes32 stateRoot, bytes32 receiptRoot,
                     int number, bool isNewHead, int score);
  event Deposit(address validator, int index);
  event Withdraw(int validatorIndex);
=======
  using RLP for RLP.RLPItem;
  using RLP for RLP.Iterator;
  using RLP for bytes;
>>>>>>> e362985d3... loops over shards and checks if eligible proposer

  struct Validator {
    // Amount of wei the validator holds
    uint deposit;
    // The validator's address
    address addr;
  }

  struct CollationHeader {
    bytes32 parentHash;
    int score;
  }

  struct Receipt {
    int shardId;
    uint txStartgas;
    uint txGasprice;
    uint value;
    bytes32 data;
    address sender;
    address to;
  }

  // Packed variables to be used in addHeader
  struct HeaderVars {
    bytes32 entireHeaderHash;
    int score;
    address validatorAddr;
    bool isNewHead;
  }

  // validatorId => Validators
  mapping (int => Validator) validators;
  // shardId => (headerHash => CollationHeader)
  mapping (int => mapping (bytes32 => CollationHeader)) collationHeaders;
  // receiptId => Receipt
  mapping (int => Receipt) receipts;
  // shardId => headerHash
  mapping (int => bytes32) shardHead;
  
  // Number of validators
  int numValidators;
  // Number of receipts
  int numReceipts;
  // Indexs of empty slots caused by the function `withdraw`
  mapping (int => int) emptySlotsStack;
  // The top index of the stack in empty_slots_stack
  int emptySlotsStackTop;
  // Has the validator deposited before?
  mapping (address => bool) isValidatorDeposited;

  // Constant values
  uint constant periodLength = 5;
  int constant public shardCount = 100;
  // The exact deposit size which you have to deposit to become a validator
<<<<<<< HEAD
  uint constant depositSize = 100 ether;
  // Number of periods ahead of current period, which the contract
  // is able to return the collator of that period
  uint constant lookAheadPeriods = 4;

=======
  uint depositSize;
  // Any given validator randomly gets allocated to some number of shards every SHUFFLING_CYCLE
  int shufflingCycleLength;
  // Gas limit of the signature validation code
  uint sigGasLimit;
  // Is a valcode addr deposited now?
  mapping (address => bool) isValcodeDeposited;
  uint periodLength;
  int numValidatorsPerCycle;
  int public shardCount;
  bytes32 addHeaderLogTopic;
  SigHasherContract sighasher;
>>>>>>> e362985d3... loops over shards and checks if eligible proposer
  // Log the latest period number of the shard
  mapping (int => int) periodHead;

  function VMC() public {
  }

  // Returns the gas limit that collations can currently have (by default make
  // this function always answer 10 million).
  function getCollationGasLimit() public pure returns(uint) {
    return 10000000;
  }

  // Uses a block hash as a seed to pseudorandomly select a signer from the validator set.
  // [TODO] Chance of being selected should be proportional to the validator's deposit.
  // Should be able to return a value for the current period or any future period up to.
  function getEligibleProposer(int _shardId, uint _period) public view returns(address) {
    require(_period >= lookAheadPeriods);
    require((_period - lookAheadPeriods) * periodLength < block.number);
    require(numValidators > 0);
    // [TODO] Should check further if this safe or not
    return validators[
      int(
        uint(
          keccak256(
            uint(
              block.blockhash(_period - lookAheadPeriods)
            ) * periodLength,
            _shardId
          )
        ) %
        uint(getValidatorsMaxIndex())
      )
    ].addr;
  }

  function deposit() public payable returns(int) {
    require(!isValidatorDeposited[msg.sender]);
    require(msg.value == depositSize);
    // Find the empty slot index in validators set
    int index;
    if (!isStackEmpty())
      index = stackPop();
    else
      index = int(numValidators);
      
    validators[index] = Validator({
      deposit: msg.value,
      addr: msg.sender
    });
    ++numValidators;
<<<<<<< HEAD
    isValidatorDeposited[msg.sender] = true;
    
    Deposit(msg.sender, index);
=======
    isValcodeDeposited[msg.sender] = true;

    log2(keccak256("deposit()"), bytes32(msg.sender), bytes32(index));
>>>>>>> e362985d3... loops over shards and checks if eligible proposer
    return index;
  }

  // Removes the validator from the validator set and refunds the deposited ether 
  function withdraw(int _validatorIndex) public {
    require(msg.sender == validators[_validatorIndex].addr);
    // [FIXME] Should consider calling the validator's contract, might be useful
    // when the validator is a contract.
    validators[_validatorIndex].addr.transfer(validators[_validatorIndex].deposit);
    isValidatorDeposited[validators[_validatorIndex].addr] = false;
    delete validators[_validatorIndex];
    stackPush(_validatorIndex);
    --numValidators;
    Withdraw(_validatorIndex);
  }

<<<<<<< HEAD
  // Attempts to process a collation header, returns true on success, reverts on failure.
  function addHeader(int _shardId, uint _expectedPeriodNumber, bytes32 _periodStartPrevHash,
                     bytes32 _parentHash, bytes32 _transactionRoot,
                     address _coinbase, bytes32 _stateRoot, bytes32 _receiptRoot,
                     int _number) public returns(bool) {
    HeaderVars memory headerVars;
=======
  function sample(int _shardId) public constant returns(address) {
    require(block.number >= periodLength);
    var cycle = int(block.number) / shufflingCycleLength;
    int cycleStartBlockNumber = cycle * shufflingCycleLength - 1;
    if (cycleStartBlockNumber < 0)
      cycleStartBlockNumber = 0;
    int cycleSeed = int(block.blockhash(uint(cycleStartBlockNumber)));
    // originally, error occurs when block.number <= 4 because
    // `seed_block_number` becomes negative in these cases.
    int seed = int(block.blockhash(block.number - (block.number % uint(periodLength)) - 1));

    uint indexInSubset = uint(keccak256(seed, bytes32(_shardId))) % uint(numValidatorsPerCycle);
    uint validatorIndex = uint(keccak256(cycleSeed, bytes32(_shardId), bytes32(indexInSubset))) % uint(getValidatorsMaxIndex());

    if (validators[int(validatorIndex)].cycle > cycle)
      return 0x0;
    else
      return validators[int(validatorIndex)].addr;
  }

  // Get all possible shard ids that the given _valcodeAddr
  // may be sampled in the current cycle
  function getShardList(address _validatorAddr) public constant returns(bool[100]) {
    bool[100] memory shardList;
    int cycle = int(block.number) / shufflingCycleLength;
    int cycleStartBlockNumber = cycle * shufflingCycleLength - 1;
    if (cycleStartBlockNumber < 0)
      cycleStartBlockNumber = 0;

    var cycleSeed = block.blockhash(uint(cycleStartBlockNumber));
    int validatorsMaxIndex = getValidatorsMaxIndex();
    if (numValidators != 0) {
      for (uint8 shardId = 0; shardId < 100; ++shardId) {
        shardList[shardId] = false;
        for (uint8 possibleIndexInSubset = 0; possibleIndexInSubset < 100; ++possibleIndexInSubset) {
          uint validatorIndex = uint(keccak256(cycleSeed, bytes32(shardId), bytes32(possibleIndexInSubset)))
                             % uint(validatorsMaxIndex);
          if (_validatorAddr == validators[int(validatorIndex)].addr) {
            shardList[shardId] = true;
            break;
          }
        }
      }
    }
    return shardList;
  }

  // function checkHeader(int _shardId, bytes32 _periodStartPrevhash, int _expectedPeriodNumber) internal {
  //   // Check if the header is valid
  //   assert(_shardId >= 0 && _shardId < shardCount);
  //   assert(block.number >= periodLength);
  //   assert(uint(_expectedPeriodNumber) == block.number / periodLength);
  //   assert(_periodStartPrevhash == block.blockhash(uint(_expectedPeriodNumber)*periodLength - 1));

  //   // Check if this header already exists
  //   var entireHeaderHash = keccak256(_header);
  //   assert(entireHeaderHash != bytes32(0));
  //   assert(collationHeaders[shardId][entireHeaderHash].score == 0);
  // }

  struct Header {
      int shardId;
      uint expectedPeriodNumber;
      bytes32 periodStartPrevhash;
      bytes32 parentCollationHash;
      bytes32 txListRoot;
      address collationCoinbase;
      bytes32 postStateRoot;
      bytes32 receiptRoot;
      int collationNumber;
      bytes sig;
    }

  function addHeader(bytes _header) public returns(bool) {
    // require(_header.length <= 4096);
    // TODO
    // values = RLPList(header, [num, num, bytes32, bytes32, bytes32, address, bytes32, bytes32, num, bytes])
    // return True
    bytes memory mHeader = _header;
    var RLPList = mHeader.toRLPItem(true).iterator();
    var header = Header({
      shardId: RLPList.next().toInt(),
      expectedPeriodNumber: RLPList.next().toUint(),
      periodStartPrevhash: RLPList.next().toBytes32(),
      parentCollationHash: RLPList.next().toBytes32(),
      txListRoot: RLPList.next().toBytes32(),
      collationCoinbase: RLPList.next().toAddress(),
      postStateRoot: RLPList.next().toBytes32(),
      receiptRoot: RLPList.next().toBytes32(),
      collationNumber: RLPList.next().toInt(),
      sig: RLPList.next().toBytes()
    });
>>>>>>> e362985d3... loops over shards and checks if eligible proposer

    // Check if the header is valid
    require((_shardId >= 0) && (_shardId < shardCount));
    require(block.number >= periodLength);
    require(_expectedPeriodNumber == block.number / periodLength);
    require(_periodStartPrevHash == block.blockhash(_expectedPeriodNumber * periodLength - 1));

    // Check if this header already exists
    headerVars.entireHeaderHash = keccak256(_shardId, _expectedPeriodNumber, _periodStartPrevHash,
                                   _parentHash, _transactionRoot, bytes32(_coinbase),
                                   _stateRoot, _receiptRoot, _number);
    assert(collationHeaders[_shardId][headerVars.entireHeaderHash].score == 0);
    // Check whether the parent exists.
    // if (parent_collation_hash == 0), i.e., is the genesis,
    // then there is no need to check.
    if (_parentHash != 0x0)
        assert(collationHeaders[_shardId][_parentHash].score > 0);
    // Check if only one collation in one period
    assert(periodHead[_shardId] < int(_expectedPeriodNumber));

    // Check the signature with validation_code_addr
<<<<<<< HEAD
    headerVars.validatorAddr = getEligibleProposer(_shardId, block.number/periodLength);
    require(headerVars.validatorAddr != 0x0);
    require(msg.sender == headerVars.validatorAddr);
=======
    var collatorValcodeAddr = sample(header.shardId);
    if (collatorValcodeAddr == 0x0)
        return false;

    // assembly {
      // TODO next block
    // }
    // sighash = extract32(raw_call(self.sighasher_addr, header, gas=200000, outsize=32), 0)
    // assert extract32(raw_call(collator_valcode_addr, concat(sighash, sig), gas=self.sig_gas_limit, outsize=32), 0) == as_bytes32(1)

    // Check score == collation_number
    var _score = collationHeaders[header.shardId][header.parentCollationHash].score + 1;
    assert(header.collationNumber == _score);
>>>>>>> e362985d3... loops over shards and checks if eligible proposer

    // Check score == collationNumber
    headerVars.score = collationHeaders[_shardId][_parentHash].score + 1;
    require(_number == headerVars.score);

    // Add the header
    collationHeaders[_shardId][headerVars.entireHeaderHash] = CollationHeader({
      parentHash: _parentHash,
      score: headerVars.score
    });

    // Update the latest period number
    periodHead[_shardId] = int(_expectedPeriodNumber);

    // Determine the head
    if (headerVars.score > collationHeaders[_shardId][shardHead[_shardId]].score) {
      shardHead[_shardId] = headerVars.entireHeaderHash;
      headerVars.isNewHead = true;
    }
<<<<<<< HEAD
=======
    // Emit log
    // TODO LOG
    // log1(addHeaderLogTopic, _header);

    return true;
  }

  function getPeriodStartPrevhash(uint _expectedPeriodNumber) public constant returns(bytes32) {
    uint blockNumber = _expectedPeriodNumber * periodLength - 1;
    require(block.number > blockNumber);
    return block.blockhash(blockNumber);
  }
>>>>>>> e362985d3... loops over shards and checks if eligible proposer

    CollationAdded(_shardId, _expectedPeriodNumber, _periodStartPrevHash,
                   _parentHash, _transactionRoot, _coinbase, _stateRoot, 
                   _receiptRoot, _number, headerVars.isNewHead, headerVars.score);

    return true;
  }

  // Records a request to deposit msg.value ETH to address to in shard shard_id
  // during a future collation. Saves a `receipt ID` for this request,
  // also saving `msg.sender`, `msg.value`, `to`, `shard_id`, `startgas`,
  // `gasprice`, and `data`.
  function txToShard(address _to, int _shardId, uint _txStartgas, uint _txGasprice, 
                     bytes12 _data) public payable returns(int) {
    receipts[numReceipts] = Receipt({
      shardId: _shardId,
      txStartgas: _txStartgas,
      txGasprice: _txGasprice,
      value: msg.value,
      sender: msg.sender,
      to: _to,
      data: _data
    });
    var receiptId = numReceipts;
    ++numReceipts;
<<<<<<< HEAD
    
    TxToShard(_to, _shardId, receiptId);
    return receiptId;
  }
  
  function updateGasPrice(int _receiptId, uint _txGasprice) public payable returns(bool) {
=======

    log3(keccak256("tx_to_shard()"), bytes32(_to), bytes32(_shardId), bytes32(receiptId));
    return receiptId;
  }

  function updataGasPrice(int _receiptId, uint _txGasprice) public payable returns(bool) {
>>>>>>> e362985d3... loops over shards and checks if eligible proposer
    require(receipts[_receiptId].sender == msg.sender);
    receipts[_receiptId].txGasprice = _txGasprice;
    return true;
  }

  function isStackEmpty() internal view returns(bool) {
    return emptySlotsStackTop == 0;
  }

  function stackPush(int index) internal {
    emptySlotsStack[emptySlotsStackTop] = index;
    ++emptySlotsStackTop;
  }
  
  function stackPop() internal returns(int) {
    if (isStackEmpty())
      return -1;
    --emptySlotsStackTop;
    return emptySlotsStack[emptySlotsStackTop];
  }

  function getValidatorsMaxIndex() internal view returns(int) {
    int activateValidatorNum = 0;
    int allValidatorSlotsNum = numValidators + emptySlotsStackTop;

    // TODO: any better way to iterate the mapping?
    for (int i = 0; i < 1024; ++i) {
        if (i >= allValidatorSlotsNum)
            break;
        if (validators[i].addr != 0x0)
            activateValidatorNum += 1;
    }
    return activateValidatorNum + emptySlotsStackTop;
  }
}
