package main

import "log"

func mustSuccess(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
