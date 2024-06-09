#!/bin/sh

# Run Go linter
echo "Running golangci-lint..."
golangci-lint run
if [ $? -ne 0 ]; then
    echo "Linting failed. Please fix the issues before committing."
    exit 1
fi

# Run Go tests
echo "Running go test..."
cd test
go test ./...go test -run TestOrder
if [ $? -ne 0 ]; then
    echo "Tests failed. Please fix the issues before committing."
    exit 1
fi

echo "All checks passed. Proceeding with commit."
exit 0
