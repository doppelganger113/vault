proto:
	protoc -I internals/pb/ internals/pb/vault.proto --go_out=plugins=grpc:internals/pb
default:
	echo "Unknown command"