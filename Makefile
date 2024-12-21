.PHONY: all
all:
	@echo "Please specify a target. e.g., make build"

.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -rf ./build

.PHONY: build
build: clean
	@echo "Building..."
	go build -o ./build/slides
