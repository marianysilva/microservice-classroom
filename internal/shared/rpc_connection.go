package shared

import (
	"fmt"

	"github.com/sumelms/microservice-classroom/pkg/config"
	"github.com/sumelms/microservice-classroom/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type RPCConnection struct {
	Connection *grpc.ClientConn
	Config     *config.RPCClient
}

func NewRPCConnection(cfg *config.RPCClient, logger logger.Logger) (*RPCConnection, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid server config")
	}

	conn, err := grpc.NewClient(cfg.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("unnable to connect with courses RPC server")
	}

	return &RPCConnection{
		Connection: conn,
		Config:     cfg,
	}, nil
}
