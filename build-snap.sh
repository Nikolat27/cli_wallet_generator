#!/bin/bash

# Go Wallet Generator Snap Build Script
# This script builds a snap package with GNOME keyring support

set -e

echo "🚀 Building Go Wallet Generator Snap..."

# Check if snapcraft is installed
if ! command -v snapcraft &> /dev/null; then
    echo "❌ snapcraft is not installed. Please install it first:"
    echo "   sudo snap install snapcraft --classic"
    exit 1
fi

# Check if we're on a supported system
if [ ! -f /etc/os-release ]; then
    echo "❌ Cannot determine OS. This build script is designed for Ubuntu/Debian systems."
    exit 1
fi

# Source OS info
source /etc/os-release

echo "📋 Build Information:"
echo "   OS: $NAME $VERSION"
echo "   Architecture: $(uname -m)"
echo "   Working Directory: $(pwd)"

# Clean any previous builds
echo "🧹 Cleaning previous build artifacts..."
snapcraft clean 2>/dev/null || true
rm -rf parts/ prime/ stage/ *.snap 2>/dev/null || true

# Ensure we have Go dependencies
echo "📦 Downloading Go dependencies..."
go mod download

# Build the snap
echo "🔨 Building snap package..."
if ! snapcraft; then
    echo "❌ Snap build failed!"
    exit 1
fi

# Find the generated snap file
SNAP_FILE=$(ls -t *.snap 2>/dev/null | head -n1)

if [ -z "$SNAP_FILE" ]; then
    echo "❌ No snap file was generated!"
    exit 1
fi

echo "✅ Snap built successfully!"
echo "📦 Snap file: $SNAP_FILE"
echo "📏 Size: $(du -h $SNAP_FILE | cut -f1)"

# Test installation locally (optional)
read -p "🤔 Do you want to install the snap locally for testing? (y/N): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "🔧 Installing snap locally..."
    sudo snap install --dangerous --devmode "$SNAP_FILE"
    echo "✅ Snap installed locally!"
    echo ""
    echo "🧪 Test commands:"
    echo "   wallet-generator-cli"
    echo "   wallet-generator-web"
    echo ""
    echo "To uninstall: sudo snap remove go-wallet-generator"
fi

echo ""
echo "🎉 Build complete!"
echo "📄 To publish to Snap Store:"
echo "   1. Create a Snap Store account at https://snapcraft.io/"
echo "   2. Run: snapcraft login"
echo "   3. Run: snapcraft upload $SNAP_FILE"
echo "   4. Visit https://snapcraft.io/snaps to manage your snap"