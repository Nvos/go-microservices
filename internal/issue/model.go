package issue

import "context"

type Issue struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NotFoundError struct {
	ID string
}

func (NotFoundError) Error() string {
	return "Issue"
}

type Store interface {
	Save(ctx context.Context, issue Issue) error
	Get(ctx context.Context, id string) (Issue, error)
	All(ctx context.Context) ([]Issue, error)
}
