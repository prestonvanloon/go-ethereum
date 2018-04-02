package sharding

import (
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Collation struct {
	header       *CollationHeader
	transactions []*types.Transaction
}

type CollationHeader struct {
	shardID              *big.Int			//the shard ID of the shard
	parentHash  		 *common.Hash		//the hash of the parent collation
	chunkRoot            *common.Hash		//the root of the chunk tree which identifies collation body
	period 				 *big.Int			//the period number in which collation to be included
	proposerAddress      *common.Address	//address of the collation proposer
	proposerBid          *big.Int			//the reward from proposer to collator for a winning proposal
	proposerSignature    []byte				//the proposer's signature as part of a proposal
}

func (c *Collation) Header() *CollationHeader           { return c.header }
func (c *Collation) Transactions() []*types.Transaction { return c.transactions }
func (c *Collation) ShardID() *big.Int           { return c.header.shardID }
func (c *Collation) ParentHash() *common.Hash { return c.header.parentHash }
func (c *Collation) Period() *big.Int           { return c.header.period }
func (c *Collation) ProposerAddress() *common.Address { return c.header.proposerAddress }
func (c *Collation) ProposerBid() *big.Int { return c.header.proposerBid }
func (c *Collation) ProposerSignature() []byte{ return c.header.proposerSignature }


func (c *Collation) SetHeader(h *CollationHeader) { c.header = h }
func (c *Collation) AddTransaction(tx *types.Transaction) {
	// TODO: Check transaction does not exceed gas limit
	c.transactions = append(c.transactions, tx)
}

func (c *Collation) GasUsed() *big.Int {
	g := uint64(0)
	for _, tx := range c.transactions {
		if g > math.MaxUint64-(g+tx.Gas()) {
			g = math.MaxUint64
			break
		}
		g += tx.Gas()
	}
	return big.NewInt(0).SetUint64(g)
}
