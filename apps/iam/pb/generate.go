package pb

//go:generate protoc -I=$PROTOC_INCLUDE -I=. --grpc-gateway_opt paths=source_relative --go_out . --go-grpc_out=. --grpc-gateway_out . --openapiv2_out .  --openapiv2_opt logtostderr=true *.proto
