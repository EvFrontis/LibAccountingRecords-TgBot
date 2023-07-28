package app

import (
	"os"
	"reflect"
	"strconv"
	"time"

	db "github.com/EvFrontis/LibAccountingRecords-TgBot/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type currentAction int

const (
	free currentAction = iota
	add
	get
)

type Add int

const (
	name Add = iota
	age
	number
)

func TelegramBot() {

	godotenv.Load()
	//Create bot
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	//Set update timeout
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Get updates from bot
	updates := bot.GetUpdatesChan(u)

	status := free
	addStatus := name
	var person db.Person

	for update := range updates {
		if update.Message == nil {
			continue
		}
		//Check if message from user is text
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			if status == free {
				switch update.Message.Text {
				case "/start":
					//Send message
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a lib bot, I can add a new entry to the card file or get the user's number by name, choose the action you need.")
					bot.Send(msg)

					if err := db.CreateUserTable(update.Message.Chat.UserName); err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error:"+err.Error())
						bot.Send(msg)
					}

					db.AddUser(update.Message.Chat.UserName, int(update.Message.Chat.ID))

				case "/get_users":
					if os.Getenv("DB_SWITCH") == "on" {
						status = get
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter the name of the person you want to find.")
						bot.Send(msg)
					} else {
						status = free
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database not connected, so I can't give an answer.")
						bot.Send(msg)
					}

				case "/add_user":
					status = add
					addStatus = name // switching to entering a name
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter name.")
					bot.Send(msg)

				default:
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choose an action.")
					bot.Send(msg)
				}
			} else if status == get {
				people, err := db.GetPeople(update.Message.Chat.UserName, update.Message.Text)
				if err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Get: Database error: "+err.Error())
					bot.Send(msg)
				}

				//Creating string which contains people
				ans := "Answer:"
				for _, value := range people {
					ans += "\n" + value.String()
				}

				//Send message
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)
				bot.Send(msg)

				status = free

			} else if status == add {
				switch addStatus {
				case name:
					person.Name = update.Message.Text
					addStatus = age // switching to entering an age
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter age.")
					bot.Send(msg)

				case age:
					person.Birthdate, err = time.Parse("2006-01-02", update.Message.Text)
					if err != nil {
						addStatus = age
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "It's not a number. Enter a number.")
						bot.Send(msg)
					} else {
						addStatus = number // switching to entering a number
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter the person's number.")
						bot.Send(msg)
					}

				case number:
					person.Num, err = strconv.Atoi(update.Message.Text)
					if err != nil {
						addStatus = number
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "It's not a number. Enter a number.")
						bot.Send(msg)
					} else {
						//Putting name, age, number to database
						if err := db.AddPerson(update.Message.Chat.UserName, person.Name, person.Birthdate, person.Num); err != nil {
							//Send message
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Database error, but bot still working. Error: "+err.Error())
							bot.Send(msg)
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "OK. Information added.")
							bot.Send(msg)
						}
						status = free
					}
				}
			}

		} else {
			//Send message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			bot.Send(msg)
		}
	}
}
