.PHONY: test clean deps lint ${PROG}

PROG=conttest

all: test ${PROG} deps lint

deps:
	go get .

${PROG}: *.go
	go build

test:
	echo No tests yet   # go test

lint:
	golint -set_exit_status .
	staticcheck .

clean:
	rm ${PROG}

install: ${PROG}
	go install .