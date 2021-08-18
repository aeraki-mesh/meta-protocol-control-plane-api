module github.com/aeraki-framework/go-control-plane

go 1.17

replace github.com/envoyproxy/go-control-plane =>  ../go-control-plane

require (
	github.com/cncf/xds/go v0.0.0-20210805033703-aa0b78936158
	github.com/envoyproxy/go-control-plane v0.9.9
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/census-instrumentation/opencensus-proto v0.2.1 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
)
