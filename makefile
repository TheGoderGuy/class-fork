
local:
	go run app/services/sales-api/main.go

# ==============================================================================
# Building containers

# $(shell git rev-parse --short HEAD)
VERSION := 1.0

all: sales

sales:
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor