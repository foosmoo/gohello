all: build

build:
	go install github.com/foosmoo/gohello/hello

test: build
	go test github.com/foosmoo/gohello/hello

run: test
	$(GOPATH)/bin/hello

clean:
	rm -f $(GOPATH)/bin/hello
