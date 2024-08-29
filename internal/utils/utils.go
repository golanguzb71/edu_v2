package utils

import (
	"edu_v2/graph/model"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
)

func AbsResponseChecking(err error, msg string) (*model.Response, error) {
	if err != nil {
		SendMessage(err.Error(), 6805374430)
		return &model.Response{
			StatusCode: 409,
			Message:    err.Error(),
		}, nil
	}
	return &model.Response{
		StatusCode: 200,
		Message:    msg,
	}, nil
}

type Response struct {
	UserID int `json:"user_id"`
	Code   int `json:"code"`
}

func OffSetGenerator(page, size *int) int {
	if page == nil || *page < 1 {
		p := 1
		page = &p
	}
	if size == nil || *size < 1 {
		s := 10
		size = &s
	}

	return *size * (*page - 1)
}

func CheckAdminCode(code string) error {
	if code != "KEY_ADM" {
		return errors.New("you are not admin pls get your code if you admin => https:/t.me/codevanbot")
	}
	return nil
}

func SendMessage(message string, chatId int64) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("LOGGER_BOT_TOKEN"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	msg := tgbotapi.NewMessage(chatId, message)
	_, err = bot.Send(msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func CalculateTotalPages(totalRecords int, size *int) int {
	if *size == 0 {
		return 0
	}
	totalPages := totalRecords / *size
	if totalRecords%*size > 0 {
		totalPages++
	}
	return totalPages
}
