package vault

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"vault/internals/pb"
)

func NewGRPCServer(endpoints Endpoints) pb.VaultServer {
	return &grpcServer{
		hash: grpctransport.NewServer(
			endpoints.HashEndpoint,
			DecodeGRPCHashRequest,
			DecodeGRPCHashResponse,
		),
		validate: grpctransport.NewServer(
			endpoints.ValidateEndpoint,
			DecodeGRPCValidateRequest,
			DecodeGRPCValidateResponse,
		),
	}
}

type grpcServer struct {
	hash     grpctransport.Handler
	validate grpctransport.Handler
}

func (s *grpcServer) Hash(ctx context.Context, req *pb.HashRequest) (*pb.HashResponse, error) {
	_, resp, err := s.hash.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.HashResponse), nil
}

func (s *grpcServer) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	_, resp, err := s.validate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.ValidateResponse), nil
}

func EncodeGRPCHashRequest(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(hashRequest)
	return &pb.HashRequest{Password: req.Password}, nil
}

func DecodeGRPCHashRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.HashRequest)
	return hashRequest{Password: req.Password}, nil
}
func EncodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(hashResponse)
	return &pb.HashResponse{Hash: res.Hash, Err: res.Err}, nil
}
func DecodeGRPCHashResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.HashResponse)
	return hashResponse{Hash: res.Hash, Err: res.Err}, nil
}
func EncodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(validateRequest)
	return &pb.ValidateRequest{Password: req.Password, Hash: req.Hash}, nil
}
func DecodeGRPCValidateRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ValidateRequest)
	return validateRequest{Password: req.Password, Hash: req.Hash}, nil
}
func EncodeGRPCValidateResponse(ctx context.Context,
	r interface{}) (interface{}, error) {
	res := r.(validateResponse)
	return &pb.ValidateResponse{Valid: res.Valid}, nil
}
func DecodeGRPCValidateResponse(ctx context.Context, r interface{}) (interface{}, error) {
	res := r.(*pb.ValidateResponse)
	return validateResponse{Valid: res.Valid}, nil
}
