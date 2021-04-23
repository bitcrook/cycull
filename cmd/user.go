/*
Copyright © 2021 Audio <hyperaudio@protonmail.com>

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

// Package cmd ...
package cmd

import (
	"fmt"

	"github.com/audioo/goseek/helpers/cli"
	"github.com/audioo/goseek/helpers/userlookup/run"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user [username] [write (y/n)]",
	Short: "Username Lookup",
	Long: `Search across 130+ sites to see if the username is associated with any of them and
	their associated pages for account deletion`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			cmd.Usage()
		} else {
			if args[1] == "y" {
				cli.Clear()
				cli.Banner()
				cli.Dispban("Username Lookup")
				fmt.Println()
				run.SendSeeker(args[0], true)
				fmt.Println()
			} else if args[1] == "n" {
				cli.Clear()
				cli.Banner()
				cli.Dispban("Username Lookup")
				fmt.Println()
				run.SendSeeker(args[0], false)
				fmt.Println()
			} else {
				cli.Clear()
				cli.Banner()
				cli.Dispban("Username Lookup")
				fmt.Println()
				run.SendSeeker(args[0], false)
				fmt.Println()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
