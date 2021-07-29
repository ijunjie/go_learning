package cmd

import (
	"cluster-register/infra"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/spf13/viper"
)

// 子命令通用 flag names
const (
	flagType        = "type"
	flagWriteToDb   = "write-to-db"
	flagIgnoreError = "ignore-error"
	flagDbHost      = "db-host"
	flagDbPort      = "db-port"
	flagDbUsername  = "db-username"
	flagDbPassword  = "db-password"
	flagDatabase    = "database"
)

// 子命令通用 param
type dbParamStruct struct {
	writeToDB   bool
	ignoreError bool
	host        string
	port        int
	username    string
	password    string
	database    string
}

func (param *dbParamStruct) checkRequired() bool {
	if param.writeToDB {
		undefinedFlag := 0
		if param.host == "" {
			log.Printf("Error: required flag(s) \"%s\" not set\n", flagDbHost)
			undefinedFlag = undefinedFlag + 1
		}
		if param.port == 0 {
			log.Printf("Error: required flag(s) \"%s\" not set\n", flagDbPort)
			undefinedFlag = undefinedFlag + 1
		}
		if param.username == "" {
			log.Printf("Error: required flag(s) \"%s\" not set\n", flagDbUsername)
			undefinedFlag = undefinedFlag + 1
		}
		if param.password == "" {
			log.Printf("Error: required flag(s) \"%s\" not set\n", flagDbPassword)
			undefinedFlag = undefinedFlag + 1
		}
		if undefinedFlag > 0 {
			log.Println()
			fmt.Printf("If \"%s\" is true, flags below are required: \n", flagWriteToDb)
			fmt.Printf("\t --%s\n", flagDbHost)
			fmt.Printf("\t --%s\n", flagDbPort)
			fmt.Printf("\t --%s\n", flagDbUsername)
			fmt.Printf("\t --%s\n", flagDbPassword)
			return false
		} else {
			return true
		}
	}
	return true
}

func (param *dbParamStruct) toDBConnectInfo() *infra.DBConnectInfo {
	return &infra.DBConnectInfo{
		Host:     param.host,
		Port:     param.port,
		Username: param.username,
		Password: param.password,
		Database: param.database,
	}
}

var dbParam *dbParamStruct

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cluster-register",
	Short: "Register a cluster to resource-manager.",
	Long:  `cluster-reister is a CLI application for resource-manager to register clusters.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	//rootCmd.PersistentFlags().StringVar(
	//	&cfgFile,
	//	"config",
	//	"",
	//	"config file (default is $HOME/.cluster-register.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cluster-register" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cluster-register")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
