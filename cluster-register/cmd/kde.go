package cmd

import (
	"cluster-register/infra"
	"cluster-register/kde"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

type kdeConfigParam = struct {
	host     string
	port     int
	username string
	password string
}

var kdeConfig = &kdeConfigParam{}
var kdeType string
var writeToDB = false
var dbConfig = &infra.DBConfigParam{}
var ignoreError = false

var kdeCmd = &cobra.Command{
	Use:   "kde --host=kde-host --port=8080 --username=admin --password=admin --type={online|offline} [...other flags]",
	Short: "Register KDE to resource-manager.",
	Example: `
Write to DB:

./cluster-register kde --host=192.168.10.10 --port=8080 --username=admin --password=admin --type=online \ 
--write-to-db --db-host=192.168.10.10 --db-port=3306 --db-username=myusername --db-password=mypassword

Show info only:

./cluster-register kde --host=192.168.10.10 --port=8080 --username=admin --password=admin --type=online`,
	Long: `Register KDE(ambari) as a cluster_config to resource-manager.`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if kdeType != "online" && kdeType != "offline" {
			log.Printf("Error: flag \"%s\" value should be \"online\" or \"offline\"\n", "type")
			return
		}
		if writeToDB {
			undefinedFlag := 0
			if dbConfig.Host == "" {
				log.Printf("Error: required flag(s) \"%s\" not set\n", "db-host")
				undefinedFlag = undefinedFlag + 1
			}
			if dbConfig.Port == 0 {
				log.Printf("Error: required flag(s) \"%s\" not set\n", "db-port")
				undefinedFlag = undefinedFlag + 1
			}
			if dbConfig.Username == "" {
				log.Printf("Error: required flag(s) \"%s\" not set\n", "db-username")
				undefinedFlag = undefinedFlag + 1
			}
			if dbConfig.Password == "" {
				log.Printf("Error: required flag(s) \"%s\" not set\n", "db-password")
				undefinedFlag = undefinedFlag + 1
			}
			if undefinedFlag > 0 {
				log.Println()
				fmt.Printf("If \"%s\" is true, flags below are required: \n", "write-to-db")
				fmt.Println("\t --db-host")
				fmt.Println("\t --db-port")
				fmt.Println("\t --db-username")
				fmt.Println("\t --db-password")
				return
			}
		}

		// 不需要 http://, e.g: 10.69.75.29:8080

		info, err := kde.KdeInfo(kdeConfig.host, kdeConfig.port, kdeConfig.username, kdeConfig.password, kdeType)
		if err != nil {
			log.Fatal(err)
		}

		json, _ := json.MarshalIndent(*info, "", "  ")
		log.Println("kde info: ")
		fmt.Println(string(json))

		quarterMem := info.MemGB / 4
		badRatio := quarterMem < info.Vcores
		if badRatio {
			if !ignoreError {
				log.Fatalf("\033[1;37;41m%s\033[0m\n",
					"Error: TotalMemoryGB should be at least 4 times or more of vcores!")
			}
		}

		if writeToDB {
			log.Println("Write to resource-manager db...")
			data := &infra.ClusterConfigInsert{
				ClusterName:    info.ClusterName,
				Host:           info.Host,
				RootCuNum:      info.Vcores,
				BasicKey:       info.BasicKey,
				RmHost:         info.YarnResourceManager,
				NmHost:         info.NameNodeHost,
				ClusterType:    info.Env,
				HadoopMasterIp: info.HadoopMasterIp,
			}

			id, err := infra.InsertClusterConfig(dbConfig, data)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("\033[1;37;42m%s\033[0m\n", fmt.Sprintf("SUCCESS: Inserted ID=%d", id))
		}

	},
}

func init() {
	rootCmd.AddCommand(kdeCmd)

	kdeCmd.Flags().StringVarP(&kdeConfig.host, "host", "", "", "KDE host (required)")
	kdeCmd.Flags().IntVarP(&kdeConfig.port, "port", "", 0, "KDE port (required)")
	kdeCmd.Flags().StringVarP(&kdeConfig.username, "username", "", "",
		"KDE username (required)")
	kdeCmd.Flags().StringVarP(&kdeConfig.password, "password", "", "",
		"KDE password (required)")

	kdeCmd.Flags().StringVarP(&kdeType, "type", "", "",
		"KDE type: online/offline (required)")

	kdeCmd.Flags().BoolVarP(&writeToDB, "write-to-db", "", false,
		"Whether write to DB, default to false.")

	kdeCmd.Flags().BoolVarP(&ignoreError, "ignore-error", "", false,
		"Whether ignore errors, default to false.")

	kdeCmd.Flags().StringVarP(&dbConfig.Host, "db-host", "", "",
		"DB host (required if write-to-db is true)")
	kdeCmd.Flags().IntVarP(&dbConfig.Port, "db-port", "", 0,
		"DB port (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbConfig.Username, "db-username", "", "",
		"DB username (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbConfig.Password, "db-password", "", "",
		"DB password (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbConfig.Database, "database", "", "resource_manager",
		"DB database")

	_ = kdeCmd.MarkFlagRequired("host")
	_ = kdeCmd.MarkFlagRequired("port")
	_ = kdeCmd.MarkFlagRequired("username")
	_ = kdeCmd.MarkFlagRequired("password")
	_ = kdeCmd.MarkFlagRequired("type")
}
