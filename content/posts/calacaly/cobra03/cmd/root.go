/*
Copyright © 2023 calacaly

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra03",
	Short: "Save information through parameters or prompts.",
	Long: `Save information through parameters or prompts.
When parameters are not passed or the entered parameters are the default values, 
interactive input will be performed`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		age, err := cmd.Flags().GetUint("age")
		if err != nil {
			fmt.Println(err)
			return
		} else {
			if age == 20 {
				ans, err := prompt_age()
				if err != nil {
					fmt.Println(err)
					return
				}
				age = ans.Age
				viper.Set("age", age)
			}
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
			return
		} else {
			if name == "" {
				ans, err := prompt_name()
				if err != nil {
					fmt.Println(err)
					return
				}
				name = ans.Name
				viper.Set("name", name)
			}
		}

		ans, err := prompt_color()
		if err != nil {
			fmt.Println(err)
			return
		}
		color := ans.FavoriteColor
		viper.Set("color", color)

		fmt.Printf("\n我叫 %s , 今年 %d 岁了, 最喜欢的颜色是 %s\n\n", name, age, color)

		viper.WriteConfigAs("info.json")

		fmt.Printf("saved as info.json\n")

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra01.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().UintP("age", "a", 20, "set your age, or wait for the prompt to enter")
	rootCmd.Flags().StringP("name", "n", "", "set your name, or wait for the prompt to enter")

}

type ANSWERS struct {
	Name          string // survey will match the question and field names
	FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
	Age           uint   // if the types don't match, survey will convert it
}

func prompt_name() (ANSWERS, error) {
	var qs = []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "What is your name?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
	}
	answers := ANSWERS{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	return answers, err
}

func prompt_color() (ANSWERS, error) {
	var qs = []*survey.Question{
		{
			Name: "color",
			Prompt: &survey.Select{
				Message: "Choose a color:",
				Options: []string{"red", "blue", "green"},
				Default: "red",
			},
		},
	}
	answers := ANSWERS{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	return answers, err
}

func prompt_age() (ANSWERS, error) {
	var qs = []*survey.Question{
		{
			Name:   "age",
			Prompt: &survey.Input{Message: "How old are you?"},
		},
	}
	answers := ANSWERS{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	return answers, err
}
