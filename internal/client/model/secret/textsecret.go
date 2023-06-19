package secret

import "time"

type TextSecret struct {
	Id         int `json:"-"`
	Title      string
	RecordType int
	Text       string
	UpdatedAt  time.Time `json:"-"`
	IsDelited  bool      `json:"-"`
}

func (ts TextSecret) GetUpdateTime() time.Time {
	return ts.UpdatedAt
}
