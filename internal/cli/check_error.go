package cli

import "log"

// Receive an error and run a log.Fatal if err != nil.
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
