.SILENT:

generate:
    protoc -I user/ -I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis user/user.proto --go_out=plugins=grpc:user
    protoc -I user/ -I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis user/user.proto --grpc-gateway_out=logtostderr=true:user

    echo "Generate done"

run-server:
    go run ./server/main.go

run-client:
    go run ./client/main.go 

validate:
    swagger validate user/todo.swagger.json