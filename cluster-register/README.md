# cluster-register

## build

Fisrt `cd cluster-register`, then

```shell
GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' .
```

说明：
> - -w 去掉 DWARF 调试信息，得到的程序不能用 gdb 调试。
> - -s 去掉符号表, panic 时候的 stack trace 不会有任何文件名/行号信息，等价于普通 C/C++ 程序被 strip 的效果。

进一步减小 bin 体积，使用 upx

```shell
GOOS=windows GOARCH=amd64 go build -ldflags '-w -s' . && upx ./cluster-register
```