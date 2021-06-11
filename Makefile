BUILD_NAME = max.work
BIN_DIR = bin

build: bin
	go build -o ./$(BIN_DIR)/$(BUILD_NAME) .

clean: bin
	rm -f $(BIN_DIR)/**

bin:
	mkdir -p $(BIN_DIR)