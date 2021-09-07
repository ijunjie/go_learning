# cluster-register

cluster-register 用于大数据云部署后初始化 KDE 和 K8S 集群注册表。


## 1. 构建

设置 
GO111MODULE=on
GOPROXY=https://goproxy.cn,direct


### 1.1 交叉编译


```shell
cd cluster-register
go mod download
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' .
```

arm
```shell script
GOOS=linux GOARCH=arm64 GOARM=7 go build -ldflags '-w -s' .
```

说明：
> - -w 去掉 DWARF 调试信息，得到的程序不能用 gdb 调试。
> - -s 去掉符号表, panic 时候的 stack trace 不会有任何文件名/行号信息，等价于普通 C/C++ 程序被 strip 的效果。

### 1.2 优化

进一步减小 bin 体积，使用 upx

upx 是跨平台通用的，可以在 windows 使用 upx.exe 压缩 Linux ELF.

```shell
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' . && upx ./cluster-register
```


## 2. 开发指南

### 2.1 环境准备

安装 go 1.16, 配置 go modules, 下载 cobra

```shell
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

下载 cobra

```shell
go get -v github.com/spf13/cobra/cobra
```

下载完成后，将 cobra.exe 加入系统 PATH

### 2.2 创建过程

使用 cobra 生成脚手架

```shell
cobra.exe init cluster-register
```

生成的脚手架具有适合扩展的项目结构，子命令在 cmd/ 下。

### 2.3 功能扩展开发


#### 2.3.1 现有功能维护

首先导入项目，可以使用 Goland,
或使用文本编辑器也可以。

下载依赖:

```shell
go mod tidy
```

#### 2.3.2 增加子命令

使用 cobra 扩展已有程序

例如，增加一个子命令 k8s:

```shell
cobra.exe add k8s
```

执行后会生成相应代码，按需修改完善即可


## 3. 代码设计详解

### 3.1 依赖管理

go.mod 中声明依赖

- mysql 驱动用于将集群信息写入资源管理数据库
- 单元测试相关
- cobra 相关
- k8s client-go

### 3.2 代码结构

- main.go 是程序入口
- cmd 是 cobra 脚手架生成的，存放子命令, kde, k8s 都在这里
- infra 包：基础设施，如提供 http 服务的 http.go 和提供数据库服务 data.go
- kde 包：获取 KDE 集群相关配置的服务
- k8s 包：获取 k8s 集群 sa 信息

### 3.3 root flag 设计

全部 flag 名以 flag 开头，设计为常量。

db 相关 flag 是子命令共用的，因此设计为包内共享常量，设计到 root.go 中

```c
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
```

绑定变量：

```c
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
```

方法：
- func (param *commonParamStruct) checkRequired() bool 用于校验 required 等
- func (param *commonParamStruct) toDBConnectInfo() *infra.DBConnectInfo 转换为 dataobj

init 中完成通用 flag 的解析。

PersistentPreRun 是一个全局钩子，每个子命令执行前都会调用。PersistentPreRun 中执行 commonParam 的校验。


### 3.4 subcommand 设计

流程：

```
(1)kdeParam -> kdeInfoRequest -> (2)kdeInfoResult -> dataobj -> (3)insertDB
(1)k8sParam -> k8sInfoRequest -> (2)k8sInfoResult -> dataobj -> (3)insertDB
```

步骤 (1) 在 cmd/kde.go 和 cmd/k8s.go 中实现
步骤 (2) 在 kde/service.go 和 k8s/service.go 中实现
步骤 (3) 在 infra/data.go 中实现

### 3.5 kde info 获取

kde/service.go 中具有关联关系的一组请求：

```c
const (
	clusterUrlTmpl  = "http://%s/api/v1/clusters"
	yarnSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=yarn-site"
	coreSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=core-site"
	metricsUrlTmpl  = "http://%s/ws/v1/cluster/metrics"
)
```

### 3.5 k8s info 获取

k8s/service.go

### 3.6 dao

infra/data.go