build:
	$(info running build text-event-service)
	@go build -o ../bin/ ..

swagger-doc:
	$(info generating swagger OpenAPI2 documentation)
	@swagger generate spec .. -o swagger.yaml --scan-models