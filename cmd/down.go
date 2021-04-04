/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this gen except in compliance with the License.
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
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Run the downward migration",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		migrationsBuild := exec.Command("go", "build")
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		migrationsBuild.Dir = wd + "/migrations"

		errBuild := migrationsBuild.Start()
		if errBuild != nil {
			log.Fatal(errBuild)
		}

		errWait := migrationsBuild.Wait()
		if errWait != nil {
			log.Fatal(errWait)
		}

		migrationsRun := exec.Command("./migrations", "down")
		migrationsRun.Dir = wd + "/migrations"
		var outBuf, errBuf bytes.Buffer
		migrationsRun.Stdout = &outBuf
		migrationsRun.Stderr = &errBuf

		errRun := migrationsRun.Run()
		if errRun != nil {
			log.Fatal(errRun)
		}

		if len(outBuf.String()) == 0 {
			color.Green("  >>> migration : complete")
			return
		}

		if errBuf.Len() > 0 {
			log.Fatal(errBuf.String())
		}

		fmt.Println(outBuf.String())

	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
