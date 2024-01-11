package model

import "time"

type Invoice struct {
	Id   string
	Date time.Time
	Due  time.Time
}
