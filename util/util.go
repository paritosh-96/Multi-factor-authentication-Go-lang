package util

import (
	"github.com/paritosh-96/RestServer/startup"
	"log"
)

func Check(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

/**
Close the database connection handler
*/
func Close() {
	if err := startup.Db.Close(); err != nil {
		log.Println("Error while closing Database: ", err)
		return
	}
	log.Println("Database handle closed successfully")
}

func Empty(str string) bool {
	if str == "" {
		return true
	}
	return false
}
