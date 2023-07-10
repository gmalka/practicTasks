/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"go.etcd.io/bbolt"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:                   "completed",
	Short:                 "View all complete todos",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			fmt.Println("expect none parameter")
			return
		}

		path, _ := homedir.Dir()
		if i := strings.IndexByte(path, '/'); i != -1 {
			path = path + "/my.db"
		} else {
			path = path + "\\my.db"
		}
		db, err := bbolt.Open(path, 0666, nil)
		if err != nil {
			log.Println(err)
		}

		db.View(func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte("todo"))
			if bucket == nil {
				fmt.Println("no todos in database")
				return nil
			}
			i := 0
			bucket.ForEach(func(k, v []byte) error {
				if v[0] == 0 {
					i++
					fmt.Printf("- %s\n", k)
				}
				return nil
			})
			if i == 0 {
				fmt.Println("no complete todos in database")
				return nil
			}

			return nil
		})

		defer db.Close()
	},
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
