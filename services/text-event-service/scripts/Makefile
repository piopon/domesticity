REPO = github.com/piopon/domesticity/services/text-event-service
PACKAGE = handlers
LD_FLAGS = -ldflags '$(API_LD) $(DATE_LD) $(SHA_LD) $(GO_LD)'
# EVENT SERVICE VERSION INFORMATION
API_URL  := $(REPO)/$(PACKAGE).VersionAPI
API_VAL  := v1.0
API_LD   := -X "$(API_URL)=$(API_VAL)"
# EVENT SERVICE LAST COMMIT HASH
SHA_URL  := $(REPO)/$(PACKAGE).CommitSHA
SHA_VAL  := $(shell git log -1 --pretty=format:"%H")
SHA_LD   := -X "$(SHA_URL)=$(SHA_VAL)"
# EVENT SERVICE GOLANG BUILD VERSION
GO_URL   := $(REPO)/$(PACKAGE).GoVersion
GO_VAL   := $(shell go version)
GO_LD    := -X "$(GO_URL)=$(GO_VAL)"
# EVENT SERVICE BUILD DATE TIME
DATE_URL := $(REPO)/$(PACKAGE).BuildTime
DATE_VAL := $(shell date "+%Y.%m.%d %H:%M:%S")
DATE_LD  := -X "$(DATE_URL)=$(DATE_VAL)"

run:
	$(info running text-event-service)
	@go run ../main.go

test:
	$(info running text-event-service unit tests)
	@go test -timeout 30s ../dataservice ../handlers ../model ../utils ..

build:
	$(info running build text-event-service)
	@go build $(LD_FLAGS) -o ../bin/ ..

swagger-doc:
	$(info generating swagger OpenAPI2 documentation)
	@swagger generate spec .. -o swagger.yaml --scan-models

swagger-sdk:
	$(info generating Golang SDK)
	@swagger -q generate client -f ./swagger.yaml -t ../sdk/ -A text-event-service