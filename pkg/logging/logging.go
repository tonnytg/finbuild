package logging

import (
	"fmt"
	"log"
	"os"
)

// Msg Message Log module for logs.
// to use this log.Msg("CRITICAL", "cannot create db")
func Msg(ErrType, Msg string) {
	file, err := os.OpenFile("activity.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("ERROR: problem to open or create file log")
	}

	// log config
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	// log registry
	log.Printf("UTC %s: %s\n", ErrType, Msg)
}
