# Simple Chatroom (Go RPC)

A simple chatroom implementation using Go's built-in RPC library.

## Files

- `server.go` - RPC server that manages chat history
- `client.go` - RPC client for sending/viewing messages
- `README.md` - This file

## Features

- Clients can send messages to the server using RPC
- Server stores all messages in memory
- Full chat history displayed after each message
- Usernames and timestamps for each message
- Graceful handling of exit ("sxit" command or Ctrl+C)
- Error handling if server goes down

## Prerequisites

- Go 1.16 or higher

## Running the Chatroom

### Deom Link : https://youtu.be/0xyTuDuHX94

### Step 1: Start the Server

```bash
go run server.go
```

You should see:

```
Chat server started on :1234
```

### Step 2: Start Client(s)

In a new terminal window:

```bash
go run client.go
```

You'll be prompted to enter a username. After that, you can start typing messages.

You can open multiple clients to simulate multiple users chatting.

### Step 3: Send Messages

Type your message and press Enter. The full chat history will be displayed after each message.

### Step 4: Exit

Type `exit` or press `Ctrl+C` to quit the client.

## Example Usage

```
$ go run client.go
Enter your username: Alice
Welcome to the chatroom, Alice!
Type 'exit' to quit or press Ctrl+C.
You: Hello everyone!

------------------------------------
CHAT HISTORY
------------------------------------
[15:30:00] Alice: Hello everyone!
------------------------------------
```
