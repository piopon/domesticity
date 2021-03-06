# BUILD SERVICE
SERVICE   := text-event-service
PORT      := 9999
# UNIT TESTS
TEST_DIRS := ./src/dataservice ./src/handlers ./src/model ./src/utils ./src/.
COVER_TMP := cover.out
# HANDLING COMPILE-TIME VARIABLES
REPO      = github.com/piopon/domesticity/services/text-event-service/src
PACKAGE   = handlers
LD_FLAGS  = '$(API_LD) $(DATE_LD) $(SHA_LD) $(GO_LD)'
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
# DOCKER
CREATE    := echo container not found - creating ...
RESTART   := echo container was found - restarting ...
MISSING   := echo container not found ...
BUILD_ARG := --build-arg LD_FLAGS=$(LD_FLAGS) --build-arg EXPOSED_PORTS=$(PORT)
ENV_FILE  := .env

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
	@go build -ldflags $(LD_FLAGS) -o ./bin/$(SERVICE) ./src

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
	@docker build --rm -t $(SERVICE):$(API_VAL) $(BUILD_ARG) --target bin .

.PHONY: docker-start
docker-start:
	$(info $(SERVICE): starting docker container ...)
	@docker ps -a --format "{{.Names}}" -f "name=$(SERVICE)" | grep $(SERVICE) > /dev/null            \
	    && { $(RESTART) ; docker start $(SERVICE) ; }                                                 \
	    || { $(CREATE)  ; docker run -d --name $(SERVICE) -p $(PORT):$(PORT) $(SERVICE):$(API_VAL) ; }

.PHONY: docker-stop
docker-stop:
	$(info $(SERVICE): stopping docker container ...)
	@docker ps -a --format "{{.Names}}" -f "name=$(SERVICE)" | grep $(SERVICE) > /dev/null            \
	    && docker stop $(SERVICE)                                                                     \
	    || $(MISSING)

.PHONY: docker-clean
docker-clean: docker-stop
	$(info $(SERVICE): cleaning docker data ...)
	@docker rm $(SERVICE) ; docker rmi $(SERVICE):$(API_VAL)

.PHONY: docker-up
docker-up:
	$(info $(SERVICE): starting docker compose network ...)
	@[ ! -f $(ENV_FILE) ] && echo LD_FLAGS=\'$(LD_FLAGS)\' > $(ENV_FILE) || true
	@docker-compose up -d

.PHONY: docker-down
docker-down:
	$(info $(SERVICE): stopping docker compose network ...)
	@docker-compose down
	@[ -f $(ENV_FILE) ] && rm $(ENV_FILE) || true
