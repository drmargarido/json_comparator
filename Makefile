BIN=
EXAMPLES_FOLDER=comparator/test_examples

all:
	go build -o bin/comparator cmd/comparator/main.go

clean:
	rm -rf bin

test: all
	@echo ""
	@echo "---- Run unit tests ----"
	go test -cover comparator/json_test.go comparator/json.go

	@echo ""
	@echo "---- Run command tests ----"
	@bash cli_test.bash

	@echo "---- Tests passed successfully ----"
