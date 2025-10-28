package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type Message struct {
	Username  string
	Content   string
	Timestamp string
}

type ChatServer struct {
	mu       sync.RWMutex
	history []Message
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

func (s *ChatServer) SendMessage(args *SendMsgArgs, reply *SendMsgReply) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	msg := Message{
		Username:  args.Username,
		Content:   args.Content,
		Timestamp: time.Now().Format("15:04:05"),
	}

	s.history = append(s.history, msg)
	reply.History = s.history

	return nil
}

func (s *ChatServer) GetHistory(args *GetHistoryArgs, reply *GetHistoryReply) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	reply.History = s.history

	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}
	defer listener.Close()

	fmt.Println("Chat server started on :1234")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
