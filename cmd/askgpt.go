/*
Copyright © 2023 SHREYANSH <shreyansh.yashi@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// askgptCmd represents the askgpt command
var askgptCmd = &cobra.Command{
	Use:   "askgpt",
	Short: "Ask anything",
	Long:  `Ask anything`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {

			inp := ""
			key := args[0]

			for i := 1; i < len(args); i++ {
				inp += args[i] + " "
			}

			getResponse(key, inp)

		} else {
			fmt.Println("try again!")
		}

	},
}

func init() {
	rootCmd.AddCommand(askgptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// askgptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// askgptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
