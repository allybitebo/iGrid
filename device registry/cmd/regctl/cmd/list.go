/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/spf13/cobra"
)

func NewListCmd(cli CLI) *cobra.Command {
	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "list users",
		Long:  `list all users in the systems`,
		Run:   cli.UsersCmd(context.Background(), List),
	}

	// usersCmd represents the users command
	regionsCmd := &cobra.Command{
		Use:   "regions",
		Short: "list regions",
		Long:  `list all available regions`,
		Run:   cli.RegionsCmd(context.Background(), List),
	}

	// usersCmd represents the users command
	nodesCmd := &cobra.Command{
		Use:   "nodes",
		Short: "list nodes",
		Long:  `list all available nodes`,
		Run:   cli.NodesCmd(context.Background(), List),
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list (users |nodes |regions )",
		Long:  `this command list all the available (users | nodes | regions)`,
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		},
	}

	listCmd.AddCommand(usersCmd, regionsCmd, nodesCmd)

	return listCmd
}
