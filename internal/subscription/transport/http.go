package transport

import (
	"net/http"

	kittransport "github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/sumelms/microservice-classroom/internal/subscription/endpoints"
	"github.com/sumelms/microservice-classroom/pkg/errors"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	"github.com/sumelms/microservice-classroom/internal/subscription/domain"
)

func NewHTTPHandler(r *mux.Router, s domain.ServiceInterface, logger log.Logger) {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(kittransport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(errors.EncodeError),
	}

	listSubscriptionHandler := endpoints.NewListSubscriptionHandler(s, opts...)
	createSubscriptionHandler := endpoints.NewCreateSubscriptionHandler(s, opts...)
	findSubscriptionHandler := endpoints.NewFindSubscriptionHandler(s, opts...)
	deleteSubscriptionHandler := endpoints.NewDeleteSubscriptionHandler(s, opts...)
	updateSubscriptionHandler := endpoints.NewUpdateSubscriptionHandler(s, opts...)

	r.Handle("/subscriptions", listSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions", createSubscriptionHandler).Methods(http.MethodPost)
	r.Handle("/subscriptions/{uuid}", findSubscriptionHandler).Methods(http.MethodGet)
	r.Handle("/subscriptions/{uuid}", deleteSubscriptionHandler).Methods(http.MethodDelete)
	r.Handle("/subscriptions/{uuid}", updateSubscriptionHandler).Methods(http.MethodPut)
}
