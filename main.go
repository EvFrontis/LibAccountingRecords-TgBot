package main

import (
	"os"
	"time"

	bot "github.com/EvFrontis/LibAccountingRecords-TgBot/cmd"
	db "github.com/EvFrontis/LibAccountingRecords-TgBot/database"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {

	time.Sleep(1 * time.Second)

	//Creating Table
	if os.Getenv("CREATE_TABLE") == "yes" {

		if os.Getenv("DB_SWITCH") == "on" {

			if err := db.CreateTable(); err != nil {

				panic(err)
			}
		}
	}

	time.Sleep(1 * time.Second)

	//Call Bot
	bot.TelegramBot()
}
