package utils

import (
	"log"
)

//HandleError will handle the error
func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
