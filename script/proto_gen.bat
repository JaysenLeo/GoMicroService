cd ../models/protos
protoc.exe --micro_out=../ --go_out=../ ./Users.proto
protoc.exe --micro_out=../ --go_out=../ ./UserService.proto
protoc-go-inject-tag.exe -input=../Users.pb.go