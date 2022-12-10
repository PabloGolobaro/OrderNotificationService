package utils

import (
	"WhatsappOrderServer/models"
	"fmt"
	tele "gopkg.in/telebot.v3"
	"log"
	"strconv"
)

func SendToAdmins(bot *tele.Bot, admins []string, order *models.Order) error {
	for _, admin := range admins {
		id, err := strconv.Atoi(admin)
		if err != nil {
			log.Println(fmt.Errorf("Ошибка отправки оповещения админам: %v", err))
			continue
		}
		message := order.String()
		_, err = bot.Send(tele.ChatID(id), message)
		if err != nil {
			return fmt.Errorf("Ошибка отправки оповещения админам: %v", err)
		}
	}
	return nil
}
