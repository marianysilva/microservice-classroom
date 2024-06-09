package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sumelms/microservice-classroom/internal/classroom/domain"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/google/uuid"

	"github.com/sumelms/microservice-classroom/pkg/validator"
)

type CreateClassroomRequest struct {
	Code         string     `json:"code" validate:"required,max=15"`
	SubjectUUID  *uuid.UUID `json:"subject_uuid"`
	CourseUUID   uuid.UUID  `json:"course_uuid" validate:"required"`
	Name         string     `json:"name" validate:"required,max=100"`
	CabSubscribe bool       `json:"cab_subscribe"`
	Description  string     `json:"description" validate:"max=255"`
	Format       string     `json:"format"`
	StartsAt     time.Time  `json:"starts_at" validate:"required"`
	EndsAt       *time.Time `json:"ends_at"`
}

type ClassroomResponse struct {
	UUID         uuid.UUID  `json:"uuid"`
	Code         string     `json:"code"`
	SubjectUUID  *uuid.UUID `json:"subject_uuid,omitempty"`
	CourseUUID   uuid.UUID  `json:"course_uuid"`
	Name         string     `json:"name"`
	CanSubscribe bool       `json:"can_subscribe"`
	Description  string     `json:"description"`
	Format       string     `json:"format"`
	StartsAt     time.Time  `json:"starts_at" validate:"required"`
	EndsAt       *time.Time `json:"ends_at,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type CreateClassroomResponse struct {
	Classroom *ClassroomResponse `json:"classroom"`
}

// NewCreateClassroomHandler creates new course handler
// @Summary      Create classroom
// @Description  Create a new classroom
// @Tags         classrooms
// @Accept       json
// @Produce      json
// @Param        classroom	  body		CreateClassroomRequest		true	"Add Classroom"
// @Success      200      {object}  CreateClassroomResponse
// @Failure      400      {object}  error
// @Failure      404      {object}  error
// @Failure      500      {object}  error
// @Router       /classrooms [post].
func NewCreateClassroomHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeCreateClassroomEndpoint(s),
		decodeCreateClassroomRequest,
		encodeCreateClassroomResponse,
		opts...,
	)
}

func makeCreateClassroomEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateClassroomRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		v := validator.NewValidator()
		if err := v.Validate(req); err != nil {
			return nil, err
		}

		classroom := domain.Classroom{}
		data, _ := json.Marshal(req)
		if err := json.Unmarshal(data, &classroom); err != nil {
			return nil, err
		}

		if err := s.CreateClassroom(ctx, &classroom); err != nil {
			return nil, err
		}

		return &CreateClassroomResponse{
			Classroom: &ClassroomResponse{
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
			},
		}, nil
	}
}

func decodeCreateClassroomRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateClassroomRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateClassroomResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
