```
Generate pb.go file :

protoc -I user/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
user/user.proto \
--go_out=plugins=grpc:user
```

###
Generate .pb.gw.go file :

```
protoc -I user/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
user/user.proto \
--grpc-gateway_out=logtostderr=true:user
```

###
Generate swagger :

```
protoc -I user/ \
-I /Users/vietwow/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.9.5/third_party/googleapis \
user/user.proto \
--swagger_out=logtostderr=true:user
```

###
Run grpc-client :
```
SERVER=localhost:50051 go run client/main.go

DB_HOST=localhost:3306 DB_USER=root DB_PASSWORD=newhacker DB_SCHEMA=grab go run server/main.go
```

Run grpc-gateway :
```
go run rest/main.go
```

###

```
swagger validate user/user.swagger.json
```
