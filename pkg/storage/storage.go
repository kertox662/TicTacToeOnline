//Package storage provides an interface for accessing storage to the server data
package storage

import "github.com/kertox662/TicTacToeOnline/pkg/data"

//Reader enables reading of the data from the storage method
type Reader interface {
	GetUserStats(username string) data.UserStats
	GetUserEmail(username string) string
	GetValidationToken(username string) int
	GetAuthToken(username string) int
}

//Writer enables the writing to the the data in server storage
type Writer interface {
	//Profile Management
	CreateUser(data.UserCredentials) error
	ChangeName(data.UserCredentials) error
	ChangeEmail(data.UserCredentials) error
	ChangePassword(data.UserCredentials) error
	DeleteUser(data.UserCredentials)
	ValidateUser(username string, token int)

	//Authentication
	GenerateAuthToken(username string) int
	PurgeAuthToken(username string, token int)

	//Game Records
	ArchiveGame(data.GameData)
}
