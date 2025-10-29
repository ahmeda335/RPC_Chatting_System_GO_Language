package main

import (
	"fmt"
	"net"
	"net/rpc"
	"rpc_assign/commons"
)

// Chat is our RPC service
type Chat struct {
	messages []commons.Message
}

// SendMessage stores the message and returns all messages
func (c *Chat) SendMessage(args commons.SendMessageArgs, reply *[]commons.Message) error {
	msg := commons.Message{
		ClientName: args.ClientName,
		Text:       args.Text,
	}
	c.messages = append(c.messages, msg)
	fmt.Printf("Message from %s: %s\n", args.ClientName, args.Text)
	*reply = c.messages
	return nil
}

func main() {
	addr := commons.Get_server_address()
	listener, _ := net.Listen("tcp", addr)
	fmt.Printf("Server running on %s...\n", addr)

	rpc.Register(new(Chat))

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
