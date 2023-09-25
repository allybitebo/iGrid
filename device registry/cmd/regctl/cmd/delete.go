package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func NewDeleteCmd(cli CLI) *cobra.Command {

	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "delete users --id <id>",
		Long:  "delete user by specifying id",
		Run:   cli.UsersCmd(context.Background(), Delete),
	}

	usersCmd.Flags().String("id", "", "user id")

	nodesCmd := &cobra.Command{
		Use:   "nodes",
		Short: "delete nodes --id <id>",
		Long:  "delete node by specifying id",
		Run:   cli.NodesCmd(context.Background(), Delete),
	}

	nodesCmd.Flags().String("id", "", "node id")

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "delete (users |nodes |regions) <id>",
		Long:  "delete by specifying id of the entity",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		},
	}

	deleteCmd.AddCommand(usersCmd, nodesCmd)

	return deleteCmd
}
