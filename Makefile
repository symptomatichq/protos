.PHONY: default
default:
	@go mod vendor
	@protoc -I=. -I=$(GOPATH)/src -I=./vendor --gogofaster_out=plugins=grpc:. \
		authentication/authentication.proto
	@protoc -I=. -I=$(GOPATH)/src -I=./vendor --gogofaster_out=plugins=grpc:. \
		customers/customers.proto
	@protoc -I=. -I=$(GOPATH)/src -I=./vendor --gogofaster_out=plugins=grpc:. \
		emails/emails.proto
