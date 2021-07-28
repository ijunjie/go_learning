# cluster-register

cluster-register 用于大数据云部署后初始化 KDE 和 K8S 集群注册表。


## 1. 构建

### 1.1 交叉编译


```shell
cd cluster-register
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' .
```

说明：
> - -w 去掉 DWARF 调试信息，得到的程序不能用 gdb 调试。
> - -s 去掉符号表, panic 时候的 stack trace 不会有任何文件名/行号信息，等价于普通 C/C++ 程序被 strip 的效果。

### 1.2 优化

进一步减小 bin 体积，使用 upx

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

### 3.2 代码结构

- main.go 是程序入口
- cmd 是 cobra 脚手架生成的，存放子命令, kde, k8s 都在这里
- infra 包：基础设施，如提供 http 服务的 http.go 和提供数据库服务 data.go
- kde 包：获取 KDE 集群相关配置的服务

### 3.3 flag 设计

全部 flag 名以 flag 开头，设计为常量。

db 相关 flag 是子命令共用的，因此设计为包内共享常量，设计到 root.go 中

```c
const (
	flagWriteToDb   = "write-to-db"
	flagIgnoreError = "ignore-error"
	flagDbHost      = "db-host"
	flagDbPort      = "db-port"
	flagDbUsername  = "db-username"
	flagDbPassword  = "db-password"
	flagDatabase    = "database"
)
```

kde.go flag 常量，仅包内使用：
```c
const (
	flagHost     = "host"
	flagPort     = "port"
	flagUsername = "username"
	flagPassword = "password"
	flagType     = "type"
)
```

### 3.4 flag 值变量设计

对应 flag 设计, param 的可见性与其保持一致.

对于 DB 相关 flag 值， 在 root.go 设计结构体：

```c
type dbParamStruct struct {
	writeToDB   bool
	ignoreError bool
	host        string
	port        int
	username    string
	password    string
	database    string
}
```

对于 kde flag 值，在 kde.go 设计结构体

```c
type kdeParamStruct struct {
	host     string
	port     int
	username string
	password string
	kdeType  string
}
```



至此，kde 的 flag 共 2 类 12 个参数，kde.go 中初始化并返回指针：
```c
var kdeParam = &kdeParamStruct{}
var dbParam = &dbParamStruct{}
```

### 3.5 init

kde 的 init 函数中，首先将 kde 挂到 root 下，然后对 12 个 flag 参数做解析。

最后将 5 个参数标记为 required:

```c
_ = kdeCmd.MarkFlagRequired(flagHost)
_ = kdeCmd.MarkFlagRequired(flagPort)
_ = kdeCmd.MarkFlagRequired(flagUsername)
_ = kdeCmd.MarkFlagRequired(flagPassword)
_ = kdeCmd.MarkFlagRequired(flagType)
```

### 3.6 run

kdeCmd 初始化参数 Run 函数中，首先对 write-to-db 为 true 时，
关联 flag 的 required 做了自定义校验，此处暂时未找到官方推荐做法。
这些逻辑应该抽取到 root.go, 用于其他子命令使用。

### 3.7 kde-info 获取

面向对象的设计精髓在于**合适的方法出现在合适的类中**，对于 go 而言，
**合适的函数、结构体要出现在合适的包和源码文件中**。

例如 service.go 中，入参和出参的设计。

service.go 中使用 http 请求和 json 解析， 将 kde 集群信息构建为 KdeInfoResult 结构体并返回。
请求参数和相应结果设计为独立的结构体。

```c
func KdeInfo(request *KdeInfoRequest) (*KdeInfoResult, error)
```

具有关联关系的一组请求：
```c
const (
	clusterUrlTmpl  = "http://%s/api/v1/clusters"
	yarnSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=yarn-site"
	coreSiteUrlTmpl = "http://%s/api/v1/clusters/%s/configurations?type=core-site"
	metricsUrlTmpl  = "http://%s/ws/v1/cluster/metrics"
)
```

### 3.8 data

data.go 设计。参考 service.go

### 3.9 对象转换

request 链路采用 toTarget 向底层转换；
response 链路采用 fromTarget 向上层传递。

### 3.10 异常处理

异常传递，log.Fatal 退出
