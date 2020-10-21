package pb

//go:generate protoc --proto_path=../vendor:$HOME/Projects/RestX:. --go_out=plugins=grpc:../internal/pb pb.proto
