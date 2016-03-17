all: test

build:
	go install github.com/foosmoo/gohello/util

clean:
	rm -f $(GOPATH)/pkg/linux_amd64/github.com/foosmoo/gohello/util.a

test: build
	go test

