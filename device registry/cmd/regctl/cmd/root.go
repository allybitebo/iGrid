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
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var address string
var port string
var uuid string
var password string

var longDesc = `
regctl is a swiss knife command-line tool for igrid system. It is mostly
used for testing and management as well as maintenance. It contains just
five main commands, which are add, get,list, delete and update that are
used for management of nodes, users and regions.
`

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "regctl",
	Short: "regctl [add |get |list |delete |update] [users |nodes |regions]",
	Long:  longDesc,
	/*Run: func(cmd *cobra.Command, args []string) {
		logUsage(cmd.Short)
	},*/

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logError(err)
		os.Exit(1)
	}

}

func init() {

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&address, "address", "http://localhost", "the address of regsvc")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.regctl.yaml)")
	rootCmd.PersistentFlags().StringVar(&port, "port", ":8080", "regsvc port")
	rootCmd.PersistentFlags().StringVar(&uuid, "uuid", "", "user unique identifier")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "user password")

	cli, err := cli(address, port)

	if err != nil {
		logError(err)
	}

	addCmd := NewAddCmd(cli)
	listCmd := NewListCmd(cli)
	getCmd := NewGetCmd(cli)
	deleteCmd := NewDeleteCmd(cli)
	updateCmd := NewUpdateCmd(cli)
	dbCmd := NewDBCmd()

	rootCmd.AddCommand(addCmd, listCmd, getCmd, deleteCmd, updateCmd, dbCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			logError(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".regctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".regctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
