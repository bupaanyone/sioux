package model

import "time"

const RootId = 1

type Account struct {
	Id        *int64
	Username  *string
	Password  *string
	LastLogin *time.Time
}
