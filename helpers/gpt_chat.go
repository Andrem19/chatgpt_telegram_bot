package helpers

import (
	"time"

	"github.com/solywsh/chatgpt"
)

func AskQuestion(question string, gpt_token string) (string, error) {
	chat := chatgpt.New(gpt_token, "user_id(not required)", 30*time.Second)
	defer chat.Close()
	answer, err := chat.Chat(question)
	if err != nil {
		AddToLog(err.Error())
	}
	return answer, nil
}