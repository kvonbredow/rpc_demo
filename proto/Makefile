PROTOC=protoc -I. --go_out=plugins=grpc:.

all: add_five.pb.go

add_five.pb.go: add_five.proto
	$(PROTOC) add_five.proto

clean:
	rm -f add_five.pb.go
