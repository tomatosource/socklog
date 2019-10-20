package example

import (
	"io/ioutil"
	"log"

	"github.com/tomatosource/socklog"
)

func main() {
	logger, err := socklog.New("localhost:8080")
	if err != nil {
		// Can't connect to socklog server, discard all logs
		log.SetOutput(ioutil.Discard)
	} else {
		defer logger.Close()
		log.SetOutput(logger)
	}

	// Log as normal
	log.Print("hello world")
}
