.PHONY: default
default:

	protoc -I=. -I=$(GOPATH)/src -I=$(GOPATH)/src/github.com/gogo/protobuf/protobuf --gogofaster_out=plugins=grpc:.\
		customers/customers.proto
		# \
		# Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
		# Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
		# Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
		# Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
		# Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:. \
