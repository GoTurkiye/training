/*
Copyright © 2021 Batuhan Apaydın <developerguyn@gmail.com>

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

package tweet

import (
	"github.com/spf13/cobra"
	"os"
)

// TweetCmd represents the tweet command
var TweetCmd = &cobra.Command{
	Use:   "tweet",
	Short: "Helps you to send tweet and show the status of the tweet",
	Long: `Basically you can send tweets and based on the command result you can query the status of the tweet to see
how many times it is retweeted, quoted, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

func init() {
	TweetCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tweetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tweetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
