package issue

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type saveIssueRequest struct {
	Title   string `json:"Title"`
	Content string `json:"Content"`
}
type saveIssueResponse struct{}

type getIssueRequest struct{}
type getIssueResponse = Issue

type allIssueRequest struct{}
type allIssueResponse struct {
	Data []Issue `json:"data"`
}

func CreateIssueSaveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		params := request.(saveIssueRequest)

		issue := Issue{
			Title:   params.Title,
			Content: params.Content,
		}

		err := s.Save(ctx, issue)

		return saveIssueResponse{}, err
	}
}

func CreateIssueAllEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		issues, err := s.All(ctx)

		if err != nil {
			return issues, err
		}

		response := allIssueResponse{Data: issues}
		return response, nil
	}
}
