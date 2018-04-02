# Prysmatic Labs Main Sharding Reference

This document serves as a main reference for Prysmatic Labs' sharding implementation for the go-ethereum client, along with our roadmap and compilation of active research and approaches to various sharding schemes.

# Table of Contents

-   [Sharding Introduction](#sharding-introduction)
    -   [Basic Sharding Idea and Design](#basic-sharding-idea-and-design)
-   [Roadmap Phases](#roadmap-phases)
    -   [The Ruby Release: Local Network](#the-ruby-release-local-network)
    -   [The Sapphire Release: Ropsten Testnet](#the-sapphire-release-ropsten-testnet)
    -   [The Diamond Release: Ethereum Mainnet](#the-diamond-release-ethereum-mainnet)
-   [Go-Ethereum Sharding Alpha Implementation](#go-ethereum-sharding-alpha-implementation)
    -   [System Architecture](#system-architecture)
    -   [System Start and User Entrypoint](#system-start-and-user-entrypoint)
    -   [The Sharding Manager Contract](#the-sharding-manager-contract)
        -   [Necessary Functionality](#necessary-functionality)
            -   [Depositing ETH and Becoming a Collator](#depositing-eth-and-becoming-a-collator)
            -   [Determining an Eligible Collator for a Period on a Shard](#determining-an-eligible-collator-for-a-period-on-a-shard)
            -   [Withdrawing From the Collator Set](#withdrawing-from-the-collator-set)
            -   [Processing and Verifying a Collation Header](#processing-and-verifying-a-collation-header)
        -   [Collator Sampling](#collator-sampling)
        -   [Collation Header Approval](#collation-header-approval)
        -   [Event Logs](#event-logs)
    -   [The Collator Client](#the-collator-client)
        -   [Local Shard Storage](#local-shard-storage)
    -   [The Proposer Client](#the-proposer-client)
        -   [Collation Headers and State Execution](#collation-headers-and-state-execution)
    -   [Peer Discovery and Shard Wire Protocol](#peer-discovery-and-shard-wire-protocol)
    -   [Protocol Modifications](#protocol-modifications)
        -   [Protocol Primitives: Collations, Blocks, Transactions, Accounts](#protocol-primitives-collations-blocks-transactions-accounts)
        -   [The EVM: What You Need to Know](#the-evm-what-you-need-to-know)
    -   [Sharding In-Practice](#sharding-in-practice)
        -   [Fork Choice Rule](#fork-choice-rule)
        -   [Use-Case Stories: Proposers](#use-case-stories-proposers)
        -   [Use-Case Stories: Collators](#use-case-stories-collators)
    -   [Current Status](#current-status)
-   [Security Considerations](#security-considerations)
    -   [Not Included in Ruby Release](#not-included-in-ruby-release)
    -   [Bribing, Coordinated Attack Models](#bribing-coordinated-attack-models)
    -   [Enforced Windback](#enforced-windback)
        -   [Explicit Finality for Stateless Clients](#explicit-finality-for-stateless-clients)
    -   [The Data Availability Problem](#the-data-availability-problem)
        -   [Introduction and Background](#introduction-and-background)
        -   [On Uniquely Attributable Faults](#on-uniquely-attributable-faults)
        -   [Erasure Codes](#erasure-codes)
-   [Beyond Phase 1](#beyond-phase-1)
    -   [Cross-Shard Communication](#cross-shard-communication)
        -   [Receipts Method](#receipts-method)
        -   [Merge Blocks](#merge-blocks)
        -   [Synchronous State Execution](#synchronous-state-execution)
    -   [Transparent Sharding](#transparent-sharding)
    -   [Tightly-Coupled Sharding (Fork-Free Sharding)](#tightly-coupled-sharding-fork-free-sharding)
-   [Active Questions and Research](#active-questions-and-research)
    -   [Separation of Proposals and Consensus](#separation-of-proposals-and-consensus)
    -   [Selecting Eligible Collators Off-Chain](#selecting-eligible-collators-off-chain)
-   [Community Updates and Contributions](#community-updates-and-contributions)
-   [Acknowledgements](#acknowledgements)
-   [References](#references)

# Sharding Introduction

Currently, every single node running the Ethereum network has to process every single transaction that goes through the network. This gives the blockchain a high amount of security because of how much validation goes into each block, but at the same time it means that an entire blockchain is only as fast as its individual nodes and not the sum of their parts. Currently, transactions on the EVM are not parallelizable, and every transaction is executed in sequence globally. The scalability problem then has to do with the idea that a blockchain can have at most 2 of these 3 properties: decentralization, security, and scalability.

If we have scalability and security, it would mean that our blockchain is centralized and that would allow it to have a faster throughput. Right now, Ethereum is decentralized and secure, but not scalable.

An approach to solving the scalability trilemma is the idea of blockchain sharding, where we split the entire state of the network into partitions called shards that contain their own independent piece of state and transaction history. In this system, certain nodes would process transactions only for certain shards, allowing the throughput of transactions processed in total across all shards to be much higher than having a single shard do all the work as the main chain does now.

## Basic Sharding Idea and Design

A sharded blockchain system is made possible by having nodes store “signed metadata” in the main chain of latest changes within each shard chain. Through this, we manage to create a layer of abstraction that tells us enough information about the global, synced state of parallel shard chains. These messages are called **collation headers**, which are specific structures that encompass important information about the chainstate of a shard in question. Collations are created by actors known as **proposer nodes** that are randomly tasked into packaging transactions and “selling” them to collator nodes that are then tasked into adding these collations into particular shards through a **proof of stake** system in a designated period of time.

These collations are holistic descriptions of the state and transactions on a certain shard.  A collation header contains the following information:

-   Information about what shard the collation corresponds to (let’s say shard 10)
-   Information about the current state of the shard before all transactions are applied
-   Information about what the state of the shard will be after all transactions are applied

For detailed information on protocol primitives including collations, see: [Protocol Primitives](#protocol-primitives). We will have two types of nodes that do the heavy lifting of our sharding logic: **proposers and collators**. The basic role of proposers is to fetch pending transactions from the txpool, execute any state logic or computation, wrap them into collations, and submit them along with an ETH deposit to a **proposals pool**.

<!--[Proposer{bg:wheat}]fetch txs-.->[TXPool], [TXPool]-.->[Proposer{bg:wheat}], [Proposer{bg:wheat}]-package txs>[Collation|header|ETH Deposit], [Collation|header|ETH Deposit]-submit>[Proposals Pool], [Collator{bg:wheat}]subscribe to-.->[Proposals Pool]-->
![proposers](https://yuml.me/6da583d7.png)

Collators add collations in the proof of work chain, throughout the document named the **canonical chain**. Collators subscribe to updates in the proposals pool and pick a collation in their best interest. Once collators are selected to add collations to the canonical chain, and do so successfully, they get paid by the deposit the proposer offered.

To recap, the role of a collator is to reach consensus through Proof of Stake on collations they receive in the period they are assigned to. This consensus will involve validation and data availability proofs of collations proposed to them by proposer nodes, along with validating collations from the immediate past (See: [Windback](#enforced-windback)).

When processing collations, proposer nodes download the merkle branches of the state that transactions within their collations need. In the case of cross-shard transactions, an access list of the state along with transaction receipts are required as part of the transaction primitive (See: [Protocol Primitives](#protocol-primitives)). Additionally, these proposers need to provide proofs of availability and validity when submitting collations for “sale” to collators. This submission process is akin to the current transaction fee open bidding market where miners accept the transactions that maximize their profits. This abstract separation of concerns between collators and proposers allows for more computational efficiency within the system, as collators will not have to do the heavy lifting of state execution and focus solely on consensus through fork-choice rules.

When deciding and signing a proposed, valid collation, collators have the responsibility of finding the **longest valid shard chain within the longest valid main chain**.

In this new protocol, a block is valid when

-   Transactions in all collations are valid
-   The state of collations after the transactions is the same as what the collation headers specified

Collators periodically get assigned to different shards, the moment between when collators get assigned to a shard and the moment they get reassigned is called a **period**.

Given that we are splitting up the global state of the Ethereum blockchain into shards, new types of attacks arise because fewer hash power is required to completely dominate a shard. This is why a **source of randomness**, and periods are critical components to ensuring the integrity of the system.

The Ethereum Wiki’s [Sharding FAQ](https://github.com/ethereum/wiki/wiki/Sharding-FAQ) suggests random sampling of collators on each shard. The goal is so that these collators will not know which shard they will get in advance. Every shard will get assigned a bunch of collators and the ones that will actually be collating transactions will be randomly sampled from that set. Otherwise, malicious actors could concentrate hash power into a single shard and try to overtake it (See: [1% Attack](https://medium.com/@icebearhww/ethereum-sharding-and-finality-65248951f649)).

Casper Proof of Stake (Casper [FFG](https://arxiv.org/abs/1710.09437) and [CBC](https://arxiv.org/abs/1710.09437)) makes this quite trivial because there is already a set of global collators that we can select collator nodes from. The source of randomness needs to be common to ensure that this sampling is entirely compulsory and can’t be gamed by the collators in question.

In practice, the first phase of sharding will not be a complete overhaul of the network, but rather an implementation through a smart contract on the main chain known as the **Sharding Manager Contract (SMC)**. Its responsibility is to manage shards and sampling proposed collators from a global collator set. As the SMC lives in the canonical chain, it will take guarantee a global state among all shard states.

Among its basic responsibilities, the SMC is be responsible for reconciling collators across all shards. It is in charge of pseudorandomly sampling collators from a collator set of accounts that have staked ETH into the SMC. The SMC is also responsible for providing immediate collation header verification that records a valid collation header hash on the canonical chain. In essence, sharding revolves around being able to store proofs of shard states in the canonical chain through this smart contract.

# Roadmap Phases

Prysmatic Labs’ implementation will follow parts of the roadmap outlined by Vitalik in his [Sharding FAQ](https://github.com/ethereum/wiki/wiki/Sharding-FAQ) to roll out a working version of quadratic sharding, with a few modifications on our releases.

1.  **Phase 1:** Basic SMC shard system with no cross-shard communication along with a proposer + collator node architecture
2.  **Phase 2:** Receipt-based, cross-shard communication
3.  **Phase 3:** Require collation headers to be added in as uncles instead of as transactions
4.  **Phase 4:** Tightly-coupled sharding with data availability proofs and robust security

To concretize these phases, we will be releasing our implementation of sharding for the geth client as follows:

## The Ruby Release: Local Network

Our current work is focused on creating a localized version of phase 1, quadratic sharding that would include the following:

-   A minimal, **collator client** system that will deploy a **Sharding Manager Contract** to a locally running geth node
-   Ability to deposit ETH into the SMC through the command line and to be selected as a collator by the local **SMC** in addition to the ability to withdraw the ETH staked
-   A **proposer node client** and Cryptoeconomic incentive system for proposer nodes to listen for pending tx’s, create collations, and submit them along with a deposit to collator nodes in the network
-   A simple command line util to **simulate pending transactions** of different types posted to the local geth node’s txpool for the local collation proposer to begin proposing collation headers
-   Ability to inspect the shard states and visualize the working system locally through the command line

We will forego many of the security considerations that will be critical for testnet and mainnet release for the purposes of demonstration and local network execution as part of the Ruby Release (See: [Security Considerations Not Included in Ruby](#not-included-in-ruby-release)).

ETA: To be determined

## The Sapphire Release: Ropsten Testnet

Part 1 of the **Sapphire Release** will focus around getting the **Ruby Release** polished enough to be live on an Ethereum testnet and manage a set of collators effectively processing collations through the **on-chain SMC**. This will require a lot more elaborate simulations around the safety of the randomness behind the collator assignments in the SMC. Futhermore we need to pass stress testing against DDoS and other sorts of byzantine attacks. Additionally, it will be the first release to have real users proposing collations concurrently along with collators that can accept these proposals and add their headers to the SMC.

Part 2 of the **Sapphire Release** will focus on implementing a cross-shard transaction mechanism via two-way pegging and the receipts system (as outlined in [Beyond Phase 1](#beyond-phase-1)) and getting that functionality ready to run on a **local, private network** as an extension to the Ruby Release.

ETA: To be determined

## The Diamond Release: Ethereum Mainnet

The **Diamond Release** will reconcile the best parts of the previous releases and deploy a full-featured, cross-shard transaction system through a Sharding Manager Contract on the Ethereum mainnet. As expected, this is the most difficult and time consuming release on the horizon for Prysmatic Labs. We plan on growing our community effort significantly over the first few releases to get all hands-on deck preparing for real ether to be staked in the SMC.

The Diamond Release should be considered the production release candidate for sharding Ethereum on the mainnet.

ETA: To Be determined

# Go-Ethereum Sharding Alpha Implementation

Prysmatic Labs will begin by focusing its implementation entirely on the **Ruby Release** from our roadmap. We plan on being as pragmatic as possible to create something that can be locally run by any developer as soon as possible. Our initial deliverable will center around a command line tool that will serve as an entrypoint into a collator client that allows staking, a proposer client that allows for simple state execution and creation of collation proposals, and processing collations through on-chain verification via the Sharding Manager Contract.

Here is a full reference spec explaining how our initial system will function:

## System Architecture

Our implementation revolves around 5 core components:

-   A **locally-running geth node** that spins up an instance of the Ethereum blockchain
-   A **Sharding Manager Contract (SMC)** that is deployed onto this blockchain instance
-   A **collator client** that connects to the running geth node through JSON-RPC, provides bindings to the SMC, and listens for incoming collation proposals
-   A **proposer client** that is tasked with state execution, processing pending tx’s from the Geth node, and creates collations that are then broadcast to collators via a wire protocol
-   A **user** that will interact with the sharding client to become a collator and automatically process transactions into shards through the sharding client’s SMC bindings.

Our initial implementation will function through simple command line arguments that will allow a user running the local geth node to deposit ETH into the SMC and join as a collator that is automatically assigned to a certain period. We will also launch a proposer client that will create collations from the geth node’s tx pool and submit them to collators for them to add their headers to the SMC.

A basic, end-to-end example of the system is as follows:

1.  _**A User starts a collator client and deposits 100ETH into the SMC:**_ the sharding client connects to a locally running geth node and asks the user to confirm a deposit from his/her personal account.

2.  _**Collator client connects & listens to incoming headers from the geth node and assigns user as collators on a shard per period:**_ The collator is selected for the current period and must accept collations from proposer nodes that offer the best prices.

3.  _**Concurrently, the proposer client processes and executes pending tx’s from the geth node:**_ the proposer client will create valid collations and submit them to collators on an open bidding system.

4.  _**Collators select collation proposals that offer highest payout:**_ Collators listen to collation headers on a certain shard with high deposit sizes and sign them.

5.  _**The collator adds collation headers through the SMC:**_ The collator client calls the `add_header` function in the SMC and to append the header to the canonical chain.

6.  _**The user is selected as collator again on the SMC in a different period or can withdraw his/her stake from the collator's Pool:**_ the user can keep staking and adding incoming collation headers and restart the process, or withdraw his/her stake and be removed from the SMC collator set.

Now, we’ll explore our architecture and implementation in detail as part of the go-ethereum repository.

## System Start and User Entrypoint

Our Ruby Release requires users to start a local geth node running a localized, private blockchain to deploy the **SMC** into. Users can spin up a collator client as a command line entrypoint into geth while the node is running as follows:

<<<<<<< HEAD
```
$ geth sharding-collator --deposit --datadir /path/to/your/datadir --password /path/to/your/password.txt --networkid 12345
```
=======
    $ geth sharding-validator --deposit 100eth --password ~/Desktop/password.txt
>>>>>>> 0d775a183... remark beautify markdown

If it is the first time the client runs, it deploys a new **SMC** into the local chain and establish a JSON-RPC connection to interact with the node directly. The `--deposit` flag tells the sharding client to automatically unlock the user’s keystore and begin depositing ETH into the SMC to become a collator.

If the initial deposit is successful, the sharding client launches a **local, transaction simulation generator**, which will queue transactions into the txpool for the geth node to process that can then be added into collations on a shard.

Concurrently, a user needs to launch a **proposer client** that starts processing transactions into collations that can then be “sold” to collators by including a cryptographic proof of an ETH deposit in their unsigned collation headers. This proof is a guarantee of a state change in the collator’s account balance for accepting to sign the incoming collation header. The proposer client can also be initialized as follows in a separate process:

<<<<<<< HEAD
<<<<<<< HEAD
```
geth sharding-proposer --datadir /path/to/your/datadir --password /path/to/your/password.txt --networkid 12345
```
=======
    geth sharding-proposer --password ~/Desktop/password.txt
>>>>>>> 0d775a183... remark beautify markdown
=======
    geth sharding-collator --password ~/Desktop/password.txt
>>>>>>> 8b8c08e58... rename sharding proposer cli entrypoint, add travis

Back to the collators, the collator client begins to work by its main loop, which involves the following steps:

1.  _**Subscribe to incoming block headers:**_ the client will begin by issuing a subscription over JSON-RPC for block headers from the running geth node.

2.  _**Check shards for eligible collator:**_ On incoming headers, the client will interact with the SMC to check if the current collator is an eligible collator for upcoming periods (only a few minutes notice)

3.  _**If the collator is selected, fetch proposals from proposal nodes and add collation headers to SMC:**_ Once a collator is selected, he/she only has a small timeframe to add collation headers through the SMC, so he/she looks for proposals from proposer nodes and accepts those that offer the highest payouts. The collator then countersigns the collation header, receives the full collation with its body after signing, and adds it to the SMC through PoS consensus.

4.  _**Supernode reconciles and adds to main chain:**_ Supernodes that are responsible for reconciling global state across shards into the canonical chain. They are tasked with observing the state across the whole galaxy of shards and adding blocks to the canonical PoW main chain. Proposers get rewarded to their coinbase address for inclusion of a block (also known as a collation subsidy).

5.  _**If user withdraws, remove from collator set:**_ A user can choose to stop being a collator and then his/her ETH is withdrawn from the collator set.

6.  _**Otherwise, collating client keeps subscribing to block headers:**_ If the user chooses to keep going,
It will be the proposer client’s responsibility to listen to any new broadcasted transactions to the node and interact with collators that have staked their ETH into the SMC through an open bidding system for collation proposals. Proposer clients are the ones responsible for **state execution** of transactions in the tx pool.

<!--[Transaction Generator]generate test txs->[Shard TXPool],[Geth Node]-deploys>[Sharding Manager Contract{bg:wheat}], [Shard TXPool]<fetch pending txs-.->[Proposer Client], [Proposer Client]-propose collation>[Collator Client],[Collator Client]add collation header->[Sharding Manager Contract{bg:wheat}]-->
![system functioning](https://yuml.me/4a7c8c5b.png)

## The Collator Manager Contract

Our solidity implementation of the Collator Manager Contract follows the reference spec outlined by Vitalik [here](https://github.com/ethereum/sharding/blob/develop/docs/doc.md).

### Necessary functionality

In our Solidity implementation, we begin with the following sensible defaults:

```javascript
// Constant values
uint constant periodLength = 5;
int constant public shardCount = 100;
// The exact deposit size which you have to deposit to become a collator
uint constant depositSize = 100 ether;
// Number of periods ahead of current period, which the contract
// is able to return the collator of that period
uint constant lookAheadPeriods = 4;
```

Then, the 4 minimal functions required by the SMC are as follows:

#### Depositing ETH and Becoming a Collator

```javascript
function deposit() public payable returns(int) {
    require(!isCollatorDeposited[msg.sender]);
    require(msg.value == depositSize);
    ...
}
```

`deposit` adds a collator to the collator set, with the collator's size being the `msg.value` (i.e., the amount of ETH deposited) in the function call. This function returns the collator's index.

#### Determining an Eligible Collator for a Period on a Shard

```javascript
function getEligibleCollator(int _shardId, int _period) public view returns(address) {
    require(_period >= lookAheadPeriod);
    require((_period - lookAheadPeriods) * periodLength < block.number);
    ...
}
```

The `getEligibleCollator` function uses a block hash as a seed to pseudorandomly select a signer from the collator set. The chance of being selected should be proportional to the collator's deposit. The function should be able to return a value for the current period or any future up to `LOOKAHEAD_PERIODS` periods ahead.

#### Withdrawing From the collator Set

```javascript
function withdraw(int _collatorIndex) public {
    require(msg.sender == collators[_collatorIndex].addr);
    ...
}
```

Authenticates the collator and removes him/her from the collator set, refunding the deposited ETH.

#### Processing and Verifying a Collation Header

```javascript
function addHeader(int _shardId, uint _expectedPeriodNumber, bytes32 _periodStartPrevHash,
                     bytes32 _parentHash, bytes32 _transactionRoot,
                     address _coinbase, bytes32 _stateRoot, bytes32 _receiptRoot,
                     int _number) public returns(bool) {
    HeaderVars memory headerVars;

    // Check if the header is valid
    require((_shardId >= 0) && (_shardId < shardCount));
    require(block.number >= periodLength);
    require(_expectedPeriodNumber == block.number / periodLength);
    require(_periodStartPrevHash == block.blockhash(_expectedPeriodNumber * periodLength - 1));
    …
}
```

The `addHeader` function is the most important function in the SMC as it provides on-chain verification of collation headers, and maintains a canonical ordering of processed collation headers.

Our current [solidity implementation](https://github.com/prysmaticlabs/geth-sharding/blob/master/sharding/contracts/sharding_manager.sol) includes all of these functions along with other utilities important for the our Ruby Release sharding scheme.

### Collator Sampling

The probability of being selected as a collator on a particular shard should be completely dependent on the stake of the collator and not on other factors. This is a key distinction. As specified in the [Sharding FAQ](https://github.com/ethereum/wiki/wiki/Sharding-FAQ) by Vitalik, “if validators [collators] could choose, then attackers with small total stake could concentrate their stake onto one shard and attack it, thereby eliminating the system’s security.”

The idea is that collators should not be able to figure out which shard they will become a collator of and during which period they will be assigned with anything more than a few minutes notice. To accomplish this, random sampling would require collators to redownload entire large parts of new shard states they get assigned to as part of the collation process in a naive approach. However, our approach separates the consensus and state execution/collation proposal mechanisms, which allows collators to not have to download shard states save for specific situations.

Ideally, we want collators to shuffle across shards very rapidly and through a trustworthy source of randomness built in-protocol.

Although this separation of consensus and state execution is an attractive way to fix the overhead of having to redownload shard states, random sampling does not help in a bribing, coordinated attack model. In Vitalik’s own words:

_"Either the attacker can bribe the great majority of the sample to do as the attacker pleases, or the attacker controls a majority of the sample directly and can direct the sample to perform arbitrary actions at low cost (O(c) cost, to be precise).
At that point, the attacker has the ability to conduct 51% attacks against that sample. The threat is further magnified because there is a risk of cross-shard contagion: if the attacker corrupts the state of a shard, the attacker can then start to send unlimited quantities of funds out to other shards and perform other cross-shard mischief. All in all, security in the bribing attacker or coordinated choice model is not much better than that of simply creating O(c) altcoins.”_

However, this problem transcends the sharding scheme itself and goes into the broader problem of fraud detection, which we have yet to comprehensively address.

### Collation Header Approval

Explains the on-chain verification of a collation header.

Work in progress.

### Event Logs

Explain how CollationAdded logs will later on be used.

Work in progress.

## The Collator Client

The main running thread of our implementation is the collator client, which serves as a bridge between users staking their ETH, proposers offering collations to these collators, and the **Sharding Manager Contract** that verifies collation headers on the canonical chain.

When we launch the client with

<<<<<<< HEAD
```
geth sharding-collator --deposit --datadir /path/to/your/datadir --password /path/to/your/password.txt --networkid 12345
```
=======
    geth sharding-validator --deposit 100eth --password ~/Desktop/password.txt
>>>>>>> 0d775a183... remark beautify markdown

The instance connects to a running geth node via JSON-RPC and calls the deposit function on a deployed, Sharding Manager Contract to insert the user into a collator set. Then, we subscribe for updates on incoming block headers and call `getEligibleCollator` on receiving each header. Once we are selected, our client fetches and “purchases” proposed, unsigned collations from a proposals pool created by proposer nodes. The collator client accepts a collation that offer the highest payout, countersigns it, and adds it to the SMC all within that period.

### Local Shard Storage

Local shard information is done through the same LevelDB, key-value store used to store the mainchain information in the local data directory specified by the running geth node. Adding a collation to a shard will effectively modify this key-value store.

Work in progress.

## The Proposer Client

In addition to launching a collator client, our system requires a user to concurrently launch a proposer client that is tasked with state execution, fetching pending tx’s from the running geth node’s txpool, and creating collations that can be sent to collators.

Users launch a proposal client as another geth entrypoint as follows:

<<<<<<< HEAD
<<<<<<< HEAD
```
geth sharding-collator --datadir /path/to/your/datadir --password /path/to/your/password.txt --networkid 12345
```
=======
    geth sharding-proposer --password ~/Desktop/password.txt
>>>>>>> 0d775a183... remark beautify markdown
=======
    geth sharding-collator --password ~/Desktop/password.txt
>>>>>>> 8b8c08e58... rename sharding proposer cli entrypoint, add travis

Launching this command connects via JSON-RPC to fetch the geth node’s tx pool and see who the currently active collator node is for the period. The proposer is tasked with running transactions to create valid collations and executing their required computations, tracking used gas, and all the heavy lifting that is usually seen in full Ethereum nodes. Once a valid collation is created, the proposer broadcasts the unsigned header **(note: the body is not broadcasted)** to a proposals pool along with a guaranteed ETH deposit that is extracted from the proposer’s account upfront. Then, the current collator assigned for the period will find proposals for him/her assigned shard and sign the one with the highest payout.

Then, the collator node calls the addHeader function on the SMC by submitting this collation header. We’ll explore the structure of collation headers in this next section along with important considerations for state execution, as this can quickly become the bottleneck of the entire sharding system.

### Collation Headers and State Execution

Work in progress.

## Peer Discovery and Shard Wire Protocol

Work in progress.

## Protocol Modifications

### Protocol Primitives: Collations, Blocks, Transactions, Accounts

(Outline the interfaces for each of these constructs, mention crucial changes in types or receiver methods in Go for each, mention transaction access lists)

Work in progress.

### The EVM: What You Need to Know

As an important aside, we’ll take a brief detour into the EVM and what we need to understand before we modify it for a sharded blockchain. At its core, the functionality of the EVM optimizes for _security_ and not for computational power with the following restrictions:

-   Every single step must be paid for upfront with gas costs to prevent DDoS
-   Programs can't interact with each other without a single byte array
    -   This also means programs can't access other programs' state
-   Sandboxed Execution - the EVM can only modify its internal state and nothing else
-   Deterministic execution guarantees

So what exactly is the EVM? The EVM was purposely designed to be a stack based machine with memory-byte arrays and key-value stores that are kept on a trie

-   Every single keys and storage values are 32 bytes
-   There are 100 total opcodes in the EVM
-   The EVM comes with a temporary memory byte-array and storage trie to hold persistent memory.

Cryptographic operations are done using pre-compiled contracts. Aside from that, the EVM provides a bunch of blockchain access-level context that allows certain opcodes to fetch useful information from the external system. For example, LOG opcodes store useful information in the log bloom filter that can be synced with light clients. This can be used as a low-gas form of storage, since LOG does not modify the state.

Additionally, the EVM contains a call-depth limit such that recursive invocations or chains of calls will eventually halt, preventing a drastic use of resources.

It is important to note that the merkle root of an Ethereum account is updated any time an `SSTORE` opcode is executed successfully by a program on the EVM that results in a key or value changing in the state merklix (merkle radix) tree.

How is this relevant to sharding? It is important to note the importance of certain opcodes in our implementation and how we will need to introduce and modify several of them for both security and scalability considerations in a sharded chain.

Work in progress.

## Sharding In-Practice

### Fork Choice Rule

In the sharding consensus mechanism, it is important to consider that we now have two layers of longest chain rules when adding a collation. When we are reaching consensus on the best shard chain, we not only have to check for the longest canonical main chain, but also the longest shard chain **within** this longest main chain. Vlad Zamfir has elaborated on this fork-choice rule in a [tweet](https://twitter.com/VladZamfir/status/945358660187893761) that is important for our sharding scheme.

### Use-Case Stories: Proposers

The primary purpose of proposers is to use their computational power for state execution of transactions and create valid collations that can then be put on an open market for collators to take. Upon offering a proposal, proposers will deposit part of their ETH as a payout to the collator that adds its collation header to the SMC, __even if the collation gets orphaned__. By forcing proposers to take on this risk, we prevent a certain degree of collation proposal spamming, albeit not without a few other security problems: (See: [Active Research](#active-questions-and-research)).

The primary incentive for proposers to generate these collations is to receive a payout to their coinbase address along with transactions fees from the ones they process once added to a block in the canonical chain.

### Use-Case Stories: Collators

The primary purpose of collators is to use Proof of Stake and reach **consensus** on valid shard chains based on the collations they process and add to the Sharding Manager Contract. They have two primary options they can choose to do:

-   They can deposit ETH into the SMC and become a collator. They then have to wait to be selected by the SMC on a particular period to add a collation header to the SMC.
-   They can accept a collation proposal from the collation pool during their eligible period.
-   They can withdraw their stake and stop being a part of the collator pool

The primary incentive for collators is to earn the payouts from the proposers offering them collations within their period.

## Current Status

Currently, Prysmatic Labs is focusing its initial implementation around the logic of the collator and proposer clients. We have built the command line entrypoints as well as the minimum, required functions of the Sharding Manager Contract that is deployed to a local Ethereum blockchain instance. Our collator client is able to subscribe for block headers from the running Geth node and determine when we are selected as an eligible collator in a given period if we have deposited ETH into the contract.

You can track our progress, open issues, and projects in our repository [here](https://github.com/prysmaticlabs/geth-sharding).

# Security Considerations

## Not Included in Ruby Release

Under the uncoordinated majority model, in order to prevent a single shard takeover, random sampling is utilized. Each shard is assigned a certain number of collators. Collators that approve the collations on that shard are sampled randomly. However for the ruby release we will not be implementing any random sampling of the collators of the shard, as the primary objective of this release is to launch an archival sharding client which deploys the Sharding Management Contract to a locally running geth node.

Also for now we will not be implementing challenge response mechanisms to mitigate instances where malicious actors are penalized and have their staked slashed for making incorrect claims regarding the veracity of collations.

We will not be considering data availability proofs as part of the ruby release we will not be implementing them as it just yet as they are an area of active research.

## Bribing, Coordinated Attack Models

Work in progress.

## Enforced Windback

When collators are extending collator chains by adding headers to the SMC, it is critical that they are able to verify some of the collation headers in the immediate past for security purposes. There have already been instances where mining blindly has led to invalid transactions that forced Bitcoin to undergo a fork (See: [BIP66 Incident](https://bitcoin.stackexchange.com/questions/38437/what-is-spv-mining-and-how-did-it-inadvertently-cause-the-fork-after-bip66-wa)).

As part of the sharding process, we want to ensure collators do two things when they look at the immediate past:

Validity of collations through checking the integrity of transactions within their body.
Checking for availability of the data within past collation bodies.

This checking process is known as **“windback”**. In a [post](https://ethresear.ch/t/enforcing-windback-validity-and-availability-and-a-proof-of-custody/949) by Justin Drake on ETHResearch, he outlines that this is necessary for security, but is counterintuitive to the end-goal of scalability as this obviously imposes more computational and network constraints on collator nodes.

One way to enforce **validity** during the windback process is for collators to produce zero-knowedge proofs of validity that can then be stored in collation headers for quick verification.

On the other hand, to enforce **availability** for the windback process, a possible approach is for collators to produce “proofs of custody” in collation headers that prove the collator was in possession of the full data of a collation when produced. Drake proposes a constant time, non-interactive zkSNARK method for collators to check these proofs of custody. In his construction, he mentions splitting up a collation body into “chunks” that are then mixed with the collator's private key through a hashing scheme. The security in this relies in the idea that a collator would not leak his/her private key without compromising him or herself, so it provides a succinct way of checking if the full data was available when a collator processed the collation body and proof was created.

### Explicit Finality for Stateless Clients

Vitalik has [mentioned](https://ethresear.ch/t/detailed-analysis-of-stateless-client-witness-size-and-gains-from-batching-and-multi-state-roots/862/5?u=rauljordan) that the average amount of windback, or how many immediate periods in the past a collator has to check before adding a collation header, is around 25. In a [medium post](https://medium.com/@icebearhww/ethereum-sharding-and-finality-65248951f649) on the value of explicit finality for sharding, Hsiao-Wei Wang mentions how the finality that Casper FFG provides would mean stateless clients would be entirely confident of blocks ahead to prevent complete reshuffling and faster collation processing. In her own words:

_“Casper FFG will provide explicit finality threshold after about 2.5 “epoch times”, i.e., 125 block times [1][7]. If validators [collators] can verify more than 125 / PERIOD_LENGTH = 25 collations during reshuffling, the shard system can benefit from explicit finality and be more confident of the 25 ahead collations from now are all finalized.”_

Casper allows us to forego some of these windback considerations and reduces the number of constraints on scalability from a collation header verification standpoint.

## The Data Availability Problem

### Introduction and Background

Work in progress.

### On Uniquely Attributable Faults

Work in progress.

### Erasure Codes

Work in progress.

# Beyond Phase 1

## Cross-Shard Communication

### Receipts Method

Work in progress.

### Merge Blocks

Work in progress.

### Synchronous State Execution

Work in progress.

## Transparent Sharding

One of the first question dApp developers ask about sharding is how much will they need to change their workflow and smart contract development to adopt the sharded blockchain scheme. An idea tangentially explored by Vitalik in his [Sharding FAQ](https://github.com/ethereum/wiki/wiki/Sharding-FAQ) was the concept of **“transparent sharding”** which means that sharding will exist exclusively at the protocol layer and will not be exposed to developers. The Ethereum state system will continue to look as it currently does, but the protocol will have a built-in system that creates shards, balances state across shards, gets rid of shards that are too small, and more. This will all be done behind the scenes, allowing devs to continue their current workflow on Ethereum. This was only briefly mentioned, but will be critical to ensure a better user experience moving forward after security considerations are addressed.

## Tightly-Coupled Sharding (Fork-Free Sharding)

A current problem with the scheme we are following for sharding is the reliance on **two fork-choice rules**. When we are reaching consensus on the best shard chain, we not only have to check for the longest canonical, main chain, but also the longest shard chain **within** this longest main chain. Fork-choice rules have long been an approach to solve the constraints that distributed systems impose on us due to factors outside of our control (Byzantine faults) and are the current standard in most public blockchains.

A problem that can occur with current distributed fork-choice ledgers is the possibility of choosing a wrong fork and continuing to do PoW on it, thereby wasting potential profits of mining on the canonical chain. Another current burden is the large amount of data that needs to be downloaded in order to validate which fork is potentially the best one to follow in any situation, opening up avenues for spam DDoS attacks.

Fortunately, there is a potential method of creating a fork-free sharding mechanism that relies on what we are currently implementing through the Sharding Manager Contract that has been explored by Justin Drake and Vitalik in [this](https://ethresear.ch/t/fork-free-sharding/1058) and this [other post](https://ethresear.ch/t/a-model-for-stage-4-tightly-coupled-sharding-plus-full-casper/1065), respectively.

The current spec of the Sharding Manager Contract __already does a canonical ordering of collation headers for us__ (i.e. we can track the timestamped logs of collation headers being added). Because the data for the SMC lives on the canonical main chain, we are able to easily extract an exact ordering and validity from headers added through the contract.

To add validity to our current SMC spec, Drake mentions that we can use a succinct zkSNARK in the collation root proving validity upon construction that can be checked directly by the `addHeader` function on the the SMC.

The other missing piece is the guarantee of data availability within collation headers submitted to the SMC which can once again be done through zero-knowledge proofs and erasure codes (See: The Data Availability Problem). By escalating this up to the SMC, we can ensure what Vitalik calls “tightly-coupled” sharding, in which case breaking a single shard would entail also breaking the progression of the canonical chain, enabling easier cross-shard communication due to having a single source of truth being the SMC and the associated collation headers it has processed. In Justin Drake’s words, “there is no fork-choice rule within the SMC”.

It is important to note that this “tightly coupled” sharding has been relegated to Phase 4 of the roadmap.

Work in progress.

# Active Questions and Research

## Separation of Proposals and Consensus

In a recent [blog post](https://ethresear.ch/t/separating-proposing-and-confirmation-of-collations/1000/7), Vitalik has outlined a novel system to the sharding mechanism through better separation of concerns. In the current reference documentation for a sharding systems, collators are responsible for proposing transactions into collations and reaching consensus on these proposals. That is, this process happens _all at once_, as proposing a collation happens in tandem with consensus.

This leads to significant computational burdens on collators that need to keep track of the state of a particular shard in the proposal process as part of the transaction packaging process. The potentially better approach outlined above is the idea of separating the transaction packaging process and the consensus mechanism into two separate nodes with different responsibilities. **Our model will be based on this separation and we will be releasing a proposer client alongside a collator client in our Ruby release**.

The two nodes would interact through a cryptoeconomic incentive system where proposers would package transactions and send unsigned collation headers (with obfuscated collation bodies) over to collators with a signed message including a deposit. If a collator chooses to accept the proposal, he/she would be rewarded by the amount specified in the deposit. This would allow proposers to focus all of their computational power solely on state execution and organizing transactions into proposals.

Along the same lines, it will make it easier for collators to constantly jump across shards to validate collations, as they no longer have the need for resyncing an entire state tree because they can simply receive collation proposals from proposer nodes. This is very important, as collator reshuffling is a crucial security consideration to prevent shard hostile takeovers.

A key question asked in the post is whether _“this makes it easy for a proposer to censor a transaction by paying a high pass-through fee for collations without a certain transaction?”_ and the answer to that is that yes, this could happen but in the current system a proposer could censor transactions by simply not including them (see: [The Data Availability Problem](#the-data-availability-problem)).

It is important to note a possible attack vector in this case: which is that a an attacker could spam proposals on a particular shard and take on the price of excluding certain transactions as a censorship mechanism while the collators would have no idea this is happening. However, given a competitive enough proposal economy, this would be very similar to the current problem of transaction spam in traditional blockchains.

In this system, collators would get paid the proposal’s deposit __even if the collation does not get appended to the shard__. Proposers would have to take on this risk to mitigate the possibilities of malicious intent to make an obfuscated-collation-body proposal system work. Only collations that are double signed and have an available body can be included in the main chain, fee F goes to the collator regardless whether collation gets into the main chain, but fee T only goes to the proposer if the collation gets included to the main chain.

In practice, we would end up with a system of 3 types of nodes to ensure the functioning of a sharded blockchain

1.  Proposer nodes that are tasked with state execution and creation of unsigned collation headers, obfuscated collation bodies, data availability proofs, and an ETH deposit to be relayed to collators.
2.  Collator nodes that accept proposals through an open auction system similar to way transactions fees currently work. These nodes then sign these collations and pass them through the SMC for inclusion into a shard through PoS.

## Selecting Eligible Collators Off-Chain

In our current implementation for the Ruby Release, we are selecting collators on-chain by calling the `getEligibleCollator` function from the SMC directly. Justin Drake proposes an alternative scheme that could potentially open up new collator selection mechanisms off-chain through a fork-choice rule in [this](https://ethresear.ch/t/fork-choice-rule-for-collation-proposal-mechanisms/922) post on ETHResearch.

In his own words, this scheme “saves gas when calling addHeader and unlocks the possibility for fancier proposer eligibility functions”. A potential way to do so would be through private collator sampling which is elaborated on below:

_“We now look at the problem of private sampling. That is, can we find a proposal mechanism which selects a single validator [collator] per period and provides “private lookahead”, i.e. it does not reveal to others which validators will be selected next?
There are various possible private sampling strategies (based on MPCs, SNARKs/STARKs, cryptoeconomic signalling, or fancy crypto) but finding a workable scheme is hard. Below we present our best attempt based on one-time ring signatures. The scheme has several nice properties:
Perfect privacy: private lookahead and private lookbehind (i.e. the scheme never matches eligible collators with specific validators)
Full lookahead: the lookahead extends to the end of the epoch (epochs are defined below, and have roughly the same size as the validator set)
Perfect fairness: within an epoch validators are selected proportionally according to deposit size, with zero variance”_  - Justin Drake

# Community Updates and Contributions

Excited by our work and want to get involved in building out our sharding releases? We created this document as a single source of reference for all things related to sharding Ethereum, and we need as much help as we can get!

You can explore our [Current Projects](https://github.com/prysmaticlabs/geth-sharding/projects) in-the works for the Ruby release. Each of the project boards contain a full collection of open and closed issues relevant to the different parts of our first implementation that we use to track our open source progress. Feel free to fork our repo and start creating PR’s after assigning yourself to an issue of interest. We are always chatting on [Gitter](https://gitter.im/prysmaticlabs/geth-sharding), so drop us a line there if you want to get more involved or have any questions on our implementation!

**Contribution Steps**

-   Create a folder in your `$GOPATH` and navigate to it `mkdir -p $GOPATH/src/github.com/ethereum && cd $GOPATH/src/github.com/ethereum`
-   Clone our repository as `go-ethereum`, `git clone https://github.com/prysmaticlabs/geth-sharding ./go-ethereum`
-   Fork the `go-ethereum` repository on Github: <https://github.com/ethereum/go-ethereum>
-   Add a remote to your fork
    \`git remote add YOURNAME <https://github.com/YOURNAME/go-ethereum>

Now you should have a remote pointing to the `origin` repo (geth-sharding) and to your forked, go-ethereum repo on Github. To commit changes and start a Pull Request, our workflow is as follows:

-   Create a new branch with a clear feature name such as `git checkout -b collations-pool`
-   Issue changes with clear commit messages
-   Push to your remote `git push YOURNAME collations-pool`
-   Go to the [geth-sharding](https://github.com/prysmaticlabs/geth-sharding) repository on Github and start a PR comparing `geth-sharding:master` with `go-ethereum:collations-pool` (your fork on your profile).
-   Add a clear PR title along with a description of what this PR encompasses, when it can be closed, and what you are currently working on. Github markdown checklists work great for this.

# Acknowledgements

A special thanks for entire [Prysmatic Labs](https://gitter.im/prysmaticlabs/geth-sharding) team for helping put this together and to Ethereum Research (Hsiao-Wei Wang) for the help and guidance in our approach.

# References

[Sharding FAQ](https://github.com/ethereum/wiki/wiki/Sharding-FAQ)

[Sharding Reference Spec](https://github.com/ethereum/sharding/blob/develop/docs/doc.md)

[Ethereum Sharding and Finality - Hsiao-Wei Wang](https://medium.com/@icebearhww/ethereum-sharding-and-finality-65248951f649)

[Data Availability and Erasure Coding](https://github.com/ethereum/research/wiki/A-note-on-data-availability-and-erasure-coding)

[Proof of Visibility for Data Availability](https://ethresear.ch/t/proof-of-visibility-for-data-availability/1073)

[Enforcing Windback and Proof of Custody](https://ethresear.ch/t/enforcing-windback-validity-and-availability-and-a-proof-of-custody/949)

[Fork-Free Sharding](https://ethresear.ch/t/fork-free-sharding/1058)

[Delayed State Execution](https://ethresear.ch/t/delayed-state-execution-finality-and-cross-chain-operations/987)

[State Execution Scalability and Cost Under DDoS Attacks](https://ethresear.ch/t/state-execution-scalability-and-cost-under-dos-attacks/1048)

[Guaranteed Collation Subsidies](https://ethresear.ch/t/guaranteed-collation-subsidies/1016)

[Fork Choice Rule for Collation Proposals](https://ethresear.ch/t/fork-choice-rule-for-collation-proposal-mechanisms/922)

[Model for Phase 4 Tightly-Coupled Sharding](https://ethresear.ch/t/a-model-for-stage-4-tightly-coupled-sharding-plus-full-casper/1065)

[History, State, and Asynchronous Accumulators in the Stateless Model](https://ethresear.ch/t/history-state-and-asynchronous-accumulators-in-the-stateless-model/287)
