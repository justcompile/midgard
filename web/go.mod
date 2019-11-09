module github.com/justcompile/midgard/web

go 1.13

require (
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-pg/pg/v9 v9.0.0-beta.7
	github.com/gorilla/websocket v1.4.1
	github.com/grafov/bcast v0.0.0-20190217190352-1447f067e08d // indirect
	github.com/justcompile/midgard/common v0.0.0
	github.com/sirupsen/logrus v1.4.2
	google.golang.org/grpc v1.23.1
)

replace github.com/justcompile/midgard/common => ../common
