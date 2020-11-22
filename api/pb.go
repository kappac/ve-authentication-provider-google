package pb

//go:generate protoc --proto_path=../vendor --proto_path=. --go_out=plugins=grpc:../internal/pb service.proto
