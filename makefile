# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=tincwebgui
BINARY_UNIX=$(BINARY_NAME)_unix

# Build logic
all: clean buildGui deps build dist

build:
	$(GOBUILD) ./...

dist:
	mkdir -p dist
	cp ../../../../bin/$(BINARY_NAME) ./dist/$(BINARY_NAME)

buildGui:
	npm --prefix ./gui install
	npm --prefix ./gui run build
	go get -u github.com/shurcooL/vfsgen
	cd gui; go run assets_generate.go 

clean:
	$(GOCLEAN)
	rm -f -r ./dist
	rm -f -r ./gui/dist
	rm -f ./gui/assets_vfsdata.go
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run: build
	./$(BINARY_NAME)

deps:
	$(GOGET)


