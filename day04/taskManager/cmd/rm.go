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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:                   "rm {todo_name}",
	Short:                 "delete rodo, named todo_name",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("expect one parameter")
			return
		}

		args[0] = strings.Join(args, " ")

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
		db.Update(func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte("todo"))
			if bucket == nil {
				bucket, err = tx.CreateBucket([]byte("todo"))
				if err != nil {
					fmt.Println(err)
					return nil
				}
			}
			b := bucket.Get([]byte(args[0]))
			if b == nil {
				fmt.Println("todo not exists")
				return nil
			}

			err = bucket.Delete([]byte(args[0]))
			if err != nil {
				fmt.Println(err)
				return nil
			}

			fmt.Println("success")
			return nil
		})

		defer db.Close()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
