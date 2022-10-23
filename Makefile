.PHONY: proto_gen
proto_gen:
	protoc -I ./proto  -I=${GOPATH}/src \
	  --go_out ./proto --go_opt paths=source_relative \
	  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
	  --openapiv2_out=logtostderr=true:./swagger-ui \
	  ./proto/company/*.proto ./proto/common/*.proto
