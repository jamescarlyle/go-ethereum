package aggregator

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/aggregator/rollupContract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
)

const (
	// rollupChanSize is the size of channel listening to ChainHeadEvent.
	rollupChanSize = 10
)

// Service implements an Obscuro aggregator daemon that pushes rollups to the L1 node connection.
type Service struct {
	backend ethapi.Backend
	client  ethclient.Client // Client to rollup through.

	connectionUrl   string         // Connection string for L1 node.
	contractAddress common.Address // Address of Rollup contract.

	l1BlockSub ethereum.Subscription
	headerCh   chan *types.Header

	privateKey  ecdsa.PrivateKey // L2 node private key.
	fromAddress common.Address   // L2 node address used for rollup submission.
	instance    *rollupContract.RollupContract
	auth        *bind.TransactOpts
	rollupCh    chan core.ChainHeadEvent
	rollupSub   event.Subscription
}

// New returns a aggregation service ready for aggregating.
func New(node *node.Node, backend ethapi.Backend, engine consensus.Engine, connectionUrl string, contractAddress common.Address) error {
	aggregator := &Service{
		backend:         backend,
		connectionUrl:   connectionUrl,
		contractAddress: contractAddress,
	}

	node.RegisterLifecycle(aggregator)
	return nil
}

func (s *Service) Start() error {
	client, err := ethclient.Dial(s.connectionUrl)
	// client, err := ethclient.Dial("wss://ropsten.infura.io/ws/v3/cb80549cbc6b4e3fa00bfa9771aa09b1")
	if err != nil {
		log.Error(err.Error())
	}

	s.client = *client
	s.headerCh = make(chan *types.Header)
	s.l1BlockSub, err = client.SubscribeNewHead(context.Background(), s.headerCh)
	defer s.l1BlockSub.Unsubscribe()
	if err != nil {
		log.Error(err.Error())
	}
	s.rollupCh = make(chan core.ChainHeadEvent, rollupChanSize)
	s.rollupSub = s.backend.SubscribeChainHeadEvent(s.rollupCh)
	defer s.rollupSub.Unsubscribe()

	// Obtain credentials for submission.
	privateKey, err := crypto.HexToECDSA("dbd6d38466c733c65194a0b1b639e54aaf7c77bc27d35a60ba9bd81e9db75f5a")
	if err != nil {
		log.Error("Can't obtain private key for rollup", "Error", err)
	}
	s.privateKey = *privateKey
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error("Error casting public key to ECDSA")
	}
	s.fromAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	// s.contractAddress = common.HexToAddress("90C2193a0DD03710f67f9D22CD060430fc15AC72")
	instance, err := rollupContract.NewRollupContract(s.contractAddress, &s.client)
	if err != nil {
		log.Error("Can't store", "Error", err)
	}
	s.instance = instance

	// s.auth = bind.NewKeyedTransactor(&s.privateKey)
	chainID, _ := client.ChainID(context.Background())
	s.auth, _ = bind.NewKeyedTransactorWithChainID(&s.privateKey, chainID)
	s.auth.Value = big.NewInt(0)     // in wei
	s.auth.GasLimit = uint64(300000) // in units

	go s.loop(s.headerCh, s.l1BlockSub, s.rollupCh, s.rollupSub)
	log.Info("Aggregator daemon started")
	return nil
}

func (s *Service) Stop() error {
	s.l1BlockSub.Unsubscribe()
	s.rollupSub.Unsubscribe()
	log.Info("Aggregator daemon stopped")
	return nil
}

func (s *Service) loop(headerCh chan *types.Header, l1BlockSub ethereum.Subscription, rollupCh chan core.ChainHeadEvent, rollupSub event.Subscription) {
	for {
		select {
		case err := <-l1BlockSub.Err():
			log.Error(err.Error())
		case header := <-headerCh:
			block, err := s.client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Error(err.Error())
			}
			log.Info("L1 block received:", "L1 block hash", block.Hash().Hex())
			ethBackend, ok := s.backend.(*eth.EthAPIBackend)
			if !ok {
				log.Error("Ethereum service not running")
			}
			ethBackend.Miner().GenerateRollup(block.Hash())
		case rollupEv := <-rollupCh:
			if rollupEv.Block.ReceivedFrom == nil { // Rollup should be submitted by the node that generated it.
				log.Info("Latest L2 rollup hash:", "L2 rollup hash", rollupEv.Block.Hash())
				s.submitRollup(rollupEv.Block.Hash(), rollupEv.Block.ParentHash(), rollupEv.Block.Transactions())
			}
		case err := <-rollupSub.Err():
			log.Error(err.Error())
		}
	}

}

func (s *Service) submitRollup(rollupHash common.Hash, parentHash common.Hash, rollupTransactions types.Transactions) {
	nonce, err := s.client.PendingNonceAt(context.Background(), s.fromAddress)
	if err != nil {
		log.Error("Can't get nonce for address", "Error", err)
	}

	gasPrice, err := s.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error("Can't get gas price", "Error", err)
	}

	s.auth.Nonce = big.NewInt(int64(nonce))
	s.auth.GasPrice = gasPrice

	transactions := [][]byte{}
	for _, r := range rollupTransactions {
		transaction, _ := r.MarshalBinary()
		transactions = append(transactions, transaction)
	}
	tx, err := s.instance.PostRollup(s.auth, rollupHash, parentHash, big.NewInt(int64(nonce)), transactions)
	if err != nil {
		log.Error("Can't submit rollup", "Error", err)
	}
	log.Info("Rollup submitted", "Transaction hash", tx.Hash().Hex(), "Rollup hash", rollupHash, "Parent hash", parentHash)
}
