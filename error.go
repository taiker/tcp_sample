package main

import (
	"log"
)

func handleServError(err error) {
	if err != nil {
		log.Fatal("Server: ", err)
	}
}
func handleHttperr(err error) []byte {
	if err != nil {
		log.Fatal("http: ", err)

	}
	return []byte(err.Error())
}
