/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/praneshsuresh/cli-uzo-golang/util"
	"github.com/spf13/cobra"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "It will open the directory in Visual Studio Code",
	Long: `This command will help to open the unzipped folder
	to Visual Studio Code.
	In order for this command to work, Visual Studio code should be installed in your system`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var file string
		var err error
		var arg string = args[0]

		//check whether the zip file exists
		fileExists, err := util.FileExists(arg)
		if err != nil {
			fmt.Println(err.Error())
		}

		//if the file exists, get the absolute path of the zip file
		if fileExists {
			file, err = filepath.Abs(arg)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Printf("File %v does not exist", arg) //return this message when file doesn't exist
			return
		}
		
		//get the current working directory (where the zip file is located in)
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		//calls the unzip function to unzip the zip file
		util.Unzip(file, wd)

		//change directory to the unzipped folder
		os.Chdir(util.FilenameWithoutExtension(file))

		//make the working directory as the newly changed directory
		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err)
		}

		//calls the exec.Command function to tell that "code" to open VS code
		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			log.Fatal("VS Code executable file not found in %PATH%")
		}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// codeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
