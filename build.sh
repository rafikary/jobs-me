#!/bin/bash

# Build script for Railway deployment

echo "🚀 Starting build process..."

# Install dependencies
echo "📦 Downloading Go modules..."
go mod download

# Build the application
echo "🔨 Building application..."
go build -o main .

echo "✅ Build completed successfully!"
