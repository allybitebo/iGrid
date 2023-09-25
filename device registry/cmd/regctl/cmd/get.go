package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func NewGetCmd(cli CLI) *cobra.Command {
	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "regctl get users --id <user-id>",
		Long:  "get a user by specifying his/her id",
		Run:   cli.UsersCmd(context.Background(), Get),
	}

	usersCmd.Flags().String("id", "", "user id")

	nodesCmd := &cobra.Command{
		Use:   "nodes",
		Short: "regctl get nodes --id <node-id>",
		Long:  "get a node by specifying its id",
		Run:   cli.NodesCmd(context.Background(), Get),
	}

	nodesCmd.Flags().String("id", "", "user id")

	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get (users |nodes |regions) <id>",
		Long:  "get a certain entity by specifying its id",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		},
	}
	getCmd.AddCommand(usersCmd, nodesCmd)
	return getCmd
}
