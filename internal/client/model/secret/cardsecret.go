package secret

import "time"

type CardSecret struct {
	Id         int `json:"-"`
	Title      string
	RecordType int
	CardNumber string
	CVV        string
	Due        string
	UpdatedAt  time.Time `json:"-"`
	IsDelited  bool      `json:"-"`
}

func (cs CardSecret) GetUpdateTime() time.Time {
	return cs.UpdatedAt
}
