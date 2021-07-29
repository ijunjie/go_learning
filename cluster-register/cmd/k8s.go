package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

const (
	flagClusterName = "name"
)

type k8sParamStruct struct {
	name    string
	k8sType string
}

var k8sParam *k8sParamStruct

//var dbParam

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Register K8S to resource-manager",
	Example: `
Write to DB:

./cluster-register k8s --name=k8s-online --type=online \ 
--write-to-db --db-host=192.168.10.10 --db-port=3306 --db-username=myusername --db-password=mypassword

Show info only:

./cluster-register k8s --name=k8s-online --type=online`,
	Long: `Register K8S as a cluster_config to resource-manager.`,
	Run: func(cmd *cobra.Command, args []string) {
		if k8sParam.k8sType != "online" && k8sParam.k8sType != "offline" && k8sParam.k8sType != "endpoint" {
			log.Printf("Error: flag \"%s\" value should be \"online|offline|endpoint\"\n", flagType)
			return
		}
		ok := dbParam.checkRequired()
		if !ok {
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(k8sCmd)

	k8sParam = &k8sParamStruct{}
	dbParam = &dbParamStruct{}

	k8sCmd.Flags().StringVarP(&k8sParam.name, flagClusterName, "", "", "K8S name (required)")
	k8sCmd.Flags().StringVarP(&k8sParam.k8sType, flagType, "", "",
		"K8S type: online/offline/endpoint (required)")

	k8sCmd.Flags().BoolVarP(&dbParam.writeToDB, flagWriteToDb, "", false,
		"Whether write to DB, default to false.")

	k8sCmd.Flags().StringVarP(&dbParam.host, flagDbHost, "", "",
		"DB host (required if write-to-db is true)")
	k8sCmd.Flags().IntVarP(&dbParam.port, flagDbPort, "", 0,
		"DB port (required if write-to-db is true)")
	k8sCmd.Flags().StringVarP(&dbParam.username, flagDbUsername, "", "",
		"DB username (required if write-to-db is true)")
	k8sCmd.Flags().StringVarP(&dbParam.password, flagDbPassword, "", "",
		"DB password (required if write-to-db is true)")
	k8sCmd.Flags().StringVarP(&dbParam.database, flagDatabase, "", "resource_manager",
		"DB database")

	_ = k8sCmd.MarkFlagRequired(flagClusterName)
	_ = k8sCmd.MarkFlagRequired(flagType)
}
