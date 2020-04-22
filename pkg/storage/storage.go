//Package storage provides an interface for accessing storage to the server data
package storage

import "github.com/kertox662/TicTacToeOnline/pkg/data"

//Storage is an interface that will be used to read and write to user data.
type Storage interface {
	Reader
	Writer
	Initialize()
}

//GStorage is the global variable for accessing the storage
var GStorage Storage

//Terms used for the interface:
//UserID - A unique ID for a user
//Credentials - A mixture of Username, Email, Password and a Auth or Validation token
//Validation Token - A token that will be used for credential based operations
//Authentication (Auth) Token - A token that will be used for login and gameplay operations

//Reader enables reading of the data from the storage method
type Reader interface {
	GetUserID(username string) int
	GetUserStats(username string) []data.UserStats
	GetLeaderboards(numPlayers, numConnect, boardSize int) []data.LeaderboardRow
	GetUserEmail(username string) string
	GetValidationToken(username string) string
	GetAuthToken(username string) string
	VerifyPassword(username, password string) bool
	VerifyAuthToken(username, token string) bool
	VerifyValidationToken(string, string) bool
	CheckValidation(username string) bool
}

//Writer enables the writing to the the data in server storage
type Writer interface {
	//Profile Management
	CreateUser(data.UserCredentials) error
	ChangeName(data.UserCredentials, string) error
	ChangeEmail(data.UserCredentials) error
	ChangePassword(data.UserCredentials) error
	DeleteUser(data.UserCredentials) error
	GenerateValidationToken(username string)
	ValidateUser(data.UserCredentials) error
	UnvalidateUser(cred data.UserCredentials) error

	//Authentication
	GenerateAuthToken(username string)
	PurgeAuthToken(username string, token string)

	//Game Records
	ArchiveGame(data.GameData)
}
