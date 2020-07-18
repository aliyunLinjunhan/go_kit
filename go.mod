module go_kit

go 1.14

require (
	github.com/go-kit/kit v0.10.0
	github.com/gorilla/mux v1.7.3
	github.com/juju/ratelimit v1.0.1
	github.com/prometheus/client_golang v1.7.1
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
)

replace (
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200625212154-ddb9806d33ae
	golang.org/x/time => github.com/golang/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/protobuf => github.com/protocolbuffers/protobuf-go v1.25.0
)
