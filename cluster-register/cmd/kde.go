package cmd

import (
	"cluster-register/infra"
	"cluster-register/kde"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

const (
	flagHost     = "host"
	flagPort     = "port"
	flagUsername = "username"
	flagPassword = "password"
)

type kdeParamStruct struct {
	host     string
	port     int
	username string
	password string
	kdeType  string
}

func (param *kdeParamStruct) toKdeInfoRequest() *kde.KdeInfoRequest {
	return &kde.KdeInfoRequest{
		KdeHost:  param.host,
		KdePort:  param.port,
		Username: param.username,
		Password: param.password,
		KdeType:  param.kdeType,
	}
}

var kdeParam *kdeParamStruct

var kdeCmd = &cobra.Command{
	Use:   "kde",
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
		if kdeParam.kdeType != "online" && kdeParam.kdeType != "offline" {
			log.Printf("Error: flag \"%s\" value should be \"online\" or \"offline\"\n", flagType)
			return
		}
		ok := dbParam.checkRequired()
		if !ok {
			return
		}

		// 不需要 http://, e.g: 10.69.75.29:8080
		kdeInfoRequest := kdeParam.toKdeInfoRequest()
		info, err := kde.KdeInfo(kdeInfoRequest)
		if err != nil {
			log.Fatal(err)
		}

		json, _ := json.MarshalIndent(*info, "", "  ")
		log.Println("kde info: ")
		fmt.Println(string(json))

		quarterMem := info.MemGB / 4
		badRatio := quarterMem < info.Vcores
		if badRatio {
			if !dbParam.ignoreError {
				log.Fatalf("\033[1;37;41m%s\033[0m\n",
					"Error: TotalMemoryGB should be at least 4 times or more of vcores!")
			}
		}

		if dbParam.writeToDB {
			log.Println("Write to resource-manager db...")
			dbConnInfo := dbParam.toDBConnectInfo()
			data := info.ToClusterConfigInsert()
			id, err := infra.InsertClusterConfig(dbConnInfo, data)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("\033[1;37;42m%s\033[0m\n", fmt.Sprintf("SUCCESS: Inserted ID=%d", id))
		}

	},
}

func init() {
	rootCmd.AddCommand(kdeCmd)

	kdeParam = &kdeParamStruct{}
	dbParam = &dbParamStruct{}

	kdeCmd.Flags().StringVarP(&kdeParam.host, flagHost, "", "", "KDE host (required)")
	kdeCmd.Flags().IntVarP(&kdeParam.port, flagPort, "", 0, "KDE port (required)")
	kdeCmd.Flags().StringVarP(&kdeParam.username, flagUsername, "", "",
		"KDE username (required)")
	kdeCmd.Flags().StringVarP(&kdeParam.password, flagPassword, "", "",
		"KDE password (required)")

	kdeCmd.Flags().StringVarP(&kdeParam.kdeType, flagType, "", "",
		"KDE type: online/offline (required)")

	kdeCmd.Flags().BoolVarP(&dbParam.writeToDB, flagWriteToDb, "", false,
		"Whether write to DB, default to false.")

	kdeCmd.Flags().BoolVarP(&dbParam.ignoreError, flagIgnoreError, "", false,
		"Whether ignore errors, default to false.")

	kdeCmd.Flags().StringVarP(&dbParam.host, flagDbHost, "", "",
		"DB host (required if write-to-db is true)")
	kdeCmd.Flags().IntVarP(&dbParam.port, flagDbPort, "", 0,
		"DB port (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbParam.username, flagDbUsername, "", "",
		"DB username (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbParam.password, flagDbPassword, "", "",
		"DB password (required if write-to-db is true)")
	kdeCmd.Flags().StringVarP(&dbParam.database, flagDatabase, "", "resource_manager",
		"DB database")

	_ = kdeCmd.MarkFlagRequired(flagHost)
	_ = kdeCmd.MarkFlagRequired(flagPort)
	_ = kdeCmd.MarkFlagRequired(flagUsername)
	_ = kdeCmd.MarkFlagRequired(flagPassword)
	_ = kdeCmd.MarkFlagRequired(flagType)
}
