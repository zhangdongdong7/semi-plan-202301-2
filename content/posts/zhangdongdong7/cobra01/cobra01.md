---
title: "作业: cobra - 01 实现编译与参数绑定"
categories:
  - 半月刊
tags:
  - 半月刊202301下
  - code
date: "2023-01-13T15:19:07+08:00"
lastmod: "2023-01-13T15:19:07+08:00"
toc: true
draft: false
hiddenFromHomePage: false

#  提交作业修改一下内容
pinTop: false
originAuthor: zhangdongdong7
originLink: ""
---


# 作业: cobra - 01 实现编译与参数绑定

要求:

1. 使用 https://github.com/spf13/cobra 实现命令工具
2. 命令具有以下参数
    1. `--name` 姓名
    2. `--age` 年龄
3. 如果年龄为空， 默认为 20 岁。
4. 完成交叉编译脚本， 编译其他平台的二进制文件

```
-rwxr-xr-x  1 franktang  staff  4220672 Jan 13 15:35 greeting-darwin-amd64
-rwxr-xr-x  1 franktang  staff  4203442 Jan 13 15:35 greeting-darwin-arm64
-rwxr-xr-x  1 franktang  staff  4215010 Jan 13 15:35 greeting-linux-amd64
-rwxr-xr-x  1 franktang  staff  4157892 Jan 13 15:35 greeting-linux-arm64
```

5. 执行输出效果如下

```bash
$ ./out/greeting-darwin-arm64
 你好, 今年 20 岁

$ ./out/greeting-darwin-arm64 --age 30 --name zhangsan
zhangsan 你好, 今年 30 岁
```

## 解题思路


### 1. 安装依赖包

```bash
$ go get -u github.com/spf13/cobra
```


### 2. 创建命令
这部分可以去搜一下这个包的使用方法https://pkg.go.dev/github.com/spf13/cobra#Command.Commands

```go
var root = &cobra.Command{
	Use:   "person",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		Person(name, age) 
	},
}
```

### 3. 指定参数


1. 定义了变量作为参数接受者。
2. 使用 init 函数， 在程序初始化的时候， 传递参数值。

```go
var (
	name string 
	age  int
)

func init() {
	root.Flags().StringVarP(&name, "name", "", "", "姓名") // 初始化姓名参数
	root.Flags().IntVarP(&age, "age", "", 20, "年龄")   // 初始化年龄参数
}
```

### 4. 指定业务逻辑
``` 
func Person(name string, age int) {
	fmt.Printf("%s 你好， 今年 %d 岁\n", name, age) // 定义person函数
}
```


### 5. 调用执行

1. 在 main 中调用 root 的执行函数 Execute()
2. 在 root 的 Run 中调用 执行逻辑入口。
3. Person是程序实际执行逻辑，及执行逻辑入口。

### 6. 编译

使用 Makefile 进行编译管理
