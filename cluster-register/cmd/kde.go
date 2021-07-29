package cmd

import (
	"cluster-register/infra"
	"cluster-register/kde"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
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
}

func (param *kdeParamStruct) toKdeInfoRequest() *kde.KdeInfoRequest {
	return &kde.KdeInfoRequest{
		KdeHost:  param.host,
		KdePort:  param.port,
		Username: param.username,
		Password: param.password,
		KdeType:  commonParam.clusterType,
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
	PreRun: func(cmd *cobra.Command, args []string) {
		if commonParam.clusterType != "online" && commonParam.clusterType != "offline" {
			fmt.Printf("Error: flag \"%s\" value should be \"{online|offline}\"\n", flagType)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		kdeInfoRequest := kdeParam.toKdeInfoRequest()
		info, err := kde.KdeInfo(kdeInfoRequest, commonParam.timeoutSeconds)
		if err != nil {
			log.Fatal(err)
		}

		json, _ := json.MarshalIndent(*info, "", "  ")
		log.Println("kde info: ")
		fmt.Println(string(json))

		quarterMem := info.MemGB / 4
		badRatio := quarterMem < info.Vcores
		if badRatio {
			if !commonParam.ignoreError {
				log.Fatalf("\033[1;37;41m%s\033[0m\n",
					"Error: TotalMemoryGB should be at least 4 times or more of vcores!")
			}
		}

		if commonParam.writeToDB {
			log.Println("Write to resource-manager db...")
			dbConnInfo := commonParam.toDBConnectInfo()
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

	kdeCmd.Flags().StringVarP(&kdeParam.host, flagHost, "", "", "KDE host (required)")
	kdeCmd.Flags().IntVarP(&kdeParam.port, flagPort, "", 0, "KDE port (required)")
	kdeCmd.Flags().StringVarP(&kdeParam.username, flagUsername, "", "",
		"KDE username (required)")
	kdeCmd.Flags().StringVarP(&kdeParam.password, flagPassword, "", "",
		"KDE password (required)")

	_ = kdeCmd.MarkFlagRequired(flagHost)
	_ = kdeCmd.MarkFlagRequired(flagPort)
	_ = kdeCmd.MarkFlagRequired(flagUsername)
	_ = kdeCmd.MarkFlagRequired(flagPassword)
}
