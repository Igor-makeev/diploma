package executor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"secretKeeper/internal/client/model"
	"strconv"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// createAuth - is executor for "create-auth" case in Execute method.
func (e *Executor) createAuth(args []string) error {
	switch len(args) - 1 {
	case 2:
		return fmt.Errorf("validation error: Password is missing")
	case 1:
		return fmt.Errorf("validation error: Login and Password is missing")
	case 0:
		return fmt.Errorf("validation error: Title, Login, Password is missing")
	}

	m := model.LoginPassSecret{
		Title:      args[1],
		RecordType: 1,
		Login:      args[2],
		Password:   args[3],
	}

	cont, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		return errMarshal
	}

	if err := e.app.SecretService.CreateSecret(m.Title, 1, string(cont)); err != nil {
		return err
	}

	return nil
}

// createText - is executor for "create-text" case in Execute method.
func (e *Executor) createText(args []string) error {
	switch len(args) - 1 {
	case 1:
		return fmt.Errorf("validation error: Text is missing")
	case 0:
		return fmt.Errorf("validation error: Title and Text is missing")
	}

	m := model.TextSecret{
		Title:      args[1],
		RecordType: 2,
		Text:       strings.Join(args[2:], " "),
	}

	marshal, errMarshal := json.Marshal(m)
	if errMarshal != nil {
		return errMarshal
	}

	if err := e.app.SecretService.CreateSecret(m.Title, 2, string(marshal)); err != nil {
		return err
	}

	return nil
}

// createBinary - is executor for "create-binary" case in Execute method.
func (e *Executor) createBinary(args []string) error {
	switch len(args) - 1 {
	case 1:
		return fmt.Errorf("validation error: Filepath is missing")
	case 0:
		return fmt.Errorf("validation error: Title and Filepath is missing")
	}

	m := model.FileSecret{
		Title:      args[1],
		RecordType: 3,
		Path:       args[2],
	}

	f, err := os.Open(m.Path)
	if err != nil {
		return err
	}
	defer f.Close()

	data, errData := ioutil.ReadAll(f)
	if errData != nil {
		return errData
	}

	errCreate := e.app.SecretService.CreateSecret(m.Title, m.RecordType, string(data))
	if errCreate != nil {
		return errCreate
	}

	return nil
}

// createCard - is executor for "create-card" case in Execute method.
func (e *Executor) createCard(args []string) error {
	switch len(args) - 1 {
	case 3:
		return fmt.Errorf("validation error: Due date is missing")
	case 2:
		return fmt.Errorf("validation error: CVV and Due date is missing")
	case 1:
		return fmt.Errorf("validation error: Card number, CVV, Due date is missing")
	case 0:
		return fmt.Errorf("validation error: Title, Card number, CVV, Due date is missing")
	}

	cardModel := model.CardSecret{
		Title:      args[1],
		RecordType: 4,
		CardNumber: args[2],
		CVV:        args[3],
		Due:        args[4],
	}

	cont, er := json.Marshal(cardModel)
	if er != nil {
		return er
	}

	if err := e.app.SecretService.CreateSecret(cardModel.Title, cardModel.RecordType, string(cont)); err != nil {
		return err
	}

	return nil
}

// deleteSecret - is executor for "delete-secret" case in Execute method.
func (e *Executor) deleteSecret(args []string) error {
	switch len(args) - 1 {
	case 0:
		return fmt.Errorf("validation error: Secret ID is missing")
	}

	id, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		return convErr
	}

	if err := e.app.SecretService.DeleteSecret(id); err != nil {
		return err
	}

	return nil
}

// getSecretsByTypeId - is executor for "get-secrets-by-type" case in Execute method.
func (e *Executor) getSecretsByTypeId(args []string) ([]model.SecretList, error) {
	switch len(args) - 1 {
	case 0:
		return nil, fmt.Errorf("validation error: Secret Type ID is missing")
	}

	id, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		return nil, convErr
	}

	list, err := e.app.SecretService.GetListOfSecretes(id)
	if err != nil {
		return nil, err
	}

	var models []model.SecretList
	for _, secret := range list {
		models = append(models, model.SecretList{
			Id:    int(secret.Id),
			Title: secret.Title,
		})
	}

	return models, nil
}

