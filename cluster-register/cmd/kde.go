/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var KdeHost string
var KdePort int
var KdeUsername string
var KdePassword string
var WriteToDB bool
var MySqlHost string
var MySqlPort int
var MySqlUsername string
var MySqlPassword string

// kdeCmd represents the kde command
var kdeCmd = &cobra.Command{
	Use:   "kde",
	Short: "Register KDE to resource-manager.",
	Long:  `Register KDE(ambari) as cluster_config to resource-manager.`,
	Args: func(cmd *cobra.Command, args []string) error {

		if len(args) != 5 || len(args) != 9 {
			return errors.New("requires at least one arg")
		}

		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	//Run: func(cmd *cobra.Command, args []string) {
	//	fmt.Printf("host: %s \n", KdeHost)
	//	fmt.Printf("port: %d \n", KdePort)
	//	fmt.Printf("username: %s \n", KdeUsername)
	//	fmt.Printf("password: %s \n", KdePassword)
	//},
}

func init() {
	rootCmd.AddCommand(kdeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kdeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kdeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	kdeCmd.Flags().StringVarP(&KdeHost, "host", "", "", "KDE host (required)")
	_ = kdeCmd.MarkFlagRequired("host")
	kdeCmd.Flags().IntVarP(&KdePort, "port", "", 8080, "KDE port (required)")
	_ = kdeCmd.MarkFlagRequired("port")
	kdeCmd.Flags().StringVarP(&KdeUsername, "username", "", "ec_admin", "KDE username (required)")
	_ = kdeCmd.MarkFlagRequired("username")
	kdeCmd.Flags().StringVarP(&KdePassword, "password", "", "", "KDE password (required)")
	_ = kdeCmd.MarkFlagRequired("password")

	kdeCmd.Flags().BoolVarP(&WriteToDB, "write-to-db", "", false, "If write to DB. Default to false.")

	if WriteToDB {
		kdeCmd.Flags().StringVarP(&MySqlHost, "mysql-host", "", "", "MySQL host (required if write-to-db is true)")
		_ = kdeCmd.MarkFlagRequired("mysql-host")
		kdeCmd.Flags().IntVarP(&MySqlPort, "mysql-port", "", 3306, "MySQL port (required if write-to-db is true)")
		_ = kdeCmd.MarkFlagRequired("mysql-port")
		kdeCmd.Flags().StringVarP(&MySqlUsername, "mysql-username", "", "", "MySQL username (required if write-to-db is true)")
		_ = kdeCmd.MarkFlagRequired("mysql-username")
		kdeCmd.Flags().StringVarP(&MySqlPassword, "mysql-password", "", "", "MySQL password (required if write-to-db is true)")
		_ = kdeCmd.MarkFlagRequired("mysql-password")
	}
}
