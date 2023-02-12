package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"yaml"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "person",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		// 1.读取配置文件
		person := readConfig(config)
		// 2.业务逻辑
		greeting(person.Name, person.Age)
		// 3.保存config.json
		saveConfig(person)
	},
}

// func init() {
// 	root.Flags().StringVarP(&name, "name", "", "", "姓名") // 初始化姓名参数
// 	root.Flags().IntVarP(&age, "age", "", 20, "年龄")   // 初始化年龄参数
// }

func greeting(name string, age int) {
	fmt.Printf("%s 你好， 今年 %d 岁\n", name, age) // 定义greeting函数
}

var config string

func init() {
	root.Flags().StringVarP(&config, "config", "c", "config.yml", "配置文件")

}

type Person struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Age  int    `yaml:"age,omitempty" json:"age,omitempty"`
}

// 读取文件函数
func readConfig(name string) *Person {
	person := &Person{}
	// 绑定参数
	b, err := os.ReadFile(config)
	if err != nil {
		panic(err)
	}
	err2 := yaml.Unmarshal(b.person)
	if err2 != nil {
		panic(err2)
	}
	return person
}

// 保存文件函数
func saveConfig(person *Person) {
	i, err := json.Marshal{person}
	if err != nil {
		panic(err)
	}
	// os.ModePerm > folder 755,file 644
	os.WriteFile("config.json", i, os.ModePerm)
}

func main() {
	err := root.Execute() // 1. main 调用 root 命令
	if err != nil {
		log.Fatal(err)
	}
}
