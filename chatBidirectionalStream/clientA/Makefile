# Makefile for a Golang program

# Biến
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Tên chương trình và thư mục đầu ra
BINARY_NAME = hihi
BINARY_UNIX = $(BINARY_NAME)_unix

# Mặc định
all: deps build

# Biên dịch chương trình
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Biên dịch chương trình cho Unix
build-unix:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

# Chạy chương trình
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)

# Dọn dẹp
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

# Cài đặt các dependencies
deps:
	$(GOGET) -u github.com/example/dependency

# Chạy các test
test:
	$(GOTEST) -v ./...

# Chạy lint
lint:
	golangci-lint run
