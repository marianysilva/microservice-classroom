package endpoints

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/endpoint"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/internal/subscription/domain"
)

type findSubscriptionRequest struct {
	ID string `json:"id"`
}

type findSubscriptionResponse struct {
	ID         uint       `json:"id"`
	UserID     string     `json:"user_id"`
	ClassroomID   string     `json:"classroom_id"`
	ValidUntil *time.Time `json:"valid_until,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

func NewFindSubscriptionHandler(s domain.ServiceInterface, opts ...kithttp.ServerOption) *kithttp.Server {
	return kithttp.NewServer(
		makeFindSubscriptionEndpoint(s),
		decodeFindSubscriptionRequest,
		encodeFindSubscriptionResponse,
		opts...,
	)
}

func makeFindSubscriptionEndpoint(s domain.ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(findSubscriptionRequest)
		if !ok {
			return nil, fmt.Errorf("invalid argument")
		}

		sub, err := s.FindSubscription(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return &findSubscriptionResponse{
			ID:         sub.ID,
			UserID:     sub.UserID,
			ClassroomID:   sub.ClassroomID,
			ValidUntil: sub.ValidUntil,
			CreatedAt:  sub.CreatedAt,
			UpdatedAt:  sub.UpdatedAt,
		}, nil
	}
}

func decodeFindSubscriptionRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, fmt.Errorf("invalid argument")
	}

	return findSubscriptionRequest{ID: id}, nil
}

func encodeFindSubscriptionResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return kithttp.EncodeJSONResponse(ctx, w, response)
}
