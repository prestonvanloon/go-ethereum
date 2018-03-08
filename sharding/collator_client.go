package sharding

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/sharding/contracts"
	cli "gopkg.in/urfave/cli.v1"
)

const (
	clientIdentifier = "geth" // Used to determine the ipc name.
)

<<<<<<< HEAD
// Client for sharding. Communicates to geth node via JSON RPC.
type Client struct {
=======
// Client for Collator. Communicates to Geth node via JSON RPC.
type collatorClient struct {
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
	endpoint string             // Endpoint to JSON RPC
	client   *ethclient.Client  // Ethereum RPC client.
	keystore *keystore.KeyStore // Keystore containing the single signer
	ctx      *cli.Context       // Command line context
	vmc      *contracts.VMC     // The deployed validator management contract
}

<<<<<<< HEAD
// MakeShardingClient for interfacing with geth full node.
func MakeShardingClient(ctx *cli.Context) *Client {
=======
// MakeCollatorClient for interfacing with Geth full node.
func MakeCollatorClient(ctx *cli.Context) *collatorClient {
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
	path := node.DefaultDataDir()
	if ctx.GlobalIsSet(utils.DataDirFlag.Name) {
		path = ctx.GlobalString(utils.DataDirFlag.Name)
	}

	endpoint := ctx.Args().First()
	if endpoint == "" {
		endpoint = fmt.Sprintf("%s/%s.ipc", path, clientIdentifier)
	}
	if ctx.GlobalIsSet(utils.IPCPathFlag.Name) {
		endpoint = ctx.GlobalString(utils.IPCPathFlag.Name)
	}

	config := &node.Config{
		DataDir: path,
	}

	scryptN, scryptP, keydir, err := config.AccountConfig()
	if err != nil {
		panic(err) // TODO(prestonvanloon): handle this
	}
	ks := keystore.NewKeyStore(keydir, scryptN, scryptP)

<<<<<<< HEAD
	return &Client{
=======
	return &collatorClient{
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
		endpoint: endpoint,
		keystore: ks,
		ctx:      ctx,
	}
}

<<<<<<< HEAD
// Start the sharding client.
// * Connects to node.
// * Verifies or deploys the validator management contract.
func (c *Client) Start() error {
	log.Info("Starting sharding client")
=======
// Start the collator client.
// * Connects to Geth node.
// * Verifies or deploys the sharding manager contract.
func (c *collatorClient) Start() error {
	log.Info("Starting collator client")
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
	rpcClient, err := dialRPC(c.endpoint)
	if err != nil {
		return err
	}
	c.client = ethclient.NewClient(rpcClient)
	defer rpcClient.Close()
	if err := initVMC(c); err != nil {
		return err
	}

	// Deposit 100ETH into the validator set in the VMC. Checks if account
	// is already a validator in the VMC (in the case the client restarted).
	// Once that's done we can subscribe to block headers
	if err := initVMCValidator(c); err != nil {
		return err
	}

	// Listens to block headers from the geth node and if we are an eligible
	// proposer, we fetch pending transactions and propose a collation
	if err := subscribeBlockHeaders(c); err != nil {
		return err
	}
	return nil
}

<<<<<<< HEAD
// Wait until sharding client is shutdown.
func (c *Client) Wait() {
	// TODO: Blocking lock.
}

// dialRPC endpoint to node.
func dialRPC(endpoint string) (*rpc.Client, error) {
	if endpoint == "" {
		endpoint = node.DefaultIPCEndpoint(clientIdentifier)
	}
	return rpc.Dial(endpoint)
}

// UnlockAccount will unlock the specified account using utils.PasswordFileFlag or empty string if unset.
func (c *Client) unlockAccount(account accounts.Account) error {
=======
// Wait until collator client is shutdown.
func (c *collatorClient) Wait() {
	log.Info("Sharding client has been shutdown...")
}

// WatchCollationHeaders checks the logs for add_header func calls
// and updates the head collation of the client. We can probably store
// this as a property of the client struct
func (c *collatorClient) WatchCollationHeaders() {

}

// UnlockAccount will unlock the specified account using utils.PasswordFileFlag or empty string if unset.
func (c *collatorClient) unlockAccount(account accounts.Account) error {
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
	pass := ""

	if c.ctx.GlobalIsSet(utils.PasswordFileFlag.Name) {
		blob, err := ioutil.ReadFile(c.ctx.GlobalString(utils.PasswordFileFlag.Name))
		if err != nil {
			return fmt.Errorf("unable to read account password contents in file %s. %v", utils.PasswordFileFlag.Value, err)
		}
		// TODO: Use bufio.Scanner or other reader that doesn't include a trailing newline character.
		pass = strings.Trim(string(blob), "\n") // Some text files end in new line, remove with strings.Trim.
	}

	return c.keystore.Unlock(account, pass)
}
<<<<<<< HEAD
=======

func (c *collatorClient) createTXOps(value *big.Int) (*bind.TransactOpts, error) {
	account := c.Account()

	return &bind.TransactOpts{
		From:  account.Address,
		Value: value,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			networkID, err := c.client.NetworkID(context.Background())
			if err != nil {
				return nil, fmt.Errorf("unable to fetch networkID: %v", err)
			}
			return c.keystore.SignTx(*account, tx, networkID /* chainID */)
		},
	}, nil
}

// Account to use for sharding transactions.
func (c *collatorClient) Account() *accounts.Account {
	accounts := c.keystore.Accounts()

	return &accounts[0]
}

// ChainReader for interacting with the chain.
func (c *collatorClient) ChainReader() ethereum.ChainReader {
	return ethereum.ChainReader(c.client)
}

// Client to interact with ethereum node.
func (c *collatorClient) Client() *ethclient.Client {
	return c.client
}

// SMCCaller to interact with the sharding manager contract.
func (c *collatorClient) SMCCaller() *contracts.SMCCaller {
	return &c.smc.SMCCaller
}

// dialRPC endpoint to node.
func dialRPC(endpoint string) (*rpc.Client, error) {
	if endpoint == "" {
		endpoint = node.DefaultIPCEndpoint(clientIdentifier)
	}
	return rpc.Dial(endpoint)
}
>>>>>>> 865b8d1d0... fixed all the typos during integration, manually tested
