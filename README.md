rm -rf ../user/user.pb.go ; protoc -I ../user ../user/user.proto --go_out=plugins=grpc:../user
go run main.go
