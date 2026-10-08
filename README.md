protoc --go_out=plugins=grpc:advcalc advcalc.proto // OLD VERSION

// NEW VERSION
protoc \
  --go_out=. \
  --go-grpc_out=. \
  advcalc.proto

go run server.go
go run client.go