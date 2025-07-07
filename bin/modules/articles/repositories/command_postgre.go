package repositories

import (
	"kumparan-backend-position-interview/bin/pkg/utils"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type PostgreCommand struct {
	db *gorm.DB
}

type DbCommandInterface interface {
	Create(params map[string]any) <-chan utils.Result
}

type CommandPayload struct {
	Table     string
	Command   string
	Parameter map[string]interface{}
	Where     map[string]interface{}
}

func NewDBCommand(db *gorm.DB) DbCommandInterface {
	return &PostgreCommand{
		db: db,
	}
}

func (c *PostgreCommand) Create(params map[string]any) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		id := rand.Intn(10000)
		title, _ := params["title"].(string)
		body, _ := params["body"].(string)
		authorID, _ := params["author_id"].(string)
		createdAt := time.Now()

		query := `
			INSERT INTO articles (id, title, body, author_id, created_at)
			VALUES (?, ?, ?, ?, ?)
		`

		err := c.db.Exec(query, id, title, body, authorID, createdAt).Error
		if err != nil {
			output <- utils.Result{Error: err}
			return
		}

		output <- utils.Result{Data: params}
	}()

	return output
}
