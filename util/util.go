package util

import "log"

func LogError(err error) {
	if err != nil {
		log.Fatalf("Error cli: %v", err)
	}
}
