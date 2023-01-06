package root

import (
	"fmt"
	"os"

	"github.com/vishnushankarsg/geth/command/backup"
	"github.com/vishnushankarsg/geth/command/genesis"
	"github.com/vishnushankarsg/geth/command/helper"
	"github.com/vishnushankarsg/geth/command/ibft"
	"github.com/vishnushankarsg/geth/command/license"
	"github.com/vishnushankarsg/geth/command/loadbot"
	"github.com/vishnushankarsg/geth/command/monitor"
	"github.com/vishnushankarsg/geth/command/peers"
	"github.com/vishnushankarsg/geth/command/secrets"
	"github.com/vishnushankarsg/geth/command/server"
	"github.com/vishnushankarsg/geth/command/status"
	"github.com/vishnushankarsg/geth/command/txpool"
	"github.com/vishnushankarsg/geth/command/version"
	"github.com/vishnushankarsg/geth/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "geth is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
