package domain

import (
	"context"

	"github.com/go-kit/log"
)

type ServiceInterface interface {
	Classrooms(ctx context.Context) ([]Classroom, error)
	CreateClassroom(ctx context.Context, classroom *Classroom) error
}

type ServiceConfiguration func(svc *Service) error

type Service struct {
	classrooms ClassroomRepository
	logger     log.Logger
}

func NewService(cfgs ...ServiceConfiguration) (*Service, error) {
	svc := &Service{}
	for _, cfg := range cfgs {
		err := cfg(svc)
		if err != nil {
			return nil, err
		}
	}

	return svc, nil
}

func WithClassroomRepository(repository ClassroomRepository) ServiceConfiguration {
	return func(svc *Service) error {
		svc.classrooms = repository

		return nil
	}
}

func WithLogger(l log.Logger) ServiceConfiguration {
	return func(svc *Service) error {
		svc.logger = l

		return nil
	}
}
