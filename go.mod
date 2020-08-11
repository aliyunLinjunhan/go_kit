module go_kit

go 1.14

require (
	github.com/coreos/etcd v3.3.22+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/go-kit/kit v0.10.0
	github.com/go-log/log v0.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/gorilla/mux v1.7.3
	github.com/juju/ratelimit v1.0.1
	github.com/lucas-clemente/quic-go v0.17.3 // indirect
	github.com/marten-seemann/qtls v0.10.0 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/miekg/dns v1.1.31 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.7.1
	github.com/streadway/amqp v1.0.0
	go.uber.org/zap v1.15.0 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200802091954-4b90ce9b60b3 // indirect
	golang.org/x/text v0.3.3 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.0.0-20200731060945-b5fad4ed8dd6 // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485 // indirect
	google.golang.org/grpc v1.31.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
	honnef.co/go/tools v0.0.1-2020.1.5 // indirect
)

replace (
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200625212154-ddb9806d33ae
	golang.org/x/time => github.com/golang/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/protobuf => github.com/protocolbuffers/protobuf-go v1.25.0
)
