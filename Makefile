dependencies:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	go install github.com/fullstorydev/grpcui/cmd/grpcui@latest

generate:
	protoc --proto_path=proto proto/*.proto --go_out=src/ --go-grpc_out=src/

grpcui:
	grpcui --plaintext 127.0.0.1:8080