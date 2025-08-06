# Go Wallet Generator

A comprehensive cryptocurrency wallet generator that provides both **Command-Line Interface (CLI)** and **Web-based GUI** for generating and managing cryptocurrency wallets with support for Bitcoin and Ethereum addresses. This tool uses BIP39 mnemonic generation and securely stores wallet data locally.

## 🚀 Features

- **Wallet Management**: Create, list, get, and delete wallets
- **Address Generation**: Generate Bitcoin and Ethereum addresses from wallets
- **Secure Storage**: Encrypted mnemonic storage using system keyring
- **BIP39 Compliance**: Standard mnemonic phrase generation
- **Dual Interface**: Both CLI and Web-based GUI for different user preferences
- **Interactive CLI**: User-friendly command-line interface for power users
- **Modern Web GUI**: Beautiful web interface for desktop users
- **Cross-platform**: Works on any device with a web browser

## 📋 Prerequisites

- **Operating System**: GNOME Ubuntu (required for keyring functionality)
- **Go**: Version 1.24 or higher
- **Dependencies**: All required packages are managed via `go.mod`

## 🛠️ Installation

### From Snap Store (Recommended)

The easiest way to install Go Wallet Generator is through the Snap Store:

```bash
# Install the snap
sudo snap install go-wallet-generator

# Run CLI version
wallet-generator-cli

# Run Web GUI version
wallet-generator-web
# Then open http://localhost:3456 in your browser
```

### From Source

1. Install dependencies:
```bash
go mod download
```

2. Run the application:

**CLI Version (Recommended for power users):**
```bash
make run
# or
go run cmd/main.go
```

**Web GUI Version (Recommended for desktop users):**
```bash
make frontend-run
# or
go run frontend/main.go
```

Then open your browser and navigate to: `http://localhost:3456`

## 📖 Usage

This project provides two interfaces to suit different user preferences:

### CLI Version (Command Line Interface)

Perfect for power users, developers, and automation. The application starts an interactive CLI prompt. Enter commands in the following format:

#### Wallet Commands

**Create a new wallet:**
```bash
wallet create -n <wallet_name>
```
Creates a new wallet with a unique name. The 12-word mnemonic phrase will be copied to your clipboard.

**List all wallets:**
```bash
wallet list
```
Displays all created wallets with their names and creation dates.

**Get wallet details:**
```bash
wallet get -n <wallet_name>
```
Retrieves and displays detailed information about a specific wallet.

**Delete a wallet:**
```bash
wallet delete -n <wallet_name>
```
Permanently deletes a wallet and all its associated addresses.

#### Address Commands

**Generate a new address:**
```bash
address create -w <wallet_name> -c <coin_symbol>
```
Generates a new address for the specified coin (btc or eth) in the given wallet.

**Supported coins:**
- `btc` - Bitcoin address
- `eth` - Ethereum address

**List wallet addresses:**
```bash
address list -w <wallet_name>
```
Displays all addresses associated with a specific wallet.

### Web GUI Version

The web interface provides a modern, user-friendly interface with the following features:

- **Left Panel**: Wallet management
  - Create new wallets
  - View existing wallets
  - Delete selected wallets

- **Right Panel**: Address management
  - Select wallet and coin type
  - Generate new addresses
  - View all addresses for selected wallet

## 💡 Examples

### CLI Examples

```bash
# Start the CLI application
> make run

# Create a new wallet
> wallet create -n my_wallet

# Generate a Bitcoin address
> address create -w my_wallet -c btc

# Generate an Ethereum address
> address create -w my_wallet -c eth

# List all wallets
> wallet list

# List addresses in a wallet
> address list -w my_wallet

# Get wallet details
> wallet get -n my_wallet

# Delete a wallet
> wallet delete -n my_wallet
```

### Web GUI Examples

**Usage:**
1. Run `make frontend-run` to start the web server
2. Open your browser and go to `http://localhost:3456`
3. Enter a wallet name and click "Create Wallet"
4. Select a wallet from the list to manage its addresses
5. Choose a coin type and click "Generate Address"
6. View generated addresses in the right panel

## 🔒 Security Features

- **Encrypted Storage**: Mnemonics are encrypted before storage
- **System Keyring**: Uses GNOME keyring for secure credential storage
- **Clipboard Integration**: Mnemonics are automatically copied to clipboard for secure handling
- **Local Storage**: All wallet data is stored locally in `wallets.json`
- **Secure Display**: Mnemonic phrases are hidden by default with toggle visibility
- **Auto-hide**: Mnemonics automatically hide after 30 seconds for security

## 📁 Project Structure

```
go_wallet_generator/
├── address/          # Address generation for different cryptocurrencies
├── bip39/           # BIP39 mnemonic generation and validation
├── cli/             # Command-line interface handlers
├── cmd/             # Main CLI application entry point
├── crypto/          # Encryption/decryption utilities
├── frontend/        # Web-based graphical user interface
│   ├── main.go      # Web server entry point
│   ├── app/         # Web application logic and API
│   └── static/      # Static HTML/CSS/JS files
├── wallet/          # Wallet management and storage
├── go.mod           # Go module dependencies
├── Makefile         # Build and run commands
└── README.md        # This file
```

## ⚠️ Important Notes

- **GNOME Ubuntu Required**: This application requires GNOME Ubuntu for keyring functionality
- **Mnemonic Safety**: The 12-word mnemonic phrase is automatically copied to your clipboard when creating a wallet. Keep it safe and secure
- **Local Storage**: Wallet data is stored in `wallets.json` in the project directory (your mnemonic is encrypted dont worry)
- **Unique Names**: Wallet names must be unique within the system
- **Web Server**: The GUI runs on `http://localhost:3456` by default
- **Backup**: Always backup your mnemonic phrases securely - they cannot be recovered if lost

## 🚨 Error Handling

The application provides clear error messages for common issues:

- Invalid commands
- Missing required parameters
- Duplicate wallet names
- Non-existent wallets
- Unsupported coin types
- Network connectivity issues
- Keyring access problems

## 🛡️ Security Best Practices

1. **Keep Mnemonics Safe**: Store your 12-word mnemonic phrases in a secure location
2. **Never Share**: Never share your mnemonic phrases with anyone
3. **Multiple Backups**: Create multiple secure backups of your mnemonic phrases
4. **Offline Storage**: Consider storing mnemonics offline for maximum security
5. **Regular Backups**: Regularly backup your wallet data
6. **Secure Environment**: Run the application in a secure environment

## 🔧 Development

To contribute to this project:

1. Ensure you're on GNOME Ubuntu
2. Install Go 1.24+
3. Run `go mod download` to install dependencies
4. Use `make run` to run the CLI application
5. Use `make frontend-run` to run the web GUI application

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📞 Support

If you encounter any issues or have questions:

1. Check the error handling section above
2. Ensure you're running on GNOME Ubuntu
3. Verify all prerequisites are installed
4. Check that the web server is accessible at `http://localhost:3456`

## 🔄 Version History

- **v1.0.0**: Initial release with CLI and Web GUI support
- Support for Bitcoin and Ethereum address generation
- BIP39 mnemonic generation
- Secure encrypted storage
- Modern web interface
