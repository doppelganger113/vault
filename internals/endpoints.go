package vault

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type Endpoints struct {
	HashEndpoint     endpoint.Endpoint
	ValidateEndpoint endpoint.Endpoint
}

func (e Endpoints) Hash(ctx context.Context, password string) (string, error) {
	req := hashRequest{Password: password}
	resp, err := e.HashEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	hashRes := resp.(hashResponse)
	if hashRes.Err != "" {
		return "", errors.New(hashRes.Err)
	}

	return hashRes.Hash, nil
}

func (e Endpoints) Validate(ctx context.Context, hash, password string) (bool, error) {
	req := validateRequest{
		Password: password,
		Hash:     hash,
	}
	resp, err := e.ValidateEndpoint(ctx, req)
	if err != nil {
		return false, err
	}

	validateRes := resp.(validateResponse)
	if validateRes.Err != "" {
		return false, errors.New(validateRes.Err)
	}

	return validateRes.Valid, nil
}

type hashRequest struct {
	Password string `json:"password"`
}
type hashResponse struct {
	Hash string `json:"hash"`
	Err  string `json:"err,omitempty"`
}

func decodeHashRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req hashRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func MakeHashEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(hashRequest)
		val, err := s.Hash(ctx, req.Password)
		if err != nil {
			return hashResponse{
				Hash: "",
				Err:  err.Error(),
			}, nil
		}

		return hashResponse{
			Hash: val,
			Err:  "",
		}, nil
	}
}

type validateRequest struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}
type validateResponse struct {
	Valid bool   `json:"valid"`
	Err   string `json:"err,omitempty"`
}

func decodeValidateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req validateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
func MakeValidateEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(validateRequest)
		val, err := s.Validate(ctx, req.Password, req.Hash)
		if err != nil {
			return validateResponse{
				Valid: false,
				Err:   err.Error(),
			}, nil
		}

		return validateResponse{
			Valid: val,
			Err:   "",
		}, nil
	}
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
