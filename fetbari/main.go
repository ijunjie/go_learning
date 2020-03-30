package main

import (
	"fetbari/clusters"
	"fetbari/common"
	"fetbari/coresite"
	"fetbari/metrics"
	"fetbari/yarnsite"
	"fmt"
	"log"
	"os"
)

const (
	hostTmpl = "http://%s:8080"
	sqlTmpl  = `INSERT INTO cluster_config (cluster_name,host,root_cu_num,basic_key,rm_host,nm_host,is_default,timestamp,file_type,cluster_type,cluster_kind,hadoop_master_ip,ingress) VALUES ('%s','%s',%d,'%s','%s','%s',1,now(),'core-site,hdfs-site,hive-site,yarn-site,spark2-defaults',%d,0,'%s','{"kdm":"http://some-addr"}');`
)

func main() {
	fmt.Printf("%21s: %-30s\n", "Fetbari Version", "1.2")
	argCount := len(os.Args[1:])
	if argCount != 4 {
		fmt.Println("Usage: fetbari {ambari server ip} {username} {password} {online/offline}")
		fmt.Println("Example: fetbari 10.10.10.10 admin admin online")
		return
	}

	ambari, username, password, env := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
	if env != "online" && env != "offline" {
		fmt.Println("The env arg should be online/offline")
		return
	}

	clusterName, err1 := clusters.ClusterName(ambari, username, password)
	if err1 != nil {
		log.Fatal(err1)
	}
	yarnResourceManager, err2 := yarnsite.YarnResourceManager(ambari, username, password, clusterName)
	if err2 != nil {
		log.Fatal(err2)
	}
	nameNodeHost, err3 := coresite.NameNodeHost(ambari, username, password, clusterName)
	if err3 != nil {
		log.Fatal(err3)
	}
	vcores, mem, err4 := metrics.VcoresAndMem(yarnResourceManager)
	if err4 != nil {
		log.Fatal(err4)
	}
	calculatedCU := vcores

	memGB := mem / 1024
	quarterMem := memGB / 4
	if quarterMem < vcores {
		fmt.Printf("%21s: %-30s\n", "### WARNING ###", "TotalMemoryGB should be at least 4 times or more of vcores!!!")
	}
	host := fmt.Sprintf(hostTmpl, ambari)
	basicKey := common.Base64Encode(username, password)
	envNumber := common.OnlineOrOffline(env)
	sql := fmt.Sprintf(sqlTmpl, clusterName, host, calculatedCU, basicKey,
		yarnResourceManager, nameNodeHost, envNumber, ambari)

	fmt.Printf("%21s: %-30s\n", "Cluster Name", clusterName)
	fmt.Printf("%21s: %-30s\n", "Host", host)
	fmt.Printf("%21s: %-30s\n", "Authorization", basicKey)
	fmt.Printf("%21s: %-30s\n", "Yarn Resource Manager", yarnResourceManager)
	fmt.Printf("%21s: %-30s\n", "Name Node Host", nameNodeHost)
	fmt.Printf("%21s: %-30d\n", "Total Vcores", vcores)
	fmt.Printf("%21s: %-30d\n", "Total MemMB", mem)
	fmt.Printf("%21s: %-30d\n", "Total MemGB", memGB)
	fmt.Printf("%21s: %-30d\n", "Total CU", calculatedCU)
	fmt.Printf("%21s: %-30d\n", "Env", envNumber)
	fmt.Printf("%21s: %-30s\n", "Insert SQL Sample", sql)
}
