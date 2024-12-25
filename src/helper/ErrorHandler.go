package helper

import "log"

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
