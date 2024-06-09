package endpoints

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/internal/classroom/domain"
)

type ListClassroomsResponse struct {
	Classrooms []ClassroomResponse `json:"classrooms"`
}

// NewListClassroomsHandler list classrooms handler
// @Summary      List classrooms
// @Description  List a new classrooms
// @Tags         classrooms
// @Produce      json
// @Success      200      {object}  ListClassroomsResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /classrooms [get].
func NewListClassroomsHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeListClassroomsEndpoint(s),
		decodeListClassroomsRequest,
		encodeListClassroomsResponse,
		opts...,
	)
}

func makeListClassroomsEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		classrooms, err := s.Classrooms(ctx)
		if err != nil {
			return nil, err
		}

		var list []ClassroomResponse
		for i := range classrooms {
			classroom := classrooms[i]
			list = append(list, ClassroomResponse{
				UUID:         classroom.UUID,
				Code:         classroom.Code,
				SubjectUUID:  classroom.SubjectUUID,
				CourseUUID:   classroom.CourseUUID,
				Name:         classroom.Name,
				CanSubscribe: classroom.CanSubscribe,
				Description:  classroom.Description,
				Format:       classroom.Format,
				StartsAt:     classroom.StartsAt,
				EndsAt:       classroom.EndsAt,
				CreatedAt:    classroom.CreatedAt,
				UpdatedAt:    classroom.UpdatedAt,
			})
		}

		return &ListClassroomsResponse{Classrooms: list}, nil
	}
}

func decodeListClassroomsRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeListClassroomsResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
