# Define the output directory
OUTPUT_DIR := ../build

# Define the source file
SOURCE_FILE := ./main.go

# Define the target executable name
TARGET := $(OUTPUT_DIR)/tla2typst

# Define the Go compiler
GO := go

# Define the build command
BUILD_CMD := $(GO) build -o $(TARGET) $(SOURCE_FILE)

# Define the clean command
CLEAN_CMD := rm -rf $(OUTPUT_DIR)

# Define the default target
.PHONY: all
all: clean build

# Define the build target
.PHONY: build
build:
	@mkdir -p $(OUTPUT_DIR)
	cd ./tla2typst && $(BUILD_CMD)

# Define the clean target
.PHONY: clean
clean:
	$(CLEAN_CMD)