# Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"). You
# may not use this file except in compliance with the License. A copy of
# the License is located at
#
# 	http://aws.amazon.com/apache2.0/
#
# or in the "license" file accompanying this file. This file is
# distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF
# ANY KIND, either express or implied. See the License for the specific
# language governing permissions and limitations under the License.

ROOT := $(shell pwd)

all: local-build

GO_VERSION := 1.15
SCRIPT_PATH := $(ROOT)/scripts/:${PATH}
SOURCES := $(shell find . -name '*.go')
BINARY_NAME := local-container-endpoints
IMAGE_REPO_NAME := amazon/amazon-ecs-local-container-endpoints
LOCAL_BINARY := bin/local/${BINARY_NAME}

# AMD_DIR and ARM_DIR correspond to ARCH_SUFFIX env var set in each CodeBuild
# project which is used in the image tags.
AMD_DIR := amd64
ARM_DIR := arm64

AMD_BINARY := bin/${AMD_DIR}/${BINARY_NAME}
ARM_BINARY := bin/${ARM_DIR}/${BINARY_NAME}
VERSION := $(shell cat VERSION)
AGENT_VERSION_COMPATIBILITY := $(shell cat AGENT_VERSION_COMPATIBILITY)
TAG := $(VERSION)-agent$(AGENT_VERSION_COMPATIBILITY)-compatible

.PHONY: generate
generate: $(SOURCES)
	PATH=$(SCRIPT_PATH) go generate ./...

.PHONY: local-build
local-build: $(LOCAL_BINARY)

.PHONY: build-local-image
build-local-image:
	docker run -v $(shell pwd):/usr/src/app/src/github.com/awslabs/amazon-ecs-local-container-endpoints \
		--workdir=/usr/src/app/src/github.com/awslabs/amazon-ecs-local-container-endpoints \
		--env GOPATH=/usr/src/app \
		--env ECS_RELEASE=cleanbuild \
		golang:$(GO_VERSION) make ${LOCAL_BINARY}
	docker build --build-arg ARCH_DIR=local -t $(IMAGE_REPO_NAME):latest-local .

# Build binaries for each architecture into their own subdirectories
$(LOCAL_BINARY): $(SOURCES)
	PATH=${PATH} golint ./local-container-endpoints/...
	./scripts/build_binary.sh ./bin/local
	@echo "Built local-container-endpoints"

$(AMD_BINARY): $(SOURCES)
	@mkdir -p ./bin/$(AMD_DIR)
	TARGET_GOOS=linux GOARCH=amd64 ./scripts/build_binary.sh ./bin/$(AMD_DIR)
	@echo "Built local-container-endpoints for linux-amd64"

$(ARM_BINARY): $(SOURCES)
	@mkdir -p ./bin/$(ARM_DIR)
	TARGET_GOOS=linux GOARCH=arm64 ./scripts/build_binary.sh ./bin/$(ARM_DIR)
	@echo "Built local-container-endpoints for linux-arm64"

# Relies on ARCH_SUFFIX environment variable which is set in the build
# environment (e.g. CodeBuild project). Value will either be amd64 or arm64.
.PHONY: build-image
build-image:
	docker run -v $(shell pwd):/usr/src/app/src/github.com/awslabs/amazon-ecs-local-container-endpoints \
		--workdir=/usr/src/app/src/github.com/awslabs/amazon-ecs-local-container-endpoints \
		--env GOPATH=/usr/src/app \
		--env ECS_RELEASE=cleanbuild \
		golang:$(GO_VERSION) make bin/${ARCH_SUFFIX}/${BINARY_NAME}
	docker build --build-arg ARCH_DIR=$(ARCH_SUFFIX) -t $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX) .
	docker tag $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX) $(IMAGE_REPO_NAME):$(TAG)-$(ARCH_SUFFIX)
	docker tag $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX) $(IMAGE_REPO_NAME):$(VERSION)-$(ARCH_SUFFIX)

.PHONY: publish-dockerhub
publish-dockerhub:
	docker push $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX)
	docker push $(IMAGE_REPO_NAME):$(TAG)-$(ARCH_SUFFIX)
	docker push $(IMAGE_REPO_NAME):$(VERSION)-$(ARCH_SUFFIX)

