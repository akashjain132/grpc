message:
	protoc --proto_path=./pb  --go_out=./server --plugin=/Users/akashjain/go/bin/protoc-gen-go product.proto

service:
	 protoc --proto_path=./pb ./pb/product.proto --plugin=/Users/akashjain/go/bin/protoc-gen-go-grpc --go-grpc_out=./server
