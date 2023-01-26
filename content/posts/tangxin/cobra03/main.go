package main

import "github.com/spf13/cobra"

func main() {
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	Use:   "aliyunx",
	Short: "aliyun 配置中心",
	Run: func(cmd *cobra.Command, args []string) {
		interactive(profile)
	},
}

var profile string

func init() {
	root.Flags().StringVarP(&profile, "profile", "p", "default", "配置名称")
}
