package logging

const ( //Logging levels
	//Unknown is the level set when setting an invalid level
	Unknown = iota
	//Debug is the level to set for displaying information for debugging
	Debug
	//Info is the level to set for displaying information
	Info
	//Warning is the level to set for when encountering warnings
	Warning
	//Error is the level to set for raising errors
	Error
	//Fatal is the level to set for unrecoverable errors
	Fatal
	//Disabled is the level to disable the logger
	Disabled
)

//GLogger is the global logger used by the program
var GLogger Logger

//Logger implements functions that are used for printing information at different logging levels
type Logger interface {
	Log(string)
	Info(string)
	Debug(string)
	Warning(string)
	Error(string) error
	Fatal(string)

	SetDefaultLevel(int)
	SetMinLevel(int)

	SetTimeFormat(string)
	SetDefaultTimeFormat()
}
