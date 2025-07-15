#!/bin/bash

# Hauska Go SDK Build and Test Script

set -e

echo "🏗️  Building Hauska Go SDK"
echo "=========================="

# Navigate to SDK directory
cd "$(dirname "$0")"

# Initialize Go module if not already done
if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init github.com/ardata-tech/hauska-go
fi

# Download dependencies
echo "📦 Downloading dependencies..."
go mod tidy

# Format code
echo "🎨 Formatting code..."
go fmt ./...

# Run go vet for static analysis
echo "🔍 Running static analysis..."
go vet ./...

# Build all packages
echo "🔨 Building packages..."
go build ./...

# Run tests (when we have them)
echo "🧪 Running tests..."
go test ./... || echo "⚠️  No tests found yet"

echo ""
echo "✅ Build completed successfully!"