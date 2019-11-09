module github.com/justcompile/midgard/worker

go 1.12

require (
	github.com/justcompile/midgard/common v0.0.0
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/grpc v1.23.1
)

replace github.com/justcompile/midgard/common => ../common
