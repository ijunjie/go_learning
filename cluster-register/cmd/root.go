package cmd

import (
	"cluster-register/infra"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

// 子命令通用 flag names
const (
	flagType           = "type"
	flagWriteToDb      = "write-to-db"
	flagIgnoreError    = "ignore-error"
	flagDbHost         = "db-host"
	flagDbPort         = "db-port"
	flagDbUsername     = "db-username"
	flagDbPassword     = "db-password"
	flagDatabase       = "database"
	flagTimeoutSeconds = "timeout-seconds"
)

// 子命令通用 param
type commonParamStruct struct {
	clusterType    string
	writeToDB      bool
	ignoreError    bool
	host           string
	port           int
	username       string
	password       string
	database       string
	timeoutSeconds int
}

func (param *commonParamStruct) checkRequired() bool {
	if param.clusterType == "" {
		log.Printf("Error: required flag(s) \"%s\" not set\n", flagType)
		return false
	}
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

func (param *commonParamStruct) toDBConnectInfo() *infra.DBConnectInfo {
	return &infra.DBConnectInfo{
		Host:     param.host,
		Port:     param.port,
		Username: param.username,
		Password: param.password,
		Database: param.database,
	}
}

var commonParam *commonParamStruct

var cfgFile string

var rootCmd = &cobra.Command{
	Use:     "cluster-register",
	Short:   "Register a cluster to resource-manager.",
	Long:    `cluster-reister is a CLI application for resource-manager to register clusters.`,
	Version: "v2021.07.1",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if ok := commonParam.checkRequired(); !ok {
			os.Exit(0)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	commonParam = &commonParamStruct{}

	rootCmd.PersistentFlags().StringVarP(&commonParam.clusterType, flagType, "", "",
		"Cluster type: online|offline(|endpoint if k8s) (required)")

	rootCmd.PersistentFlags().BoolVarP(&commonParam.writeToDB, flagWriteToDb, "", false,
		"Whether write to DB, default to false.")

	rootCmd.PersistentFlags().BoolVarP(&commonParam.ignoreError, flagIgnoreError, "", false,
		"Whether ignore errors, default to false.")

	rootCmd.PersistentFlags().StringVarP(&commonParam.host, flagDbHost, "", "",
		"DB host (required if write-to-db is true)")
	rootCmd.PersistentFlags().IntVarP(&commonParam.port, flagDbPort, "", 0,
		"DB port (required if write-to-db is true)")
	rootCmd.PersistentFlags().StringVarP(&commonParam.username, flagDbUsername, "", "",
		"DB username (required if write-to-db is true)")
	rootCmd.PersistentFlags().StringVarP(&commonParam.password, flagDbPassword, "", "",
		"DB password (required if write-to-db is true)")
	rootCmd.PersistentFlags().StringVarP(&commonParam.database, flagDatabase, "", "resource_manager",
		"DB database")
	rootCmd.PersistentFlags().IntVarP(&commonParam.timeoutSeconds, flagTimeoutSeconds, "", 5,
		"Http request timeout seconds. Default to 5.")
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
