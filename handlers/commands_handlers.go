package handlers

import (
	"fmt"
	tele "gopkg.in/telebot.v3"
)

var StartCommand = func(ctx tele.Context) error {
	var answer string
	user := ctx.Sender()
	answer = fmt.Sprintf("Здравствуй %v.\n", user.Username)
	return ctx.Send(answer)
}
