package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"rpc_assign/commons"
	"strings"
)

func main() {
	// Connect to server
	client, _ := rpc.Dial("tcp", commons.Get_server_address())

	var name string
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	name = strings.TrimSpace(input)

	for {
		var text string
		fmt.Print("> ")
		input, _ = reader.ReadString('\n')
		text = strings.TrimSpace(input)

		if text == "quit" {
			break
		}

		// Send message and get all messages back
		args := commons.SendMessageArgs{ClientName: name, Text: text}
		var messages []commons.Message
		client.Call("Chat.SendMessage", args, &messages)

		// Print all messages
		for _, m := range messages {
			fmt.Printf("%s: %s\n", m.ClientName, m.Text)
		}
	
	}
}
