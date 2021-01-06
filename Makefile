.PHONY: proto
proto:
	protoc -I=. --go_out=. --go-grpc_out=. internal/portdomainsvc/grpc/port.proto
