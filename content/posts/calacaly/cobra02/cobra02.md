---
title: "作业: cobra - 02 读取配置配置文件"
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
pinTop: true
originAuthor: calacaly
originLink: ""
---


# 作业: cobra - 02 读取配置配置文件

要求:

1. 使用 https://github.com/spf13/cobra 实现命令工具
2. 命令具有以下参数
    1. `--config` , `-c` 配置文件

**配置文件如下**

```yaml
# config.yml
name: zhangsan
age: 20
```

3. 将配置文件保存为 `JSON` 格式 

```bash
$ cat config.json
```

输出结果

```json
{
    "name":"zhangsan",
    "age": 20
}
```
