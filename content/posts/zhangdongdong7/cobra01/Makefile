## 默认变量， 获取当前 go 的环境变量
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

## build 在指定环境变异，直接执行make build使用默认系统和架构编译
build:
	go build -o out/greeting-$(GOOS)-$(GOARCH) .


## 通过指定环境变量， 执行交叉编译,直接执行make buildx交叉编译不同系统和架构的二进制文件
buildx:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build
