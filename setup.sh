#!/usr/bin/env bash
# Install script for Actio (Linux / macOS) after downloading a release package.

set -euo pipefail

BIN_NAME="actio"
if [[ "$(uname -s)" == "MINGW"* || "$(uname -s)" == "MSYS"* || "$(uname -s)" == "CYGWIN"* ]]; then
  echo "This script is intended for Linux/macOS. On Windows, use setup.bat."
  exit 1
fi

if [[ ! -f "$BIN_NAME" ]]; then
  echo "error: binary '$BIN_NAME' not found in current directory."
  echo "Make sure you've extracted the release archive into this folder."
  exit 1
fi

echo "Found binary: $(pwd)/$BIN_NAME"
echo
read -r -p "Install to /usr/local/bin so it's on your PATH? (requires sudo) [y/N]: " yn
if [[ "$yn" =~ ^[Yy]$ ]]; then
  if [[ ! -w /usr/local/bin ]]; then
    echo "Requesting sudo permission to move the binary into /usr/local/bin..."
    sudo mv "$BIN_NAME" /usr/local/bin/
  else
    mv "$BIN_NAME" /usr/local/bin/
  fi
  echo "Installed to /usr/local/bin/$BIN_NAME"
  echo "You can now run it via: actio"
else
  echo "No changes made. You can run it from this directory with: ./$(basename "$BIN_NAME")"
  echo "Or add this directory to your PATH (for the current shell):" 
  echo "  export PATH=\"\$PATH:$(pwd)\""
fi
