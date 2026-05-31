#!/bin/bash
# Setup script for Linux/Mac

echo "🚀 Job Application Tracker - Setup Script"
echo "=========================================="
echo ""

# Check if Go is installed
echo "Checking prerequisites..."
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install from: https://go.dev/dl/"
    exit 1
fi
echo "✅ Go version: $(go version)"

# Check if .env exists
if [ ! -f ".env" ]; then
    echo "⚠️  .env file not found. Creating from .env.example..."
    cp .env.example .env
    echo "✅ .env file created. Please edit .env with your database credentials."
    echo "Press any key to continue..."
    read -n 1 -s
fi

# Install dependencies
echo ""
echo "📦 Installing dependencies..."
go mod download
if [ $? -ne 0 ]; then
    echo "❌ Failed to download dependencies"
    exit 1
fi
echo "✅ Dependencies installed"

# Build application
echo ""
echo "🔨 Building application..."
go build -o main main.go
if [ $? -ne 0 ]; then
    echo "❌ Build failed"
    exit 1
fi
echo "✅ Build successful"

# Make executable
chmod +x main

# Run application
echo ""
echo "🎯 Starting server..."
echo "Application will be available at: http://localhost:3000"
echo "Press Ctrl+C to stop the server"
echo ""

./main
