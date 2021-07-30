package cmd

import (
	"cluster-register/infra"
	"cluster-register/k8s"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const (
	flagClusterName = "name"
)

type k8sParamStruct struct {
	name string
}

func (param *k8sParamStruct) toK8sInfoRequest() *k8s.K8sInfoRequest {
	return &k8s.K8sInfoRequest{
		ClusterName: param.name,
		K8sType:     commonParam.clusterType,
	}
}

var k8sParam *k8sParamStruct

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
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		if commonParam.clusterType != "online" && commonParam.clusterType != "offline" {
			log.Printf("Error: flag \"%s\" value should be \"{online|offline|endpoint}\"\n", flagType)
			os.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		request := k8sParam.toK8sInfoRequest()
		info, err := k8s.K8sInfo(request, commonParam.timeoutSeconds)
		if err != nil {
			log.Fatal(err)
		}

		json, _ := json.MarshalIndent(*info, "", "  ")
		log.Println("k8s info: ")
		log.Println(string(json))

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
	rootCmd.AddCommand(k8sCmd)

	k8sParam = &k8sParamStruct{}

	k8sCmd.Flags().StringVarP(&k8sParam.name, flagClusterName, "", "", "K8S name (required)")

	_ = k8sCmd.MarkFlagRequired(flagClusterName)
}
