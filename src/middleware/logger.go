package middleware

import "log"

var loglevel int

func init() {
	log.Printf("logger says hi\n")
	loglevel = 4
}

func WriteLog(msg string) {
	log.Printf("%s\n", msg)
}

func GetLogLevel() int {
	return loglevel
}
