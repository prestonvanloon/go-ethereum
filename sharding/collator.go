package sharding

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
)

<<<<<<< HEAD
=======
type collatorClient interface {
	Account() *accounts.Account
	ChainReader() ethereum.ChainReader
	VMCCaller() *contracts.VMCCaller
}

>>>>>>> 6ae47790b... refactor account, and unlock account when client is created
// SubscribeBlockHeaders checks incoming block headers and determines if
// we are an eligible proposer for collations. Then, it finds the pending tx's
// from the running geth node and sorts them by descending order of gas price,
// eliminates those that ask for too much gas, and routes them over
// to the VMC to create a collation
func subscribeBlockHeaders(c *Client) error {
	headerChan := make(chan *types.Header, 16)

<<<<<<< HEAD
	_, err := c.client.SubscribeNewHead(context.Background(), headerChan)
=======
	account := c.Account()

	_, err := c.ChainReader().SubscribeNewHead(context.Background(), headerChan)
>>>>>>> 6ae47790b... refactor account, and unlock account when client is created
	if err != nil {
		return fmt.Errorf("unable to subscribe to incoming headers. %v", err)
	}

	log.Info("listening for new headers...")

	for {
		select {
		case head := <-headerChan:
			// Query the current state to see if we are an eligible proposer
			log.Info(fmt.Sprintf("received new header %v", head.Number.String()))
			// TODO: Only run this code on certain periods?
			err := watchShards(c, head)
			if err != nil {
				return fmt.Errorf("unable to watch shards. %v", err)
			}
		}
	}
}

<<<<<<< HEAD
// watchShards checks if we are an eligible proposer for collation for
// the available shards in the VMC. The function calls getEligibleProposer from
// the VMC and proposes a collation if conditions are met
func watchShards(c *Client, head *types.Header) error {

	accounts := c.keystore.Accounts()
	if len(accounts) == 0 {
		return fmt.Errorf("no accounts found")
	}
=======
// checkShardsForProposal checks if we are an eligible proposer for
// collation for the available shards in the VMC. The function calls
// getEligibleProposer from the VMC and proposes a collation if
// conditions are met
func checkShardsForProposal(c collatorClient, head *types.Header) error {
	account := c.Account()
>>>>>>> 6ae47790b... refactor account, and unlock account when client is created

	if err := c.unlockAccount(accounts[0]); err != nil {
		return fmt.Errorf("cannot unlock account. %v", err)
	}

	log.Info(fmt.Sprint("watching shards..."))
	s := 0
	for s < shardCount {
		// Checks if we are an eligible proposer according to the VMC
		ops := bind.CallOpts{}
		period := head.Number.Div(head.Number, big.NewInt(int64(periodLength)))
		addr, err := c.vmc.VMCCaller.GetEligibleProposer(&ops, big.NewInt(int64(s)), period)

		// If output is non-empty and the addr == coinbase
		if err == nil && addr == accounts[0].Address {
			log.Info(fmt.Sprintf("selected as collator on shard %d", s))
			err := proposeCollation(s)
			if err != nil {
				return fmt.Errorf("could not propose collation. %v", err)
			}
		}

<<<<<<< HEAD
		s++
	}
=======
// isAccountInValidatorSet checks if the client is in the validator pool because
// we can't guarantee our tx for deposit will be in the next block header we receive.
// The function calls IsValidatorDeposited from the VMC and returns true if
// the client is in the validator pool
func isAccountInValidatorSet(c collatorClient) (bool, error) {
	account := c.Account()
>>>>>>> 6ae47790b... refactor account, and unlock account when client is created

	return nil
}

<<<<<<< HEAD
func proposeCollation() error {
=======
// proposeCollation interacts with the VMC directly to add a collation header
func proposeCollation(shardID int) error {
	// TODO: Adds a collation header to the VMC with the following fields:
	// [
	//  shard_id: uint256,
	//  expected_period_number: uint256,
	//  period_start_prevhash: bytes32,
	//  parent_hash: bytes32,
	//  transactions_root: bytes32,
	//  coinbase: address,
	//  state_root: bytes32,
	//  receipts_root: bytes32,
	//  number: uint256,
	//  sig: bytes
	// ]
	//
	// Before calling this, we would need to have access to the state of
	// the period_start_prevhash. Refer to the comments in:
	// https://github.com/ethereum/py-evm/issues/258#issuecomment-359879350
	//
	// This function will call FetchCandidateHead() of the VMC to obtain
	// more necessary information.
	//
	// This functions will fetch the transactions in the txpool and and apply
	// them to finish up the collation. It will then need to broadcast the
	// collation to the main chain using JSON-RPC.
	log.Info(fmt.Sprint("propose collation called"))
>>>>>>> e754f7c3c... propose collation called on geteligibleproposer
	return nil
}
