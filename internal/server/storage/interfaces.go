//go:generate mockgen -source=./interfaces.go -destination=./mock/storage.go -package=storagemock
package storage

import (
	"context"

	"secretKeeper/internal/server/model"
)

type UserServerStorage interface {
	// Create - create a new model.User in storage.
	Create(ctx context.Context, user model.User) (model.User, error)
	// GetByLoginAndPassword - returns model.User from storage.
	GetByLoginAndPassword(ctx context.Context, user model.User) (model.User, error)
	// DeleteUser - deletes a user from storage.
	DeleteUser(ctx context.Context, user model.User) (model.User, error)
}

type SecretTypeServerStorage interface {
	// GetSecretTypes - returns list of model.SecretType from storage.
	GetSecretTypes(ctx context.Context) ([]model.SecretType, error)
}

type SecretServerStorage interface {
	// CreateSecret - creates new model.Secret in storage.
	CreateSecret(ctx context.Context, secret model.Secret) (model.Secret, error)
	// GetSecret - gets a model.Secret from storage.
	GetSecret(ctx context.Context, secret model.Secret) (model.Secret, error)
	// DeleteSecret - deletes a model.Secret from storage.
	DeleteSecret(ctx context.Context, secret model.Secret) (model.Secret, error)
	// EditSecret - updates a model.Secret in storage.
	EditSecret(ctx context.Context, secret model.Secret, isForce bool) (model.Secret, error)
	// GetListOfSecretByType - returns a list of []model.Secret from storage.
	GetListOfSecretByType(ctx context.Context, secretType model.SecretType, user model.User) ([]model.Secret, error)
}
