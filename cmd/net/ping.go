/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"
	"github.com/spf13/cobra"
)

var (
	urlPath string

	client = &http.Client{
		Timeout: time.Second * 10,
	}
)

func ping(domain string)(int, error){
	url := "http://" + domain

	req,err := http.NewRequest("HEAD", url, nil)

	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil


}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This pings a remote url",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		
		// logic
		if resp, err := ping(urlPath); err != nil {
			panic(err)
		} else {
			fmt.Println(resp)
		}
	},

}

func init() {
	// rootCmd.AddCommand(pingCmd)


	// Here you will define your flags and configuration settings.

	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The url to ping.")

	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		panic(err)
	}

	NetCmd.AddCommand(pingCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}