// getSecret - is executor for "get-secret" case in Execute method.
func (e *Executor) getSecret(args []string) (interface{}, error) {
	switch len(args) - 1 {
	case 0:
		return nil, fmt.Errorf("validation error: Secret ID is missing")
	}

	id, convErr := strconv.Atoi(args[1])
	if convErr != nil {
		return nil, convErr
	}

	secret, err := e.app.SecretService.GetSecret(id)
	if err != nil {
		st, _ := status.FromError(err)
		switch st.Code() {
		case codes.NotFound:
			return nil, fmt.Errorf(st.Message())
		default:
			return nil, err
		}
	}

	return secret, nil
}

// getSecretBinary - is executor for "get-secret-binary" case in Execute method.
func (e *Executor) getSecretBinary(args []string) error {
	switch len(args) - 1 {
	case 0:
		return fmt.Errorf("validation error: Secret ID and Path is missing")
	case 1:
		return fmt.Errorf("validation error: Path is missing")
	}

	id, errConv := strconv.Atoi(args[1])
	if errConv != nil {
		return errConv
	}

	err := e.app.SecretService.GetBinarySecret(id, args[2])
	if err != nil {
		return err
	}

	return nil
}

// editSecret - is executor for "edit-secret" case in Execute method.
func (e *Executor) editSecret(args []string, isForce bool) error {
	var (
		recordType int
		id         int
		converted  []byte
	)

	numArgs := len(args) - 1
	if numArgs >= 3 {
		recordType, errConv := strconv.Atoi(args[3])
		if errConv != nil {
			return errConv
		}

		id, errConv = strconv.Atoi(args[1])
		if errConv != nil {
			return errConv
		}

		switch recordType {
		case 1:
			switch numArgs {
			case 4:
				return fmt.Errorf("validation error: Password is missing")
			case 3:
				return fmt.Errorf("validation error: Login and Password is missing")
			default:
				converted, errConv = json.Marshal(model.LoginPassSecret{
					Id:         id,
					Title:      args[2],
					RecordType: 1,
					Login:      args[4],
					Password:   args[5],
				})
				if errConv != nil {
					return errConv
				}
			}
		case 2:
			switch numArgs {
			case 3:
				return fmt.Errorf("validation error: Text is missing")
			default:
				converted, errConv = json.Marshal(model.TextSecret{
					Id:         id,
					Title:      args[2],
					RecordType: 2,
					Text:       strings.Join(args[4:], " "),
				})
				if errConv != nil {
					return errConv
				}
			}
		case 3:
			switch numArgs {
			case 3:
				return fmt.Errorf("validation error: Filepath is missing")
			default:
				converted, errConv = json.Marshal(model.FileSecret{
					Id:         id,
					Title:      args[2],
					RecordType: 3,
					Path:       args[3],
				})
				if errConv != nil {
					return errConv
				}
			}
		case 4:
			switch numArgs {
			case 5:
				return fmt.Errorf("validation error: Due date is missing")
			case 4:
				return fmt.Errorf("validation error: CVV and Due date is missing")
			case 3:
				return fmt.Errorf("validation error: Card number, CVV and Due date is missing")
			default:
				converted, errConv = json.Marshal(model.CardSecret{
					Id:         id,
					Title:      args[2],
					RecordType: 4,
					CardNumber: args[4],
					CVV:        args[5],
					Due:        args[6],
				})
				if errConv != nil {
					return errConv
				}
			}
		}
	} else {
		return fmt.Errorf("validation error: Secret ID, Title, Secret Type ID and secret fields is missing")
	}

	if err := e.app.SecretService.EditSecret(id, args[3], recordType, string(converted), isForce); err != nil {
		st, _ := status.FromError(err)

		fmt.Println(st.Message())

		if st.Code() == codes.FailedPrecondition {
			fmt.Println("starting re-sync")

			e.app.Syncer.SyncAll()

			fmt.Println("re-sync ended")
		}

		return nil
	}

	return nil
}
