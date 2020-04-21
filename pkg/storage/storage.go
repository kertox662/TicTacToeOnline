//Package storage provides an interface for accessing storage to the server data
package storage

import "github.com/kertox662/TicTacToeOnline/pkg/data"

//Storage is an interface that will be used to read and write to user data.
type Storage interface {
	Reader
	Writer
}

//GStorage is the global variable for accessing the storage
var GStorage Storage

//Reader enables reading of the data from the storage method
type Reader interface {
	Initialize()

	GetUserStats(username string) []data.UserStats
	GetLeaderboards(numPlayers, numConnect, boardSize int) []data.LeaderboardRow
	GetUserEmail(username string) string
	GetValidationToken(username string) string
	GetAuthToken(username string) string
	VerifyPassword(username, password string) bool
}

//Writer enables the writing to the the data in server storage
type Writer interface {
	Initialize()

	//Profile Management
	CreateUser(data.UserCredentials) error
	ChangeName(data.UserCredentials) error
	ChangeEmail(data.UserCredentials) error
	ChangePassword(data.UserCredentials) error
	DeleteUser(data.UserCredentials) error
	GenerateValidationToken(username string)
	ValidateUser(username string, token string) bool

	//Authentication
	GenerateAuthToken(username string) int
	PurgeAuthToken(username string, token int)

	//Game Records
	ArchiveGame(data.GameData)
}
