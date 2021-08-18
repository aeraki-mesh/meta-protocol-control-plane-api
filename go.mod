module github.com/aeraki-framework/go-control-plane

go 1.17

replace github.com/envoyproxy/go-control-plane => ../go-control-plane

require (
	github.com/cncf/xds/go v0.0.0-20210805033703-aa0b78936158
	github.com/envoyproxy/go-control-plane v0.9.9
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.27.1
)
