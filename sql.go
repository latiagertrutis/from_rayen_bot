// ///////////////////////////////////////////////////////////////////
// Filename: sql.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Mon Oct 12 16:39:31 2020 (+0200)
// Last-Updated: Mon Oct 12 18:43:00 2020 (+0200)
//////////////////////////////////////////////////////////////////////

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	Db            *sql.DB
	AddChatUser   *sql.Stmt
	RmChatUser    *sql.Stmt
	CheckChatUser *sql.Stmt
	GetChatUsers  *sql.Stmt
}

var (
	GlobalDB *Db
)

const (
	ChatMembersTable = `
CREATE TABLE IF NOT EXISTS "chat_members" (
	"chat_id" BIGINT NOT NULL,
	"user_id" INTEGER NOT NULL,
	UNIQUE(chat_id, user_id)
);
`
	AddChatUser = `
INSERT INTO "chat_members" ("chat_id", "user_id")
VALUES (?1, ?2);
`
	RmChatUSer = `
DELETE FROM "chat_members" WHERE (chat_id=?1 AND user_id=?2);
`
	CheckChatUser = `
SELECT EXISTS(SELECT 1 FROM "chat_members" WHERE (chat_id=?1 AND user_id=?2) LIMIT 1);
`
	GetChatUsers = `
SELECT user_id FROM "chat_members" WHERE chat_id=?;
`
)

func NewDb(Path string) (*Db, error) {
	var err error
	var d Db

	d.Db, err = sql.Open("sqlite3", Path)
	if err != nil {
		return nil, err
	}

	err = d.Db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = d.Db.Exec(ChatMembersTable)
	if err != nil {
		log.Error("Creating the table")
		return nil, err
	}

	d.AddChatUser, err = d.Db.Prepare(AddChatUser)
	if err != nil {
		log.Error("Initializing AddChatUser")
		return nil, err
	}
	d.RmChatUser, err = d.Db.Prepare(RmChatUSer)
	if err != nil {
		log.Error("Initializing RmChatUser")
		return nil, err
	}
	d.CheckChatUser, err = d.Db.Prepare(CheckChatUser)
	if err != nil {
		return nil, err
	}
	d.GetChatUsers, err = d.Db.Prepare(GetChatUsers)
	if err != nil {
		return nil, err
	}
	return &d, nil
}

//
// sql.go ends here
