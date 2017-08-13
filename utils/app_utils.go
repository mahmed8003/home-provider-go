package utils

import "github.com/rs/zerolog"

/*
GetLogLevelByString :
*/
func GetLogLevelByString(level string) zerolog.Level {

	d := zerolog.Disabled

	if level == "debug" {
		d = zerolog.DebugLevel

	} else if level == "info" {
		d = zerolog.InfoLevel

	} else if level == "warn" {
		d = zerolog.WarnLevel

	} else if level == "error" {
		d = zerolog.ErrorLevel

	} else if level == "fatal" {
		d = zerolog.FatalLevel

	} else if level == "panic" {
		d = zerolog.PanicLevel

	} else if level == "disabled" {
		d = zerolog.Disabled
	}

	return d
}
