.PHONY: proto
proto:
	protoc -I=. --go_out=. --go-grpc_out=. internal/portdomain/proto/port.proto
