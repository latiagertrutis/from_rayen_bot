// ///////////////////////////////////////////////////////////////////
// Filename: main.go
// Description:
// Author: Mateo Rodriguez Ripolles (teorodrip@posteo.net)
// Maintainer:
// Created: Sun May 31 13:08:06 2020 (+0200)
// ///////////////////////////////////////////////////////////////////

package main

import (
	"github.com/op/go-logging"
	"os"
)

var log = logging.MustGetLogger("rayen_log")

func InitLog() {
	format := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	backend1 := logging.NewLogBackend(os.Stdout, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	backend1Formatter := logging.NewBackendFormatter(backend1, format)

	backend2Leveled := logging.AddModuleLevel(backend2)
	backend2Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Formatter, backend2Leveled)
}

func main() {
	InitLog()

	app := InitCli()

	app.Run(os.Args)
}
