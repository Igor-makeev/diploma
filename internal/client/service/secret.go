package service

import (
	"errors"
	"fmt"
	"os"

	"google.golang.org/protobuf/types/known/timestamppb"

	"secretKeeper/internal/client/model"
	"secretKeeper/internal/client/model/secret"
	"secretKeeper/internal/client/storage"
	"secretKeeper/pkg/apperr"
	"secretKeeper/pkg/crypt"
	pb "secretKeeper/proto"
)

type SecretClientService struct {
	glCtx   *model.GlobalContext
	client  pb.SecretClient
	storage storage.Memorier
	crypt   crypt.Crypter
	syncer  storage.Syncer
}

// NewSecretClientService - creates new SecretClientService.
func NewSecretClientService(
	glCtx *model.GlobalContext, client pb.SecretClient, st storage.Memorier, cr crypt.Crypter, sr storage.Syncer,
) *SecretClientService {
	return &SecretClientService{
		glCtx:   glCtx,
		client:  client,
		storage: st,
		crypt:   cr,
		syncer:  sr,
	}
}

// GetListOfSecretes - attempts to return list of secrets from memory, if nothing is found then makes gRPC request
// to server.
func (s *SecretClientService) GetListOfSecretes(id int) ([]*pb.SecretList, error) {
	var list []*pb.SecretList
	list = s.storage.GetSecretList(id)
	if len(list) > 0 {
		return list, nil
	}

	result, err := s.client.GetListOfSecretsByType(s.glCtx.Ctx, &pb.GetListOfSecretsByTypeRequest{TypeId: uint32(id)})
	if err != nil {
		return nil, err
	}

	return result.SecretLists, nil
}

// GetBinarySecret - get binary data from server and stores it into file.
func (s *SecretClientService) GetBinarySecret(id int, location string) error {
	res, err := s.client.GetSecret(s.glCtx.Ctx, &pb.GetSecretRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	if res.Type != 3 {
		return errors.New("this method only works with binary data, please appropriate method next time")
	}

	decoded, errDecode := s.crypt.Decode(string(res.Content))
	if errDecode != nil {
		return errDecode
	}

	f, openErr := os.OpenFile(location, os.O_CREATE|os.O_WRONLY, 0644)
	if openErr != nil {
		return openErr
	}
	defer f.Close()

	_, wrError := f.Write([]byte(decoded))
	if wrError != nil {
		return wrError
	}

	fmt.Println("data written to file")

	return nil
}

// GetSecret -  makes gRPC request to server.
func (s *SecretClientService) GetSecret(id int) (secret.ResSecret, error) {

	result, err := s.client.GetSecret(s.glCtx.Ctx, &pb.GetSecretRequest{Id: int32(id)})
	if err != nil {
		return secret.ResSecret{}, err
	}

	if result.Type == 3 {
		return secret.ResSecret{}, errors.New("to get binary data, pleas use proper method")
	}

	decoded, errDecode := s.crypt.Decode(string(result.Content))
	if errDecode != nil {
		return secret.ResSecret{}, errDecode
	}

	if result.IsDelited {

		return secret.ResSecret{}, apperr.ErrSecretNotFound

	}

	return secret.ResSecret{
		UpdatedAt: result.UpdatedAt.AsTime(),
		Content:   decoded,
	}, nil
}

// CreateSecret - creates new secret on the server and then makes re-sync memory storage.
func (s *SecretClientService) CreateSecret(title string, recordType int, content string) error {
	contentT := []byte(s.crypt.Encode(content))

	result, err := s.client.CreateSecret(s.glCtx.Ctx, &pb.CreateSecretRequest{
		Title:   title,
		Type:    uint32(recordType),
		Content: contentT,
	})

	if err != nil {
		return err
	}

	fmt.Println("created new secret with ID:", result.Id)

	s.syncer.SyncAll()

	return nil
}

// DeleteSecret - deletes a secrete from server and then makes re-sync memory storage.
func (s *SecretClientService) DeleteSecret(id int) error {
	s.storage.DeleteSecret(id)
	_, err := s.client.DeleteSecret(s.glCtx.Ctx, &pb.DeleteSecretRequest{Id: uint32(id)})
	if err != nil {
		return err
	}

	fmt.Println("successfully deleted secret")

	// s.storage.ResetStorage()
	// s.syncer.SyncAll()

	return nil
}

// EditSecret - edits secret on the server and then makes re-sync memory storage.
func (s *SecretClientService) EditSecret(id int, title string, recordType int, content string, isForce bool) error {

	localSecret, _ := s.GetSecret(id)

	contentT := []byte(s.crypt.Encode(content))

	_, err := s.client.EditSecret(
		s.glCtx.Ctx, &pb.EditSecretRequest{
			Id:        uint32(id),
			Title:     title,
			Type:      uint32(recordType),
			Content:   contentT,
			UpdatedAt: timestamppb.New(localSecret.UpdatedAt),
			IsForce:   isForce,
		},
	)
	if err != nil {
		return err
	}

	fmt.Println("successfully edited secret")

	s.syncer.SyncAll()

	return nil
}
