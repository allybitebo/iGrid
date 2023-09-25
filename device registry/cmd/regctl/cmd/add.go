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

func NewAddCmd(cli CLI) *cobra.Command {

	usersCmd := &cobra.Command{
		Use:   "users",
		Short: "users -n <name> -e <email> -p <password> -r <region-id>",
		Long:  `add new user to the system`,
		Run:   cli.UsersCmd(context.Background(), Add),
	}

	usersCmd.Flags().StringP("name", "n", "", "user full name")
	usersCmd.Flags().StringP("email", "e", "", "email address")
	usersCmd.Flags().StringP("password", "p", "", "passphrase")
	usersCmd.Flags().StringP("region", "r", "", "region-id")

	regionsCmd := &cobra.Command{
		Use:   "regions",
		Short: "regions --id <id> --name <name> --desc <description>",
		Long:  `add new region to the system`,
		Run:   cli.RegionsCmd(context.Background(), Add),
	}

	regionsCmd.Flags().StringP("id", "i", "", "region id")
	regionsCmd.Flags().StringP("name", "n", "", "region name")
	regionsCmd.Flags().StringP("desc", "d", "", "region description")

	nodesCmd := &cobra.Command{
		Use:     "nodes",
		Short:   "add new node",
		Long:    "add new node by specifying addr,name,region, location,type and master",
		Example: "regctl add nodes --ad <address> --name <>",
		Run:     cli.NodesCmd(context.Background(), Add),
	}

	//adr,name,region,lat,long,master,type

	nodesCmd.Flags().StringP("adr", "d", "", "mac address of the node")
	nodesCmd.Flags().StringP("name", "n", "", "name of the node")
	nodesCmd.Flags().StringP("region", "r", "", "region where the node is installed")
	nodesCmd.Flags().StringP("lat", "l", "", "latitude")
	nodesCmd.Flags().StringP("long", "g", "", "longitude")
	nodesCmd.Flags().StringP("master", "m", "", "master node")
	nodesCmd.Flags().IntP("type", "t", 0, "the type of the node")

	addCmd := &cobra.Command{
		Use:   "add",
		Short: "add (users |nodes |regions)",
		Long:  `add a new entity to the network (users |nodes |regions )`,
		Run: func(cmd *cobra.Command, args []string) {
			logUsage(cmd.Short)
		},
	}

	addCmd.AddCommand(usersCmd, regionsCmd, nodesCmd)

	return addCmd
}
