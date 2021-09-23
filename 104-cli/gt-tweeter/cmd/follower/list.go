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

package follower

import (
	"fmt"
	"github.com/GoTurkiye/training/104-cli/gt-tweeter/pkg/client"
	"github.com/dghubble/go-twitter/twitter"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var count int

const defaultFollowerFormatTemplate = `Follower: {{ .Name }}@{{ .ID }}, following: {{ .Following }}
👉 https://twitter.com/@{{ .ScreenName }}`

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the followers",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.New(client.Options{
			ConsumerKey:    viper.GetString("twitter.consumerKey"),
			ConsumerSecret: viper.GetString("twitter.consumerSecret"),
			AccessToken:    viper.GetString("twitter.accessToken"),
			AccessSecret:   viper.GetString("twitter.accessSecret"),
		})

		s, _ := cmd.Flags().GetInt("count")

		followers, _, err := c.Followers.List(&twitter.FollowerListParams{
			Count: s,
		})

		if err != nil {
			log.Fatalf("could not list followers: %v\n", err)
		}

		t := template.New("follower_template")

		for index, follower := range followers.Users {
			tmpl, err := t.Parse(defaultFollowerFormatTemplate)
			if err != nil {
				log.Fatalf("failed to parse template: %v\n", err)
			}

			fmt.Fprintf(os.Stdout, "%d. ", index)
			err = tmpl.Execute(os.Stdout, follower)
			if err != nil {
				log.Fatalf("failed to execute template: %v\n", err)
			}
			fmt.Fprintln(os.Stdout, "")
			fmt.Fprintln(os.Stdout, "")
		}
	},
}

func init() {
	listCmd.Flags().IntVarP(&count, "count", "c", 20, "specify the count of the followers")
	//followerCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
