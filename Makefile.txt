############################# Main targets #############################
# Run all checks and build.
install: clean staticcheck errcheck workflowcheck bins
########################################################################

##### Variables ######
UNIT_TEST_DIRS := $(sort $(dir $(shell find . -name "*_test.go")))
MAIN_FILES := $(shell find . -name "main.go")
MOD_FILES_DIR := $(sort $(dir $(shell find . -name "go.mod")))
TEST_TIMEOUT := 20s
COLOR := "\e[1;36m%s\e[0m\n"
BIN_DIR := $(shell pwd)/bin

define NEWLINE


endef

##### Targets ######
staticcheck:
	@printf $(COLOR) "Run static check..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@staticcheck ./...

errcheck:
	@printf $(COLOR) "Run error check..."
	@go install github.com/kisielk/errcheck@latest
	@errcheck ./...

workflowcheck:
	@printf $(COLOR) "Run workflow check..."
	@go install go.temporal.io/sdk/contrib/tools/workflowcheck@latest
	@workflowcheck -show-pos ./...
	
clean:
	rm -rf bin
	
ci-build: staticcheck errcheck workflowcheck
