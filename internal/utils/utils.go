package utils

import (
	"edu_v2/graph/model"
	"errors"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"runtime"
	"time"
)

func AbsResponseChecking(err error, msg string) (*model.Response, error) {
	if err != nil {
		ErrorLogger(err)
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
		errTxt := "\"unauthorized access: you are not an admin, please obtain the correct code\""
		ErrorLogger(errors.New(errTxt))
		return errors.New(errTxt)
	}

	return nil
}
func ErrorLogger(err error) {
	pc, fn, line, _ := runtime.Caller(1)
	detailedError := fmt.Sprintf("[ERROR] %s\nOccurred at %s:%d in %s\nTime: %s",
		err,
		fn,
		line,
		runtime.FuncForPC(pc).Name(),
		time.Now().Format(time.RFC3339),
	)
	SendMessage(detailedError, 6805374430)
}

func SendMessage(message string, chatId int64) {
	botToken := os.Getenv("LOGGER_BOT_TOKEN")
	if botToken == "" {
		fmt.Println("LOGGER_BOT_TOKEN is not set")
		return
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		fmt.Println("Failed to create bot:", err.Error())
		return
	}

	msg := tgbotapi.NewMessage(chatId, message)
	_, err = bot.Send(msg)
	if err != nil {
		fmt.Println("Failed to send message:", err.Error())
		return
	}

	fmt.Println("Message sent successfully:", message)
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
