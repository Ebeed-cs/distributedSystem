package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Message struct {
	Username  string
	Content   string
	Timestamp string
}

type SendMsgArgs struct {
	Username string
	Content  string
}

type SendMsgReply struct {
	History []Message
}

type GetHistoryArgs struct{}

type GetHistoryReply struct {
	History []Message
}

func printHistory(history []Message) {
	fmt.Println("\n------------------------------------")
	fmt.Println("CHAT HISTORY")
	fmt.Println("------------------------------------")
	if len(history) == 0 {
		fmt.Println("No messages yet.")
	} else {
		for _, msg := range history {
			fmt.Printf("[%s] %s: %s\n", msg.Timestamp, msg.Username, msg.Content)
		}
	}
	fmt.Println("------------------------------------\n")
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer client.Close()

	fmt.Print("Enter your username: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	username := strings.TrimSpace(scanner.Text())

	if username == "" {
		fmt.Println("Username cannot be empty.")
		os.Exit(1)
	}

	fmt.Println("Welcome to the chatroom, " + username + "!")
	fmt.Println("Type 'exit' to quit or press Ctrl+C.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n\nExiting chatroom...")
		os.Exit(0)
	}()

	for {
		fmt.Print("You: ")
		scanner.Scan()
		message := strings.TrimSpace(scanner.Text())

		if message == "exit" {
			fmt.Println("Exiting chatroom...")
			break
		}

		if message == "" {
			continue
		}

		args := &SendMsgArgs{
			Username: username,
			Content:  message,
		}
		var reply SendMsgReply

		err := client.Call("ChatServer.SendMessage", args, &reply)
		if err != nil {
			fmt.Println("Error sending message:", err)
			fmt.Println("Server may be down. Exiting...")
			break
		}

		printHistory(reply.History)
	}
}
