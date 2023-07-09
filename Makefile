submodule:
	git submodule add git@github.com:apinanyogaratnam/jwt-protobuf

proto:
	protoc --go_out=plugins=grpc:jwt jwt-protobuf/protos/*.proto
