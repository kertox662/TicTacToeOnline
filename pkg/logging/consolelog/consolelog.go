package consolelog

import (
	"fmt"
	"time"

	"github.com/kertox662/TicTacToeOnline/pkg/logging"
)

//Console implements the Logger interface and holds data for the format
type Console struct {
	defaultLevel int
	minLevel     int
	dateFormat   string
}

var format = map[int]string{
	logging.Info:    "[\033[1;37mINFO\033[0m]:",
	logging.Debug:   "[\033[1;34mDEBUG\033[0m]:",
	logging.Warning: "[\033[1;33mWARNING\033[0m]:",
	logging.Error:   "[\033[1;31mERROR\033[0m]:",
	logging.Fatal:   "[\033[1;31mFATAL\033[0m]:",
}

const defaultDateFormat = "2006-01-02 15:04:05"

//Log prints out to the console in the format based on the default logging level
func (cnsl *Console) Log(msg string) {
	if cnsl.defaultLevel < cnsl.minLevel {
		return
	}
	if cnsl.defaultLevel == logging.Unknown {
		return
	}
	// date := cnsl.getTime()
	// msgBytes := []byte(strings.Join([]string{date, format[cnsl.defaultLevel], msg}, " "))
	fmt.Println(cnsl.getTime(), format[cnsl.defaultLevel], msg)
	// os.Stdout.Write(msgBytes)
	// os.Stdout.Write([]byte("\n"))
}

//Info prints out to the console at the Info level
func (cnsl *Console) Info(msg string) {
	if cnsl.minLevel > logging.Info {
		return
	}
	fmt.Println(cnsl.getTime(), format[logging.Info], msg)
}

//Debug prints out to the console at the Debug level
func (cnsl *Console) Debug(msg string) {
	if cnsl.minLevel > logging.Debug {
		return
	}
	date := cnsl.getTime()
	println(date, format[logging.Debug], msg)
}

//Warning prints out to the at the Warning level
func (cnsl *Console) Warning(msg string) {
	if cnsl.minLevel > logging.Warning {
		return
	}
	date := cnsl.getTime()
	fmt.Println(date, format[logging.Warning], msg)
}

//Error prints out to the at the Error level and returns an error
func (cnsl *Console) Error(msg string) error {
	if cnsl.minLevel > logging.Error {
		return fmt.Errorf("Error Logging level disabled")
	}
	date := cnsl.getTime()
	err := date + " " + format[logging.Error] + " " + msg
	fmt.Println(err)
	return fmt.Errorf(err)
}

//Fatal prints out to the at the Fatal level and exits the program
func (cnsl *Console) Fatal(msg string) {
	if cnsl.minLevel > logging.Fatal {
		return
	}
	date := cnsl.getTime()
	panic(date + " " + format[logging.Fatal] + " " + msg)
}

//SetDefaultLevel sets the level for the Log function to print at.
func (cnsl *Console) SetDefaultLevel(level int) {
	if level < 0 || level > 5 {
		cnsl.defaultLevel = logging.Unknown
	} else {
		cnsl.defaultLevel = level
	}
}

//SetMinLevel sets the level for the Log function to print at.
func (cnsl *Console) SetMinLevel(level int) {
	if level < logging.Unknown || level > logging.Disabled {
		cnsl.minLevel = logging.Unknown
	} else {
		cnsl.minLevel = level
	}
}

//SetTimeFormat sets the string format for the loggin functions to use in time.Format()
func (cnsl *Console) SetTimeFormat(format string) {
	cnsl.dateFormat = format
}

//SetDefaultTimeFormat sets a time format defined by the package
func (cnsl *Console) SetDefaultTimeFormat() {
	cnsl.dateFormat = defaultDateFormat
}

func (cnsl *Console) getTime() string {
	t := time.Now()
	return t.Format(cnsl.dateFormat)
}

func parseArgs(args ...interface{}) []interface{} {
	for i, v := range args {
		fmt.Println(i, v)
	}
	return nil
}
