#    \\ SPIKE: Secure your secrets with SPIFFE. — https://spike.ist/
#  \\\\\ Copyright 2024-present SPIKE contributors.
# \\\\\\\ SPDX-License-Identifier: Apache-2.0

# Run tests with coverage report and open HTML visualization
# Usage: make test/cover
# Executes all tests with race detection and coverage profiling
# Generates an HTML coverage report and opens it in the default browser
# Coverage data is temporarily stored in /tmp/coverage.out
# Flags: -v (verbose), -race (race detection), -buildvcs (include VCS info)
test/cover:
	go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...
	go tool cover -html=/tmp/coverage.out

# Run all tests with race detection
# Usage: make test
# Executes all tests in the project with verbose output and race detection
# Does not generate coverage reports (use test/cover for that)
# Flags: -v (verbose), -race (race detection), -buildvcs (include VCS info)
test:
	go test -v -race -buildvcs ./...

# Comprehensive code quality audit
# Usage: make audit
# Prerequisite: runs 'test' target first to ensure tests pass
# Performs multiple quality checks:
#   1. go mod tidy -diff: checks if go.mod needs tidying
#      (fails if changes needed)
#   2. go mod verify: verifies module dependencies haven't been tampered with
#   3. gofmt check: ensures all Go files are properly formatted
#   4. go vet: runs Go's built-in static analysis
#   5. staticcheck: runs advanced static analysis
#      (excluding ST1000, U1000 checks)
#   6. govulncheck: scans for known security vulnerabilities
# Designed for CI/CD pipelines to ensure code quality and security
audit: test
	go mod tidy -diff
	go mod verify
	test -z "$(shell gofmt -l .)"
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@latest run

# Interactive confirmation prompt
# Usage: make confirm && make some-destructive-action
# Prompts user for confirmation before proceeding with potentially destructive
# operations
# Returns successfully only if user explicitly types 'y', defaults to 'N' on
# empty input
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# Check for uncommitted changes in git repository
# Usage: make no-dirty
# Ensures the working directory is clean with no uncommitted changes
# Useful as a prerequisite for deployment or release targets
# Exits with error code if there are any modified, added, or untracked files
no-dirty:
	@test -z "$(shell git status --porcelain)"

# Check for available Go module upgrades
# Usage: make upgradeable
# Downloads and runs go-mod-upgrade tool to display available dependency updates
# Does not actually upgrade anything, only shows what could be upgraded
# Requires internet connection to fetch the tool and check for updates
upgradeable:
	@go run github.com/oligot/go-mod-upgrade@latest

# Clean up Go module dependencies and format code
# Usage: make tidy
# Performs two operations:
#   1. go mod tidy -v: removes unused dependencies and adds missing ones
#   2. go fmt ./...: formats all Go source files in the project
# Should be run before committing code changes
tidy:
	go mod tidy -v
	go fmt ./...
