package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("SSL_MODE"),)
	db, err := sqlx.Open(os.Getenv("DB_DRIVER"), config)
	if err != nil {
		log.Printf("Failed database connect: %v\n", err)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Failed database Begin: %v\n", err)
		return
	}
	defer tx.Rollback()

	rows, err := tx.Query("DELETE FROM phonebook RETURNING phonenumber")
	if err != nil {
		log.Printf("Failed database Delete: %v\n", err)
		return
	}
	defer rows.Close()
	phones := make(map[string]interface{}, 10)
	for rows.Next() {
		var str string

		err := rows.Scan(&str)
		if err != nil {
			log.Printf("Failed database returning data Scan: %v\n", err)
			return
		}
		fmt.Printf("\033[32m#-> Served number: \033[0m%s\n", str)
		str = normalizeNumber(str)
		if _, ok := phones[str]; !ok {
			phones[str] = nil
			_, err := db.Exec("INSERT INTO phonebook(phonenumber) VALUES ($1)", str)
			if err != nil {
				log.Printf("Failed database insert Data: %v\n", err)
				return
			}
			fmt.Printf("\033[33m Recorded number: \033[0m%s\n", str)
		} else {
			fmt.Printf("\033[0;31m Already exists: \033[0m%s\n", str)
		}
	}

	tx.Commit()
}

func normalizeNumber(str string) string {
	result := []rune(str)
	left, right := 0, 0

	for ; right < len(result); right++ {
		if result[right] >= '0' && result[right] <= '9' {
			result[left] = result[right]
			left++
		}
	}

	return string(result[:left])
}