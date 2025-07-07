package usecases

import (
	"context"
	"fmt"
	"kumparan-backend-position-interview/bin/modules/articles/models/binding"
	"kumparan-backend-position-interview/bin/modules/articles/repositories"
	"kumparan-backend-position-interview/bin/pkg/utils"
)

type ArticleQueryUsecase struct {
	commandDb repositories.DbCommandInterface
	queryDb   repositories.DbQueryInterface
}

type ArticleQueryUsecaseInterface interface {
	GetList(ctx context.Context, filtered map[string]any) utils.Result
	Search(ctx context.Context, payload binding.Search) utils.Result
}

// NewArticleQueryUsecase creates a new instance of ArticleQueryUsecase.
// It requires two parameters:
//   - QueryPg: An implementation of the PostgreQueryInterface, which is responsible for executing Querys
//     against the PostgreSQL database, such as inserting or updating articles.
//   - queryPg: An implementation of the PostgreQueryInterface, which is responsible for querying the PostgreSQL
//     database, allowing for retrieval of articles or related data before performing insert operations.
//
// This function returns a pointer to an ArticleQueryUsecase instance, which can be used to handle article
// Querys in the application.
func NewArticleQueryUsecase(command repositories.DbCommandInterface, query repositories.DbQueryInterface) *ArticleQueryUsecase {
	return &ArticleQueryUsecase{commandDb: command, queryDb: query}
}

func (c *ArticleQueryUsecase) GetList(ctx context.Context, filtered map[string]any) (result utils.Result) {
	QueryPy := &repositories.QueryPayload{
		Table: "articles",
		Query: "SELECT * FROM articles",
		Where: filtered,
	}

	result = <-c.queryDb.GetList(QueryPy)
	if result.Error != nil {
		return
	}
	return
}

func (c *ArticleQueryUsecase) Search(ctx context.Context, payload binding.Search) (result utils.Result) {
	var filtered = make(map[string]any)

	if payload.Title != "" {
		filtered["title"] = payload.Title
	}
	if payload.Body != "" {
		filtered["body"] = payload.Body
	}
	
	QueryPy := &repositories.QueryPayload{
		Table:     "articles",
		Query:     "SELECT * FROM articles",
		Parameter: filtered,
	}

	fmt.Println("Filtered:", filtered)

	result = <-c.queryDb.Search(QueryPy)
	if result.Error != nil {
		return
	}
	return
}
