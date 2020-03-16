package vault

import (
	"context"
	"testing"
)

func TestService_Hash(t *testing.T) {
	srv := NewService()
	ctx := context.Background()
	hashed, err := srv.Hash(ctx, "password")
	if err != nil {
		t.Fatalf("Hash: %s", err)
	}
	ok, err := srv.Validate(ctx, "password", hashed)
	if err != nil {
		t.Fatalf("Validate: %s", err)
	}
	if !ok {
		t.Fatal("Expected true from Validate")
	}
	ok, err = srv.Validate(ctx, "wrong password", hashed)
	if err != nil {
		t.Fatalf("Validate: %s", err)
	}
	if ok {
		t.Error("expected false from Validate")
	}
}
