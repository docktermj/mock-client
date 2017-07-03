package port

// Inspirations:
//  - https://gist.github.com/hakobe/6f70d69b8c5243117787fd488ae7fbf2

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/docktermj/mock-client/common/help"
	"github.com/docopt/docopt-go"
)

// Read a message from the network and respond.
func reader(reader io.Reader) {
	byteBuffer := make([]byte, 1024)
	for {
		numberOfBytesRead, err := reader.Read(byteBuffer)
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
    mock-client port [options] 

Options:
   -h, --help
   --host=<hostname>     Optional: hostname Default: 127.0.0.1
   --port=<port_number>  Port number to connect
   --debug               Log debugging messages
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)

	// Test for required commandline options.

	message := ""

	if args["--port"] == nil {
		message += "Missing '--port' parameter;"
	}

	if len(message) > 0 {
		help.ShowHelp(usage)
		fmt.Println(strings.Replace(message, ";", "\n", -1))
		log.Fatalln(strings.Replace(message, ";", "; ", -1))
	}

	// Get commandline options.

	port := args["--port"].(string)
	isDebug := args["--debug"].(bool)

	hostname := "127.0.0.1"
	if args["--address"] != nil {
		hostname = args["--address"].(string)
	}
	address := fmt.Sprintf("%s:%s", hostname, port)

	// Debugging information.

	if isDebug {
		log.Printf("Sending to '%s' network with address '%s'", "tcp", address)
	}

	// Create a network connection to a service.

	networkConnection, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer networkConnection.Close()

	// Start asynchronous Reader.

	go reader(networkConnection)

	// Loop through Writer.

	loopNumber := 0
	for {
		loopNumber += 1
		outboundMessage := fmt.Sprintf("Sending #%d", loopNumber)
		_, err := networkConnection.Write([]byte(outboundMessage))
		if err != nil {
			log.Fatal("Write error:", err)
			break
		}
		fmt.Println(">>>", outboundMessage)
		time.Sleep(2 * time.Second)
	}
}
