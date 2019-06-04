package issue

import "context"

type service struct {
	store Store
}

type Service interface {
	Get(ctx context.Context, id string) (Issue, error)
	Save(ctx context.Context, issue Issue) error
	All(ctx context.Context) ([]Issue, error)
}

func NewIssueService(store Store) *service {
	return &service{
		store: store,
	}
}

func (s service) Get(ctx context.Context, id string) (Issue, error) {
	return s.store.Get(ctx, id)
}

func (s service) Save(ctx context.Context, issue Issue) error {
	return s.store.Save(ctx, issue)
}

func (s service) All(ctx context.Context) ([]Issue, error) {
	return s.store.All(ctx)
}
