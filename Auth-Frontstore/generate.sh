protoc services/auth/authpb/auth.proto --go-grpc_out=. --go_out=.

#Server Start
go run services/auth/server/server.go

#Client Start
go run grpc/client/main.go

#MongoDB Start
mongod --dbpath="C:\Program Files\MongoDB\Server\4.2\data".

#DynamoDB start
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb