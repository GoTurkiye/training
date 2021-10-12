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
	"fmt"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var message string

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a Tweet with the given message",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		// Send a Tweet
		tweet, _, err := c.Statuses.Update(message, nil)

		if err != nil {
			log.Fatalf("coult not send tweet: %v\n", err)
		}

		fmt.Printf(`Tweet created successfully
You can check the status of the Tweet by running the following command:
$ gt-tweeter tweet show --id %d`, tweet.ID)
	},
}

func init() {
	sendCmd.Flags().StringVarP(&message, "message", "m", "", "specify the message")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
