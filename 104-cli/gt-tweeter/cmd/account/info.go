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

package account

import (
	"fmt"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"
)

const UserInfoTemplate string = `Name: %s
Handle: @%s
Description: %s 
Followers Count: %d
Location: %s
`

// infoCmd represents the get command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get the detail account information",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		user, _, err := c.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
			IncludeEntities: twitter.Bool(true), SkipStatus: twitter.Bool(false), IncludeEmail: twitter.Bool(true),
		})

		if err != nil {
			log.Fatalf("could not get account info: %v\n", err)
		}

		fmt.Printf(UserInfoTemplate, user.Name,
			user.ScreenName,
			user.Description,
			user.FollowersCount,
			user.Location)
	},
}

func init() {
	//followerCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
