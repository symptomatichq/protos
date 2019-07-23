.PHONY: default
default:
	@go mod vendor
	@protoc -I=. -I=$(GOPATH)/src -I=./vendor --gogofaster_out=plugins=grpc:. \
		customers/customers.proto
