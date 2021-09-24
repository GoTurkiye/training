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
	"github.com/spf13/viper"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var id int64

const defaultTweetFormatTemplate = `
Created At: {{ .CreatedAt }}
Favorite Count: {{ .FavoriteCount }}
Retweet  Count: {{ .RetweetCount }}
Quote Count: {{ .QuoteCount }}
`

// showCmd represents the list command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the status of the Tweet",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		tweet, _, err := c.Statuses.Show(id, nil)

		if err != nil {
			log.Fatalf("could not show tweet: %v\n", err)
		}

		t := template.New("tweet_template")

		tmpl, err := t.Parse(defaultTweetFormatTemplate)
		if err != nil {
			log.Fatalf("failed to parse template: %v\n", err)
		}

		err = tmpl.Execute(os.Stdout, tweet)
		if err != nil {
			log.Fatalf("failed to execute template: %v\n", err)
		}
		fmt.Fprintln(os.Stdout, "")
		fmt.Fprintln(os.Stdout, "")
	},
}

func init() {
	TweetCmd.AddCommand(showCmd)
	showCmd.Flags().Int64VarP(&id, "id", "i", 0, "specify the id of the tweet")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
