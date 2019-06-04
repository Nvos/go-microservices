package issue

import (
	"context"
	"github.com/satori/go.uuid"
	"sort"
)

type InMemoryStore struct {
	issues map[string]Issue
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		issues: map[string]Issue{},
	}
}

func (store *InMemoryStore) Save(ctx context.Context, issue Issue) error {
	issue.ID = uuid.NewV4().String()
	store.issues[issue.ID] = issue

	return nil
}

func (store *InMemoryStore) Get(ctx context.Context, id string) (Issue, error) {
	issue, ok := store.issues[id]

	if !ok {
		return issue, NotFoundError{ID: id}
	}

	return issue, nil
}

func (store *InMemoryStore) All(ctx context.Context) ([]Issue, error) {
	issues := make([]Issue, len(store.issues))
	keys := make([]string, len(store.issues))
	i := 0

	for key := range store.issues {
		keys[i] = key
		i++
	}

	sort.Strings(keys)

	for i, key := range keys {
		issues[i] = store.issues[key]
	}

	return issues, nil
}
