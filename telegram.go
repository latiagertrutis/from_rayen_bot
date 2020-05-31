// ///////////////////////////////////////////////////////////////////
// Filename: telegram.go
// Description: telegram main file
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun May 31 13:59:15 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI

func InitBot(token string) error {
	var err error

	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	return nil
}

func TelegramMain() error {
	var update tgbotapi.Update

	upConfig := tgbotapi.UpdateConfig{
		Timeout: 500,
	}

	updates, err := bot.GetUpdatesChan(upConfig)
	if err != nil {
		return err
	}

	updates.Clear()
	for {
		select {
		case update = <-updates:
			if update.Message == nil {
				continue
			}
			log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Tu te callas")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
		log.Info("HOLA")
	}
}
