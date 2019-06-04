package issue

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	transport "github.com/go-kit/kit/transport/http"
	"net/http"
)

type errorer interface {
	error() error
}

func MakeHandler(service *service) http.Handler {
	//opts := []kithttp.ServerOption{
	//	kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	//	kithttp.ServerErrorEncoder(encodeError),
	//}

	issueSaveHandler := transport.NewServer(
		CreateIssueSaveEndpoint(service),
		decodeSaveIssueRequest,
		encodeResponse,
	)

	issueAllHandler := transport.NewServer(
		CreateIssueAllEndpoint(service),
		decodeRequest,
		encodeResponse,
	)

	r := chi.NewRouter()

	r.Method(http.MethodPost, "/", issueSaveHandler)
	r.Method(http.MethodGet, "/", issueAllHandler)

	return r
}

func decodeSaveIssueRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var params saveIssueRequest
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return nil, err
	}

	return params, nil
}

func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
