#!/bin/bash

# go2ban Automated Installation Script
# This script automates the installation of go2ban service

set -e  # Exit on any error

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root
check_root() {
    if [[ $EUID -eq 0 ]]; then
        print_error "This script should not be run as root. Please run as a regular user."
        exit 1
    fi
}

# Detect system architecture
detect_architecture() {
    ARCH=$(uname -m)
    case $ARCH in
        x86_64)
            echo "amd64"
            ;;
        aarch64|arm64)
            echo "arm64"
            ;;
        armv7l)
            echo "arm"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
}

# Simple installation from release
install_from_release() {
    print_status "Installing go2ban from release binary..."
    
    ARCH=$(detect_architecture)
    VERSION="v1.1.7"
    BINARY_URL="https://raw.githubusercontent.com/vv198x/go2ban/main/releases/${VERSION}/go2ban-linux-${ARCH}"
    
    print_status "Detected architecture: $ARCH"
    print_status "Downloading go2ban ${VERSION} for Linux ${ARCH}..."
    
    # Download binary
    if wget -q --spider "$BINARY_URL"; then
        wget -O go2ban "$BINARY_URL"
        chmod +x go2ban
        print_success "Binary downloaded successfully"
    else
        print_error "Binary not found for architecture $ARCH"
        print_status "Falling back to build from source..."
        return 1
    fi
    
    # Install binary
    sudo install -m 755 go2ban /usr/local/bin/
    rm go2ban
    
    print_success "go2ban binary installed to /usr/local/bin/go2ban"
    return 0
}

# Check and install Go
check_go() {
    print_status "Checking Go installation..."
    
    if command -v go &> /dev/null; then
        GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
        print_success "Go is installed: $GO_VERSION"
        
        # Check if version is >= 1.16 (required for io/fs package)
        MAJOR=$(echo $GO_VERSION | cut -d. -f1)
        MINOR=$(echo $GO_VERSION | cut -d. -f2)
        
        if [[ $MAJOR -eq 1 && $MINOR -lt 16 ]]; then
            print_warning "Go version $GO_VERSION is too old. go2ban requires Go 1.16 or higher."
            print_status "Installing Go 1.21.6..."
            install_go
        else
            print_success "Go version is compatible"
        fi
    else
        print_warning "Go is not installed. Installing Go 1.21.6..."
        install_go
    fi
}

# Install Go 1.21.6
install_go() {
    print_status "Downloading Go 1.21.6..."
    wget -q https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
    
    print_status "Installing Go..."
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
    
    # Add Go to PATH
    export PATH=$PATH:/usr/local/go/bin
    
    # Add to shell profile
    if [[ -f ~/.bashrc ]]; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    fi
    if [[ -f ~/.zshrc ]]; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
    fi
    
    # Clean up
    rm go1.21.6.linux-amd64.tar.gz
    
    print_success "Go 1.21.6 installed successfully"
    
    # Verify installation
    /usr/local/go/bin/go version
}

# Install required packages
install_dependencies() {
    print_status "Installing required packages..."
    
    # Detect package manager
    if command -v apt-get &> /dev/null; then
        sudo apt-get update
        sudo apt-get install -y make git wget
    elif command -v yum &> /dev/null; then
        sudo yum install -y make git wget
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y make git wget
    else
        print_error "Unsupported package manager. Please install make, git, and wget manually."
        exit 1
    fi
    
    print_success "Dependencies installed"
}

# Build go2ban
build_go2ban() {
    print_status "Building go2ban from source..."
    
    # Ensure we're using the correct Go version
    export PATH=/usr/local/go/bin:$PATH
    
    # Clean and tidy modules
    print_status "Cleaning Go modules cache..."
    go clean -modcache
    
    print_status "Running go mod tidy..."
    go mod tidy
    
    print_status "Building go2ban binary..."
    go build -o go2ban
    
    if [[ -f go2ban ]]; then
        print_success "go2ban built successfully"
    else
        print_error "Failed to build go2ban"
        exit 1
    fi
}

