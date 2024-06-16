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

func (s *Service) CreateClassroom(ctx context.Context, classroom *Classroom) error {
	isActiveCourse, err := s.courses.IsActiveCourse(ctx, classroom.CourseUUID)

	if err != nil {
		return fmt.Errorf("can't verify course uuid: %w", err)
	}

	if !isActiveCourse {
		return fmt.Errorf("course uuid not found or inactive")
	}

	if err := s.classrooms.CreateClassroom(classroom); err != nil {
		return fmt.Errorf("service can't create classroom: %w", err)
	}

	return nil
}