.PHONY: create-manifests
create-manifests:
	docker manifest create $(IMAGE_NAME):latest $(IMAGE_NAME):latest-$(ARCH_SUFFIX) $(IMAGE_NAME):latest-$(ARCH_SUFFIX)
	docker manifest create $(IMAGE_NAME):$(TAG) $(IMAGE_NAME):$(TAG)-$(ARCH_SUFFIX) $(IMAGE_NAME):$(TAG)-$(ARCH_SUFFIX)
	docker manifest create $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):$(VERSION)-$(ARCH_SUFFIX) $(IMAGE_NAME):$(VERSION)-$(ARCH_SUFFIX)

.PHONY: annotate-manifests
annotate-manifests:
	docker manifest annotate --arch $(ARCH_SUFFIX) $(IMAGE_NAME):latest-$(ARCH_SUFFIX)

.PHONY: push-manifests
push-manifests:
	docker manifest push $(IMAGE_NAME):latest
	docker manifest push $(IMAGE_NAME):$(TAG)
	docker manifest push $(IMAGE_NAME):$(VERSION)

.PHONY: inspect-manifests
push-manifests:
	docker manifest inspect $(IMAGE_NAME):latest
	docker manifest inspect $(IMAGE_NAME):$(TAG)
	docker manifest inspect $(IMAGE_NAME):$(VERSION)

.PHONY: test
test:
	go test -mod=vendor -timeout=120s -v -cover ./local-container-endpoints/...

.PHONY: functional-test
functional-test:
	go test -mod=vendor -timeout=120s -v -tags functional -cover ./local-container-endpoints/handlers/functional_tests/...

.PHONY: integ
integ: build-local-image
	docker build -t amazon-ecs-local-container-endpoints-integ-test:latest -f ./integ/Dockerfile .
	docker-compose --file ./integ/docker-compose.yml up --abort-on-container-exit

.PHONY: verify
verify:
	docker pull $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX)
	docker run -d -p 8000:80 -v /var/run:/var/run -v $(HOME)/.aws/:/home/.aws/ -e "ECS_LOCAL_METADATA_PORT=80" -e "HOME=/home" -e "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=${AWS_CONTAINER_CREDENTIALS_RELATIVE_URI}" --name endpoints $(IMAGE_REPO_NAME):latest-$(ARCH_SUFFIX)
	curl -s localhost:8000/creds
	curl -s localhost:8000/v2/stats
	curl -s localhost:8000/v2/metadata
	curl -s localhost:8000/v3
	curl -s localhost:8000/v4
	docker stop endpoints && docker rm endpoints
	docker pull $(IMAGE_REPO_NAME):$(TAG)-$(ARCH_SUFFIX)
	docker run -d -p 8000:80 -v /var/run:/var/run -v $(HOME)/.aws/:/home/.aws/ -e "ECS_LOCAL_METADATA_PORT=80" -e "HOME=/home" -e "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=${AWS_CONTAINER_CREDENTIALS_RELATIVE_URI}" --name endpoints $(IMAGE_REPO_NAME):$(TAG)-$(ARCH_SUFFIX)
	curl -s localhost:8000/creds
	curl -s localhost:8000/v2/stats
	curl -s localhost:8000/v3
	curl -s localhost:8000/v4
	docker stop endpoints && docker rm endpoints
	docker pull $(IMAGE_REPO_NAME):$(VERSION)-$(ARCH_SUFFIX)
	docker run -d -p 8000:80 -v /var/run:/var/run -v $(HOME)/.aws/:/home/.aws/ -e "ECS_LOCAL_METADATA_PORT=80" -e "HOME=/home" -e "AWS_CONTAINER_CREDENTIALS_RELATIVE_URI=${AWS_CONTAINER_CREDENTIALS_RELATIVE_URI}" --name endpoints $(IMAGE_REPO_NAME):$(VERSION)-$(ARCH_SUFFIX)
	curl -s localhost:8000/creds
	curl -s localhost:8000/v2/stats
	curl -s localhost:8000/v2/metadata
	curl -s localhost:8000/v3
	curl -s localhost:8000/v4
	docker stop endpoints && docker rm endpoints

.PHONY: clean
clean:
	rm bin/local/local-container-endpoints
