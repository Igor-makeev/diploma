package storage

import (
	"errors"
	"sync"

	"secretKeeper/internal/client/model/secret"
	"secretKeeper/proto"
)

type Memorier interface {
	GetLoginPassSecret(id int) (secret.LoginPassSecret, bool, error)
	GetCardSecret(id int) (secret.CardSecret, bool, error)
	GetTextSecret(id int) (secret.TextSecret, bool, error)
	FindInStorage(id int) (interface{}, bool)
	GetSecretList(id int) []*proto.SecretList
	DeleteSecret(id int)
	ResetStorage()
}

type DataEditor interface {
	SetLoginPassSecrets([]secret.LoginPassSecret)
	SetCardSecrets([]secret.CardSecret)
	SetTextSecrets([]secret.TextSecret)
}

type MemoryStorage struct {
	mu               sync.RWMutex
	LoginPassSecrets map[int]secret.LoginPassSecret
	TextSecrets      map[int]secret.TextSecret
	CardSecrets      map[int]secret.CardSecret
}

// NewMemoryStorage - creates new MemoryStorage.
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		LoginPassSecrets: make(map[int]secret.LoginPassSecret, 0),
		TextSecrets:      make(map[int]secret.TextSecret, 0),
		CardSecrets:      make(map[int]secret.CardSecret, 0),
	}

}

// GetLoginPassSecret - attempts to return model.LoginPassSecret from MemoryStorage.
//
// If record is found, then returns model.LoginPassSecret and true as representation of found record.
//
// If nothing is found, then returns empty model.LoginPassSecret and false as representation of not found record.
func (ms *MemoryStorage) GetLoginPassSecret(id int) (secret.LoginPassSecret, bool, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	data, ok := ms.LoginPassSecrets[id]
	if !ok {
		return secret.LoginPassSecret{}, ok, errors.New("Login/Pass not found")
	}

	return data, ok, nil
}

// SetLoginPassSecrets - Sets []model.LoginPassSecret to MemoryStorage.
func (ms *MemoryStorage) SetLoginPassSecrets(models []secret.LoginPassSecret) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	for _, m := range models {
		ms.LoginPassSecrets[m.Id] = m
	}
}

// ResetStorage - removes all records from MemoryStorage.
func (ms *MemoryStorage) ResetStorage() {
	ms.LoginPassSecrets = make(map[int]secret.LoginPassSecret, 0)
	ms.TextSecrets = make(map[int]secret.TextSecret, 0)
	ms.CardSecrets = make(map[int]secret.CardSecret, 0)
}

// SetCardSecrets - Sets []model.CardSecret to MemoryStorage.
func (ms *MemoryStorage) SetCardSecrets(models []secret.CardSecret) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	for _, m := range models {
		ms.CardSecrets[m.Id] = m
	}
}

// GetCardSecret - attempts to return model.CardSecret from MemoryStorage.
//
// If record is found, then returns model.CardSecret and true as representation of found record.
//
// If nothing is found, then returns empty model.CardSecret and false as representation of not found record.
func (ms *MemoryStorage) GetCardSecret(id int) (secret.CardSecret, bool, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	data, ok := ms.CardSecrets[id]
	if !ok {
		return secret.CardSecret{}, ok, errors.New("card data not found")
	}

	return data, ok, nil
}

// SetTextSecrets - Sets []model.TextSecret to MemoryStorage.
func (ms *MemoryStorage) SetTextSecrets(models []secret.TextSecret) {
	ms.mu.RLock()
	defer ms.mu.RUnlock()

	for _, m := range models {
		ms.TextSecrets[m.Id] = m
	}
}

// GetTextSecret - attempts to return model.TextSecret from MemoryStorage.
//
// If record is found, then returns model.TextSecret and true as representation of found record.
//
// If nothing is found, then returns empty model.TextSecret and false as representation of not found record.
func (ms *MemoryStorage) GetTextSecret(id int) (secret.TextSecret, bool, error) {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	data, ok := ms.TextSecrets[id]
	if !ok {
		return secret.TextSecret{}, ok, errors.New("text not found")
	}

	return data, ok, nil
}

// FindInStorage - attempts to find record in MemoryStorage by provided id.
//
// If record is found in LoginPassSecrets then returns model.LoginPassSecret and true as representation of found record.
//
// If record is found in TextSecrets then returns model.TextSecret and true as representation of found record.
//
// If record is found in CardSecrets then returns model.CardSecret and true as representation of found record.
//
// Otherwise, returns nil and false as representation of not found record.
func (ms *MemoryStorage) FindInStorage(id int) (interface{}, bool) {
	data, ok, _ := ms.GetLoginPassSecret(id)
	if ok {
		if data.IsDelited {
			return "sorry secret has been delited", true
		}
		return data, true
	}

	dataCard, okCard, _ := ms.GetCardSecret(id)
	if okCard {
		if dataCard.IsDelited {
			return "sorry secret has been delited", true
		}
		return dataCard, true
	}

	textData, okText, _ := ms.GetTextSecret(id)
	if okText {
		if textData.IsDelited {
			return "sorry secret has been delited", true
		}
		return textData, true
	}

	return nil, false
}

// GetSecretList - goes through all MemoryStorage records and returns []*GetSecretList.
func (ms *MemoryStorage) GetSecretList(id int) []*proto.SecretList {
	ms.mu.Lock()
	defer ms.mu.Unlock()

	var list []*proto.SecretList
	switch id {
	case 1:
		for _, data := range ms.LoginPassSecrets {
			list = append(list, &proto.SecretList{Id: uint32(data.Id), Title: data.Title})
		}
	case 2:
		for _, data := range ms.TextSecrets {
			list = append(list, &proto.SecretList{Id: uint32(data.Id), Title: data.Title})
		}
	case 4:
		for _, data := range ms.CardSecrets {
			list = append(list, &proto.SecretList{Id: uint32(data.Id), Title: data.Title})
		}
	}

	return list
}

func (ms *MemoryStorage) DeleteSecret(id int) {
	_, ok, _ := ms.GetLoginPassSecret(id)
	if ok {
		ms.LoginPassSecrets[id] = secret.LoginPassSecret{}
	}

	_, okCard, _ := ms.GetCardSecret(id)
	if okCard {
		ms.CardSecrets[id] = secret.CardSecret{}
	}

	_, okText, _ := ms.GetTextSecret(id)
	if okText {
		ms.TextSecrets[id] = secret.TextSecret{}
	}

}
