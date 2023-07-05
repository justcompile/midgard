module github.com/justcompile/midgard/worker

go 1.12

require (
	github.com/justcompile/midgard/common v0.0.0
	google.golang.org/grpc v1.53.0
)

replace github.com/justcompile/midgard/common => ../common
