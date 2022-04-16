protoc --go_out=paths=source_relative:./ ./Products/products.proto

 protoc --go_out=plugins=grpc:. ./Products/products.proto