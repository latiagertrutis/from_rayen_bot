// ///////////////////////////////////////////////////////////////////
// Filename: telegram.go
// Description: telegram main file
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun May 31 13:59:15 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
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

func AddChatMember(chatID, userID int64) error {
	var exists bool

	exists = false
	row := GlobalDB.CheckChatUser.QueryRow(chatID, userID)
	err := row.Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		_, err = GlobalDB.AddChatUser.Exec(chatID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func RmChatMember(chatID, userID int64) error {
	var exists bool

	exists = false
	row := GlobalDB.CheckChatUser.QueryRow(chatID, userID)
	err := row.Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		_, err = GlobalDB.RmChatUser.Exec(chatID, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

func UpdateChatMembers(update tgbotapi.Update) error {
	var err error
	m := update.Message
	err = nil

	if m.LeftChatMember != nil {
		log.Infof("Removing %s, from %s\n", m.LeftChatMember.UserName, m.Chat.Title)
		err = RmChatMember(m.Chat.ID, int64(m.LeftChatMember.ID))
	} else if m.NewChatMembers != nil {
		for _, u := range *m.NewChatMembers {
			log.Infof("Adding %s to %s\n", u.UserName, m.Chat.Title)
			err = AddChatMember(m.Chat.ID, int64(u.ID))
		}
	} else if m.Text != "" && !m.From.IsBot {
		err = AddChatMember(m.Chat.ID, int64(m.From.ID))
	}
	return err
}

func MentionAllUsers(chatId int64) (tgbotapi.MessageConfig, error) {
	var user int
	mentions := ""

	rows, err := GlobalDB.GetChatUsers.Query(chatId)
	if err != nil {
		return tgbotapi.MessageConfig{}, err
	}

	for rows.Next() {
		err := rows.Scan(&user)
		if err != nil {
			log.Warning(err)
		} else {
			chatConfig := tgbotapi.ChatConfigWithUser{
				ChatID: chatId,
				UserID: user,
			}
			chatMember, err := bot.GetChatMember(chatConfig)
			if err != nil {
				log.Warning(err)
			}
			mentions += "@" + chatMember.User.UserName + " "
		}
	}
	rows.Close()
	return tgbotapi.NewMessage(chatId, mentions), nil
}

func RouteUpdate(update tgbotapi.Update) error {
	err := UpdateChatMembers(update)
	if err != nil {
		return err
	}

	switch update.Message.Text {
	case "/gente":
		m, err := MentionAllUsers(update.Message.Chat.ID)
		if err != nil {
			return err
		}
		m.ReplyToMessageID = update.Message.MessageID
		bot.Send(m)
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
			err := RouteUpdate(update)
			if err != nil {
				log.Error(err)
			}
		}
	}
}
