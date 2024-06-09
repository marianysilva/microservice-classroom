package http

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
	"github.com/sumelms/microservice-classroom/internal/classroom/endpoints"
	"github.com/sumelms/microservice-classroom/pkg/errors"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	classroomRouter := NewClassroomRouter(s, opts...)

	r.PathPrefix("/classrooms").Handler(classroomRouter)
}

func NewClassroomRouter(s domain.ServiceInterface, opts ...kithttp.ServerOption) *mux.Router {
	r := mux.NewRouter().PathPrefix("/classrooms").Subrouter().StrictSlash(true)

	createClassroomHandler := endpoints.NewCreateClassroomHandler(s, opts...)
	r.Handle("", createClassroomHandler).Methods(http.MethodPost)

	listClassroomHandler := endpoints.NewListClassroomsHandler(s, opts...)
	r.Handle("", listClassroomHandler).Methods(http.MethodGet)

	return r
}
