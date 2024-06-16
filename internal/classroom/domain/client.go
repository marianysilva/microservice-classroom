package domain

import (
	"context"

	"github.com/google/uuid"
)

type CourseClientService interface {
	IsActiveCourse(ctx context.Context, courseUUID uuid.UUID) (bool, error)
}
