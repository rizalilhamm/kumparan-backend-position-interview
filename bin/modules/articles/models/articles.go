package models

import "time"

var Table = "articles"

type Articles struct {
	ID            string    `json:"id" gorm:"primaryKey;type:varchar(36)"`
	CreatedAt     int64     `json:"created_at"`
	CreatedBy     string    `json:"created_by" gorm:"type:varchar(36)"`
	UpdatedAt     int64     `json:"updated_at"`
	UpdatedBy     string    `json:"updated_by" gorm:"type:varchar(36)"`
	DeletedAt     int64     `json:"deleted_at"`
	DeletedBy     string    `json:"deleted_by" gorm:"type:varchar(36)"`
	Title         string    `json:"title" gorm:"not null"`
	Subtitle      string    `json:"subtitle"`
	PublishedAt   time.Time `json:"published_at"`
	Body          string    `json:"body" gorm:"not null"`
	Status        string    `json:"status" gorm:"type:varchar(36)"`
	StatusMessage string    `json:"status_message"`
}

func (a *Articles) TableName() string {
	return Table
}
