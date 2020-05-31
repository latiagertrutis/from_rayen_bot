// ///////////////////////////////////////////////////////////////////
// Filename: cli.go
// Description: command line interface
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun May 31 13:44:31 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/urfave/cli"
	"time"
)

func InitCli() *cli.App {
	return &cli.App{
		Name:     "from_rayen_bot",
		Usage:    "Awesome telegram bot",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Mateo Rodriguez",
				Email: "teorodrip@posteo.net",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Usage:    "Telebram bot token",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			TelegramMain(c.String("token"))
			return nil
		},
	}
}