# Install go2ban service
install_service() {
    print_status "Installing go2ban service..."
    
    # Stop existing service if running
    sudo systemctl stop go2ban 2>/dev/null || true
    
    # Create directories
    sudo mkdir -p /var/log/go2ban /etc/go2ban
    
    # Install binary (if not already installed from release)
    if [[ ! -f /usr/local/bin/go2ban ]]; then
        if [[ -f go2ban ]]; then
            sudo install -m 755 go2ban /usr/local/bin/
        else
            print_error "go2ban binary not found"
            exit 1
        fi
    fi
    
    # Install configuration
    if [[ -f deploy/go2ban.conf ]]; then
        sudo install -m 644 deploy/go2ban.conf /etc/go2ban/
        print_success "Configuration installed to /etc/go2ban/go2ban.conf"
    else
        print_warning "Configuration file not found. You'll need to create it manually."
    fi
    
    # Install systemd service
    if [[ -f deploy/go2ban.service ]]; then
        sudo install -m 644 deploy/go2ban.service /etc/systemd/system/
        sudo systemctl daemon-reload
        print_success "Systemd service installed"
    else
        print_warning "Systemd service file not found. You'll need to create it manually."
    fi
    
    print_success "go2ban service installed successfully"
}

# Configure go2ban
configure_go2ban() {
    print_status "Opening configuration file for editing..."
    
    if [[ -f /etc/go2ban/go2ban.conf ]]; then
        echo -e "${YELLOW}The configuration file will open in nano editor.${NC}"
        echo -e "${YELLOW}Press Ctrl+X, then Y, then Enter to save and exit.${NC}"
        echo -e "${YELLOW}Press Ctrl+C to skip configuration editing.${NC}"
        
        read -p "Press Enter to continue or Ctrl+C to skip..."
        
        sudo nano /etc/go2ban/go2ban.conf
    else
        print_warning "Configuration file not found at /etc/go2ban/go2ban.conf"
    fi
}

# Start service
start_service() {
    echo -e "${YELLOW}Do you want to start and enable the go2ban service? (y/n)${NC}"
    read -p "Your choice: " choice
    
    case $choice in
        [Yy]* )
            print_status "Starting go2ban service..."
            sudo systemctl enable go2ban
            sudo systemctl start go2ban
            
            # Check service status
            if sudo systemctl is-active --quiet go2ban; then
                print_success "go2ban service started successfully"
                print_status "Service status:"
                sudo systemctl status go2ban --no-pager -l
            else
                print_error "Failed to start go2ban service"
                print_status "Check logs with: sudo journalctl -u go2ban -f"
            fi
            ;;
        [Nn]* )
            print_status "Service not started. You can start it later with:"
            echo "  sudo systemctl enable go2ban"
            echo "  sudo systemctl start go2ban"
            ;;
        * )
            print_warning "Invalid choice. Service not started."
            ;;
    esac
}

# Show installation options
show_options() {
    echo -e "${GREEN}================================${NC}"
    echo -e "${GREEN}  go2ban Installation Script${NC}"
    echo -e "${GREEN}================================${NC}"
    echo ""
    echo -e "${BLUE}Choose installation method:${NC}"
    echo "  1) Quick install from release (recommended)"
    echo "     - Downloads pre-built binary"
    echo "     - Fastest installation"
    echo "     - No Go compilation required"
    echo ""
    echo "  2) Build from source"
    echo "     - Compiles from source code"
    echo "     - Requires Go 1.16+"
    echo "     - Takes longer but ensures latest code"
    echo ""
    echo "  3) Exit"
    echo ""
}

# Main installation function
main() {
    show_options
    
    read -p "Enter your choice (1-3): " choice
    
    case $choice in
        1)
            echo ""
            print_status "Starting quick installation from release..."
            check_root
            install_dependencies
            if install_from_release; then
                install_service
                configure_go2ban
                start_service
            else
                print_warning "Release installation failed, falling back to build from source..."
                check_go
                build_go2ban
                install_service
                configure_go2ban
                start_service
            fi
            ;;
        2)
            echo ""
            print_status "Starting build from source..."
            check_root
            install_dependencies
            check_go
            build_go2ban
            install_service
            configure_go2ban
            start_service
            ;;
        3)
            print_status "Installation cancelled"
            exit 0
            ;;
        *)
            print_error "Invalid choice. Please run the script again."
            exit 1
            ;;
    esac
    
    echo ""
    echo -e "${GREEN}================================${NC}"
    echo -e "${GREEN}  Installation Complete!${NC}"
    echo -e "${GREEN}================================${NC}"
    echo ""
    echo -e "${BLUE}Useful commands:${NC}"
    echo "  sudo systemctl status go2ban    # Check service status"
    echo "  sudo systemctl start go2ban     # Start service"
    echo "  sudo systemctl stop go2ban      # Stop service"
    echo "  sudo systemctl restart go2ban   # Restart service"
    echo "  sudo journalctl -u go2ban -f    # View logs"
    echo ""
    echo -e "${BLUE}Configuration:${NC}"
    echo "  sudo nano /etc/go2ban/go2ban.conf"
    echo ""
}

# Run main function
main "$@" 