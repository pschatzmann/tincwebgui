# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=tincwebgui
BINARY_UNIX=$(BINARY_NAME)_unix

# Build logic
all: clean deps buildGui build dist

build:
	cd gui; go run assets_generate.go 
	$(GOBUILD) ./...

dist:
	mkdir -p dist
	cp ../../../../bin/$(BINARY_NAME) ./dist/$(BINARY_NAME)

buildGui:
	npm --prefix ./gui install
	npm --prefix ./gui run build

clean:
	$(GOCLEAN)
	rm -f -r ./dist
	rm -f -r ./gui/dist
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET)
	go get -u github.com/shurcooL/vfsgen
	go get -u github.com/shurcooL/vfsgen/cmd/vfsgendev


