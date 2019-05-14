# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=tincwebgui
BINARY_UNIX=$(BINARY_NAME)_unix

# Build logic
all: clean deps dist build

build:
	$(GOBUILD) ./...
	cp ../../../../bin/$(BINARY_NAME) ./dist/$(BINARY_NAME)

dist:
	npm --prefix ./gui install
	npm --prefix ./gui run build
	cp -R ./gui/dist/ ./dist/

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


