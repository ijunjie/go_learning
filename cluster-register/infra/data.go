package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DBConnectInfo struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type ClusterConfigInsert struct {
	ClusterName    string
	Host           string
	RootCuNum      int
	BasicKey       string
	RmHost         string
	NmHost         string
	ClusterType    int
	ClusterKind    int
	HadoopMasterIp string
}

const (
	insertsql = `INSERT INTO cluster_config 
    (cluster_name,host,root_cu_num,basic_key,rm_host,nm_host,is_default,timestamp,file_type,cluster_type,cluster_kind,hadoop_master_ip,ingress) 
VALUES ( ?, ?, ?, ?, ?, ?, 1, now(), 'core-site,hdfs-site,hive-site,yarn-site,spark2-defaults', ?, ?, ?, '{"kdm":"http://some-addr"}')`
)

func InsertClusterConfig(dbConfig *DBConnectInfo, data *ClusterConfigInsert) (int64, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
	db, _ := sql.Open("mysql", dataSource)
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(5 * time.Second)

	if err := db.Ping(); err != nil {
		return -1, err
	}

	log.Printf("connect to DB succeeded, datasource: %s \n", dataSource)

	// start tx
	tx, err := db.Begin()
	if err != nil {
		return -1, err
	}

	// sql
	stmt, err1 := tx.Prepare(insertsql)
	if err1 != nil {
		return -1, err1
	}

	res, err2 := stmt.Exec(
		data.ClusterName,
		data.Host,
		data.RootCuNum,
		data.BasicKey,
		data.RmHost,
		data.NmHost,
		data.ClusterType,
		data.ClusterKind,
		data.HadoopMasterIp,
	)
	if err2 != nil {
		return -1, err2
	}
	_ = tx.Commit()

	id, err3 := res.LastInsertId()
	if err3 != nil {
		return -1, err3
	}

	return id, nil

}

func TypeNumber(env string) int {
	if env == "online" {
		return 0
	} else if env == "offline" {
		return 1
	} else if env == "endpoint" {
		return 2
	} else {
		return -1
	}
}
