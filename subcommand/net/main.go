package net

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
    mock-client net [options] 

Options:
   -h, --help
   --network=<network_type>  Type of network used for communication
   --address=<address>       Address for specific network_type.
   --debug                   Log debugging messages
   
Where:
   network_type   Examples: 'unix', 'tcp'
   address        Examples: '/tmp/test.sock', '127.0.0.1:12345'
`

	// DocOpt processing.

	args, _ := docopt.Parse(usage, nil, true, "", false)

	// Test for required commandline options.

	message := ""

	if args["--network"] == nil {
		message += "Missing '--network' parameter;"
	}

	if args["--address"] == nil {
		message += "Missing '--address' parameter;"

	}

	if len(message) > 0 {
		help.ShowHelp(usage)
		fmt.Println(strings.Replace(message, ";", "\n", -1))
		log.Fatalln(strings.Replace(message, ";", "; ", -1))
	}

	// Get commandline options.

	network := args["--network"].(string)
	address := args["--address"].(string)
	isDebug := args["--debug"].(bool)

	// Create a network connection.

	if isDebug {
	}

	networkConnection, err := net.Dial(network, address)
	if err != nil {
		log.Fatal("Dial error", err)
	}
	defer networkConnection.Close()

	// Start asynchronous Reader.

	go reader(networkConnection)

	// Loop through Writer.

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
