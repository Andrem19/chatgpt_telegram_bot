package helpers

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	db "github.com/Andrem19/telegramGPT/db/sqlc"
	_ "github.com/lib/pq"
)

var queries *db.Queries
var database *sql.DB

func StartWithDb(config Config) {
	database, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	queries = db.New(database)
}

func Switcher(message string, chat_id int64) (string, error) {
	if message == "/start" {
		return "Provide your GPT token!\nExample:\nsk-56jOrOP1HD1uxcrTrh0CT3BlbkTFcmPLUSH3xvj7JYOnZMSz", nil
	} else if message == "/info" {
		return "to use our bot you need to input your gpt token, then you can talk with bot", nil
	} else if len([]byte(message)) == 51 &&  message[0:2] == "sk" {
		userExist, err := queries.GetUsers(context.Background(), fmt.Sprintf("%d", chat_id))
		if err != nil {
			AddToLog(err.Error())
		}
		if userExist.ID >0 {

			param := db.UpdateUserTokenParams{
				ChatID: fmt.Sprintf("%d", chat_id),
				GptToken: message,
			}
			userWithNewToken, err := queries.UpdateUserToken(context.Background(), param)
			if err != nil {
				AddToLog(err.Error())
			}
			if userWithNewToken.ID > 0 && userWithNewToken.GptToken == message {
				return "Your token was changed", nil
			}
		} else {
			user := db.CreateUserParams{
				ChatID: fmt.Sprintf("%d", chat_id),
				GptToken: message,
				Step: 1,
				LastAnswer: "",
			}
			id, err := queries.CreateUser(context.Background(), user)
			if err != nil {
				fmt.Println("Error:", err.Error())
			}
			if id > 0 {
				return "Account was created. You can comunicate with bot", nil
			}
		}
	}else if message == "/deleteMe" {
		queries.DeleteAccount(context.Background(), fmt.Sprintf("%d", chat_id))
		return "Your account successfuly deleted", nil
	} else if message == "/reset" {
		param := db.UpdateStepAndAnswerParams{
			ChatID: fmt.Sprintf("%d", chat_id),
			Step: 1,
			LastAnswer: "",
		}
		_, err := queries.UpdateStepAndAnswer(context.Background(), param)
		if err !=nil {
			AddToLog(err.Error())
		}
		return "Bot was reseted", nil
	} else if len(message) > 0 {
		var answer string
		user, err := queries.GetUsers(context.Background(), fmt.Sprintf("%d", chat_id))
		if err != nil {
			AddToLog(err.Error())
			return "User does not Exist. First provide your GPT token!\nExample:\nsk-56jOrOP1HD1uxcrTrh0CT3BlbkTFcmPLUSH3xvj7JYOnZMSz", nil
		}
		var message2 string
		if user.Step > 1 {
			message2 = fmt.Sprintf("ChatGPT said: %s. Me said: %s", user.LastAnswer, message)
		} else if user.Step == 1 {
			message2 = message
		}

		answer, err = AskQuestion(message2, user.GptToken)
		if err != nil {
			AddToLog(err.Error())
		}

		if user.Step < 4 {
			nextStep := user.Step +1
			param := db.UpdateStepAndAnswerParams{
				ChatID: fmt.Sprintf("%d", chat_id),
				Step: nextStep,
				LastAnswer: answer,
			}
			_, err := queries.UpdateStepAndAnswer(context.Background(), param)
			if err !=nil {
				AddToLog(err.Error())
			}
			return answer, nil
		} else {
			param := db.UpdateStepAndAnswerParams{
				ChatID: fmt.Sprintf("%d", chat_id),
				Step: 1,
				LastAnswer: "",
			}
			_, err := queries.UpdateStepAndAnswer(context.Background(), param)
			if err !=nil {
				AddToLog(err.Error())
			}
			return answer, nil
		}
	} 
	return "No one case suit", nil
}