package domain

import (
	"context"
	"fmt"
)

func (s *Service) Classrooms(_ context.Context) ([]Classroom, error) {
	cc, err := s.classrooms.Classrooms()
	if err != nil {
		return []Classroom{}, fmt.Errorf("service didn't found any classroom: %w", err)
	}

	return cc, nil
}

func (s *Service) CreateClassroom(_ context.Context, classroom *Classroom) error {
	if err := s.classrooms.CreateClassroom(classroom); err != nil {
		return fmt.Errorf("service can't create classroom: %w", err)
	}

	return nil
}
