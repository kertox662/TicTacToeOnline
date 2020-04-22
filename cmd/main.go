package main

import (
	"strconv"

	"github.com/kertox662/TicTacToeOnline/internal/postgres"
	"github.com/kertox662/TicTacToeOnline/pkg/logging"
	"github.com/kertox662/TicTacToeOnline/pkg/logging/consolelog"
	"github.com/kertox662/TicTacToeOnline/pkg/storage"
)

func main() {
	logging.GLogger = &consolelog.Console{}
	logging.GLogger.SetDefaultTimeFormat()
	// logging.GLogger.SetMinLevel(logging.Info)

	logging.GLogger.Info("Starting up storage")
	storage.GStorage = &postgres.DBStorage{}
	storage.GStorage.Initialize()
	logging.GLogger.Info("Started Storage")

	logging.GLogger.Info("Admin Validated: " + strconv.FormatBool(storage.GStorage.CheckValidation("admin")))
	logging.GLogger.Info("User  Validated: " + strconv.FormatBool(storage.GStorage.CheckValidation("user")))
}
