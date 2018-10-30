// Package tictactoe implements interactive tictactoe game for two players
package tictactoe

import (
	fmt "fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

// ServerPort defines port server is running on
const ServerPort = 7000

type ticTacToeServer struct {
}

// StartServer starts game server
func StartServer() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", ServerPort))
	if err != nil {
		return fmt.Errorf("unable to start server: %v", err)
	}
	grpcServer := grpc.NewServer()
	RegisterTicTacToeServerServer(grpcServer, &ticTacToeServer{})
	log.Printf("Starting game server on port %v", ServerPort)
	grpcServer.Serve(lis)
	return nil
}

func (s *ticTacToeServer) Play(stream TicTacToeServer_PlayServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("error while receiving: %v", err)
		}
		action := in.GetText()
		log.Printf("Received command: %v", action)
		result := Result{
			Message: fmt.Sprintf("%s action received", action),
		}
		if err := stream.Send(&result); err != nil {
			log.Printf("Error while sending to client: %v", err)
		}
	}
}
