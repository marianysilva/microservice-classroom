package clients

import (
	log "github.com/sumelms/microservice-classroom/pkg/logger"

	"github.com/sumelms/microservice-classroom/internal/shared"
	courses "github.com/sumelms/microservice-classroom/internal/shared/clients/rpc"
	"github.com/sumelms/microservice-classroom/internal/shared/clients/rpc/pb"
	"github.com/sumelms/microservice-classroom/pkg/config"
)

type ClientServices struct {
	Courses courses.CoursesService
}

func NewCoursesClient(cfg *config.Config, logger log.Logger) (pb.CoursesClient, error) {
	// Create RPC Server
	coursesRPCConnection, err := shared.NewRPCConnection(cfg.RPCClients.Courses, logger)
	if err != nil {
		logger.Log("msg", "unable to connect with courses RPC server: ", "error", err)
		return nil, err
	}

	client := pb.NewCoursesClient(coursesRPCConnection.Connection)
	return client, err
}
