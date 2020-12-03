# BUILD SERVICE
SERVICE   := text-event-service
# UNIT TESTS
TEST_DIRS := ./src/dataservice ./src/handlers ./src/model ./src/utils ./src/.
COVER_TMP := cover.out
# HANDLING COMPILE-TIME VARIABLES
REPO      = github.com/piopon/domesticity/services/text-event-service/src
PACKAGE   = handlers
LD_FLAGS  = -ldflags '$(API_LD) $(DATE_LD) $(SHA_LD) $(GO_LD)'
# EVENT SERVICE VERSION INFORMATION
API_URL   := $(REPO)/$(PACKAGE).VersionAPI
API_VAL   := v1.0
API_LD    := -X "$(API_URL)=$(API_VAL)"
# EVENT SERVICE LAST COMMIT HASH
SHA_URL   := $(REPO)/$(PACKAGE).CommitSHA
SHA_VAL   := $(shell git log -1 --pretty=format:"%H")
SHA_LD    := -X "$(SHA_URL)=$(SHA_VAL)"
# EVENT SERVICE GOLANG BUILD VERSION
GO_URL    := $(REPO)/$(PACKAGE).GoVersion
GO_VAL    := $(shell go version)
GO_LD     := -X "$(GO_URL)=$(GO_VAL)"
# EVENT SERVICE BUILD DATE TIME
DATE_URL  := $(REPO)/$(PACKAGE).BuildTime
DATE_VAL  := $(shell date "+%Y.%m.%d %H:%M:%S")
DATE_LD   := -X "$(DATE_URL)=$(DATE_VAL)"

.PHONY: run
run:
	$(info $(SERVICE): compiling and running ...)
	@go run src/main.go

.PHONY: test
test: test-setup test-run test-clean

.PHONY: test-setup
test-setup:
	$(info $(SERVICE): preparing workspace for unit tests ...)
	@cp -r ./test/. ./src

.PHONY: test-run
.IGNORE: test-run
test-run:
	$(info $(SERVICE): running unit tests ...)
	@go test -count 1 -coverprofile $(COVER_TMP) -timeout 30s $(TEST_DIRS)
	@go tool cover -func $(COVER_TMP) | grep total:

.PHONY: test-clean
test-clean:
	$(info $(SERVICE): cleaning workspace after unit tests ...)
	@[ -f $(COVER_TMP) ] && rm $(COVER_TMP)
	@find ./src/ -name "*_test.go" -type f -delete

.PHONY: build
build:
	$(info $(SERVICE): building output file ...)
	@go build $(LD_FLAGS) -o ./bin/$(SERVICE) ./src

.PHONY: swagger-doc
swagger-doc:
	$(info $(SERVICE): generating swagger documentation ...)
	@swagger generate spec ./src -o ./resources/swagger.yaml --scan-models

.PHONY: swagger-sdk
swagger-sdk:
	$(info $(SERVICE): generating Golang SDK ...)
	@swagger -q generate client -f ./resources/swagger.yaml -t ./sdk/ -A $(SERVICE)

.PHONY: docker-build
docker-build:
	$(info $(SERVICE): building docker image ...)
	@docker build --rm -t $(SERVICE):$(API_VAL) --target bin .

.PHONY: docker-start
docker-start:
	$(info $(SERVICE): starting docker container ...)
	@docker run -d $(SERVICE):$(API_VAL)

.PHONY: docker-clean
docker-clean:
	$(info $(SERVICE): cleaning docker data ...)
	@docker rmi $(SERVICE):$(API_VAL)
