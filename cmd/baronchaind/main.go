package main

import (
	"fmt"
	"os"

	"github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"baronchain/app"
	"baronchain/cmd/baronchaind/cmd"
)

func main() {
	// Set prefixes
	accountAddressPrefix := app.AccountAddressPrefix
	
	// Set address prefixes
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(accountAddressPrefix, accountAddressPrefix+"pub")
	config.SetBech32PrefixForValidator(accountAddressPrefix+"valoper", accountAddressPrefix+"valoperpub")
	config.SetBech32PrefixForConsensusNode(accountAddressPrefix+"valcons", accountAddressPrefix+"valconspub")
	config.Seal()

	// Initialize the root command
	rootCmd, _ := cmd.NewRootCmd()

	// Set flags
	rootCmd.PersistentFlags().String(flags.FlagChainID, "", "Chain ID of tendermint node")

	// Prepare and execute the root command
	executor := cli.PrepareBaseCmd(rootCmd, "BC", app.DefaultNodeHome)
	err := svrcmd.Execute(executor, app.DefaultNodeHome)

	if err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			fmt.Fprintf(os.Stderr, "Error executing root command: %v\n", err)
			os.Exit(e.Code)
		default:
			fmt.Fprintf(os.Stderr, "Error executing root command: %v\n", err)
			os.Exit(1)
		}
	}
}
