package data

const ( //BoardSize Enum represents three separate sizes for a tic-tac-toe board
	//BoardStd is the standard size for a board
	BoardStd = iota + 1
	//BoardBig is the bigger size for a board
	BoardBig
	//BoardHuge is the biggest size for a board
	BoardHuge
)

//UserCredentials holds the basic credentials that are needed for user sign-up
type UserCredentials struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//UserStats hold the game stats for a user
type UserStats struct {
	NumPlayers      int
	BoardSize       int
	NumberToConnect int
	GamesPlayed     int
	NumWins         int
	NumLosses       int
	Rating          int
}

//LeaderboardRow holds the data for a row when the leaderboards are requested
type LeaderboardRow struct {
	Username    string
	GamesPlayed int
	NumWins     int
	NumLosses   int
	Rating      int
}

//GameData represents a generic game that has data
type GameData interface {
	GetGameType() int
}

//TicTacToeData represents the data for a game of TicTacToe
type TicTacToeData struct {
	NumPlayers int
	BoardSize  int //Enum
	NumConnect int
	UserIDs    []int
	WinnerID   int
	IsRated    bool
}

//GetGameType returns the gametype of a TicTacToe game
func (gData TicTacToeData) GetGameType() int {
	return 1
}
