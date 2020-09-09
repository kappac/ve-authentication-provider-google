package api

//go:generate protoc --proto_path=$GOPATH/src:$HOME/Projects/RestX:. --micro_out=../internal/api --go_out=../internal/api {file_name}.proto
