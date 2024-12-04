package chain

import (
	"time"

	"go.thebigfile.com/core/consensus"
	"go.thebigfile.com/core/types"
)

func parseAddr(s string) types.Address {
	addr, err := types.ParseAddress(s)
	if err != nil {
		panic(err)
	}
	return addr
}

// Mainnet returns the network parameters and genesis block for the mainnet Sia
// blockchain.
func Mainnet() (*consensus.Network, types.Block) {
	n := &consensus.Network{
		Name: "mainnet",

		InitialCoinbase: types.Siacoins(300000),
		MinimumCoinbase: types.Siacoins(30000),
		InitialTarget:   types.BlockID{4: 32},
		BlockInterval:   10 * time.Minute,
		MaturityDelay:   144,
	}
	n.HardforkDevAddr.Height = 10000
	n.HardforkDevAddr.OldAddress = types.Address{}
	n.HardforkDevAddr.NewAddress = parseAddr("c10f0fc99f2ac2d6268a3427394cef5f419e858d4152309e9f8d4abbe8c495eeb804e05a961c")

	n.HardforkTax.Height = 21000

	n.HardforkStorageProof.Height = 100000

	n.HardforkOak.Height = 135000
	n.HardforkOak.FixHeight = 139000
	n.HardforkOak.GenesisTimestamp = time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC)

	n.HardforkASIC.Height = 179000
	n.HardforkASIC.OakTime = 120000 * time.Second
	n.HardforkASIC.OakTarget = types.BlockID{8: 32}

	n.HardforkFoundation.Height = 298000
	n.HardforkFoundation.PrimaryAddress = parseAddr("c10f0fc99f2ac2d6268a3427394cef5f419e858d4152309e9f8d4abbe8c495eeb804e05a961c")
	n.HardforkFoundation.FailsafeAddress = parseAddr("c10f0fc99f2ac2d6268a3427394cef5f419e858d4152309e9f8d4abbe8c495eeb804e05a961c")

	n.HardforkV2.AllowHeight = 1000000   // TBD
	n.HardforkV2.RequireHeight = 1025000 // ~six months later

	b := types.Block{
		Timestamp: n.HardforkOak.GenesisTimestamp,
		Transactions: []types.Transaction{{
			SiafundOutputs: []types.SiafundOutput{
				{Address: parseAddr("c10f0fc99f2ac2d6268a3427394cef5f419e858d4152309e9f8d4abbe8c495eeb804e05a961c"), Value: 10000},
			},
		}},
	}

	return n, b
}

// TestnetZen returns the chain parameters and genesis block for the "Zen"
// testnet chain.
func TestnetZen() (*consensus.Network, types.Block) {
	n := &consensus.Network{
		Name: "zen",

		InitialCoinbase: types.Siacoins(300000),
		MinimumCoinbase: types.Siacoins(300000),
		InitialTarget:   types.BlockID{3: 1},
		BlockInterval:   10 * time.Minute,
		MaturityDelay:   144,
	}

	n.HardforkDevAddr.Height = 1
	n.HardforkDevAddr.OldAddress = types.Address{}
	n.HardforkDevAddr.NewAddress = types.Address{}

	n.HardforkTax.Height = 2

	n.HardforkStorageProof.Height = 5

	n.HardforkOak.Height = 10
	n.HardforkOak.FixHeight = 12
	n.HardforkOak.GenesisTimestamp = time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC)

	n.HardforkASIC.Height = 20
	n.HardforkASIC.OakTime = 10000 * time.Second
	n.HardforkASIC.OakTarget = types.BlockID{3: 1}

	n.HardforkFoundation.Height = 30
	n.HardforkFoundation.PrimaryAddress = parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e")
	n.HardforkFoundation.FailsafeAddress = types.VoidAddress

	n.HardforkV2.AllowHeight = 112000   // March 1, 2025 @ 7:00:00 UTC
	n.HardforkV2.RequireHeight = 116380 // 1 month later

	b := types.Block{
		Timestamp: n.HardforkOak.GenesisTimestamp,
		Transactions: []types.Transaction{{
			SiacoinOutputs: []types.SiacoinOutput{{
				Address: parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e"),
				Value:   types.Siacoins(1).Mul64(1e12),
			}},
			SiafundOutputs: []types.SiafundOutput{{
				Address: parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e"),
				Value:   10000,
			}},
		}},
	}

	return n, b
}

// TestnetAnagami returns the chain parameters and genesis block for the "Anagami"
// testnet chain.
func TestnetAnagami() (*consensus.Network, types.Block) {
	// use a modified version of Zen
	n, genesis := TestnetZen()

	n.Name = "anagami"
	n.HardforkOak.GenesisTimestamp = time.Date(2024, time.December, 1, 0, 0, 0, 0, time.UTC)
	n.HardforkV2.AllowHeight = 2016         // ~2 weeks in
	n.HardforkV2.RequireHeight = 2016 + 288 // ~2 days later

	n.HardforkFoundation.PrimaryAddress = parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e")
	n.HardforkFoundation.FailsafeAddress = types.VoidAddress

	// move the genesis airdrops for easier testing
	genesis.Transactions[0].SiacoinOutputs = []types.SiacoinOutput{{
		Address: parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e"),
		Value:   types.Siacoins(1).Mul64(1e12),
	}}
	genesis.Transactions[0].SiafundOutputs = []types.SiafundOutput{{
		Address: parseAddr("c1b783407f81502e0cd7bf5a12a32dd62017f30a64cd6bf86eca7a379caa4fe7fe2e356d2f4e"),
		Value:   10000,
	}}
	return n, genesis
}
