package usecases

import (
	"context"
	"kumparan-backend-position-interview/bin/modules/articles/models/binding"
	"kumparan-backend-position-interview/bin/modules/articles/repositories"
	"kumparan-backend-position-interview/bin/pkg/utils"
	"time"

	"github.com/google/uuid"
)

type ArticleCommandUsecase struct {
	commandDb repositories.DbCommandInterface
	queryDb   repositories.DbQueryInterface
}

type ArticleCommandUsecaseInterface interface {
	Create(ctx context.Context, payload *binding.Create) utils.Result
}

// NewArticleCommandUsecase creates a new instance of ArticleCommandUsecase.
// It requires two parameters:
//   - commandPg: An implementation of the PostgreCommandInterface, which is responsible for executing commands
//     against the PostgreSQL database, such as inserting or updating articles.
//   - queryPg: An implementation of the PostgreQueryInterface, which is responsible for querying the PostgreSQL
//     database, allowing for retrieval of articles or related data before performing insert operations.
//
// This function returns a pointer to an ArticleCommandUsecase instance, which can be used to handle article
// commands in the application.
func NewArticleCommandUsecase(commandPg repositories.DbCommandInterface, queryPg repositories.DbQueryInterface) *ArticleCommandUsecase {
	return &ArticleCommandUsecase{commandDb: commandPg, queryDb: queryPg}
}

func (c *ArticleCommandUsecase) Create(ctx context.Context, payload *binding.Create) (result utils.Result) {
	
	params := map[string]any{
		"id":         uuid.NewString(),
		"title":      payload.Title,
		"body":       payload.Body,
		"author_id":  payload.AuthorID,
		"created_at": time.Now(),
	}

	result = <-c.commandDb.Create(params)
	if result.Error != nil {
		return
	}
	return
}
