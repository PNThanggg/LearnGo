APP_NAME=fs
BIN_DIR=bin
MAIN_PATH=./cmd/fs

.PHONY: build run test clean fmt lint install

build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(MAIN_PATH)

run: build
	@echo "🚀 Running $(APP_NAME)..."
	@./$(BIN_DIR)/$(APP_NAME)

test:
	@echo "🧪 Running tests..."
	@go test ./... -v -cover

fmt:
	@echo "🎨 Formatting code..."
	@go fmt ./...

lint:
	@echo "Linting..."
	@golangci-lint run

install:
	@echo "Installing binary..."
	@go install $(MAIN_PATH)

clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BIN_DIR)