package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func NewUpdateCmd(cli CLI) *cobra.Command {

	usersCmd := &cobra.Command{
		Use:     "users",
		Short:   "update users (group | region)",
		Long:    "used to update either the region or group or both",
		Example: "regctl update users --id \"uegeteg\" -r \"AB001\" -g 1",
		Run:     cli.UsersCmd(context.Background(), Update),
	}

	usersCmd.Flags().String("id", "", "user id")
	usersCmd.Flags().IntP("group", "g", 0, "new user group")
	usersCmd.Flags().StringP("region", "r", "", "new region-id")

	updateCmd := &cobra.Command{
		Use:     "update",
		Short:   "update (user |node |region)",
		Long:    "used for modifying the details of the entity in igrid system",
		Example: "regctl update [users |nodes |regions]",
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Example)
		},
	}

	updateCmd.AddCommand(usersCmd)

	return updateCmd
}
