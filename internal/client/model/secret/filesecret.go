package secret

import "time"

type FileSecret struct {
	Id         int `json:"-"`
	Title      string
	RecordType int
	Path       string
	Binary     []byte
	UpdatedAt  time.Time `json:"-"`
	IsDelited  bool      `json:"-"`
}

func (fs FileSecret) GetUpdateTime() time.Time {
	return fs.UpdatedAt
}
