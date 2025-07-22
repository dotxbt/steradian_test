all: build
run: 
	@go run cmd/core/main.go

build: 
	@go build -o ./build/core cmd/core/main.go

clean:
	@echo "📦 Removing..."
	@rm -rf build/
	
.PHONY: all run build