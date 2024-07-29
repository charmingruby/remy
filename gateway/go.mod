module github.com/charmingruby/remy-gateway

go 1.22.3

require (
	github.com/charmingruby/remy-common v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	google.golang.org/grpc v1.65.0
)

require (
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240722135656-d784300faade // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/charmingruby/remy-common => ../common
