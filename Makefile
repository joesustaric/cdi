# Borrowed from:
# https://github.com/silven/go-example/blob/master/Makefile

BINARY = cdi
VET_REPORT = vet.report
GOARCH = amd64

VERSION?=0.1.1
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=joesustaric
BUILD_DIR=${GOPATH}/cdi/cmd/cdi
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.VERSION=${VERSION} -X main.COMMIT=${COMMIT} -X main.BRANCH=${BRANCH}"

# Build the project
all: clean test vet linux darwin windows

linux:
	cd ${BUILD_DIR}; \
	GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ../../bin/${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ../../bin/${BINARY}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ../../bin/${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

update:
	go get -u ./...

test:
	# Compile the binary into the GOPATH
	cd cmd/cdi/; go install

	cd ../..
	# Just doing basic test now
	go clean -testcache
	go test ./... -v

test-with-report:
	# Compile the binary into the GOPATH
	cd cmd/cdi/; go install

	cd ../..
	# Just doing basic test now
	go clean -testcache
	go test ./... -v -json | go-test-report

vet:
	# To fix later
	# -cd ${BUILD_DIR}; \
	# go vet ./... > ${VET_REPORT} 2>&1 ; \
	# cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VET_REPORT}
	-rm -f ${BINARY}-*

.PHONY: linux darwin windows test vet fmt clean
