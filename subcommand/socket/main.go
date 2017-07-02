package socket

// Inspirations:
//  - https://gist.github.com/hakobe/6f70d69b8c5243117787fd488ae7fbf2

import (
	"io"
	"log"
	"fmt"
	"net"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/docktermj/mock-client/common/help"
)

// Read a message from the network and respond.
func reader(reader io.Reader) {
	byteBuffer := make([]byte, 1024)
	for {
		numberOfBytesRead, err := reader.Read(byteBuffer[:])
		if err != nil {
			return
		}
		fmt.Println("<<<", string(byteBuffer[0:numberOfBytesRead]))
	}
}

// Function for the "command pattern".
func Command(argv []string) {

	usage := `
Usage:
    mock-server socket [options] 

Options:
   -h, --help
   --socket-file=<file>        Socket file
   --debug               Log debugging messages
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)

	// Test for required commandline options.

	if args["--socket-file"] == nil {
		message := "Missing '--socket-file' parameter"
		fmt.Println(message)
		help.ShowHelp(usage)
		log.Fatalln(message)
	}

	// Get commandline options.

	socketFile := args["--socket-file"].(string)
	isDebug := args["--debug"].(bool)

	// Listen on the Unix Domain Socket

	if isDebug {

	}

	networkConnection, err := net.Dial("unix", socketFile)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer networkConnection.Close()

    // Start asynchronous Reader
	
	go reader(networkConnection)
	
	// Loop through Writer
	
	loopNumber := 1
	for {
	    loopNumber += 1
		outboundMessage := fmt.Sprintf("Sending #%d", loopNumber)
		_, err := networkConnection.Write([]byte(outboundMessage))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		fmt.Println(">>>", outboundMessage)
		time.Sleep(1e9)
	}
}
