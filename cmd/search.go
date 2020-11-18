/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"log"
	"fmt"
	//"vilya/pkg/utils"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var (
 searchCmd = &cobra.Command{
	Use:   "search",
	Short: "searches for maintenane updates affecting the product",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		Search()
	},
}
team string
)


func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.PersistentFlags().StringVar(&team, "team", "qam-manager", "what is the abrev. of the team in maintenance qam-caasp/qam-sle/qam-suma")
	//rootCmd.PersistentFlags().IntVar(&workers, "workers", 0, "number of workers in the cluster")
	//rootCmd.PersistentFlags().IntVar(&masters, "masters", 0, "number of masters in the cluster")
	//rootCmd.PersistentFlags().StringVarP(&pool, "pool", "p", "", "name of pool for the project")
	//rootCmd.PersistentFlags().StringVarP(&pool, "distro", "d", "", "name of distro in the cluster")
}

func Search() {
	updates, err := Config.CheckForUpd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for _, elem := range updates {
		log.Printf("Package: %+v  Repos: %+v\n", elem.SRCRPMS, elem.Repository)
	}
}
