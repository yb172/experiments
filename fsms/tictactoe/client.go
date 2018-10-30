package tictactoe

import (
	"bufio"
	"context"
	fmt "fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

// ConnectAndPlay connects to game server
func ConnectAndPlay() error {
	serverAddr := fmt.Sprintf("127.0.0.1:%v", ServerPort)
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("unable to connect to game server: %v", err)
	}
	defer conn.Close()

	client := NewTicTacToeServerClient(conn)
	stream, err := client.Play(context.Background())
	if err != nil {
		return fmt.Errorf("error while establishing stream with game server: %v", err)
	}

	waitc := make(chan struct{})

	// Receiver loop - receive and print messages from server
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
			}
			if err != nil {
				log.Fatalf("Failed to receive: %v", err)
			}
			log.Printf("Got response %v", in)
		}
	}()

	// Sender loop - send actions to server
	log.Println("Connection to game server established. Send commands")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		if text == "exit" {
			if err := stream.CloseSend(); err != nil {
				return fmt.Errorf("error while closing stream: %v", err)
			}
		}
		action := Action{
			Text: text,
		}
		if err := stream.Send(&action); err != nil {
			return fmt.Errorf("error while sending action: %v", err)
		}
	}
}
