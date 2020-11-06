build:
	$(info running build text-event-service)
	@go build -o ../bin/ ..

swagger-doc:
	$(info generating swagger OpenAPI2 documentation)
	@swagger generate spec .. -o swagger.yaml --scan-models

swagger-sdk:
	$(info generating Golang SDK)
	@swagger -q generate client -f ./swagger.yaml -t ../sdk/ -A text-event-service