package models

import "time"

// make sure u install right package github.com/dmarkham/enumer
//
//go:generate enumer -type=UserStatus -sql -json
type UserStatus int8

const (
	// USER_INIT when start pipeline
	USER_INIT UserStatus = iota + 1

	// USER_SUCCESS when pipeline success
	USER_SUCCESS

	// USER_FAIL when pipeline fail
	USER_FAIL

	// USER_MUST_CLEAN wait a long time but can't get response from third party,
	// response error to client and need make cron job job remove from third party
	// after clean will change to USER_FAIL
	USER_MUST_CLEAN
)

type User struct {
	Id        int64      `gorm:"primaryKey, column:id" json:"id,omitempty"`
	Name      string     `gorm:"column:name" json:"name,omitempty"`
	ReqId     string     `gorm:"column:req_id" json:"req_id,omitempty"`
	RetryTime int        `gorm:"column:retry_time" json:"retry_time,omitempty"`
	Status    UserStatus `gorm:"column:status" json:"status,omitempty"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (User) TabaleName() string {
	return "users"
}
