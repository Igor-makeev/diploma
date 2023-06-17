package secret

import "time"

type ResSecret struct {
	UpdatedAt time.Time

	Content string
}

type Secret interface {
	GetUpdateTime() time.Time
}

type SecretList struct {
	Id    int
	Title string
}
