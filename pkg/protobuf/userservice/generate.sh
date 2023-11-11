rm -rf *.pb.go
rm -rf ./userservice
rm -rf ./gw/*

go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

protoc --go_out . --go_opt paths=source_relative \
   --proto_path=../../ \
   --proto_path=./ \
   user.proto

protoc --go_out ./gw --go_opt paths=source_relative \
   --go-grpc_out ./gw --go-grpc_opt paths=source_relative \
   --grpc-gateway_out=./gw \
   --grpc-gateway_opt logtostderr=true \
   --grpc-gateway_opt generate_unbound_methods=true \
   --openapiv2_out=./gw \
   --openapiv2_opt logtostderr=true \
   --openapiv2_opt use_go_templates=true \
   --openapiv2_opt grpc_api_configuration=rules.yaml \
   --openapiv2_opt openapi_configuration=swagger.yaml \
   --proto_path=../../ \
   --proto_path=./ \
   user.proto

