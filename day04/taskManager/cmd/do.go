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

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:                   "do {todo_name}",
	Short:                 "Mark todo named todo_name as done",
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
			if b[0] == 0 {
				fmt.Println("todo already done")
				return nil
			}

			err = bucket.Put([]byte(args[0]), []byte{0})
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
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
