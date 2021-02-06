GOPATH:=$(shell go env GOPATH)

# ----- Install -----

.PHONY: install
install:
	go mod download

# ----- Test -----

BUILDENV := CGO_ENABLED=0
TESTFLAGS := -short -cover -v
SERVICE=fpn-openvpn

.PHONY: test
test:
	$(BUILDENV) go test $(TESTFLAGS) ./...


# ----- Build -----

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(SERVICE)


.PHONY: dev
dev:
	go run .

# ----- CI -----

DOCKER_REGISTRY=freemiumvpn
DOCKER_CONTAINER_NAME=$(SERVICE)
DOCKER_REPOSITORY=$(DOCKER_REGISTRY)/$(DOCKER_CONTAINER_NAME)
SHA8=$(shell echo $(GITHUB_SHA) | cut -c1-8)

docker-image:
	docker build --rm \
		--build-arg SERVICE=$(SERVICE) \
		--tag $(DOCKER_REPOSITORY):local .
		

ci-docker-auth:
	@echo "${DOCKER_PASSWORD}" | docker login --username "${DOCKER_USERNAME}" --password-stdin

ci-docker-build:
	@docker build --rm \
		--build-arg SERVICE=$(SERVICE) \
		--tag $(DOCKER_REPOSITORY):$(SHA8) \
		--tag $(DOCKER_REPOSITORY):latest .

ci-docker-build-push: ci-docker-build
	@docker push $(DOCKER_REPOSITORY):$(SHA8)
	@docker push $(DOCKER_REPOSITORY):latest
