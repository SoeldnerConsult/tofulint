#!/usr/bin/env bash

set -eo pipefail


detect_arch() {
    case "$(uname -m)" in
        i386|i686) echo "386" ;;
        x86_64)    echo "amd64" ;;
        arm64|aarch64) echo "arm64" ;;
        *) echo "unknown" ;;
    esac
}

ARCH=$(detect_arch)
echo "Detected architecture: $ARCH"


detect_os() {
    case "$(uname -s)" in
        Darwin*)  echo "darwin_${ARCH}" ;;
        MINGW64*|MSYS_NT*) echo "windows_${ARCH}" ;;
        *) echo "linux_${ARCH}" ;;
    esac
}

OS_TARGET=$(detect_os)
echo "Detected OS target: $OS_TARGET"

echo -e "\n====================================================\n"


TMP_DIR=$(mktemp -d -t tofulint.XXXXXXXX)
ZIP_FILE="$TMP_DIR/tofulint.zip"
EXECUTABLE="$TMP_DIR/tofulint"


if [ -z "$TFLINT_VERSION" ] || [ "$TFLINT_VERSION" = "latest" ]; then
    echo "Fetching latest TofuLint..."
    DOWNLOAD_URL="https://github.com/SoeldnerConsult/tofulint/releases/latest/download/tofulint_${OS_TARGET}.zip"
else
    echo "Fetching TofuLint version $TFLINT_VERSION..."
    DOWNLOAD_URL="https://github.com/SoeldnerConsult/tofulint/releases/download/${TFLINT_VERSION}/tofulint_${OS_TARGET}.zip"
fi

curl -sSL -o "$ZIP_FILE" "$DOWNLOAD_URL" || { echo "Download failed"; exit 1; }
echo "Download completed."

echo -e "\nUnpacking $ZIP_FILE..."
unzip -o "$ZIP_FILE" -d "$TMP_DIR"


if [[ "$OS_TARGET" == windows* ]]; then
    DEST_DIR="${TFLINT_INSTALL_PATH:-/bin}"
else
    DEST_DIR="${TFLINT_INSTALL_PATH:-/usr/local/bin}"
fi

echo "Installing TofuLint to $DEST_DIR ..."


if [[ ! -w "$DEST_DIR" && "$OS_TARGET" != windows* ]]; then
    SUDO="sudo"
else
    SUDO=""
fi

$SUDO mkdir -p "$DEST_DIR"
$SUDO install -m 0755 "$EXECUTABLE" "$DEST_DIR" || { echo "Installation failed"; exit 1; }


echo "Cleaning up temporary files..."
rm -rf "$TMP_DIR"

echo -e "\n====================================================\n"
echo "Installed Tofulint version:"
"$DEST_DIR/tofulint" -v