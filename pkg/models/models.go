package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: suitable quote doesn't found")

type Quote struct {
	ID int
	Author string
	Quote string
	Source string
	SourceType string
	Created time.Time
}