```
protoc -I user/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
user/user.proto \
--go_out=plugins=grpc:user
```

###

```
protoc -I user/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
user/user.proto \
--grpc-gateway_out=logtostderr=true:user
```

###

```
SERVER=localhost:50051 go run client/main.go

DB_HOST=localhost:3306 DB_USER=root DB_PASSWORD=newhacker DB_SCHEMA=grab go run server/main.go

go run rest/main.go
```

###

```
swagger validate todo.swagger.json
```
