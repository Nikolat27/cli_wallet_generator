# 📦 Snap Publishing Guide for Go Wallet Generator

This guide will walk you through publishing your Go Wallet Generator to the Snap Store with full GNOME keyring support.

## 🔧 Prerequisites

### System Requirements
- **Ubuntu 20.04+** or **Ubuntu-based distribution**
- **GNOME desktop environment** (for keyring support)
- **snapcraft** installed
- **Snap Store developer account**

### Install Snapcraft
```bash
sudo snap install snapcraft --classic
```

## 🏗️ Building the Snap

### 1. Build Locally
```bash
# Make the build script executable (if not already)
chmod +x build-snap.sh

# Run the build script
./build-snap.sh
```

The script will:
- ✅ Check prerequisites
- ✅ Clean previous builds
- ✅ Download Go dependencies
- ✅ Build the snap with GNOME keyring support
- ✅ Offer local installation for testing

### 2. Manual Build (Alternative)
```bash
# Clean previous builds
snapcraft clean

# Build the snap
snapcraft

# The snap file will be generated (e.g., go-wallet-generator_1.0.0_amd64.snap)
```

## 🧪 Local Testing

### Install for Testing
```bash
# Install the snap in devmode for testing
sudo snap install --dangerous --devmode go-wallet-generator_*.snap
```

### Test Both Interfaces
```bash
# Test CLI version
wallet-generator-cli

# Test Web GUI version (runs on localhost:3456)
wallet-generator-web

# Then open your browser to: http://localhost:3456
```

### Test Keyring Integration
1. Create a wallet using either interface
2. Verify the mnemonic is stored securely (no password prompts)
3. Check that wallets persist across app restarts
4. Verify fallback to file-based storage if keyring unavailable

### Uninstall Test Version
```bash
sudo snap remove go-wallet-generator
```

## 🌐 Publishing to Snap Store

### 1. Create Snap Store Account
1. Go to [snapcraft.io](https://snapcraft.io/)
2. Click "Sign up" and create an account
3. Verify your email address

### 2. Register Snap Name
```bash
# Login to snapcraft
snapcraft login

# Register your snap name (must be unique)
snapcraft register go-wallet-generator
```

### 3. Upload Your Snap
```bash
# Upload the snap file
snapcraft upload go-wallet-generator_1.0.0_amd64.snap

# Set release channels
snapcraft release go-wallet-generator <revision> stable
```

### 4. Configure Snap Store Listing
1. Go to [snapcraft.io/snaps](https://snapcraft.io/snaps)
2. Find your snap and click "Manage"
3. Update:
   - **Description** (already set in snapcraft.yaml)
   - **Screenshots** (recommended)
   - **Contact information**
   - **Website** (if you have one)

## 🔒 Security & Permissions

### Required Snap Interfaces
Your snap requests these permissions for full functionality:

- **`gnome-keyring`** - Access to GNOME keyring for secure mnemonic storage
- **`password-manager-service`** - Alternative keyring access
- **`desktop`** / **`desktop-legacy`** - Desktop integration
- **`network`** / **`network-bind`** - Web GUI server functionality
- **`home`** - Access to user's home directory for wallet storage
- **`browser-support`** - Web GUI browser integration

### Security Review
- Snaps undergo automatic security review
- Manual review may be required for sensitive interfaces
- Your snap uses `strict` confinement for security

## 📋 Snap Store Listing Best Practices

### 1. Optimize Description
```yaml
summary: Cryptocurrency wallet generator with CLI and Web GUI
description: |
  A comprehensive cryptocurrency wallet generator that provides both Command-Line 
  Interface (CLI) and Web-based GUI for generating and managing cryptocurrency 
  wallets with support for Bitcoin and Ethereum addresses. This tool uses BIP39 
  mnemonic generation and securely stores wallet data using GNOME keyring.
```

### 2. Add Screenshots
- CLI interface showing wallet creation
- Web GUI showing the wallet management interface
- Address generation examples

### 3. Set Categories
- **Finance** (primary)
- **Utilities** (secondary)

### 4. Add Keywords
- cryptocurrency
- wallet
- bitcoin
- ethereum
- BIP39
- mnemonic
- keyring

## 🚀 Release Process

### 1. Version Management
```bash
# Update version in snapcraft.yaml before building
version: '1.0.1'
```

### 2. Release Channels
- **`edge`** - Development builds
- **`beta`** - Pre-release testing
- **`candidate`** - Release candidates
- **`stable`** - Production releases

### 3. Release Command
```bash
# Release to specific channel
snapcraft release go-wallet-generator <revision> stable
```

## 🔍 Monitoring & Maintenance

### 1. Check Snap Metrics
- Visit [snapcraft.io/snaps](https://snapcraft.io/snaps)
- Monitor downloads, ratings, and reviews

### 2. Handle User Feedback
- Respond to reviews and issues
- Update documentation based on common questions

### 3. Regular Updates
- Keep dependencies updated
- Release security patches promptly
- Add new features based on user requests

## ⚠️ Important Notes

### GNOME Keyring Support
- Your snap is configured for **strict confinement** with proper keyring access
- Users on non-GNOME systems will fall back to file-based storage
- Test thoroughly on different desktop environments

### File Locations in Snap
- **Wallet data**: `~/snap/go-wallet-generator/current/wallets.json`
- **Secret keys**: GNOME keyring or `~/snap/go-wallet-generator/current/.wallet_secret.key`

### Security Considerations
- Never store sensitive data in logs
- Ensure mnemonic phrases are properly encrypted
- Follow cryptocurrency security best practices

## 🆘 Troubleshooting

### Build Issues
```bash
# If build fails, clean and retry
snapcraft clean
rm -rf parts/ prime/ stage/
./build-snap.sh
```

### Keyring Access Issues
- Ensure `gnome-keyring` interface is connected
- Check if GNOME desktop is running
- Verify D-Bus session is available

### Publishing Issues
- Ensure snap name is registered
- Check account permissions
- Verify snap passes security review

## 📞 Support

- **Snapcraft Documentation**: https://snapcraft.io/docs
- **Snap Store**: https://snapcraft.io/
- **Ubuntu Forums**: https://discourse.ubuntu.com/c/snapcraft

---

**Happy Publishing! 🎉**

Your Go Wallet Generator will provide secure cryptocurrency wallet management to users across the Ubuntu ecosystem!