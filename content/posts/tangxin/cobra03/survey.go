package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		// 1. Input 输入框
		Name: "id",
		Prompt: &survey.Input{
			Message: "Access Secret ID: ",
		},
		Validate: survey.Required,
		// Transform: survey.Title,
	},
	{
		// 2. Password 密码输入框
		Name: "key",
		Prompt: &survey.Password{
			Message: "Access Secret Key: ",
		},
		Validate: survey.Required,
		// Transform: survey.Title,
	},
	{
		// 3. Select 单选框
		Name: "region",
		Prompt: &survey.Select{
			Message: "Choose a region: ",
			Options: []string{"cn-shanghai", "cn-hangzhou"},
			Default: "cn-hangzhou",
		},
	},
	{
		// 4. MultiSelect 多选框
		Name: "language",
		Prompt: &survey.MultiSelect{
			Message: "Supported Configure Language: ",
			Options: []string{"zh", "en", "jp"},
		},
	},
}

func interactive(profile string) {

	answers := struct {
		ID          string
		Key         string
		ChinaRegion string `survey:"region"`
		Language    []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if !confirm() {
		fmt.Println("用户取消文件保存!")
	}

	dumpConfig(profile, answers)
}

func confirm() bool {
	ok := false
	// 5. Confirm 确认框
	prompt := &survey.Confirm{
		Message: "是否保存文件?",
	}
	survey.AskOne(prompt, &ok)

	return ok
}

func dumpConfig(profile string, answer any) {
	b, err := json.MarshalIndent(answer, "", "  ")
	if err != nil {
		panic(err)
	}

	name := fmt.Sprintf("%s.config.json", profile)
	err2 := os.WriteFile(name, b, os.ModePerm)
	if err2 != nil {
		panic(err2)
	}
}
