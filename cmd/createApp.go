// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/phd/cmd/FileInputOutput"

	"github.com/spf13/cobra"
)

// createAppCmd represents the createApp command
var createAppCmd = &cobra.Command{
	Use:   "create-app",
	Short: "create some project with specification project like website using js,php,go etc",
	Long:  `before you start this app you must install depedence like npm,composser,go get`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("createApp called")
		setApp, err := cmd.Flags().GetString("SetApp")
		name, err := cmd.Flags().GetString("name")
		framework, err := cmd.Flags().GetString("framework")

		if err != nil {
			fmt.Println("error arguement")
		}
		if setApp == "" {
			fmt.Println(" empty set app please using -s to set")
		} else {
			if setApp == "website" {
				if framework == "" {
					fmt.Println(" empty set framework")
				} else if framework == "codeigniter" || framework == "laravel" {
					if name == "" {
						fmt.Println(" empty set name")
					} else {
						if err := FileInputOutput.DownloadFile(framework); err != nil {
							panic(err)
						}
						fmt.Println("download successfull")
						fmt.Println("unziping")
						FileInputOutput.Unzip(framework)
						fmt.Println("unziping successfull")
						fmt.Println("delete zip file")
						FileInputOutput.DeleteFile(framework)
						fmt.Println("Delete zip successfull")
						fmt.Println("rename directory")
						FileInputOutput.RenameFile(framework, name)
						fmt.Println("done")
					}
				}
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(createAppCmd)

	// Here you will define your flags and configuration settings.
	createAppCmd.Flags().StringP("SetApp", "s", "", "set application you want")
	createAppCmd.Flags().StringP("framework", "f", "", "set framework")
	createAppCmd.Flags().StringP("name", "n", "", "Set your name")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createAppCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createAppCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
