package repositories

import (
	"kumparan-backend-position-interview/bin/pkg/utils"
	"strings"

	"gorm.io/gorm"
)

type PostgreQuery struct {
	db *gorm.DB
}

type DbQueryInterface interface {
	GetList(payload *QueryPayload) <-chan utils.Result
	Search(payload *QueryPayload) <-chan utils.Result
}

type QueryPayload struct {
	Table     string
	Query     string
	Parameter map[string]interface{}
	Where     map[string]interface{}
	Select    string
	Join      string
	Limit     int
	Offset    int
}

func NewDBQuery(db *gorm.DB) DbQueryInterface {
	return &PostgreQuery{
		db: db,
	}
}

func (c *PostgreQuery) GetList(payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var query = c.db.Table(payload.Table)

		if payload.Select != "" {
			query = query.Select(payload.Select)
		}

		if payload.Join != "" {
			query = query.Joins(payload.Join)
		}

		if payload.Where != nil {
			query = query.Where(payload.Where)
		}

		
		var result []map[string]interface{}
		if err := query.Find(&result).Error; err != nil {
			output <- utils.Result{Error: err}
			return
		}

		output <- utils.Result{Data: result}
	}()

	return output
}

func (c *PostgreQuery) Search(payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		query := c.db.Debug().Table(payload.Table)

		// Apply SELECT
		if payload.Select != "" {
			query = query.Select(payload.Select)
		}

		// Apply WHERE map
		if payload.Where != nil {
			query = query.Where(payload.Where)
		}

		// Apply LIMIT & OFFSET
		if payload.Limit > 0 {
			query = query.Limit(payload.Limit)
		}
		if payload.Offset > 0 {
			query = query.Offset(payload.Offset)
		}

		// Prepare LIKE conditions if title/body present
		conditions := []string{}
		args := []interface{}{}

		if titleRaw, ok := payload.Parameter["title"]; ok {
			if titleStr, ok := titleRaw.(string); ok && titleStr != "" {
				conditions = append(conditions, "title LIKE ?")
				args = append(args, "%"+titleStr+"%")
			}
		}

		if bodyRaw, ok := payload.Parameter["body"]; ok {
			if bodyStr, ok := bodyRaw.(string); ok && bodyStr != "" {
				conditions = append(conditions, "body LIKE ?")
				args = append(args, "%"+bodyStr+"%")
			}
		}

		// Combine conditions with OR
		if len(conditions) > 0 {
			query = query.Where(strings.Join(conditions, " OR "), args...)
		}

		// Execute query
		var result []map[string]interface{}
		if err := query.Find(&result).Error; err != nil {
			output <- utils.Result{Error: err}
			return
		}

		output <- utils.Result{Data: result}
	}()
	return output
}
