package errHandler

import (
	"log"
	"os"
)

// Handler Just Handler
func Handler(err error) {
	if !(err == nil) {
		log.Println("Error : ", err.Error())
	}
}

//HandlerExit Handler & Exit with code 1
func HandlerExit(err error) {
	if !(err == nil) {
		log.Println("Error : ", err.Error(), "Exiting (1)")
		os.Exit(1)
	}
}

//HandlerBool if handlel return true
func HandlerBool(err error) bool {
	if !(err == nil) {
		log.Println("Error : ", err.Error())
		return true
	}
	return false
}
