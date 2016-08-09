package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/ilijamt/postfixbeat/beater"
)

func main() {
	err := beat.Run("postfixbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
