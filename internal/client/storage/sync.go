package storage

import (
	"encoding/json"
	"fmt"

	"secretKeeper/internal/client/model"
	"secretKeeper/internal/client/model/secret"
	"secretKeeper/pkg/crypt"
	pb "secretKeeper/proto"
)

type Syncer interface {
	SyncAll()
	SyncPassLoginData() error
	SyncCardData() error
	SyncTextData() error
}

type Sync struct {
	storage      DataEditor
	secretClient pb.SecretClient
	glCtx        *model.GlobalContext
	cr           crypt.Crypter
}

// NewSync - creates new Sync.
func NewSync(de DataEditor, sc pb.SecretClient, ctx *model.GlobalContext, cr crypt.Crypter) *Sync {
	return &Sync{storage: de, secretClient: sc, glCtx: ctx, cr: cr}
}

// SyncAll - runs SyncTextData, SyncPassLoginData, SyncCardData under the hood.
func (s *Sync) SyncAll() {
	if err := s.SyncTextData(); err != nil {
		fmt.Println(err)
	}

	if err := s.SyncPassLoginData(); err != nil {
		fmt.Println(err)
	}

	if err := s.SyncCardData(); err != nil {
		fmt.Println(err)
	}
}

// SyncTextData - makes gRPC request to server and on success sets acquired records to MemoryStorage.TextSecrets.
func (s *Sync) SyncTextData() error {
	texts, err := s.secretClient.GetListOfSecretsByType(s.glCtx.Ctx, &pb.GetListOfSecretsByTypeRequest{TypeId: 2})
	if err != nil {
		panic(err)
	}

	var list []secret.TextSecret
	for _, text := range texts.SecretLists {
		id := int(text.Id)
		m := secret.TextSecret{}

		decoded, errDecode := s.cr.Decode(string(text.Content))
		if errDecode != nil {
			return errDecode
		}

		errUnmarshal := json.Unmarshal([]byte(decoded), &m)
		if errUnmarshal != nil {
			return errUnmarshal
		}
		m.Id = id
		m.UpdatedAt = text.UpdatedAt.AsTime()

		list = append(list, m)
	}

	s.storage.SetTextSecrets(list)

	return nil
}

// SyncCardData - makes gRPC request to server and on success sets acquired records to MemoryStorage.CardSecrets.
func (s *Sync) SyncCardData() error {
	cards, err := s.secretClient.GetListOfSecretsByType(s.glCtx.Ctx, &pb.GetListOfSecretsByTypeRequest{TypeId: 4})
	if err != nil {
		panic(err)
	}

	var list []secret.CardSecret
	for _, card := range cards.SecretLists {
		id := int(card.Id)
		m := secret.CardSecret{}

		decoded, errDecode := s.cr.Decode(string(card.Content))
		if errDecode != nil {
			return errDecode
		}

		errUnmarshal := json.Unmarshal([]byte(decoded), &m)
		if errUnmarshal != nil {
			return errUnmarshal
		}
		m.Id = id
		m.UpdatedAt = card.UpdatedAt.AsTime()

		list = append(list, m)
	}

	s.storage.SetCardSecrets(list)

	return nil
}

// SyncPassLoginData - makes gRPC request to server and on success sets acquired records to
// MemoryStorage.LoginPassSecrets.
func (s *Sync) SyncPassLoginData() error {
	lists, err := s.secretClient.GetListOfSecretsByType(s.glCtx.Ctx, &pb.GetListOfSecretsByTypeRequest{TypeId: 1})
	if err != nil {
		panic(err)
	}

	var list []secret.LoginPassSecret
	for _, sList := range lists.SecretLists {
		id := int(sList.Id)
		m := secret.LoginPassSecret{}

		decoded, errDecode := s.cr.Decode(string(sList.Content))
		if errDecode != nil {
			return errDecode
		}

		errUnmarshal := json.Unmarshal([]byte(decoded), &m)
		if errUnmarshal != nil {
			return errUnmarshal
		}
		m.Id = id
		m.UpdatedAt = sList.UpdatedAt.AsTime()

		list = append(list, m)
	}

	s.storage.SetLoginPassSecrets(list)

	return nil
}
