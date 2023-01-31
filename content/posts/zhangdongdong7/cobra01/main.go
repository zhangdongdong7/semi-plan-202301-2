package main
import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
)

var (
	name string 
	age  int
)

var root = &cobra.Command{
	Use:   "person",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		Person(name, age) // 2. root 命令调用person函数
	},
}

func init() {
	root.Flags().StringVarP(&name, "name", "", "", "姓名") // 初始化姓名参数
	root.Flags().IntVarP(&age, "age", "", 20, "年龄")   // 初始化年龄参数
}

func Person(name string, age int) {
	fmt.Printf("%s 你好， 今年 %d 岁\n", name, age) // 定义person函数
}

// func Execute() error {
// 	return root.Execute() // 定义Execute函数
// }

func main() {
	err := root.Execute() // 1. main 调用 root 命令
	if err != nil {
		log.Fatal(err)
	}
}