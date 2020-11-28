run:
	$(info running text-event-service)
	@go run ../main.go

test:
	$(info running text-event-service unit tests)
	@go test -timeout 30s ../dataservice ../handlers ../model ../utils ..

build:
	$(info running build text-event-service)
	@go build \
		-ldflags \
			"-X github.com/piopon/domesticity/services/text-event-service/handlers.VersionAPI=v1.0 \
			 -X github.com/piopon/domesticity/services/text-event-service/handlers.BuildTime=time \
			 -X github.com/piopon/domesticity/services/text-event-service/handlers.CommitSHA=sha \
			 -X github.com/piopon/domesticity/services/text-event-service/handlers.GoVersion=golang" \
		-o ../bin/ ..

swagger-doc:
	$(info generating swagger OpenAPI2 documentation)
	@swagger generate spec .. -o swagger.yaml --scan-models

swagger-sdk:
	$(info generating Golang SDK)
	@swagger -q generate client -f ./swagger.yaml -t ../sdk/ -A text-event-service