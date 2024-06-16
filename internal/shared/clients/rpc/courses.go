package courses

import (
	"context"

	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/shared/clients/rpc/pb"
)

type CoursesService struct {
	Client pb.CoursesClient
}

func NewCoursesService(client pb.CoursesClient) *CoursesService {
	return &CoursesService{Client: client}
}

func (courses CoursesService) IsActiveCourse(ctx context.Context, courseUUID uuid.UUID) (bool, error) {

	request := &pb.CourseRequest{
		CourseUUID: courseUUID.String(),
	}

	res, err := courses.Client.IsActiveCourse(context.Background(), request)

	return res.GetIsActive(), err
}
