package vault

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Hash(ctx context.Context, pass string) (string, error)
	Validate(ctx context.Context, password, hash string) (bool, error)
}

// SERVICE
// Business logic
type vaultService struct {
}

func (v *vaultService) Hash(_ context.Context, pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}
func (v *vaultService) Validate(_ context.Context, password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, nil
	}
	return true, nil
}

func NewService() Service {
	return &vaultService{}
}
