BINARY_NAME=osc-utility
VERSION=$(shell grep -o '"[0-9]\+\.[0-9]\+\.[0-9]\+"' main.go | tr -d '"')
BUILD_DIR=build
LDFLAGS=-ldflags "-s -w"

.PHONY: all clean darwin linux windows

all: darwin linux windows

clean:
	rm -rf $(BUILD_DIR)

# Create build directory
$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

# Darwin (macOS) builds
darwin: $(BUILD_DIR)
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_amd64
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_darwin_arm64

# Linux builds
linux: $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_amd64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_linux_arm64

# Windows builds
windows: $(BUILD_DIR)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_amd64.exe
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)_$(VERSION)_windows_arm64.exe

# Create compressed archives for each binary
compress: all
	cd $(BUILD_DIR) && for file in $(BINARY_NAME)_*; do \
		if [[ $$file != *.zip && $$file != *.tar.gz ]]; then \
			if [[ $$file == *windows* ]]; then \
				zip "$$file.zip" "$$file"; \
			else \
				tar czf "$$file.tar.gz" "$$file"; \
			fi; \
		fi; \
	done 