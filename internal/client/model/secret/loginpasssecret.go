package secret

import "time"

type LoginPassSecret struct {
	Id         int `json:"-"`
	Title      string
	RecordType int
	Login      string
	Password   string
	UpdatedAt  time.Time `json:"-"`
	IsDelited  bool      `json:"-"`
}

func (lps LoginPassSecret) GetUpdateTime() time.Time {
	return lps.UpdatedAt
}
