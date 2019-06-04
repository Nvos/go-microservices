package main

import (
	"github.com/go-chi/chi"
	"go-microservices/internal/issue"
	"net/http"
)

func main() {
	issueStore := issue.NewInMemoryStore()
	issueService := issue.NewIssueService(issueStore)

	issueHandler := issue.MakeHandler(issueService)

	r := chi.NewMux()
	r.Mount("/v1/issue", issueHandler)

	http.ListenAndServe(":8080", r)
}
