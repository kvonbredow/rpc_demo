GOCC=go build
MAKE=make

all: client server

client: client.go proto/add_five.pb.go
	$(GOCC) client.go

server: server.go proto/add_five.pb.go
	$(GOCC) server.go

proto/add_five.pb.go: proto/add_five.proto
	$(MAKE) -C proto/

clean:
	$(MAKE) clean -C proto/
	rm -f client server
